package main

import (
	log "github.com/Sirupsen/logrus"
	"os"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	//log.SetFormatter(&log.JSONFormatter{})

	// Output to stderr instead of stdout, could also be a file.
	log.SetOutput(os.Stderr)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)
}

func main() {

	nc := NoiseMaker{}
	nc.init(os.Args)
	nc.validate()

	dcBuffer := make(chan DeploymentConfig, 10)

	go nc.InitGenerator(dcBuffer)
	nc.InitCreator(dcBuffer)

}
