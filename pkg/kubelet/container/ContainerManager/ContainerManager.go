// 容器运行时，维护从Pod名称到Pod中容器ID的映射
package ContainerManager

import (
	"minik8s/entity"
	"sync"
)

type ContainerManager struct {
	mtx                   sync.RWMutex
	PodNameToContainerIDs map[string][]string
}

func NewContainerManager() *ContainerManager {
	return &ContainerManager{
		PodNameToContainerIDs: map[string][]string{},
	}
}

// func (rm *ContainerManager) SetContainerIDByPodName(pod *entity.Pod, containerID string) error {
// 	rm.PodNameToContainerIDs[pod.Metadata.Name] = append(rm.PodNameToContainerIDs[pod.Metadata.Name], containerID)
//     return nil
// }

func (rm *ContainerManager) SetContainerIDsByPodName(pod *entity.Pod, containerIdMap []string) error {
	rm.mtx.Lock()
	defer rm.mtx.Unlock()
	for _, containerId := range containerIdMap {
		rm.PodNameToContainerIDs[pod.Metadata.Name] = append(rm.PodNameToContainerIDs[pod.Metadata.Name], containerId)
	}
	return nil
}

//TODO 是否需要单独对deployment中容器管理还是使用pod的管理抽象

func (rm *ContainerManager) GetContainerIDsByPodName(PodName string) []string {
	rm.mtx.RLock()
	defer rm.mtx.RUnlock()
	return rm.PodNameToContainerIDs[PodName]
}

func (rm *ContainerManager) DeletePodNameToContainerIds(PodName string) error {
	rm.mtx.Lock()
	defer rm.mtx.Unlock()
	delete(rm.PodNameToContainerIDs, PodName)
	return nil
}
