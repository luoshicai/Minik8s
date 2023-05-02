// grpc server is API Server and client is Kubelet

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.9
// source: apiserver_for_kubelet.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_apiserver_for_kubelet_proto protoreflect.FileDescriptor

var file_apiserver_for_kubelet_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x61, 0x70, 0x69, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x66, 0x6f, 0x72, 0x5f,
	0x6b, 0x75, 0x62, 0x65, 0x6c, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x61,
	0x70, 0x69, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x66, 0x6f, 0x72, 0x5f, 0x6b, 0x75, 0x62,
	0x65, 0x6c, 0x65, 0x74, 0x1a, 0x0b, 0x73, 0x68, 0x61, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x32, 0xdc, 0x01, 0x0a, 0x17, 0x41, 0x70, 0x69, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4b,
	0x75, 0x62, 0x65, 0x6c, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x35, 0x0a,
	0x08, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x13, 0x2e, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14,
	0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x0c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x4e, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x2e, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x15, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x50, 0x6f, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x2e, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x64, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x13, 0x5a, 0x11, 0x6d, 0x69, 0x6e, 0x69, 0x6b, 0x38, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_apiserver_for_kubelet_proto_goTypes = []interface{}{
	(*HelloRequest)(nil),           // 0: share.HelloRequest
	(*RegisterNodeRequest)(nil),    // 1: share.RegisterNodeRequest
	(*UpdatePodStatusRequest)(nil), // 2: share.UpdatePodStatusRequest
	(*HelloResponse)(nil),          // 3: share.HelloResponse
	(*StatusResponse)(nil),         // 4: share.StatusResponse
}
var file_apiserver_for_kubelet_proto_depIdxs = []int32{
	0, // 0: apiserver_for_kubelet.ApiServerKubeletService.SayHello:input_type -> share.HelloRequest
	1, // 1: apiserver_for_kubelet.ApiServerKubeletService.RegisterNode:input_type -> share.RegisterNodeRequest
	2, // 2: apiserver_for_kubelet.ApiServerKubeletService.UpdatePodStatus:input_type -> share.UpdatePodStatusRequest
	3, // 3: apiserver_for_kubelet.ApiServerKubeletService.SayHello:output_type -> share.HelloResponse
	4, // 4: apiserver_for_kubelet.ApiServerKubeletService.RegisterNode:output_type -> share.StatusResponse
	4, // 5: apiserver_for_kubelet.ApiServerKubeletService.UpdatePodStatus:output_type -> share.StatusResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_apiserver_for_kubelet_proto_init() }
func file_apiserver_for_kubelet_proto_init() {
	if File_apiserver_for_kubelet_proto != nil {
		return
	}
	file_share_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_apiserver_for_kubelet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apiserver_for_kubelet_proto_goTypes,
		DependencyIndexes: file_apiserver_for_kubelet_proto_depIdxs,
	}.Build()
	File_apiserver_for_kubelet_proto = out.File
	file_apiserver_for_kubelet_proto_rawDesc = nil
	file_apiserver_for_kubelet_proto_goTypes = nil
	file_apiserver_for_kubelet_proto_depIdxs = nil
}
