// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v3.21.12
// source: book.proto

package pb

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

type GetAllBooksRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        string                 `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAllBooksRequest) Reset() {
	*x = GetAllBooksRequest{}
	mi := &file_book_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAllBooksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllBooksRequest) ProtoMessage() {}

func (x *GetAllBooksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_book_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllBooksRequest.ProtoReflect.Descriptor instead.
func (*GetAllBooksRequest) Descriptor() ([]byte, []int) {
	return file_book_proto_rawDescGZIP(), []int{0}
}

func (x *GetAllBooksRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type GetAllBooksResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Books         []*Book                `protobuf:"bytes,1,rep,name=books,proto3" json:"books,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAllBooksResponse) Reset() {
	*x = GetAllBooksResponse{}
	mi := &file_book_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAllBooksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllBooksResponse) ProtoMessage() {}

func (x *GetAllBooksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_book_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllBooksResponse.ProtoReflect.Descriptor instead.
func (*GetAllBooksResponse) Descriptor() ([]byte, []int) {
	return file_book_proto_rawDescGZIP(), []int{1}
}

func (x *GetAllBooksResponse) GetBooks() []*Book {
	if x != nil {
		return x.Books
	}
	return nil
}

type BorrowBookRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	BookId        string                 `protobuf:"bytes,2,opt,name=book_id,json=bookId,proto3" json:"book_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BorrowBookRequest) Reset() {
	*x = BorrowBookRequest{}
	mi := &file_book_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BorrowBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BorrowBookRequest) ProtoMessage() {}

func (x *BorrowBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_book_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BorrowBookRequest.ProtoReflect.Descriptor instead.
func (*BorrowBookRequest) Descriptor() ([]byte, []int) {
	return file_book_proto_rawDescGZIP(), []int{2}
}

func (x *BorrowBookRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *BorrowBookRequest) GetBookId() string {
	if x != nil {
		return x.BookId
	}
	return ""
}

type BorrowBookResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BorrowBookResponse) Reset() {
	*x = BorrowBookResponse{}
	mi := &file_book_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BorrowBookResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BorrowBookResponse) ProtoMessage() {}

func (x *BorrowBookResponse) ProtoReflect() protoreflect.Message {
	mi := &file_book_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BorrowBookResponse.ProtoReflect.Descriptor instead.
func (*BorrowBookResponse) Descriptor() ([]byte, []int) {
	return file_book_proto_rawDescGZIP(), []int{3}
}

func (x *BorrowBookResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ReturnBookRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	BookId        string                 `protobuf:"bytes,2,opt,name=book_id,json=bookId,proto3" json:"book_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReturnBookRequest) Reset() {
	*x = ReturnBookRequest{}
	mi := &file_book_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReturnBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReturnBookRequest) ProtoMessage() {}

func (x *ReturnBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_book_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReturnBookRequest.ProtoReflect.Descriptor instead.
func (*ReturnBookRequest) Descriptor() ([]byte, []int) {
	return file_book_proto_rawDescGZIP(), []int{4}
}

func (x *ReturnBookRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ReturnBookRequest) GetBookId() string {
	if x != nil {
		return x.BookId
	}
	return ""
}

type ReturnBookResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReturnBookResponse) Reset() {
	*x = ReturnBookResponse{}
	mi := &file_book_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReturnBookResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReturnBookResponse) ProtoMessage() {}

