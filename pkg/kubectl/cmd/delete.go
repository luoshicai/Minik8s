package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	pb "minik8s/pkg/proto"
	//"minik8s/tools/log"
	"log"
	"time"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete resources cluster",
	Long: `delete resources cluster
		   for example:
		   kubectl delete deployment [deployment name] [-n namespace] 	delete deployment`,
	Run: doDelete,
}

func doDelete(cmd *cobra.Command, args []string) {
	if len(args) != 2 {
		//log.LOG("describe err must have 2 args")
		log.Println("describe err must have 2 args")
		return
	}
	name := args[1]
	switch args[0] {
	case "po":
	case "pod":
	case "pods":
		log.Printf("[CMD]exec delete Pod %s", name)
		deletePod(name)
	case "node":
	case "nodes":
		deleteNode(name)
	case "service":
		deleteService(name)
	case "function":
		deleteFunction(name)
	case "deployment":
	case "deploy":
		log.Printf("[CMD]exec delete Deployment %s", name)
		deleteDeployment(name)
	}
}

func deletePod(name string) {
	// 通过 rpc 连接 apiserver
	cli := NewClient()
	if cli == nil {
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := cli.DeletePod(ctx, &pb.DeletePodRequest{
		Data: []byte(name),
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Delete Pod, response ", res)
}

func deleteNode(name string) {

}

func deleteDeployment(name string) {
	cli := NewClient()
	if cli == nil {
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	log.Println("begin delete Deployment", name)
	res, err := cli.DeleteDeployment(ctx, &pb.DeleteDeploymentRequest{
		DeploymentName: name,
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Delete Pod, response ", res)
}

func deleteFunction(name string) {

}

func deleteService(name string) {

}
