// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.4
// source: pbmodels/pbModels.proto

package pbmodels

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

//* Definición de Modelos de Objetos.
type Person struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName     string `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	SecondName    string `protobuf:"bytes,3,opt,name=second_name,json=secondName,proto3" json:"second_name,omitempty"`
	FirstSurname  string `protobuf:"bytes,4,opt,name=first_surname,json=firstSurname,proto3" json:"first_surname,omitempty"`
	SecondSurname string `protobuf:"bytes,5,opt,name=second_surname,json=secondSurname,proto3" json:"second_surname,omitempty"`
	DateOfBirth   string `protobuf:"bytes,6,opt,name=date_of_birth,json=dateOfBirth,proto3" json:"date_of_birth,omitempty"`
	Gender        string `protobuf:"bytes,7,opt,name=gender,proto3" json:"gender,omitempty"`
	Alias         string `protobuf:"bytes,8,opt,name=alias,proto3" json:"alias,omitempty"`
}

func (x *Person) Reset() {
	*x = Person{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbmodels_pbModels_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Person) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Person) ProtoMessage() {}

func (x *Person) ProtoReflect() protoreflect.Message {
	mi := &file_pbmodels_pbModels_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Person.ProtoReflect.Descriptor instead.
func (*Person) Descriptor() ([]byte, []int) {
	return file_pbmodels_pbModels_proto_rawDescGZIP(), []int{0}
}

func (x *Person) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Person) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *Person) GetSecondName() string {
	if x != nil {
		return x.SecondName
	}
	return ""
}

func (x *Person) GetFirstSurname() string {
	if x != nil {
		return x.FirstSurname
	}
	return ""
}

func (x *Person) GetSecondSurname() string {
	if x != nil {
		return x.SecondSurname
	}
	return ""
}

func (x *Person) GetDateOfBirth() string {
	if x != nil {
		return x.DateOfBirth
	}
	return ""
}

func (x *Person) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *Person) GetAlias() string {
	if x != nil {
		return x.Alias
	}
	return ""
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbmodels_pbModels_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_pbmodels_pbModels_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_pbmodels_pbModels_proto_rawDescGZIP(), []int{1}
}

type ListPersonsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//string parent = 1;                      //* No entiendo que vamos a poner aquí, es decir, porque existe ésta propiedad.
	PageSize  int32  `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`   //* Cantidad de objetos, en este caso 'personas', mostradas en cada página.
	PageToken string `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"` //* Número de la página actual. Ésta numeración siempre comienza en CERO.
}

func (x *ListPersonsRequest) Reset() {
	*x = ListPersonsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbmodels_pbModels_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPersonsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPersonsRequest) ProtoMessage() {}

