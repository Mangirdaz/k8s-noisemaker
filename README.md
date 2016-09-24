# Kubernetes/Openshift Noise Maker
Working on Kubernetes/Openshift external integrations like Loadbalancers, CMDB, Storage, Logging, Alerts etc where was a need to have "Noise" in our development environments as we needed API events to be generated for our integrations to pick up. 

Utility is able to generate PODS in multiple namespaces, and from multiple images. By providing list of different images you can achieve different outcome. 

### Build
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

### Run
    ./kubernetes-noisemaker <flags>
```
NAME:
   K8S Noise Maker - For Infra Testing
USAGE:
   kubernetes-noisemaker [global options] command [command options] [arguments...]
VERSION:
   0.1
COMMANDS:
     help, h  Shows a list of commands or help for one command
GLOBAL OPTIONS:
   --concurrentcy value   How many Deployment Runs ant the same time (default: 10)
   --minttl value         Minimal TTL for Deployment (default: 10)
   --maxreolicas value    Max Replicas Count for Deployment (default: 10)
   --maxttl value         Maximal TTL for Deployment (default: 60)
   --host value           Master API host (default: "localhost:8443")
   --token value          K8S/OSE token (default: "...")
   --namespacelist value  list of namespaces where to create deployments Example --namespacelist='sample,example' (default: "sample")
   --imagelist value      list of images to use for deployments Example --imagelist='mangirdas/test-go-app:v0.1,mangirdas/test-go-app:v0.2' (default: "mangirdas/test-go-app:v0.1,mangirdas/test-go-app:v0.2")
   --help, -h             show help
   --version, -v          print the version
```

