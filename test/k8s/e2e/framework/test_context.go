/*
Copyright 2016 The Kubernetes Authors.

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

package framework

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/golang/glog"
	"github.com/onsi/ginkgo/config"
	"github.com/spf13/viper"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
	"k8s.io/kubernetes/pkg/cloudprovider"
)

const defaultHost = "http://127.0.0.1:8080"

type TestContextType struct {
	KubeConfig         string
	KubeContext        string
	KubeAPIContentType string
	KubeVolumeDir      string
	CertDir            string
	Host               string
	// TODO: Deprecating this over time... instead just use gobindata_util.go , see #23987.
	RepoRoot string

	Provider       string
	CloudConfig    CloudConfig
	KubectlPath    string
	OutputDir      string
	ReportDir      string
	ReportPrefix   string
	Prefix         string
	MinStartupPods int
	// Timeout for waiting for system pods to be running
	SystemPodsStartupTimeout time.Duration
	UpgradeTarget            string
	EtcdUpgradeStorage       string
	EtcdUpgradeVersion       string
	UpgradeImage             string
	GCEUpgradeScript         string
	// SystemdServices are comma separated list of systemd services the test framework
	// will dump logs for.
	SystemdServices          string
	ImageServiceEndpoint     string
	MasterOSDistro           string
	NodeOSDistro             string
	VerifyServiceAccount     bool
	DeleteNamespace          bool
	DeleteNamespaceOnFailure bool
	AllowedNotReadyNodes     int
	CleanStart               bool
	// Currently supported values are 'hr' for human-readable and 'json'. It's a comma separated list.
	OutputPrintType string
	// NodeSchedulableTimeout is the timeout for waiting for all nodes to be schedulable.
	NodeSchedulableTimeout time.Duration
	// CreateTestingNS is responsible for creating namespace used for executing e2e tests.
	// It accepts namespace base name, which will be prepended with e2e prefix, kube client
	// and labels to be applied to a namespace.
	CreateTestingNS CreateTestingNSFn
	// If set to true test will dump data about the namespace in which test was running.
	DumpLogsOnFailure bool
	// Disables dumping cluster log from master and nodes after all tests.
	DisableLogDump bool
	// Path to the GCS artifacts directory to dump logs from nodes. Logexporter gets enabled if this is non-empty.
	LogexporterGCSPath string
	// If the garbage collector is enabled in the kube-apiserver and kube-controller-manager.
	GarbageCollectorEnabled bool
	// FeatureGates is a set of key=value pairs that describe feature gates for alpha/experimental features.
	FeatureGates string

	// Indicates what path the kubernetes-anywhere is installed on
	KubernetesAnywherePath string

	// Viper-only parameters.  These will in time replace all flags.

	// Example: Create a file 'e2e.json' with the following:
	// 	"Cadvisor":{
	// 		"MaxRetries":"6"
	// 	}

	Viper string
}

type CloudConfig struct {
	ApiEndpoint       string
	ProjectID         string
	Zone              string // for multizone tests, arbitrarily chosen zone
	Region            string
	MultiZone         bool
	Cluster           string
	MasterName        string
	NodeInstanceGroup string // comma-delimited list of groups' names
	NumNodes          int
	ClusterIPRange    string
	ClusterTag        string
	Network           string
	ConfigFile        string // for azure and openstack
	NodeTag           string
	MasterTag         string

	Provider cloudprovider.Interface
}

var TestContext TestContextType

// Register flags common to all e2e test suites.
func RegisterCommonFlags() {
	// Turn on verbose by default to get spec names
	config.DefaultReporterConfig.Verbose = true

	// Turn on EmitSpecProgress to get spec progress (especially on interrupt)
	config.GinkgoConfig.EmitSpecProgress = true

	// Randomize specs as well as suites
	config.GinkgoConfig.RandomizeAllSpecs = true

	flag.StringVar(&TestContext.OutputPrintType, "output-print-type", "json", "Format in which summaries should be printed: 'hr' for human readable, 'json' for JSON ones.")
	flag.BoolVar(&TestContext.DumpLogsOnFailure, "dump-logs-on-failure", true, "If set to true test will dump data about the namespace in which test was running.")
	flag.BoolVar(&TestContext.DisableLogDump, "disable-log-dump", false, "If set to true, logs from master and nodes won't be gathered after test run.")
	flag.StringVar(&TestContext.LogexporterGCSPath, "logexporter-gcs-path", "", "Path to the GCS artifacts directory to dump logs from nodes. Logexporter gets enabled if this is non-empty.")
	flag.BoolVar(&TestContext.DeleteNamespace, "delete-namespace", true, "If true tests will delete namespace after completion. It is only designed to make debugging easier, DO NOT turn it off by default.")
	flag.BoolVar(&TestContext.DeleteNamespaceOnFailure, "delete-namespace-on-failure", true, "If true, framework will delete test namespace on failure. Used only during test debugging.")
	flag.IntVar(&TestContext.AllowedNotReadyNodes, "allowed-not-ready-nodes", 0, "If non-zero, framework will allow for that many non-ready nodes when checking for all ready nodes.")

	flag.StringVar(&TestContext.Host, "host", "", fmt.Sprintf("The host, or apiserver, to connect to. Will default to %s if this argument and --kubeconfig are not set", defaultHost))
	flag.StringVar(&TestContext.ReportPrefix, "report-prefix", "", "Optional prefix for JUnit XML reports. Default is empty, which doesn't prepend anything to the default name.")
	flag.StringVar(&TestContext.ReportDir, "report-dir", "", "Path to the directory where the JUnit XML reports should be saved. Default is empty, which doesn't generate these reports.")
	flag.StringVar(&TestContext.FeatureGates, "feature-gates", "", "A set of key=value pairs that describe feature gates for alpha/experimental features.")
	flag.StringVar(&TestContext.Viper, "viper-config", "e2e", "The name of the viper config i.e. 'e2e' will read values from 'e2e.json' locally.  All e2e parameters are meant to be configurable by viper.")
	flag.StringVar(&TestContext.SystemdServices, "systemd-services", "docker", "The comma separated list of systemd services the framework will dump logs for.")
	flag.StringVar(&TestContext.ImageServiceEndpoint, "image-service-endpoint", "", "The image service endpoint of cluster VM instances.")
	flag.StringVar(&TestContext.KubernetesAnywherePath, "kubernetes-anywhere-path", "/workspace/kubernetes-anywhere", "Which directory kubernetes-anywhere is installed to.")
}

// Register flags specific to the cluster e2e test suite.
func RegisterClusterFlags() {
	flag.BoolVar(&TestContext.VerifyServiceAccount, "e2e-verify-service-account", true, "If true tests will verify the service account before running.")
	flag.StringVar(&TestContext.KubeConfig, clientcmd.RecommendedConfigPathFlag, os.Getenv(clientcmd.RecommendedConfigPathEnvVar), "Path to kubeconfig containing embedded authinfo.")
	flag.StringVar(&TestContext.KubeContext, clientcmd.FlagContext, "", "kubeconfig context to use/override. If unset, will use value from 'current-context'")
	flag.StringVar(&TestContext.KubeAPIContentType, "kube-api-content-type", "application/vnd.kubernetes.protobuf", "ContentType used to communicate with apiserver")

	flag.StringVar(&TestContext.KubeVolumeDir, "volume-dir", "/var/lib/kubelet", "Path to the directory containing the kubelet volumes.")
	flag.StringVar(&TestContext.CertDir, "cert-dir", "", "Path to the directory containing the certs. Default is empty, which doesn't use certs.")
	flag.StringVar(&TestContext.RepoRoot, "repo-root", "../../", "Root directory of kubernetes repository, for finding test files.")
	flag.StringVar(&TestContext.Provider, "provider", "", "The name of the Kubernetes provider (gce, gke, local, etc.)")
	flag.StringVar(&TestContext.KubectlPath, "kubectl-path", "kubectl", "The kubectl binary to use. For development, you might use 'cluster/kubectl.sh' here.")
	flag.StringVar(&TestContext.OutputDir, "e2e-output-dir", "/tmp", "Output directory for interesting/useful test data, like performance data, benchmarks, and other metrics.")
	flag.StringVar(&TestContext.Prefix, "prefix", "e2e", "A prefix to be added to cloud resources created during testing.")
	flag.StringVar(&TestContext.MasterOSDistro, "master-os-distro", "debian", "The OS distribution of cluster master (debian, trusty, or coreos).")
	flag.StringVar(&TestContext.NodeOSDistro, "node-os-distro", "debian", "The OS distribution of cluster VM instances (debian, trusty, or coreos).")

	// TODO: Flags per provider?  Rename gce-project/gce-zone?
	cloudConfig := &TestContext.CloudConfig
	flag.StringVar(&cloudConfig.MasterName, "kube-master", "", "Name of the kubernetes master. Only required if provider is gce or gke")
	flag.StringVar(&cloudConfig.ApiEndpoint, "gce-api-endpoint", "", "The GCE ApiEndpoint being used, if applicable")
	flag.StringVar(&cloudConfig.ProjectID, "gce-project", "", "The GCE project being used, if applicable")
	flag.StringVar(&cloudConfig.Zone, "gce-zone", "", "GCE zone being used, if applicable")
	flag.StringVar(&cloudConfig.Region, "gce-region", "", "GCE region being used, if applicable")
	flag.BoolVar(&cloudConfig.MultiZone, "gce-multizone", false, "If true, start GCE cloud provider with multizone support.")
	flag.StringVar(&cloudConfig.Cluster, "gke-cluster", "", "GKE name of cluster being used, if applicable")
	flag.StringVar(&cloudConfig.NodeInstanceGroup, "node-instance-group", "", "Name of the managed instance group for nodes. Valid only for gce, gke or aws. If there is more than one group: comma separated list of groups.")
	flag.StringVar(&cloudConfig.Network, "network", "e2e", "The cloud provider network for this e2e cluster.")
	flag.IntVar(&cloudConfig.NumNodes, "num-nodes", -1, "Number of nodes in the cluster")
	flag.StringVar(&cloudConfig.ClusterIPRange, "cluster-ip-range", "10.64.0.0/14", "A CIDR notation IP range from which to assign IPs in the cluster.")
	flag.StringVar(&cloudConfig.NodeTag, "node-tag", "", "Network tags used on node instances. Valid only for gce, gke")
	flag.StringVar(&cloudConfig.MasterTag, "master-tag", "", "Network tags used on master instances. Valid only for gce, gke")

	flag.StringVar(&cloudConfig.ClusterTag, "cluster-tag", "", "Tag used to identify resources.  Only required if provider is aws.")
	flag.StringVar(&cloudConfig.ConfigFile, "cloud-config-file", "", "Cloud config file.  Only required if provider is azure.")
	flag.IntVar(&TestContext.MinStartupPods, "minStartupPods", 0, "The number of pods which we need to see in 'Running' state with a 'Ready' condition of true, before we try running tests. This is useful in any cluster which needs some base pod-based services running before it can be used.")
	flag.DurationVar(&TestContext.SystemPodsStartupTimeout, "system-pods-startup-timeout", 10*time.Minute, "Timeout for waiting for all system pods to be running before starting tests.")
	flag.DurationVar(&TestContext.NodeSchedulableTimeout, "node-schedulable-timeout", 4*time.Hour, "Timeout for waiting for all nodes to be schedulable.")
	flag.StringVar(&TestContext.UpgradeTarget, "upgrade-target", "ci/latest", "Version to upgrade to (e.g. 'release/stable', 'release/latest', 'ci/latest', '0.19.1', '0.19.1-669-gabac8c8') if doing an upgrade test.")
	flag.StringVar(&TestContext.EtcdUpgradeStorage, "etcd-upgrade-storage", "", "The storage version to upgrade to (either 'etcdv2' or 'etcdv3') if doing an etcd upgrade test.")
	flag.StringVar(&TestContext.EtcdUpgradeVersion, "etcd-upgrade-version", "", "The etcd binary version to upgrade to (e.g., '3.0.14', '2.3.7') if doing an etcd upgrade test.")
	flag.StringVar(&TestContext.UpgradeImage, "upgrade-image", "", "Image to upgrade to (e.g. 'container_vm' or 'gci') if doing an upgrade test.")
	flag.StringVar(&TestContext.GCEUpgradeScript, "gce-upgrade-script", "", "Script to use to upgrade a GCE cluster.")
	flag.BoolVar(&TestContext.CleanStart, "clean-start", false, "If true, purge all namespaces except default and system before running tests. This serves to Cleanup test namespaces from failed/interrupted e2e runs in a long-lived cluster.")
	flag.BoolVar(&TestContext.GarbageCollectorEnabled, "garbage-collector-enabled", true, "Set to true if the garbage collector is enabled in the kube-apiserver and kube-controller-manager, then some tests will rely on the garbage collector to delete dependent resources.")
}

// ViperizeFlags sets up all flag and config processing. Future configuration info should be added to viper, not to flags.
func ViperizeFlags() {

	// Part 1: Set regular flags.
	// TODO: Future, lets eliminate e2e 'flag' deps entirely in favor of viper only,
	// since go test 'flag's are sort of incompatible w/ flag, glog, etc.
	RegisterCommonFlags()
	RegisterClusterFlags()
	flag.Parse()

	// Part 2: Set Viper provided flags.
	// This must be done after common flags are registered, since Viper is a flag option.
	viper.SetConfigName(TestContext.Viper)
	viper.AddConfigPath(".")
	viper.ReadInConfig()

	// TODO Consider wether or not we want to use overwriteFlagsWithViperConfig().
	viper.Unmarshal(&TestContext)

	AfterReadingAllFlags(&TestContext)
}

func createKubeConfig(clientCfg *restclient.Config) *clientcmdapi.Config {
	clusterNick := "cluster"
	userNick := "user"
	contextNick := "context"

	config := clientcmdapi.NewConfig()

	credentials := clientcmdapi.NewAuthInfo()
	credentials.Token = clientCfg.BearerToken
	credentials.ClientCertificate = clientCfg.TLSClientConfig.CertFile
	if len(credentials.ClientCertificate) == 0 {
		credentials.ClientCertificateData = clientCfg.TLSClientConfig.CertData
	}
	credentials.ClientKey = clientCfg.TLSClientConfig.KeyFile
	if len(credentials.ClientKey) == 0 {
		credentials.ClientKeyData = clientCfg.TLSClientConfig.KeyData
	}
	config.AuthInfos[userNick] = credentials

	cluster := clientcmdapi.NewCluster()
	cluster.Server = clientCfg.Host
	cluster.CertificateAuthority = clientCfg.CAFile
	if len(cluster.CertificateAuthority) == 0 {
		cluster.CertificateAuthorityData = clientCfg.CAData
	}
	cluster.InsecureSkipTLSVerify = clientCfg.Insecure
	config.Clusters[clusterNick] = cluster

	context := clientcmdapi.NewContext()
	context.Cluster = clusterNick
	context.AuthInfo = userNick
	config.Contexts[contextNick] = context
	config.CurrentContext = contextNick

	return config
}

// AfterReadingAllFlags makes changes to the context after all flags
// have been read.
func AfterReadingAllFlags(t *TestContextType) {
	// Only set a default host if one won't be supplied via kubeconfig
	if len(t.Host) == 0 && len(t.KubeConfig) == 0 {
		// Check if we can use the in-cluster config
		if clusterConfig, err := restclient.InClusterConfig(); err == nil {
			if tempFile, err := ioutil.TempFile(os.TempDir(), "kubeconfig-"); err == nil {
				kubeConfig := createKubeConfig(clusterConfig)
				clientcmd.WriteToFile(*kubeConfig, tempFile.Name())
				t.KubeConfig = tempFile.Name()
				glog.Infof("Using a temporary kubeconfig file from in-cluster config : %s", tempFile.Name())
			}
		}
		if len(t.KubeConfig) == 0 {
			glog.Warningf("Unable to find in-cluster config, using default host : %s", defaultHost)
			t.Host = defaultHost
		}
	}
}
