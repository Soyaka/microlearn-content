// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.21.12
// source: tutorial.proto

package content

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

var File_tutorial_proto protoreflect.FileDescriptor

var file_tutorial_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x74, 0x75, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xfb, 0x01,
	0x0a, 0x0f, 0x54, 0x75, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x39, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x75, 0x74, 0x6f, 0x72,
	0x69, 0x61, 0x6c, 0x12, 0x09, 0x2e, 0x54, 0x75, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c, 0x1a, 0x06,
	0x2e, 0x52, 0x65, 0x73, 0x4f, 0x4b, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x3a, 0x01,
	0x2a, 0x22, 0x09, 0x2f, 0x74, 0x75, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c, 0x12, 0x39, 0x0a, 0x0e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x75, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c, 0x12, 0x09,
	0x2e, 0x54, 0x75, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c, 0x1a, 0x06, 0x2e, 0x52, 0x65, 0x73, 0x4f,
	0x4b, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x3a, 0x01, 0x2a, 0x1a, 0x09, 0x2f, 0x74,
	0x75, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c, 0x12, 0x38, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x54, 0x75, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c, 0x12, 0x06, 0x2e, 0x52, 0x65, 0x71, 0x49,
	0x44, 0x1a, 0x06, 0x2e, 0x52, 0x65, 0x73, 0x4f, 0x4b, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x10, 0x2a, 0x0e, 0x2f, 0x74, 0x75, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c, 0x2f, 0x7b, 0x49, 0x44,
	0x7d, 0x12, 0x38, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x54, 0x75, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c,
	0x12, 0x06, 0x2e, 0x52, 0x65, 0x71, 0x49, 0x44, 0x1a, 0x09, 0x2e, 0x54, 0x75, 0x74, 0x6f, 0x72,
	0x69, 0x61, 0x6c, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x12, 0x0e, 0x2f, 0x74, 0x75,
	0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c, 0x2f, 0x7b, 0x49, 0x44, 0x7d, 0x42, 0x1b, 0x5a, 0x19, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x6f, 0x79, 0x61, 0x6b, 0x61,
	0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_tutorial_proto_goTypes = []interface{}{
	(*Tutorial)(nil), // 0: Tutorial
	(*ReqID)(nil),    // 1: ReqID
	(*ResOK)(nil),    // 2: ResOK
}
var file_tutorial_proto_depIdxs = []int32{
	0, // 0: TutorialService.CreateTutorial:input_type -> Tutorial
	0, // 1: TutorialService.UpdateTutorial:input_type -> Tutorial
	1, // 2: TutorialService.DeleteTutorial:input_type -> ReqID
	1, // 3: TutorialService.GetTutorial:input_type -> ReqID
	2, // 4: TutorialService.CreateTutorial:output_type -> ResOK
	2, // 5: TutorialService.UpdateTutorial:output_type -> ResOK
	2, // 6: TutorialService.DeleteTutorial:output_type -> ResOK
	0, // 7: TutorialService.GetTutorial:output_type -> Tutorial
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tutorial_proto_init() }
func file_tutorial_proto_init() {
	if File_tutorial_proto != nil {
		return
	}
	file_content_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tutorial_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tutorial_proto_goTypes,
		DependencyIndexes: file_tutorial_proto_depIdxs,
	}.Build()
	File_tutorial_proto = out.File
	file_tutorial_proto_rawDesc = nil
	file_tutorial_proto_goTypes = nil
	file_tutorial_proto_depIdxs = nil
}