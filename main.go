package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ghodss/yaml"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdlatest "k8s.io/client-go/tools/clientcmd/api/latest"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "usage: kubeconfig_merge config1 config2 ...\n")
		os.Exit(1)
	}

	loadingRules := clientcmd.ClientConfigLoadingRules{
		Precedence: os.Args[1:],
	}

	mergedConfig, err := loadingRules.Load()
	if err != nil {
		log.Fatal(err)
	}

	json, err := runtime.Encode(clientcmdlatest.Codec, mergedConfig)
	if err != nil {
		log.Fatal(err)
	}
	output, err := yaml.JSONToYAML(json)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v", string(output))
}
