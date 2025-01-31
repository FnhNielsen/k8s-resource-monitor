package main

import (
	"log"

	"github.com/Fnhnielsen/k8s-resource-monitor/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatalf("Error: %v", err)
	}
}