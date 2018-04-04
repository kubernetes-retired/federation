# End to end tests for Kubernetes Federation

General information for e2e tests: see [e2e-tests](https://git.k8s.io/community/contributors/devel/e2e-tests.md#federation-e2e-tests)

## Running locally (not using a cloud provider)


### Simple setup - no DNS or LBs
The below is greatly helped by [this](https://github.com/emaildanwilson/minikube-federation). Thank you emaildanwilson.

1. First you need to install:
* Minikube: https://github.com/kubernetes/minikube/releases/
* Git clone the federation repo in your local GO dir.

2. Build Kubernetes, push fcp image

```
cd $GOPATH/src/k8s.io/federation
make quick-release	# Will take awhile
docker load -i _output/release-images/amd64/fcp-amd64.tar
docker tag gcr.io/google_containers/fcp-amd64:<tag> <your dockerhub username>/fcp-amd64:<tag>
docker push <your dockerhub username>/fcp-amd64:<tag>
```

3. Start minikube

```
minikube start -p minikube
```

4. Install FCP and join clusters
```
_output/dockerized/bin/linux/amd64/kubefed init myfed --host-cluster-context=minikube --api-server-service-type=NodePort --image=<your dockerhub id>/fcp-amd64:<tag> --controllermanager-arg-overrides="--controllers=service-dns=false" --dns-provider=dummy
```
Wait for the control pane to come up. The shell output is self-explanatory. If it hangs on "waiting for control pane to come up........." it's probably a docker problem - might not be able to find the correct image. Make sure you pushed it in step 2. and that the machine can pull it properly. Use `kubectl proxy` to diagnose.

```
minikube start -p us
minikube start -p europe
kubefed join us --host-cluster-context=minikube --context=myfed
kubefed join europe --host-cluster-context=minikube --context=myfed
```

5. Run e2e tests

```
_output/dockerized/bin/linux/amd64/e2e.test --kubeconfig ~/.kube/config --federated-kube-context myfed -ginkgo.focus="\[Feature:Federation\].Features.Preferences"
```
You can also add `--e2e-output-dir "./e2eres"` to put the results into a local folder (default is `/tmp`).

6. Clean up

```
kubectl delete ns federation-system --context=minikube
minikube stop -p minikube
minikube delete -p us
minikube delete -p europe
minikube delete -p minikube
```

## Instruction Dump below - will be restructured into normal docs above.

TODO: Remove

	git clone https://github.com/emaildanwilson/minikube-federation
	cd ./minikube-federation

Create minikube clusters:

	kubectl run etcd --image=quay.io/coreos/etcd:v2.3.7 --env="ETCD_LISTEN_CLIENT_URLS=http://0.0.0.0:2379" --env="ETCD_ADVERTISE_CLIENT_URLS=http://etcd.default:2379" --port=2379 --expose --context=minikube

	helm init --kube-context minikube
	helm version  # should work

	cat ./coredns-values.yaml
	helm install --namespace default --name coredns --kube-context minikube -f ./coredns-values.yaml stable/coredns

	kubectl get svc,pods -l app=coredns-coredns --context=minikube  # Should be running

	minikube start -p us
	minikube start -p europe
	minikube addons disable dashboard -p minikube
	minikube addons disable dashboard -p us
	minikube addons disable dashboard -p europe

	kubectl label node us failure-domain.beta.kubernetes.io/zone=us1 failure-domain.beta.kubernetes.io/region=us --context=us
	kubectl label node europe failure-domain.beta.kubernetes.io/zone=europe1 failure-domain.beta.kubernetes.io/region=europe --context=europe

	kubectl create configmap ingress-uid --from-literal=uid=us1 -n kube-system --context=us
	kubectl create configmap ingress-uid --from-literal=uid=europe1 -n kube-system --context=europe

	kubefed init myfed --host-cluster-context=minikube --dns-provider="coredns" --dns-zone-name="myzone." --api-server-service-type=NodePort --api-server-advertise-address=$(minikube ip -p minikube) --apiserver-enable-basic-auth=true --apiserver-enable-token-auth=true --apiserver-arg-overrides="--anonymous-auth=true,--v=4" --dns-provider-config="./coredns-provider.conf"

	kubectl get pods --context=minikube -n federation-system
	kubectl logs -l module=federation-controller-manager --context=minikube -n federation-system

	kubectl config use-context myfed
	kubectl create ns default

	kubefed join us --host-cluster-context=minikube
	kubefed join europe --host-cluster-context=minikube

	kubectl label cluster us environment=test location=us pciCompliant=false
	kubectl label cluster europe environment=prod location=europe pciCompliant=true

	kubectl get clusters -L environment -L location -L pciCompliant

#Launch App time

	kubectl apply -f ./hello.yaml --context=myfed
	kubectl get svc,rs,po,ing --context=us
	kubectl get svc,rs,po,ing --context=europe

	kubectl run -it --rm --restart=Never --image=infoblox/dnstools:latest dnstools --context=minikube
	# Execute DNS Lookup inside the machine
	nslookup hello.default.myfed.svc.myzone coredns-coredns.default