func (x *ReturnBookResponse) ProtoReflect() protoreflect.Message {
	mi := &file_book_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReturnBookResponse.ProtoReflect.Descriptor instead.
func (*ReturnBookResponse) Descriptor() ([]byte, []int) {
	return file_book_proto_rawDescGZIP(), []int{5}
}

func (x *ReturnBookResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type UpdateBookStatusRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateBookStatusRequest) Reset() {
	*x = UpdateBookStatusRequest{}
	mi := &file_book_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateBookStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBookStatusRequest) ProtoMessage() {}

func (x *UpdateBookStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_book_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBookStatusRequest.ProtoReflect.Descriptor instead.
func (*UpdateBookStatusRequest) Descriptor() ([]byte, []int) {
	return file_book_proto_rawDescGZIP(), []int{6}
}

type UpdateBookStatusResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateBookStatusResponse) Reset() {
	*x = UpdateBookStatusResponse{}
	mi := &file_book_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateBookStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBookStatusResponse) ProtoMessage() {}

func (x *UpdateBookStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_book_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBookStatusResponse.ProtoReflect.Descriptor instead.
func (*UpdateBookStatusResponse) Descriptor() ([]byte, []int) {
	return file_book_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateBookStatusResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Book struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title         string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Author        string                 `protobuf:"bytes,3,opt,name=author,proto3" json:"author,omitempty"`
	UserId        string                 `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Status        string                 `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Book) Reset() {
	*x = Book{}
	mi := &file_book_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Book) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Book) ProtoMessage() {}

func (x *Book) ProtoReflect() protoreflect.Message {
	mi := &file_book_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Book.ProtoReflect.Descriptor instead.
func (*Book) Descriptor() ([]byte, []int) {
	return file_book_proto_rawDescGZIP(), []int{8}
}

func (x *Book) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Book) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Book) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *Book) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Book) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type BorrowedBook struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	BookId        string                 `protobuf:"bytes,2,opt,name=book_id,json=bookId,proto3" json:"book_id,omitempty"`
	UserId        string                 `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	BorrowedDate  string                 `protobuf:"bytes,4,opt,name=borrowed_date,json=borrowedDate,proto3" json:"borrowed_date,omitempty"`
	ReturnDate    string                 `protobuf:"bytes,5,opt,name=return_date,json=returnDate,proto3" json:"return_date,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BorrowedBook) Reset() {
	*x = BorrowedBook{}
	mi := &file_book_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BorrowedBook) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BorrowedBook) ProtoMessage() {}

func (x *BorrowedBook) ProtoReflect() protoreflect.Message {
	mi := &file_book_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BorrowedBook.ProtoReflect.Descriptor instead.
func (*BorrowedBook) Descriptor() ([]byte, []int) {
	return file_book_proto_rawDescGZIP(), []int{9}
}

func (x *BorrowedBook) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *BorrowedBook) GetBookId() string {
	if x != nil {
		return x.BookId
	}
	return ""
}

func (x *BorrowedBook) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *BorrowedBook) GetBorrowedDate() string {
	if x != nil {
		return x.BorrowedDate
	}
	return ""
}

func (x *BorrowedBook) GetReturnDate() string {
	if x != nil {
		return x.ReturnDate
	}
	return ""
}

