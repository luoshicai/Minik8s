package scale

import (
	"context"
	"fmt"
	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
	"minik8s/entity"
	"minik8s/tools/log"
	"strings"
	"time"
)

const (
	PrometheusAddress    string        = "http://localhost:9090"
	QueryTimeout         time.Duration = 5 * time.Second
	UsageComputeDuration string        = "10s"
)

// MetricsManager monitors the CPU and memory usage of all the ready pods at set intervals.
type MetricsManager interface {
	// PodCPUUsage queries the average CPU usage of a given pod in the past certain seconds.
	PodCPUUsage(pod *entity.Pod) (float64, error)
	// PodMemoryUsage queries the average memory usage of a given pod in the past certain seconds.
	PodMemoryUsage(pod *entity.Pod) (uint64, error)
}

type metricsManagerInner struct {
	prometheusAPI v1.API
}

func NewMetricsManager() MetricsManager {
	client, err := api.NewClient(api.Config{
		Address: PrometheusAddress,
	})
	if err != nil {
		log.PrintE(err)
	}
	return &metricsManagerInner{
		prometheusAPI: v1.NewAPI(client),
	}
}

func (mm *metricsManagerInner) PodCPUUsage(pod *entity.Pod) (float64, error) {
	var queryBuilder strings.Builder

	// Pause container
	pauseName := pod.Metadata.Name + "_pause"
	containerQuery := containerCPUUsageQuery(pauseName)
	queryBuilder.WriteString(containerQuery)

	// Other containers
	for _, container := range pod.Spec.Containers {
		containerName := pod.Metadata.Name + "_" + container.Name
		containerQuery = containerCPUUsageQuery(containerName)
		queryBuilder.WriteString(" or ")
		queryBuilder.WriteString(containerQuery)
	}

	// Sum the results
	query := "sum(" + queryBuilder.String() + ")"

	// Query Promethus
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeout)
	defer cancel()
	result, warnings, err := mm.prometheusAPI.Query(ctx, query, time.Now())
	if err != nil {
		log.PrintE("fail to get cpu usage from prometheus:", err)
		return 0.0, err
	}
	if len(warnings) > 0 {
		log.PrintE("warnings from prometheus", warnings)
	}
	if result.(model.Vector).Len() == 0 {
		return 0.0, fmt.Errorf("fail to get cpu usage for pod %s: no data from prometheus", pod.Metadata.Name)
	}

	fmt.Printf("pod %s cpu usage: %f\n", pod.Metadata.Name, float64(result.(model.Vector)[0].Value))
	return float64(result.(model.Vector)[0].Value), nil
}

func (mm *metricsManagerInner) PodMemoryUsage(pod *entity.Pod) (uint64, error) {
	var queryBuilder strings.Builder

	// Pause container
	pauseName := pod.Metadata.Name + "_pause"
	containerQuery := containerMemoryUsageQuery(pauseName)
	queryBuilder.WriteString(containerQuery)

	// Other containers
	for _, container := range pod.Spec.Containers {
		containerName := pod.Metadata.Name + "_" + container.Name
		containerQuery = containerMemoryUsageQuery(containerName)
		queryBuilder.WriteString(" or ")
		queryBuilder.WriteString(containerQuery)
	}

	// Sum the results
	query := "sum(" + queryBuilder.String() + ")"

	// Query Promethus
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeout)
	defer cancel()
	result, warnings, err := mm.prometheusAPI.Query(ctx, query, time.Now())
	if err != nil {
		log.PrintE("fail to get memory usage from prometheus:", err)
		return 0, err
	}
	if len(warnings) > 0 {
		log.PrintW("warnings from prometheus:", warnings)
	}
	if result.(model.Vector).Len() == 0 {
		return 0, fmt.Errorf("fail to get memory usage for pod %s", pod.Metadata.Name)
	}

	fmt.Printf("pod %s memory usage: %d bytes\n", pod.Metadata.Name, uint64(result.(model.Vector)[0].Value))
	return uint64(result.(model.Vector)[0].Value), nil
}

// containerCPUUsageQuery is a helper function that generates the PromQL to query a container's
// CPU usage.
func containerCPUUsageQuery(containerName string) string {
	var query strings.Builder
	query.WriteString("sum(rate(container_cpu_usage_seconds_total{name=\"")
	query.WriteString(containerName)
	query.WriteString("\"}[")
	query.WriteString(UsageComputeDuration)
	query.WriteString("])) by (name)")
	return query.String()
}

// containerMemoryUsageQuery is a helper function that generates the PromQL to query a container's
// memory usage.
func containerMemoryUsageQuery(containerName string) string {
	var query strings.Builder
	query.WriteString("avg_over_time(container_memory_usage_bytes{name=\"")
	query.WriteString(containerName)
	query.WriteString("\"}[")
	query.WriteString(UsageComputeDuration)
	query.WriteString("])")
	return query.String()
}

// GeneratePrometheusTargets writes all the endpoints that prometheus needs to listen on into a config file.
// Prometheus will check this file regularly to get the latest info.
//func GeneratePrometheusTargets(nodes []*entity.Node) error {
//	// Create the directory of prometheus target file if not exists
//	rootPath, err := os.Getwd()
//	if err != nil {
//		return err
//	}
//	dirName := filepath.Join(rootPath, core.PROMETHEUS_TARGET_DIR)
//	err = os.MkdirAll(dirName, os.ModePerm)
//	if err != nil {
//		return err
//	}
//
//	// Construct file contents
//	targets := make([]core.PrometheusTargetObject, 0, len(nodes))
//	for _, node := range nodes {
//		targetAddr := fmt.Sprintf("%s:%d", node.Status.Address, core.CADVISOR_PORT)
//		targetObj := core.PrometheusTargetObject{
//			Targets: []string{targetAddr},
//			Label: core.PrometheusTargetLabel{
//				Job: node.Name,
//			},
//		}
//		targets = append(targets, targetObj)
//	}
//	data, err := json.MarshalIndent(targets, "", "  ")
//	if err != nil {
//		return err
//	}
//
//	// Write to prometheus target file
//	fileName := filepath.Join(dirName, core.PROMETHEUS_TARGET_FILE)
//	err = os.WriteFile(fileName, data, os.ModePerm)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
