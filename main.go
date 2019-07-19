package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	homeDir, _ := os.UserHomeDir()
	configPath := homeDir + "/.kube/config"
	config, err := clientcmd.BuildConfigFromFlags("", configPath)
	if err != nil {
		fmt.Print(err)
		return
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	namespaces, err := clientset.CoreV1().Namespaces().List(v1.ListOptions{})
	if err != nil {
		panic(err)
	}
	table := make([][]string, 0)

	for _, namespace := range namespaces.Items {
		deploymentList, err := clientset.AppsV1beta1().Deployments(namespace.Name).List(v1.ListOptions{})
		if err != nil {
			panic(err)
		}

		for _, deployment := range deploymentList.Items {
			for _, container := range deployment.Spec.Template.Spec.Containers {
				item := make([]string, 7)

				item[0] = namespace.Name
				item[1] = deployment.Name
				item[2] = container.Name

				requestCPU := container.Resources.Requests.Cpu().Value()
				requestMem := container.Resources.Requests.Memory().Value()
				limitCPU := container.Resources.Limits.Cpu().Value()
				limitMem := container.Resources.Limits.Memory().Value()

				item[3] = strconv.FormatInt(requestCPU, 10)
				item[4] = strconv.FormatInt(limitCPU, 10)
				item[5] = strconv.FormatInt(requestMem/1024/1024, 10)
				item[6] = strconv.FormatInt(limitMem/1024/1024, 10)

				table = append(table, item)
			}
		}
	}

	tw := tablewriter.NewWriter(os.Stdout)
	tw.SetHeader([]string{"namespace", "deployment", "container", "CPU\nREQUEST", "CPU\nLIMIT", "Mem\nREQUEST", "Mem\nLIMIT"})

	tw.AppendBulk(table)

	tw.Render()
}
