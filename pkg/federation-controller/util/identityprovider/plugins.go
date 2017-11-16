/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package identityprovider

import (
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/golang/glog"
)

// factory is a function that returns a Interface interface.
// The config parameter provides an io.Reader handler to the factory in
// order to load specific configurations. If no configuration is provided
// the parameter is nil.
type factory func(config io.Reader) (Interface, error)

// All registered identity providers.
var providersMutex sync.Mutex
var providers = make(map[string]factory)

// RegisterIdentityProvider registers a identityprovider.factory by name.  This
// is expected to happen during startup.
func RegisterIdentityProvider(name string, f factory) {
	providersMutex.Lock()
	defer providersMutex.Unlock()
	if _, found := providers[name]; found {
		glog.Fatalf("Identity provider %q was registered twice", name)
	}
	glog.V(1).Infof("Registered Identity provider %q", name)
	providers[name] = f
}

// GetIdentityProvider creates an instance of the named identity provider, or nil if
// the name is not known.  The error return is only used if the named provider
// was known but failed to initialize. The config parameter specifies the
// io.Reader handler of the configuration file for the identity provider, or nil
// for no configuration.
func GetIdentityProvider(name string, config io.Reader) (Interface, error) {
	providersMutex.Lock()
	defer providersMutex.Unlock()
	f, found := providers[name]
	if !found {
		return nil, nil
	}
	return f(config)
}

// Returns a list of identity identity providers.
func RegisteredIdentityProviders() []string {
	registeredProviders := make([]string, len(providers))
	i := 0
	for provider := range providers {
		registeredProviders[i] = provider
		i = i + 1
	}
	return registeredProviders
}

// InitIdentityProvider creates an instance of the named identity provider.
func InitIdentityProvider(name string, configFilePath string) (Interface, error) {
	var provider Interface
	var err error

	if name == "" {
		glog.Info("No identity provider specified.")
		return nil, nil
	}

	if configFilePath != "" {
		var config *os.File
		config, err = os.Open(configFilePath)
		if err != nil {
			return nil, fmt.Errorf("Couldn't open identity provider configuration %s: %#v", configFilePath, err)
		}

		defer config.Close()
		provider, err = GetIdentityProvider(name, config)
	} else {
		// Pass explicit nil so plugins can actually check for nil. See
		// "Why is my nil error value not equal to nil?" in golang.org/doc/faq.
		provider, err = GetIdentityProvider(name, nil)
	}

	if err != nil {
		return nil, fmt.Errorf("could not init identity provider %q: %v", name, err)
	}
	if provider == nil {
		return nil, fmt.Errorf("unknown DNS provider %q", name)
	}

	return provider, nil
}
