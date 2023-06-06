package dtos

type SmallType struct {
	proto

	Id     int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Field1 string `protobuf:"bytes,2,opt,name=field1,proto3" json:"field1,omitempty"`
	Field2 string `protobuf:"bytes,3,opt,name=field2,proto3" json:"field2,omitempty"`
}

type MediumType struct {
	proto

	Id     int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Field1 string `protobuf:"bytes,2,opt,name=field1,proto3" json:"field1,omitempty"`
	Field2 string `protobuf:"bytes,3,opt,name=field2,proto3" json:"field2,omitempty"`
	Field3 string `protobuf:"bytes,4,opt,name=field3,proto3" json:"field3,omitempty"`
	Field4 string `protobuf:"bytes,5,opt,name=field4,proto3" json:"field4,omitempty"`
	Field5 string `protobuf:"bytes,6,opt,name=field5,proto3" json:"field5,omitempty"`
	Field6 string `protobuf:"bytes,7,opt,name=field6,proto3" json:"field6,omitempty"`
	Field7 string `protobuf:"bytes,8,opt,name=field7,proto3" json:"field7,omitempty"`
	Field8 string `protobuf:"bytes,9,opt,name=field8,proto3" json:"field8,omitempty"`
}

type LargeType struct {
	proto

	Id     int32         `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Field1 string        `protobuf:"bytes,2,opt,name=field1,proto3" json:"field1,omitempty"`
	Field2 string        `protobuf:"bytes,3,opt,name=field2,proto3" json:"field2,omitempty"`
	Field3 string        `protobuf:"bytes,4,opt,name=field3,proto3" json:"field3,omitempty"`
	Field4 string        `protobuf:"bytes,5,opt,name=field4,proto3" json:"field4,omitempty"`
	Field5 string        `protobuf:"bytes,6,opt,name=field5,proto3" json:"field5,omitempty"`
	Field6 string        `protobuf:"bytes,7,opt,name=field6,proto3" json:"field6,omitempty"`
	Field7 string        `protobuf:"bytes,8,opt,name=field7,proto3" json:"field7,omitempty"`
	Field8 string        `protobuf:"bytes,9,opt,name=field8,proto3" json:"field8,omitempty"`
	Nested []*MediumType `protobuf:"bytes,10,rep,name=nested,proto3" json:"nested,omitempty"`
}
