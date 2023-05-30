package prettyprint

import (
	"fmt"
	"minik8s/entity"
	"minik8s/tools/yamlParser"
	"testing"
)

func TestPrettyPrint(t *testing.T) {
	type args struct {
		titles []string
		data   [][]string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test PrettyPrint Deployment",
			args: args{
				titles: []string{"Name", "Replicas"},
				data:   [][]string{},
			},
		},
	}
	deployment := &entity.Deployment{}
	yamlParser.ParseYaml(deployment, "../../test/autoscale_deployment.yaml")
	tests[0].args.data = append(tests[0].args.data, []string{deployment.Metadata.Name, fmt.Sprintf("%d", deployment.Status.Replicas)})
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PrettyPrint(tt.args.titles, tt.args.data)
		})
	}
}
