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
			//replica := deployment.Spec.Replicas
			item := make([]string, 0)
			item = append(item, namespace.Name)
			item = append(item, deployment.Name)

			var requestCPU int64
			var requestMem int64
			var limitCPU int64
			var limitMem int64

			for _, container := range deployment.Spec.Template.Spec.Containers {
				requestCPU += container.Resources.Requests.Cpu().MilliValue()
				requestMem += container.Resources.Requests.Memory().Value()
				limitCPU += container.Resources.Limits.Cpu().MilliValue()
				limitMem += container.Resources.Limits.Memory().Value()
			}
			item = append(item, strconv.FormatInt(requestCPU, 10))
			item = append(item, strconv.FormatInt(limitCPU, 10))
			item = append(item, strconv.FormatInt(requestMem/1024/1024, 10))
			item = append(item, strconv.FormatInt(limitMem/1024/1024, 10))
			item = append(item, fmt.Sprintf("%d", *(deployment.Spec.Replicas)))
			table = append(table, item)
		}
	}

	tw := tablewriter.NewWriter(os.Stdout)
	tw.SetHeader([]string{"namespace", "deployment", "CPU\nREQUEST", "CPU\nLIMIT", "Mem\nREQUEST", "Mem\nLIMIT", "replica"})
	tw.AppendBulk(table)
	tw.Render()
}
