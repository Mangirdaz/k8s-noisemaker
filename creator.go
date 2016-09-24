package main

import (
	log "github.com/Sirupsen/logrus"
	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/client/restclient"
	kclient "k8s.io/kubernetes/pkg/client/unversioned"
	"strings"
	"time"
)

var deploymentCount int = 0

func (nc NoiseMaker) InitCreator(queue chan DeploymentConfig) {
	log.Info("InitCreator")
	//setup client

	log.Debugf("Client config. Host: [%s]. Token: [%s...]", nc.KConfig.Host, nc.KConfig.Token[0:9])
	config := &restclient.Config{
		Host:        nc.KConfig.Host,
		BearerToken: nc.KConfig.Token,
		Insecure:    true,
	}
	log.Debug("Create client")
	kubeClient, err := kclient.New(config)
	if err != nil {
		log.Fatalln("Client not created kubeClient with %s", err)
	}

	log.Debugf("APIVersion: %s", kubeClient.APIVersion())

	//init Creator wathing
	log.Debug("Init DC Watch")
	for {
		if deploymentCount < nc.ConcurrentDeployments {
			dc := <-queue
			go nc.createDeployment(dc, kubeClient)
			time.Sleep(1 * time.Second)
		}
	}

}

func (nc NoiseMaker) createDeployment(dc DeploymentConfig, kubeClient *kclient.Client) {
	log.Debugf("Deployment Creating. TTL: %d, Image: %s, Replicas: %d, Namespace: %s, Name: %s", dc.TTL, dc.Image, dc.Replicas, dc.Namespace, dc.Name)
	pod, err := kubeClient.Pods(dc.Namespace).Create(&api.Pod{
		ObjectMeta: api.ObjectMeta{
			Name: strings.ToLower(dc.Name),
		},
		Spec: api.PodSpec{
			Containers: []api.Container{
				{
					Name:  strings.ToLower(dc.Name),
					Image: dc.Image,
					Ports: []api.ContainerPort{
						{
							ContainerPort: 8888,
						},
					},
				},
			},
		},
	})
	if err == nil {
		deploymentCount++
	} else {
		log.Errorf("Deployment failed with error %s", err)
	}

	log.Infof("Pod Created, name: %s, namespace: %s", pod.Name, pod.Namespace)
	time.Sleep(time.Duration(dc.TTL) * time.Second)

	defer deletePod(pod, kubeClient, dc.Namespace)

}

func deletePod(pod *api.Pod, kubeClient *kclient.Client, namespace string) {
	log.Infof("Delete Pod %s, namespace %s", pod.Name, namespace)
	options := api.DeleteOptions{}
	err := kubeClient.Pods(namespace).Delete(pod.Name, &options)
	if err != nil {
		log.Error("Delete pod error %s", err)
	} else {
		deploymentCount--
	}

}
