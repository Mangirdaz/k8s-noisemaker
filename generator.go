package main

import (
	"github.com/Pallinder/go-randomdata"
	log "github.com/Sirupsen/logrus"
	"math/rand"
	"strings"
	"time"
)

const apiVersion = "extensions/v1beta1"

func (nc NoiseMaker) InitGenerator(dcBuffer chan DeploymentConfig) {

	//loop to keep generating

	for {
		if len(dcBuffer) < 10 {
			ttl := random(nc.MinTTL, nc.MaxTTL)
			image := getRandomString(nc.ImageList)
			replicas := random(1, nc.MaxReplicas)
			namespace := getRandomString(nc.NamespaceList)
			name := randomdata.SillyName()

			dc := DeploymentConfig{
				ApiVersion: apiVersion,
				Image:      image,
				Name:       name,
				Namespace:  namespace,
				Replicas:   replicas,
				TTL:        ttl,
			}
			log.Debugf("Deployment Generated and Passed to buffer. TTL: %d, Image: %s, Replicas: %d, Namespace: %s, Name: %s", ttl, image, replicas, namespace, name)
			dcBuffer <- dc
			time.Sleep(50 * time.Millisecond)
		}
	}

}

func random(min, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(max-min) + min
}

func getRandomString(list string) string {
	array := strings.Split(list, ",")
	return array[rand.Intn(len(array))]
}
