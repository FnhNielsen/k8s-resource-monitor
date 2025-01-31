package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:  "k8s-resource-monitor",
	Short: "Kubernetes resource monitor CLI",
	Long:  "A CLI tool to monitor Kubernetes cluster resources such as Pods, Nodes, and Deployments.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}