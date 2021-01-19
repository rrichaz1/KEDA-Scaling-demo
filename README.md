# KEDA-Scaling-demo

This is a small project to demonstrate the scaling in Azure Kubernetes Service with the help of Kubernetes Event Driven Scaling.

It uses the number of messages in the service bus queue as an external factor for scaling.

This project can be divided into following sub-modules:

1. Golang-receiver: A module which consumes messages from the service bus queue and waits for a random time ranging from (0-60) seconds to mimic some message processing. This is the module which will be deployed to AKS.

2. Golang-sender: This will be used to send messages to the service bus to queue to test scaling, it has a compiled version (for Linux) as well, so if you don't want to run the compilation just use following:
```sh
$ ./sender

```

3. charts: This folder contains the helm charts used to deploy the service to AKS

## Scaling parameters
The scaling parameters are mentioned in the `values.yaml` file under charts folder. The main specific things to note are following:
```sh
autoscaling:
  enabled: true         # Checks whether to deploy KEDA scaled object or not
  minReplicas: 2        # Specifies the min replica count
  maxReplicas: 20       # Specifies the max replica count
  pollingInterval: 5    # Polling Interval tells the amount of time KEDA should check external metric
  cooldownPeriod: 10    # This specifies the amount time to wait after scaling has occured
  messageCount: 10      # It specifies the number of messages per pod to used for calculating scaling param

envVariables:
  sb_namespace: ""      # It speciifes the namespace of service bus queue
  sb_key_name: ""       # SB key name
  sb_key_value: ""      # SB key value
  sb_queue: "keda-demo" # SB queue to be used by application and KEDA to check for messages
  queueConnString: ""   # The entire connection string to be used by KEDA

```
## Prerequisites to run this project.
1. A running Kubernetes cluster
2. kubectl with connection to your cluster (valid .kubeconfig)
3. helm 3
4. docker (if you want to build a new version of reciever otherwise not)

## Runing this project
This project has multiple separate blocks which can run separtely.

### 1. Docker Image
There is a prebuilt image of docker available on dockerhub, you can just use this image directly. It requires a few environment variables to be set, which we are setting through the helm chart using a config map, but same can be achieved in multiple ways.
Docker link: [https://hub.docker.com/repository/docker/dvkcool/az_service_bus_consumer_golang](https://hub.docker.com/repository/docker/dvkcool/az_service_bus_consumer_golang)

The environment variables required are:
```sh
  sb_namespace: ""      # The queue namespace
  sb_key_name: ""       # The queue key name
  sb_key_value: ""      # The access key value
  sb_queue: "keda-demo" # The queue name to listen to

```
### 2. Receiver / Sender (Manual run)
If you wish to modify the reciever / sender and run it manually, you can do same with help of golang. 

Just follow official docs to install Go on your machine and happy hacking, but remember it will still need those environment variables to be set.

### 3. Deploying it through helm
1. Firstly make sure you have KEDA installed in your AKS cluster, you can follow docs at [keda.sh](https://keda.sh) , but following steps are gist of it:
```sh
$ helm repo add kedacore https://kedacore.github.io/charts
$ helm repo update
$ kubectl create namespace keda
$ helm install keda kedacore/keda --namespace keda
```

2. Change directory to charts folder and run following command to install the receiver (or consumer) to your cluster

```sh
$ helm install az-sb-consumer-poc ./az-sb-consumer-poc -n {{Desired Namespace in AKS }}
```
### Uninstalling the project
1. Remove the application from AKS
```sh
$ helm uninstall az-sb-consumer-poc -n {{Namespace name}}
```
2. Remove KEDA from your cluster.
```sh
$ helm uninstall -n keda keda
$ kubectl delete -f https://raw.githubusercontent.com/kedacore/keda/main/config/crd/bases/keda.sh_scaledobjects.yaml
$ kubectl delete -f https://raw.githubusercontent.com/kedacore/keda/main/config/crd/bases/keda.sh_scaledjobs.yaml
$ kubectl delete -f https://raw.githubusercontent.com/kedacore/keda/main/config/crd/bases/keda.sh_triggerauthentications.yaml
```
## References
The go module used to interact with service bus can be found here: [github.com/michaelbironneau/asbclient](https://github.com/michaelbironneau/asbclient)


## Found an Issue ?
Feel free to open an issue, or send me a PR with the fix.
I will be happy to merge the PR.

___________________________________________________________________________________________________
Happy Coding
____________________________________________________________________________________________________
Divyanshu Kumar
