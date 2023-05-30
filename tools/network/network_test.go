package network

import "testing"

func TestGetNetInterfaceIPv4Addr(t *testing.T) {
	type args struct {
		interfaceName string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "Get IP Frome local",
			args:    args{interfaceName: "ens0"},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetNetInterfaceIPv4Addr(tt.args.interfaceName)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetNetInterfaceIPv4Addr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetNetInterfaceIPv4Addr() got = %v, want %v", got, tt.want)
			}
		})
	}
}
