package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"minik8s/entity"
	"minik8s/pkg/kubectl/utils"

	// "minik8s/tools/log"
	"minik8s/tools/yamlParser"
	"strings"
	"time"

	pb "minik8s/pkg/proto"

	"github.com/spf13/cobra"
	// "google.golang.org/grpc/channelz/service"
)

var (
	filename string
	applyCmd = &cobra.Command{
		Use:   "apply",
		Short: "kubectl apply is to create object",
		Long: `kubectl apply is to create object
		   for example:
		   kubectl apply -f *.yaml`,
		Run: doApply,
	}
)

func init() {
	applyCmd.Flags().StringVarP(&filename, "filename", "f", "", "yaml name")
	err := applyCmd.MarkFlagRequired("filename")
	if err != nil {
		fmt.Println("required filename")
		return
	}
}

func doApply(cmd *cobra.Command, args []string) {
	var (
		dirname                  string
		filenameWithoutExtention string
	)

	b, err := yamlParser.FileExists(filename)
	if !b || err != nil {
		fmt.Println("file does not exist")
		return
	}

	// 把路径按照 ‘/’ 拆分开，获取没有 .yaml 后缀的文件名
	arr := strings.Split(filename, "/")
	for i := 0; i < len(arr)-1; i++ {
		dirname = dirname + arr[i] + "/"
	}
	if len(dirname) == 0 {
		dirname = "."
	}
	

	filenameWithoutExtention = strings.Split(arr[len(arr)-1], ".")[0]

	obj, err := utils.GetField(dirname, filenameWithoutExtention, "kind")
	if err != nil {
		fmt.Println("file has no such field")
	}

	fmt.Println("dirname: ", dirname, "filename without extention: ", filenameWithoutExtention)

	switch obj {
	case "Pod", "pod":
		err := applyPod(filename)
		if err != nil {
			fmt.Println(err)
		}


	case "Deployment", "deployment":
		applyDeployment(filename)
		if err != nil {
			fmt.Println(err)
		}

	case "Service", "service":
		applyService(filename)
		if err != nil {
			fmt.Println(err)
		}

	case "Node", "node":
	// TODO
	
	default:
		fmt.Println("there is no object named ")
	}
}


func applyPod(filename string) error {
			// 先 parse yaml 文件
			pod := &entity.Pod{}
			_, err := yamlParser.ParseYaml(pod, filename)
			if err != nil {
				fmt.Println("parse pod failed")
				return err
			}
			fmt.Println(pod)
	
			// 通过 rpc 连接 apiserver
			cli := NewClient()
			if cli == nil {
				return fmt.Errorf("fail to connect to apiserver")
			}
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
	
			// 把 pod 序列化成 string 传给 apiserver
			podByte, err := json.Marshal(pod)
			if err != nil {
				fmt.Println("parse pod error")
				return err
			}
	
			res, err := cli.ApplyPod(ctx, &pb.ApplyPodRequest{
				Data: podByte,
			})
	
			fmt.Println("Create Pod, response ", res)
			return nil
}

func applyDeployment(filename string) error {
	deploy := &entity.Deployment{}
	_, err := yamlParser.ParseYaml(deploy, filename)
	if err != nil {
		fmt.Println("parse deploy failed")
		return err
	}

	cli := NewClient()
	if cli == nil {
		return fmt.Errorf("fail to connect to apiserver")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 把 pod 序列化成 string 传给 apiserver
	podByte, err := json.Marshal(deploy)
	if err != nil {
		fmt.Println("parse deploy error")
		return err
	}

	res, err := cli.ApplyDeployment(ctx, &pb.ApplyDeploymentRequest{
		Data: podByte,
	})

	fmt.Println("Create Deployment, response ", res)
	return nil
}	

func applyService(filename string) error {
	service := &entity.Service{}
	_, err := yamlParser.ParseYaml(service, filename)
	if err != nil {
		fmt.Println("parse service failed")
		return err
	}

	cli := NewClient()
	if cli == nil {
		return fmt.Errorf("fail to connect to apiserver")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 把 pod 序列化成 string 传给 apiserver
	podByte, err := json.Marshal(service)
	if err != nil {
		fmt.Println("parse service error")
		return err
	}

	res, err := cli.ApplyService(ctx, &pb.ApplyServiceRequest{
		Data: podByte,
	})

	fmt.Println("Create Service, response ", res)
	return nil
}	