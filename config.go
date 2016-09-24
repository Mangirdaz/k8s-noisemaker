package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"k8s.io/kubernetes/pkg/client/restclient"
	"os"
)

// apiVersion: extensions/v1beta1
// kind: Deployment
// metadata:
//   name: nginx-deployment
// spec:
//   replicas: 3
//   template:
//     metadata:
//       labels:
//         app: nginx
//     spec:
//       containers:
//       - name: nginx
//         image: nginx:1.7.9
//         ports:
//        - containerPort: 80

//DeploymentConfig for creating artifact on K8S
type DeploymentConfig struct {
	ApiVersion string
	Kind       string
	Name       string
	Replicas   int
	Image      string
	Namespace  string
	TTL        int
}

type KubernetesConfig struct {
	RestClient restclient.Config
	Host       string
	Token      string
}

type NoiseMaker struct {
	KConfig               KubernetesConfig
	ConcurrentDeployments int
	MinTTL                int
	MaxTTL                int
	MaxReplicas           int
	NamespaceList         string
	ImageList             string
}

func (nm *NoiseMaker) init(args []string) {
	log.Debug("Init NoiseMaker")

	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:        "concurrentcy",
			Value:       10,
			Usage:       "How many Deployment Runs ant the same time",
			Destination: &nm.ConcurrentDeployments,
		},
		cli.IntFlag{
			Name:        "minttl",
			Value:       10,
			Usage:       "Minimal TTL for Deployment",
			Destination: &nm.MinTTL,
		},
		cli.IntFlag{
			Name:        "maxreolicas",
			Value:       10,
			Usage:       "Max Replicas Count for Deployment",
			Destination: &nm.MaxReplicas,
		},
		cli.IntFlag{
			Name:        "maxttl",
			Value:       60,
			Usage:       "Maximal TTL for Deployment",
			Destination: &nm.MaxTTL,
		},
		cli.StringFlag{
			Name:        "host",
			Value:       "localhost:8443",
			Usage:       "Master API host",
			Destination: &nm.KConfig.Host,
		},
		cli.StringFlag{
			Name:        "token",
			Value:       "...",
			Usage:       "K8S/OSE token",
			Destination: &nm.KConfig.Token,
		},
		cli.StringFlag{
			Name:        "namespacelist",
			Value:       "sample",
			Usage:       "list of namespaces where to create deployments Example --namespacelist='sample,example'",
			Destination: &nm.NamespaceList,
		},
		cli.StringFlag{
			Name:        "imagelist",
			Value:       "mangirdas/test-go-app:v0.1,mangirdas/test-go-app:v0.2",
			Usage:       "list of images to use for deployments Example --imagelist='mangirdas/test-go-app:v0.1,mangirdas/test-go-app:v0.2'",
			Destination: &nm.ImageList,
		},
	}

	app.Name = "K8S Noise Maker"
	app.Usage = "For Infra Testing"
	app.Version = "0.1"
	app.Action = func(c *cli.Context) error {
		fmt.Println(":)")
		return nil
	}

	app.Run(os.Args)
}

func (nm *NoiseMaker) validate() {

	if len(nm.KConfig.Token) < 4 {
		log.Fatal("Token is MANDATORY with --token")
	}

}
