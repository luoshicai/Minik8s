package main

import (
	"fmt"
	"minik8s/entity"
	"minik8s/pkg/kubelet/pod/podfunc"
	HASH "minik8s/tools/hash"
	PORT "minik8s/tools/port"
	UUID "minik8s/tools/uuid"
	"minik8s/tools/yamlParser"
	"strconv"
	"time"
)

func createDeployment() error {
	deployment := &entity.Deployment{}
	path := "/home/zhaoxi/go/src/minik8s/test/nginx-deployment.yaml"
	_, err := yamlParser.ParseYaml(deployment, path)
	if err != nil {
		fmt.Println(path + "parse error!")
	}

	Replicas := 0
	Replicas = int(deployment.Spec.Replicas)
	if Replicas == 0 {
		fmt.Println("[yaml ERROR]Replicas==0")
	}
	//根据template获得template hash
	templateHash := strconv.Itoa(int(HASH.HASH([]byte(deployment.Metadata.Name + strconv.Itoa(int(deployment.Spec.Replicas))))))

	var DeployContainerIDMap []string
	Pods := make(map[string]*entity.Pod, Replicas)
	for i := 0; i < Replicas; i++ {

		//创建replicas份Pod
		pod := &entity.Pod{}
		pod.Metadata = deployment.Spec.Template.Metadata
		pod.Metadata.Uid = UUID.UUID()
		//组合产生Deployment pod的名字
		pod.Metadata.Name = deployment.Metadata.Name + "-" + templateHash + "-" + pod.Metadata.Uid[:5]

		pod.Spec = deployment.Spec.Template.Spec
		Pods[pod.Metadata.Name+pod.Metadata.Uid] = pod

		//使用模板启动时，之后的replica port使用递增号,更新container中的端口映射信息
		for j, con := range pod.Spec.Containers {
			for m, port := range con.Ports {
				//port.ContainerPort = oldportToNewport(port.ContainerPort, i)
				if port.HostPort != "" {
					port.HostPort = oldportToNewport(port.HostPort, i)
				} else {
					port.HostPort = strconv.Itoa(PORT.GetFreePort())
				}
				pod.Spec.Containers[j].Ports[m] = port
			}
		}
		//启动创建的pod
		var ContainerIDMap []string
		ContainerIDMap, err = podfunc.CreatePod(pod)
		if err != nil {
			fmt.Println("create Deployment error!")
			podfunc.DeletePod(ContainerIDMap)
			pod.Status.Phase = entity.Failed
		} else {
			pod.Status.Phase = entity.Running
		}
		for _, contianerId := range ContainerIDMap {
			DeployContainerIDMap = append(DeployContainerIDMap, contianerId)

		}
	}
	//启动成功
	fmt.Println("*************create deployment success!************")
	deployment.Status.StartTime = time.Now()
	deployment.Status.Replicas = int32(Replicas)
	return nil
}

func oldportToNewport(oldContainerPort string, i int) string {
	num, _ := strconv.Atoi(oldContainerPort)
	newContainerPort := num + i
	return strconv.Itoa(newContainerPort)
}