type Status struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        string                 `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Status) Reset() {
	*x = Status{}
	mi := &file_book_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_book_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_book_proto_rawDescGZIP(), []int{10}
}

func (x *Status) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_book_proto protoreflect.FileDescriptor

var file_book_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x72, 0x65,
	0x70, 0x6f, 0x22, 0x2c, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x42, 0x6f, 0x6f, 0x6b,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0x37, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x05, 0x62, 0x6f, 0x6f, 0x6b, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x2e, 0x42, 0x6f,
	0x6f, 0x6b, 0x52, 0x05, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x22, 0x45, 0x0a, 0x11, 0x42, 0x6f, 0x72,
	0x72, 0x6f, 0x77, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x62, 0x6f, 0x6f, 0x6b, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x64,
	0x22, 0x2e, 0x0a, 0x12, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x45, 0x0a, 0x11, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17,
	0x0a, 0x07, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x22, 0x2e, 0x0a, 0x12, 0x52, 0x65, 0x74, 0x75, 0x72,
	0x6e, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x19, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x34, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x75, 0x0a, 0x04, 0x42, 0x6f, 0x6f, 0x6b,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22,
	0x96, 0x01, 0x0a, 0x0c, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x65, 0x64, 0x42, 0x6f, 0x6f, 0x6b,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x17, 0x0a, 0x07, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x65, 0x64, 0x5f, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x62, 0x6f, 0x72, 0x72, 0x6f,
	0x77, 0x65, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x74, 0x75, 0x72,
	0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65,
	0x74, 0x75, 0x72, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x22, 0x20, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0xa6, 0x02, 0x0a, 0x0b, 0x42,
	0x6f, 0x6f, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x0b, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x12, 0x18, 0x2e, 0x72, 0x65, 0x70, 0x6f,
	0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f,
	0x0a, 0x0a, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x17, 0x2e, 0x72,
	0x65, 0x70, 0x6f, 0x2e, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x2e, 0x42, 0x6f, 0x72,
	0x72, 0x6f, 0x77, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3f, 0x0a, 0x0a, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x17, 0x2e,
	0x72, 0x65, 0x70, 0x6f, 0x2e, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x42, 0x6f, 0x6f, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x2e, 0x52, 0x65,
	0x74, 0x75, 0x72, 0x6e, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x51, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x08, 0x5a, 0x06, 0x2f, 0x70, 0x62, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_book_proto_rawDescOnce sync.Once
	file_book_proto_rawDescData = file_book_proto_rawDesc
)

func file_book_proto_rawDescGZIP() []byte {
	file_book_proto_rawDescOnce.Do(func() {
		file_book_proto_rawDescData = protoimpl.X.CompressGZIP(file_book_proto_rawDescData)
	})
	return file_book_proto_rawDescData
}

var file_book_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_book_proto_goTypes = []any{
	(*GetAllBooksRequest)(nil),       // 0: repo.GetAllBooksRequest
	(*GetAllBooksResponse)(nil),      // 1: repo.GetAllBooksResponse
	(*BorrowBookRequest)(nil),        // 2: repo.BorrowBookRequest
	(*BorrowBookResponse)(nil),       // 3: repo.BorrowBookResponse
	(*ReturnBookRequest)(nil),        // 4: repo.ReturnBookRequest
	(*ReturnBookResponse)(nil),       // 5: repo.ReturnBookResponse
	(*UpdateBookStatusRequest)(nil),  // 6: repo.UpdateBookStatusRequest
	(*UpdateBookStatusResponse)(nil), // 7: repo.UpdateBookStatusResponse
	(*Book)(nil),                     // 8: repo.Book
	(*BorrowedBook)(nil),             // 9: repo.BorrowedBook
	(*Status)(nil),                   // 10: repo.Status
}
var file_book_proto_depIdxs = []int32{
	8, // 0: repo.GetAllBooksResponse.books:type_name -> repo.Book
	0, // 1: repo.BookService.GetAllBooks:input_type -> repo.GetAllBooksRequest
	2, // 2: repo.BookService.BorrowBook:input_type -> repo.BorrowBookRequest
	4, // 3: repo.BookService.ReturnBook:input_type -> repo.ReturnBookRequest
	6, // 4: repo.BookService.UpdateBookStatus:input_type -> repo.UpdateBookStatusRequest
	1, // 5: repo.BookService.GetAllBooks:output_type -> repo.GetAllBooksResponse
	3, // 6: repo.BookService.BorrowBook:output_type -> repo.BorrowBookResponse
	5, // 7: repo.BookService.ReturnBook:output_type -> repo.ReturnBookResponse
	7, // 8: repo.BookService.UpdateBookStatus:output_type -> repo.UpdateBookStatusResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_book_proto_init() }
func file_book_proto_init() {
	if File_book_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_book_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_book_proto_goTypes,
		DependencyIndexes: file_book_proto_depIdxs,
		MessageInfos:      file_book_proto_msgTypes,
	}.Build()
	File_book_proto = out.File
	file_book_proto_rawDesc = nil
	file_book_proto_goTypes = nil
	file_book_proto_depIdxs = nil
}