func (x *ListPersonsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pbmodels_pbModels_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPersonsRequest.ProtoReflect.Descriptor instead.
func (*ListPersonsRequest) Descriptor() ([]byte, []int) {
	return file_pbmodels_pbModels_proto_rawDescGZIP(), []int{2}
}

func (x *ListPersonsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListPersonsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type ListPersonsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Person        []*Person `protobuf:"bytes,1,rep,name=person,proto3" json:"person,omitempty"`                                      //* Array, o slice, de objetos, en este caso de tipo 'Person', ya que se definió el modificador 'repeated'.
	NextPageToken string    `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"` //* Número de la página siguiente a la actualmente mostrada.
}

func (x *ListPersonsResponse) Reset() {
	*x = ListPersonsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbmodels_pbModels_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPersonsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPersonsResponse) ProtoMessage() {}

func (x *ListPersonsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pbmodels_pbModels_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPersonsResponse.ProtoReflect.Descriptor instead.
func (*ListPersonsResponse) Descriptor() ([]byte, []int) {
	return file_pbmodels_pbModels_proto_rawDescGZIP(), []int{3}
}

func (x *ListPersonsResponse) GetPerson() []*Person {
	if x != nil {
		return x.Person
	}
	return nil
}

func (x *ListPersonsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type GetPersonRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` //* Para identificar de forma inequívoca un objeto 'person', únicamente necesitamos su 'id'.
}

func (x *GetPersonRequest) Reset() {
	*x = GetPersonRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbmodels_pbModels_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPersonRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPersonRequest) ProtoMessage() {}

func (x *GetPersonRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pbmodels_pbModels_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPersonRequest.ProtoReflect.Descriptor instead.
func (*GetPersonRequest) Descriptor() ([]byte, []int) {
	return file_pbmodels_pbModels_proto_rawDescGZIP(), []int{4}
}

func (x *GetPersonRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CreatePersonRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//string parent = 1;                      //* Lo mismo que antes en 'ListPersonsRequest', no entiendo el porqué de su existencia.
	//string person_id = 2;                   //* En el ejemplo, entiendo que el campo 'id' viene generado ya en la 'request', pero lo generaremos más tarde.
	Person *Person `protobuf:"bytes,1,opt,name=person,proto3" json:"person,omitempty"`
}

func (x *CreatePersonRequest) Reset() {
	*x = CreatePersonRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbmodels_pbModels_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePersonRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePersonRequest) ProtoMessage() {}

func (x *CreatePersonRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pbmodels_pbModels_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePersonRequest.ProtoReflect.Descriptor instead.
func (*CreatePersonRequest) Descriptor() ([]byte, []int) {
	return file_pbmodels_pbModels_proto_rawDescGZIP(), []int{5}
}

func (x *CreatePersonRequest) GetPerson() *Person {
	if x != nil {
		return x.Person
	}
	return nil
}

type UpdatePersonRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Person *Person `protobuf:"bytes,1,opt,name=person,proto3" json:"person,omitempty"` //FieldMask update_mask = 2;              //* No entiendo que vamos a poner aquí, es decir, porque existe ésta propiedad-
}

func (x *UpdatePersonRequest) Reset() {
	*x = UpdatePersonRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbmodels_pbModels_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePersonRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePersonRequest) ProtoMessage() {}

func (x *UpdatePersonRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pbmodels_pbModels_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePersonRequest.ProtoReflect.Descriptor instead.
func (*UpdatePersonRequest) Descriptor() ([]byte, []int) {
	return file_pbmodels_pbModels_proto_rawDescGZIP(), []int{6}
}

func (x *UpdatePersonRequest) GetPerson() *Person {
	if x != nil {
		return x.Person
	}
	return nil
}

type DeletePersonRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PersonId string `protobuf:"bytes,1,opt,name=person_id,json=personId,proto3" json:"person_id,omitempty"`
}

func (x *DeletePersonRequest) Reset() {
	*x = DeletePersonRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbmodels_pbModels_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePersonRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePersonRequest) ProtoMessage() {}

func (x *DeletePersonRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pbmodels_pbModels_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePersonRequest.ProtoReflect.Descriptor instead.
func (*DeletePersonRequest) Descriptor() ([]byte, []int) {
	return file_pbmodels_pbModels_proto_rawDescGZIP(), []int{7}
}

func (x *DeletePersonRequest) GetPersonId() string {
	if x != nil {
		return x.PersonId
	}
	return ""
}

var File_pbmodels_pbModels_proto protoreflect.FileDescriptor

var file_pbmodels_pbModels_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x62, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x70, 0x62, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x22, 0xf6, 0x01, 0x0a, 0x06, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73,
	0x65, 0x63, 0x6f, 0x6e, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d,
	0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x73, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x69, 0x72, 0x73, 0x74, 0x53, 0x75, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x5f, 0x73, 0x75, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x65, 0x63, 0x6f, 0x6e,
	0x64, 0x53, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x64, 0x61, 0x74, 0x65,
	0x5f, 0x6f, 0x66, 0x5f, 0x62, 0x69, 0x72, 0x74, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x66, 0x42, 0x69, 0x72, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06,
	0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x50, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61,
	0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x65, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x06,
	0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x06, 0x70, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e,
	0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x22, 0x0a, 0x10,
	0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x3d, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x06, 0x70, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73,
	0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x06, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x22,
	0x3d, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x06, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e,
	0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x06, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x22, 0x32,
	0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x49, 0x64, 0x32, 0xc4, 0x02, 0x0a, 0x0d, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x73, 0x12, 0x1a, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1b, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x09,
	0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x18, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x50, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x12, 0x3b, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x12, 0x1b, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x12, 0x3b, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x12, 0x1b, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x3a, 0x0a,
	0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x1b, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x65, 0x72, 0x6f, 0x64, 0x69, 0x6e, 0x61,
	0x6d, 0x69, 0x63, 0x61, 0x74, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x72, 0x65, 0x73,
	0x74, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x62, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pbmodels_pbModels_proto_rawDescOnce sync.Once
	file_pbmodels_pbModels_proto_rawDescData = file_pbmodels_pbModels_proto_rawDesc
)

func file_pbmodels_pbModels_proto_rawDescGZIP() []byte {
	file_pbmodels_pbModels_proto_rawDescOnce.Do(func() {
		file_pbmodels_pbModels_proto_rawDescData = protoimpl.X.CompressGZIP(file_pbmodels_pbModels_proto_rawDescData)
	})
	return file_pbmodels_pbModels_proto_rawDescData
}

var file_pbmodels_pbModels_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_pbmodels_pbModels_proto_goTypes = []interface{}{
	(*Person)(nil),              // 0: models.Person
	(*Empty)(nil),               // 1: models.Empty
	(*ListPersonsRequest)(nil),  // 2: models.ListPersonsRequest
	(*ListPersonsResponse)(nil), // 3: models.ListPersonsResponse
	(*GetPersonRequest)(nil),    // 4: models.GetPersonRequest
	(*CreatePersonRequest)(nil), // 5: models.CreatePersonRequest
	(*UpdatePersonRequest)(nil), // 6: models.UpdatePersonRequest
	(*DeletePersonRequest)(nil), // 7: models.DeletePersonRequest
}
var file_pbmodels_pbModels_proto_depIdxs = []int32{
	0, // 0: models.ListPersonsResponse.person:type_name -> models.Person
	0, // 1: models.CreatePersonRequest.person:type_name -> models.Person
	0, // 2: models.UpdatePersonRequest.person:type_name -> models.Person
	2, // 3: models.PersonService.ListPersons:input_type -> models.ListPersonsRequest
	4, // 4: models.PersonService.GetPerson:input_type -> models.GetPersonRequest
	5, // 5: models.PersonService.CreatePerson:input_type -> models.CreatePersonRequest
	6, // 6: models.PersonService.UpdatePerson:input_type -> models.UpdatePersonRequest
	7, // 7: models.PersonService.DeletePerson:input_type -> models.DeletePersonRequest
	3, // 8: models.PersonService.ListPersons:output_type -> models.ListPersonsResponse
	0, // 9: models.PersonService.GetPerson:output_type -> models.Person
	0, // 10: models.PersonService.CreatePerson:output_type -> models.Person
	0, // 11: models.PersonService.UpdatePerson:output_type -> models.Person
	1, // 12: models.PersonService.DeletePerson:output_type -> models.Empty
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pbmodels_pbModels_proto_init() }
func file_pbmodels_pbModels_proto_init() {
	if File_pbmodels_pbModels_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pbmodels_pbModels_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Person); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pbmodels_pbModels_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pbmodels_pbModels_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPersonsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pbmodels_pbModels_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPersonsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pbmodels_pbModels_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPersonRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pbmodels_pbModels_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePersonRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pbmodels_pbModels_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePersonRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pbmodels_pbModels_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePersonRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pbmodels_pbModels_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pbmodels_pbModels_proto_goTypes,
		DependencyIndexes: file_pbmodels_pbModels_proto_depIdxs,
		MessageInfos:      file_pbmodels_pbModels_proto_msgTypes,
	}.Build()
	File_pbmodels_pbModels_proto = out.File
	file_pbmodels_pbModels_proto_rawDesc = nil
	file_pbmodels_pbModels_proto_goTypes = nil
	file_pbmodels_pbModels_proto_depIdxs = nil
}
