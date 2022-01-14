package vestingv1beta1

import (
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	v1beta11 "github.com/cosmos/cosmos-sdk/api/cosmos/auth/v1beta1"
	v1beta1 "github.com/cosmos/cosmos-sdk/api/cosmos/base/v1beta1"
	_ "github.com/gogo/protobuf/gogoproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	reflect "reflect"
	sync "sync"
)

var _ protoreflect.List = (*_BaseVestingAccount_2_list)(nil)

type _BaseVestingAccount_2_list struct {
	list *[]*v1beta1.Coin
}

func (x *_BaseVestingAccount_2_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_BaseVestingAccount_2_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_BaseVestingAccount_2_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*v1beta1.Coin)
	(*x.list)[i] = concreteValue
}

func (x *_BaseVestingAccount_2_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*v1beta1.Coin)
	*x.list = append(*x.list, concreteValue)
}

func (x *_BaseVestingAccount_2_list) AppendMutable() protoreflect.Value {
	v := new(v1beta1.Coin)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_BaseVestingAccount_2_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_BaseVestingAccount_2_list) NewElement() protoreflect.Value {
	v := new(v1beta1.Coin)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_BaseVestingAccount_2_list) IsValid() bool {
	return x.list != nil
}

var _ protoreflect.List = (*_BaseVestingAccount_3_list)(nil)

type _BaseVestingAccount_3_list struct {
	list *[]*v1beta1.Coin
}

func (x *_BaseVestingAccount_3_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_BaseVestingAccount_3_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_BaseVestingAccount_3_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*v1beta1.Coin)
	(*x.list)[i] = concreteValue
}

func (x *_BaseVestingAccount_3_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*v1beta1.Coin)
	*x.list = append(*x.list, concreteValue)
}

func (x *_BaseVestingAccount_3_list) AppendMutable() protoreflect.Value {
	v := new(v1beta1.Coin)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_BaseVestingAccount_3_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_BaseVestingAccount_3_list) NewElement() protoreflect.Value {
	v := new(v1beta1.Coin)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_BaseVestingAccount_3_list) IsValid() bool {
	return x.list != nil
}

var _ protoreflect.List = (*_BaseVestingAccount_4_list)(nil)

type _BaseVestingAccount_4_list struct {
	list *[]*v1beta1.Coin
}

func (x *_BaseVestingAccount_4_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_BaseVestingAccount_4_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_BaseVestingAccount_4_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*v1beta1.Coin)
	(*x.list)[i] = concreteValue
}

func (x *_BaseVestingAccount_4_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*v1beta1.Coin)
	*x.list = append(*x.list, concreteValue)
}

func (x *_BaseVestingAccount_4_list) AppendMutable() protoreflect.Value {
	v := new(v1beta1.Coin)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_BaseVestingAccount_4_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_BaseVestingAccount_4_list) NewElement() protoreflect.Value {
	v := new(v1beta1.Coin)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_BaseVestingAccount_4_list) IsValid() bool {
	return x.list != nil
}

var (
	md_BaseVestingAccount                   protoreflect.MessageDescriptor
	fd_BaseVestingAccount_base_account      protoreflect.FieldDescriptor
	fd_BaseVestingAccount_original_vesting  protoreflect.FieldDescriptor
	fd_BaseVestingAccount_delegated_free    protoreflect.FieldDescriptor
	fd_BaseVestingAccount_delegated_vesting protoreflect.FieldDescriptor
	fd_BaseVestingAccount_end_time          protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_vesting_v1beta1_vesting_proto_init()
	md_BaseVestingAccount = File_cosmos_vesting_v1beta1_vesting_proto.Messages().ByName("BaseVestingAccount")
	fd_BaseVestingAccount_base_account = md_BaseVestingAccount.Fields().ByName("base_account")
	fd_BaseVestingAccount_original_vesting = md_BaseVestingAccount.Fields().ByName("original_vesting")
	fd_BaseVestingAccount_delegated_free = md_BaseVestingAccount.Fields().ByName("delegated_free")
	fd_BaseVestingAccount_delegated_vesting = md_BaseVestingAccount.Fields().ByName("delegated_vesting")
	fd_BaseVestingAccount_end_time = md_BaseVestingAccount.Fields().ByName("end_time")
}

var _ protoreflect.Message = (*fastReflection_BaseVestingAccount)(nil)

type fastReflection_BaseVestingAccount BaseVestingAccount

func (x *BaseVestingAccount) ProtoReflect() protoreflect.Message {
	return (*fastReflection_BaseVestingAccount)(x)
}

func (x *BaseVestingAccount) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_vesting_v1beta1_vesting_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_BaseVestingAccount_messageType fastReflection_BaseVestingAccount_messageType
var _ protoreflect.MessageType = fastReflection_BaseVestingAccount_messageType{}

type fastReflection_BaseVestingAccount_messageType struct{}

func (x fastReflection_BaseVestingAccount_messageType) Zero() protoreflect.Message {
	return (*fastReflection_BaseVestingAccount)(nil)
}
func (x fastReflection_BaseVestingAccount_messageType) New() protoreflect.Message {
	return new(fastReflection_BaseVestingAccount)
}
func (x fastReflection_BaseVestingAccount_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_BaseVestingAccount
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_BaseVestingAccount) Descriptor() protoreflect.MessageDescriptor {
	return md_BaseVestingAccount
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_BaseVestingAccount) Type() protoreflect.MessageType {
	return _fastReflection_BaseVestingAccount_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_BaseVestingAccount) New() protoreflect.Message {
	return new(fastReflection_BaseVestingAccount)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_BaseVestingAccount) Interface() protoreflect.ProtoMessage {
	return (*BaseVestingAccount)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_BaseVestingAccount) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.BaseAccount != nil {
		value := protoreflect.ValueOfMessage(x.BaseAccount.ProtoReflect())
		if !f(fd_BaseVestingAccount_base_account, value) {
			return
		}
	}
	if len(x.OriginalVesting) != 0 {
		value := protoreflect.ValueOfList(&_BaseVestingAccount_2_list{list: &x.OriginalVesting})
		if !f(fd_BaseVestingAccount_original_vesting, value) {
			return
		}
	}
	if len(x.DelegatedFree) != 0 {
		value := protoreflect.ValueOfList(&_BaseVestingAccount_3_list{list: &x.DelegatedFree})
		if !f(fd_BaseVestingAccount_delegated_free, value) {
			return
		}
	}
	if len(x.DelegatedVesting) != 0 {
		value := protoreflect.ValueOfList(&_BaseVestingAccount_4_list{list: &x.DelegatedVesting})
		if !f(fd_BaseVestingAccount_delegated_vesting, value) {
			return
		}
	}
	if x.EndTime != int64(0) {
		value := protoreflect.ValueOfInt64(x.EndTime)
		if !f(fd_BaseVestingAccount_end_time, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_BaseVestingAccount) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.vesting.v1beta1.BaseVestingAccount.base_account":
		return x.BaseAccount != nil
	case "cosmos.vesting.v1beta1.BaseVestingAccount.original_vesting":
		return len(x.OriginalVesting) != 0
	case "cosmos.vesting.v1beta1.BaseVestingAccount.delegated_free":
		return len(x.DelegatedFree) != 0
	case "cosmos.vesting.v1beta1.BaseVestingAccount.delegated_vesting":
		return len(x.DelegatedVesting) != 0
	case "cosmos.vesting.v1beta1.BaseVestingAccount.end_time":
		return x.EndTime != int64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.vesting.v1beta1.BaseVestingAccount"))
		}
		panic(fmt.Errorf("message cosmos.vesting.v1beta1.BaseVestingAccount does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_BaseVestingAccount) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.vesting.v1beta1.BaseVestingAccount.base_account":
		x.BaseAccount = nil
	case "cosmos.vesting.v1beta1.BaseVestingAccount.original_vesting":
		x.OriginalVesting = nil
	case "cosmos.vesting.v1beta1.BaseVestingAccount.delegated_free":
		x.DelegatedFree = nil
	case "cosmos.vesting.v1beta1.BaseVestingAccount.delegated_vesting":
		x.DelegatedVesting = nil
	case "cosmos.vesting.v1beta1.BaseVestingAccount.end_time":
		x.EndTime = int64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.vesting.v1beta1.BaseVestingAccount"))
		}
		panic(fmt.Errorf("message cosmos.vesting.v1beta1.BaseVestingAccount does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_BaseVestingAccount) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.vesting.v1beta1.BaseVestingAccount.base_account":
		value := x.BaseAccount
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "cosmos.vesting.v1beta1.BaseVestingAccount.original_vesting":
		if len(x.OriginalVesting) == 0 {
			return protoreflect.ValueOfList(&_BaseVestingAccount_2_list{})
		}
		listValue := &_BaseVestingAccount_2_list{list: &x.OriginalVesting}
		return protoreflect.ValueOfList(listValue)
	case "cosmos.vesting.v1beta1.BaseVestingAccount.delegated_free":
		if len(x.DelegatedFree) == 0 {
			return protoreflect.ValueOfList(&_BaseVestingAccount_3_list{})
		}
		listValue := &_BaseVestingAccount_3_list{list: &x.DelegatedFree}
		return protoreflect.ValueOfList(listValue)
	case "cosmos.vesting.v1beta1.BaseVestingAccount.delegated_vesting":
		if len(x.DelegatedVesting) == 0 {
			return protoreflect.ValueOfList(&_BaseVestingAccount_4_list{})
		}
		listValue := &_BaseVestingAccount_4_list{list: &x.DelegatedVesting}
		return protoreflect.ValueOfList(listValue)
	case "cosmos.vesting.v1beta1.BaseVestingAccount.end_time":
		value := x.EndTime
		return protoreflect.ValueOfInt64(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.vesting.v1beta1.BaseVestingAccount"))
		}
		panic(fmt.Errorf("message cosmos.vesting.v1beta1.BaseVestingAccount does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_BaseVestingAccount) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.vesting.v1beta1.BaseVestingAccount.base_account":
		x.BaseAccount = value.Message().Interface().(*v1beta11.BaseAccount)
	case "cosmos.vesting.v1beta1.BaseVestingAccount.original_vesting":
		lv := value.List()
		clv := lv.(*_BaseVestingAccount_2_list)
		x.OriginalVesting = *clv.list
	case "cosmos.vesting.v1beta1.BaseVestingAccount.delegated_free":
		lv := value.List()
		clv := lv.(*_BaseVestingAccount_3_list)
		x.DelegatedFree = *clv.list
	case "cosmos.vesting.v1beta1.BaseVestingAccount.delegated_vesting":
		lv := value.List()
		clv := lv.(*_BaseVestingAccount_4_list)
		x.DelegatedVesting = *clv.list
	case "cosmos.vesting.v1beta1.BaseVestingAccount.end_time":
		x.EndTime = value.Int()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.vesting.v1beta1.BaseVestingAccount"))
		}
		panic(fmt.Errorf("message cosmos.vesting.v1beta1.BaseVestingAccount does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_BaseVestingAccount) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.vesting.v1beta1.BaseVestingAccount.base_account":
		if x.BaseAccount == nil {
			x.BaseAccount = new(v1beta11.BaseAccount)
		}
		return protoreflect.ValueOfMessage(x.BaseAccount.ProtoReflect())
	case "cosmos.vesting.v1beta1.BaseVestingAccount.original_vesting":
		if x.OriginalVesting == nil {
			x.OriginalVesting = []*v1beta1.Coin{}
		}
		value := &_BaseVestingAccount_2_list{list: &x.OriginalVesting}
		return protoreflect.ValueOfList(value)
	case "cosmos.vesting.v1beta1.BaseVestingAccount.delegated_free":
		if x.DelegatedFree == nil {
			x.DelegatedFree = []*v1beta1.Coin{}
		}
		value := &_BaseVestingAccount_3_list{list: &x.DelegatedFree}
		return protoreflect.ValueOfList(value)
	case "cosmos.vesting.v1beta1.BaseVestingAccount.delegated_vesting":
		if x.DelegatedVesting == nil {
			x.DelegatedVesting = []*v1beta1.Coin{}
		}
		value := &_BaseVestingAccount_4_list{list: &x.DelegatedVesting}
		return protoreflect.ValueOfList(value)
	case "cosmos.vesting.v1beta1.BaseVestingAccount.end_time":
		panic(fmt.Errorf("field end_time of message cosmos.vesting.v1beta1.BaseVestingAccount is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.vesting.v1beta1.BaseVestingAccount"))
		}
		panic(fmt.Errorf("message cosmos.vesting.v1beta1.BaseVestingAccount does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_BaseVestingAccount) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.vesting.v1beta1.BaseVestingAccount.base_account":
		m := new(v1beta11.BaseAccount)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "cosmos.vesting.v1beta1.BaseVestingAccount.original_vesting":
		list := []*v1beta1.Coin{}
		return protoreflect.ValueOfList(&_BaseVestingAccount_2_list{list: &list})
	case "cosmos.vesting.v1beta1.BaseVestingAccount.delegated_free":
		list := []*v1beta1.Coin{}
		return protoreflect.ValueOfList(&_BaseVestingAccount_3_list{list: &list})
	case "cosmos.vesting.v1beta1.BaseVestingAccount.delegated_vesting":
		list := []*v1beta1.Coin{}
		return protoreflect.ValueOfList(&_BaseVestingAccount_4_list{list: &list})
	case "cosmos.vesting.v1beta1.BaseVestingAccount.end_time":
		return protoreflect.ValueOfInt64(int64(0))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.vesting.v1beta1.BaseVestingAccount"))
		}
		panic(fmt.Errorf("message cosmos.vesting.v1beta1.BaseVestingAccount does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_BaseVestingAccount) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.vesting.v1beta1.BaseVestingAccount", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_BaseVestingAccount) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_BaseVestingAccount) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_BaseVestingAccount) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_BaseVestingAccount) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*BaseVestingAccount)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.BaseAccount != nil {
			l = options.Size(x.BaseAccount)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if len(x.OriginalVesting) > 0 {
			for _, e := range x.OriginalVesting {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if len(x.DelegatedFree) > 0 {
			for _, e := range x.DelegatedFree {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if len(x.DelegatedVesting) > 0 {
			for _, e := range x.DelegatedVesting {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if x.EndTime != 0 {
			n += 1 + runtime.Sov(uint64(x.EndTime))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*BaseVestingAccount)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.EndTime != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.EndTime))
			i--
			dAtA[i] = 0x28
		}
		if len(x.DelegatedVesting) > 0 {
			for iNdEx := len(x.DelegatedVesting) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.DelegatedVesting[iNdEx])
				if err != nil {
					return protoiface.MarshalOutput{
						NoUnkeyedLiterals: input.NoUnkeyedLiterals,
						Buf:               input.Buf,
					}, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
				i--
				dAtA[i] = 0x22
			}
		}
		if len(x.DelegatedFree) > 0 {
			for iNdEx := len(x.DelegatedFree) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.DelegatedFree[iNdEx])
				if err != nil {
					return protoiface.MarshalOutput{
						NoUnkeyedLiterals: input.NoUnkeyedLiterals,
						Buf:               input.Buf,
					}, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
				i--
				dAtA[i] = 0x1a
			}
		}
		if len(x.OriginalVesting) > 0 {
			for iNdEx := len(x.OriginalVesting) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.OriginalVesting[iNdEx])
				if err != nil {
					return protoiface.MarshalOutput{
						NoUnkeyedLiterals: input.NoUnkeyedLiterals,
						Buf:               input.Buf,
					}, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
				i--
				dAtA[i] = 0x12
			}
		}
		if x.BaseAccount != nil {
			encoded, err := options.Marshal(x.BaseAccount)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*BaseVestingAccount)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: BaseVestingAccount: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: BaseVestingAccount: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field BaseAccount", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.BaseAccount == nil {
					x.BaseAccount = &v1beta11.BaseAccount{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.BaseAccount); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field OriginalVesting", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.OriginalVesting = append(x.OriginalVesting, &v1beta1.Coin{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.OriginalVesting[len(x.OriginalVesting)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field DelegatedFree", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.DelegatedFree = append(x.DelegatedFree, &v1beta1.Coin{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.DelegatedFree[len(x.DelegatedFree)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field DelegatedVesting", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.DelegatedVesting = append(x.DelegatedVesting, &v1beta1.Coin{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.DelegatedVesting[len(x.DelegatedVesting)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 5:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field EndTime", wireType)
				}
				x.EndTime = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.EndTime |= int64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_ContinuousVestingAccount                      protoreflect.MessageDescriptor
	fd_ContinuousVestingAccount_base_vesting_account protoreflect.FieldDescriptor
	fd_ContinuousVestingAccount_start_time           protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_vesting_v1beta1_vesting_proto_init()
	md_ContinuousVestingAccount = File_cosmos_vesting_v1beta1_vesting_proto.Messages().ByName("ContinuousVestingAccount")
	fd_ContinuousVestingAccount_base_vesting_account = md_ContinuousVestingAccount.Fields().ByName("base_vesting_account")
	fd_ContinuousVestingAccount_start_time = md_ContinuousVestingAccount.Fields().ByName("start_time")
}

var _ protoreflect.Message = (*fastReflection_ContinuousVestingAccount)(nil)

type fastReflection_ContinuousVestingAccount ContinuousVestingAccount

func (x *ContinuousVestingAccount) ProtoReflect() protoreflect.Message {
	return (*fastReflection_ContinuousVestingAccount)(x)
}

func (x *ContinuousVestingAccount) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_vesting_v1beta1_vesting_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_ContinuousVestingAccount_messageType fastReflection_ContinuousVestingAccount_messageType
var _ protoreflect.MessageType = fastReflection_ContinuousVestingAccount_messageType{}

type fastReflection_ContinuousVestingAccount_messageType struct{}

func (x fastReflection_ContinuousVestingAccount_messageType) Zero() protoreflect.Message {
	return (*fastReflection_ContinuousVestingAccount)(nil)
}
func (x fastReflection_ContinuousVestingAccount_messageType) New() protoreflect.Message {
	return new(fastReflection_ContinuousVestingAccount)
}
func (x fastReflection_ContinuousVestingAccount_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_ContinuousVestingAccount
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_ContinuousVestingAccount) Descriptor() protoreflect.MessageDescriptor {
	return md_ContinuousVestingAccount
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_ContinuousVestingAccount) Type() protoreflect.MessageType {
	return _fastReflection_ContinuousVestingAccount_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_ContinuousVestingAccount) New() protoreflect.Message {
	return new(fastReflection_ContinuousVestingAccount)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_ContinuousVestingAccount) Interface() protoreflect.ProtoMessage {
	return (*ContinuousVestingAccount)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_ContinuousVestingAccount) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.BaseVestingAccount != nil {
		value := protoreflect.ValueOfMessage(x.BaseVestingAccount.ProtoReflect())
		if !f(fd_ContinuousVestingAccount_base_vesting_account, value) {
			return
		}
	}
	if x.StartTime != int64(0) {
		value := protoreflect.ValueOfInt64(x.StartTime)
		if !f(fd_ContinuousVestingAccount_start_time, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_ContinuousVestingAccount) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.vesting.v1beta1.ContinuousVestingAccount.base_vesting_account":
		return x.BaseVestingAccount != nil
	case "cosmos.vesting.v1beta1.ContinuousVestingAccount.start_time":
		return x.StartTime != int64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.vesting.v1beta1.ContinuousVestingAccount"))
		}
		panic(fmt.Errorf("message cosmos.vesting.v1beta1.ContinuousVestingAccount does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ContinuousVestingAccount) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.vesting.v1beta1.ContinuousVestingAccount.base_vesting_account":
		x.BaseVestingAccount = nil
	case "cosmos.vesting.v1beta1.ContinuousVestingAccount.start_time":
		x.StartTime = int64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.vesting.v1beta1.ContinuousVestingAccount"))
		}
		panic(fmt.Errorf("message cosmos.vesting.v1beta1.ContinuousVestingAccount does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_ContinuousVestingAccount) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.vesting.v1beta1.ContinuousVestingAccount.base_vesting_account":
		value := x.BaseVestingAccount
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "cosmos.vesting.v1beta1.ContinuousVestingAccount.start_time":
		value := x.StartTime
		return protoreflect.ValueOfInt64(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.vesting.v1beta1.ContinuousVestingAccount"))
		}
		panic(fmt.Errorf("message cosmos.vesting.v1beta1.ContinuousVestingAccount does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ContinuousVestingAccount) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.vesting.v1beta1.ContinuousVestingAccount.base_vesting_account":
		x.BaseVestingAccount = value.Message().Interface().(*BaseVestingAccount)
	case "cosmos.vesting.v1beta1.ContinuousVestingAccount.start_time":
		x.StartTime = value.Int()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.vesting.v1beta1.ContinuousVestingAccount"))
		}
		panic(fmt.Errorf("message cosmos.vesting.v1beta1.ContinuousVestingAccount does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ContinuousVestingAccount) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.vesting.v1beta1.ContinuousVestingAccount.base_vesting_account":
		if x.BaseVestingAccount == nil {
			x.BaseVestingAccount = new(BaseVestingAccount)
		}
		return protoreflect.ValueOfMessage(x.BaseVestingAccount.ProtoReflect())
	case "cosmos.vesting.v1beta1.ContinuousVestingAccount.start_time":
		panic(fmt.Errorf("field start_time of message cosmos.vesting.v1beta1.ContinuousVestingAccount is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.vesting.v1beta1.ContinuousVestingAccount"))
		}
		panic(fmt.Errorf("message cosmos.vesting.v1beta1.ContinuousVestingAccount does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_ContinuousVestingAccount) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.vesting.v1beta1.ContinuousVestingAccount.base_vesting_account":
		m := new(BaseVestingAccount)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "cosmos.vesting.v1beta1.ContinuousVestingAccount.start_time":
		return protoreflect.ValueOfInt64(int64(0))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.vesting.v1beta1.ContinuousVestingAccount"))
		}
		panic(fmt.Errorf("message cosmos.vesting.v1beta1.ContinuousVestingAccount does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_ContinuousVestingAccount) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.vesting.v1beta1.ContinuousVestingAccount", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_ContinuousVestingAccount) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ContinuousVestingAccount) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_ContinuousVestingAccount) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_ContinuousVestingAccount) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*ContinuousVestingAccount)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.BaseVestingAccount != nil {
			l = options.Size(x.BaseVestingAccount)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.StartTime != 0 {
			n += 1 + runtime.Sov(uint64(x.StartTime))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*ContinuousVestingAccount)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.StartTime != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.StartTime))
			i--
			dAtA[i] = 0x10
		}
		if x.BaseVestingAccount != nil {
			encoded, err := options.Marshal(x.BaseVestingAccount)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*ContinuousVestingAccount)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ContinuousVestingAccount: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ContinuousVestingAccount: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field BaseVestingAccount", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.BaseVestingAccount == nil {
					x.BaseVestingAccount = &BaseVestingAccount{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.BaseVestingAccount); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 2:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field StartTime", wireType)
				}
				x.StartTime = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.StartTime |= int64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_DelayedVestingAccount                      protoreflect.MessageDescriptor
	fd_DelayedVestingAccount_base_vesting_account protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_vesting_v1beta1_vesting_proto_init()
	md_DelayedVestingAccount = File_cosmos_vesting_v1beta1_vesting_proto.Messages().ByName("DelayedVestingAccount")
	fd_DelayedVestingAccount_base_vesting_account = md_DelayedVestingAccount.Fields().ByName("base_vesting_account")
}

var _ protoreflect.Message = (*fastReflection_DelayedVestingAccount)(nil)

type fastReflection_DelayedVestingAccount DelayedVestingAccount

func (x *DelayedVestingAccount) ProtoReflect() protoreflect.Message {
	return (*fastReflection_DelayedVestingAccount)(x)
}

func (x *DelayedVestingAccount) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_vesting_v1beta1_vesting_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_DelayedVestingAccount_messageType fastReflection_DelayedVestingAccount_messageType
var _ protoreflect.MessageType = fastReflection_DelayedVestingAccount_messageType{}

type fastReflection_DelayedVestingAccount_messageType struct{}

func (x fastReflection_DelayedVestingAccount_messageType) Zero() protoreflect.Message {
	return (*fastReflection_DelayedVestingAccount)(nil)
}
func (x fastReflection_DelayedVestingAccount_messageType) New() protoreflect.Message {
	return new(fastReflection_DelayedVestingAccount)
}
func (x fastReflection_DelayedVestingAccount_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_DelayedVestingAccount
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_DelayedVestingAccount) Descriptor() protoreflect.MessageDescriptor {
	return md_DelayedVestingAccount
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_DelayedVestingAccount) Type() protoreflect.MessageType {
	return _fastReflection_DelayedVestingAccount_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_DelayedVestingAccount) New() protoreflect.Message {
	return new(fastReflection_DelayedVestingAccount)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_DelayedVestingAccount) Interface() protoreflect.ProtoMessage {
	return (*DelayedVestingAccount)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_DelayedVestingAccount) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.BaseVestingAccount != nil {
		value := protoreflect.ValueOfMessage(x.BaseVestingAccount.ProtoReflect())
		if !f(fd_DelayedVestingAccount_base_vesting_account, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_DelayedVestingAccount) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.vesting.v1beta1.DelayedVestingAccount.base_vesting_account":
		return x.BaseVestingAccount != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.vesting.v1beta1.DelayedVestingAccount"))
		}
		panic(fmt.Errorf("message cosmos.vesting.v1beta1.DelayedVestingAccount does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_DelayedVestingAccount) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.vesting.v1beta1.DelayedVestingAccount.base_vesting_account":
		x.BaseVestingAccount = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.vesting.v1beta1.DelayedVestingAccount"))
		}
		panic(fmt.Errorf("message cosmos.vesting.v1beta1.DelayedVestingAccount does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_DelayedVestingAccount) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.vesting.v1beta1.DelayedVestingAccount.base_vesting_account":
		value := x.BaseVestingAccount
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.vesting.v1beta1.DelayedVestingAccount"))
		}
		panic(fmt.Errorf("message cosmos.vesting.v1beta1.DelayedVestingAccount does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_DelayedVestingAccount) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.vesting.v1beta1.DelayedVestingAccount.base_vesting_account":
		x.BaseVestingAccount = value.Message().Interface().(*BaseVestingAccount)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.vesting.v1beta1.DelayedVestingAccount"))
		}
		panic(fmt.Errorf("message cosmos.vesting.v1beta1.DelayedVestingAccount does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_DelayedVestingAccount) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.vesting.v1beta1.DelayedVestingAccount.base_vesting_account":
		if x.BaseVestingAccount == nil {
			x.BaseVestingAccount = new(BaseVestingAccount)
		}
		return protoreflect.ValueOfMessage(x.BaseVestingAccount.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.vesting.v1beta1.DelayedVestingAccount"))
		}
		panic(fmt.Errorf("message cosmos.vesting.v1beta1.DelayedVestingAccount does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_DelayedVestingAccount) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.vesting.v1beta1.DelayedVestingAccount.base_vesting_account":
		m := new(BaseVestingAccount)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.vesting.v1beta1.DelayedVestingAccount"))
		}
		panic(fmt.Errorf("message cosmos.vesting.v1beta1.DelayedVestingAccount does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_DelayedVestingAccount) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.vesting.v1beta1.DelayedVestingAccount", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_DelayedVestingAccount) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_DelayedVestingAccount) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_DelayedVestingAccount) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_DelayedVestingAccount) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*DelayedVestingAccount)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.BaseVestingAccount != nil {
			l = options.Size(x.BaseVestingAccount)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*DelayedVestingAccount)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.BaseVestingAccount != nil {
			encoded, err := options.Marshal(x.BaseVestingAccount)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*DelayedVestingAccount)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: DelayedVestingAccount: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: DelayedVestingAccount: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field BaseVestingAccount", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.BaseVestingAccount == nil {
					x.BaseVestingAccount = &BaseVestingAccount{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.BaseVestingAccount); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var _ protoreflect.List = (*_Period_2_list)(nil)

type _Period_2_list struct {
	list *[]*v1beta1.Coin
}

func (x *_Period_2_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_Period_2_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_Period_2_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*v1beta1.Coin)
	(*x.list)[i] = concreteValue
}

func (x *_Period_2_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*v1beta1.Coin)
	*x.list = append(*x.list, concreteValue)
}

func (x *_Period_2_list) AppendMutable() protoreflect.Value {
	v := new(v1beta1.Coin)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_Period_2_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_Period_2_list) NewElement() protoreflect.Value {
	v := new(v1beta1.Coin)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_Period_2_list) IsValid() bool {
	return x.list != nil
}

var (
	md_Period        protoreflect.MessageDescriptor
	fd_Period_length protoreflect.FieldDescriptor
	fd_Period_amount protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_vesting_v1beta1_vesting_proto_init()
	md_Period = File_cosmos_vesting_v1beta1_vesting_proto.Messages().ByName("Period")
	fd_Period_length = md_Period.Fields().ByName("length")
	fd_Period_amount = md_Period.Fields().ByName("amount")
}

var _ protoreflect.Message = (*fastReflection_Period)(nil)

type fastReflection_Period Period

func (x *Period) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Period)(x)
}

func (x *Period) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_vesting_v1beta1_vesting_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Period_messageType fastReflection_Period_messageType
var _ protoreflect.MessageType = fastReflection_Period_messageType{}

type fastReflection_Period_messageType struct{}

func (x fastReflection_Period_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Period)(nil)
}
func (x fastReflection_Period_messageType) New() protoreflect.Message {
	return new(fastReflection_Period)
}
func (x fastReflection_Period_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Period
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Period) Descriptor() protoreflect.MessageDescriptor {
	return md_Period
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Period) Type() protoreflect.MessageType {
	return _fastReflection_Period_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Period) New() protoreflect.Message {
	return new(fastReflection_Period)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Period) Interface() protoreflect.ProtoMessage {
	return (*Period)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Period) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Length != int64(0) {
		value := protoreflect.ValueOfInt64(x.Length)
		if !f(fd_Period_length, value) {
			return
		}
	}
	if len(x.Amount) != 0 {
		value := protoreflect.ValueOfList(&_Period_2_list{list: &x.Amount})
		if !f(fd_Period_amount, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_Period) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.vesting.v1beta1.Period.length":
		return x.Length != int64(0)
	case "cosmos.vesting.v1beta1.Period.amount":
		return len(x.Amount) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.vesting.v1beta1.Period"))
		}
		panic(fmt.Errorf("message cosmos.vesting.v1beta1.Period does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Period) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.vesting.v1beta1.Period.length":
		x.Length = int64(0)
	case "cosmos.vesting.v1beta1.Period.amount":
		x.Amount = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.vesting.v1beta1.Period"))
		}
		panic(fmt.Errorf("message cosmos.vesting.v1beta1.Period does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Period) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.vesting.v1beta1.Period.length":
		value := x.Length
		return protoreflect.ValueOfInt64(value)
	case "cosmos.vesting.v1beta1.Period.amount":
		if len(x.Amount) == 0 {
			return protoreflect.ValueOfList(&_Period_2_list{})
		}
		listValue := &_Period_2_list{list: &x.Amount}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.vesting.v1beta1.Period"))
		}
		panic(fmt.Errorf("message cosmos.vesting.v1beta1.Period does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Period) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.vesting.v1beta1.Period.length":
		x.Length = value.Int()
	case "cosmos.vesting.v1beta1.Period.amount":
		lv := value.List()
		clv := lv.(*_Period_2_list)
		x.Amount = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.vesting.v1beta1.Period"))
		}
		panic(fmt.Errorf("message cosmos.vesting.v1beta1.Period does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Period) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.vesting.v1beta1.Period.amount":
		if x.Amount == nil {
			x.Amount = []*v1beta1.Coin{}
		}
		value := &_Period_2_list{list: &x.Amount}
		return protoreflect.ValueOfList(value)
	case "cosmos.vesting.v1beta1.Period.length":
		panic(fmt.Errorf("field length of message cosmos.vesting.v1beta1.Period is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.vesting.v1beta1.Period"))
		}
		panic(fmt.Errorf("message cosmos.vesting.v1beta1.Period does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Period) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.vesting.v1beta1.Period.length":
		return protoreflect.ValueOfInt64(int64(0))
	case "cosmos.vesting.v1beta1.Period.amount":
		list := []*v1beta1.Coin{}
		return protoreflect.ValueOfList(&_Period_2_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.vesting.v1beta1.Period"))
		}
		panic(fmt.Errorf("message cosmos.vesting.v1beta1.Period does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Period) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.vesting.v1beta1.Period", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Period) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Period) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_Period) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Period) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Period)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.Length != 0 {
			n += 1 + runtime.Sov(uint64(x.Length))
		}
		if len(x.Amount) > 0 {
			for _, e := range x.Amount {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*Period)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.Amount) > 0 {
			for iNdEx := len(x.Amount) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.Amount[iNdEx])
				if err != nil {
					return protoiface.MarshalOutput{
						NoUnkeyedLiterals: input.NoUnkeyedLiterals,
						Buf:               input.Buf,
					}, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
				i--
				dAtA[i] = 0x12
			}
		}
		if x.Length != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Length))
			i--
			dAtA[i] = 0x8
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*Period)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Period: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Period: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Length", wireType)
				}
				x.Length = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Length |= int64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Amount = append(x.Amount, &v1beta1.Coin{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Amount[len(x.Amount)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var _ protoreflect.List = (*_PeriodicVestingAccount_3_list)(nil)

type _PeriodicVestingAccount_3_list struct {
	list *[]*Period
}

func (x *_PeriodicVestingAccount_3_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_PeriodicVestingAccount_3_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_PeriodicVestingAccount_3_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*Period)
	(*x.list)[i] = concreteValue
}

func (x *_PeriodicVestingAccount_3_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*Period)
	*x.list = append(*x.list, concreteValue)
}

func (x *_PeriodicVestingAccount_3_list) AppendMutable() protoreflect.Value {
	v := new(Period)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_PeriodicVestingAccount_3_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_PeriodicVestingAccount_3_list) NewElement() protoreflect.Value {
	v := new(Period)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_PeriodicVestingAccount_3_list) IsValid() bool {
	return x.list != nil
}

var (
	md_PeriodicVestingAccount                      protoreflect.MessageDescriptor
	fd_PeriodicVestingAccount_base_vesting_account protoreflect.FieldDescriptor
	fd_PeriodicVestingAccount_start_time           protoreflect.FieldDescriptor
	fd_PeriodicVestingAccount_vesting_periods      protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_vesting_v1beta1_vesting_proto_init()
	md_PeriodicVestingAccount = File_cosmos_vesting_v1beta1_vesting_proto.Messages().ByName("PeriodicVestingAccount")
	fd_PeriodicVestingAccount_base_vesting_account = md_PeriodicVestingAccount.Fields().ByName("base_vesting_account")
	fd_PeriodicVestingAccount_start_time = md_PeriodicVestingAccount.Fields().ByName("start_time")
	fd_PeriodicVestingAccount_vesting_periods = md_PeriodicVestingAccount.Fields().ByName("vesting_periods")
}

var _ protoreflect.Message = (*fastReflection_PeriodicVestingAccount)(nil)

type fastReflection_PeriodicVestingAccount PeriodicVestingAccount

func (x *PeriodicVestingAccount) ProtoReflect() protoreflect.Message {
	return (*fastReflection_PeriodicVestingAccount)(x)
}

func (x *PeriodicVestingAccount) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_vesting_v1beta1_vesting_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_PeriodicVestingAccount_messageType fastReflection_PeriodicVestingAccount_messageType
var _ protoreflect.MessageType = fastReflection_PeriodicVestingAccount_messageType{}

type fastReflection_PeriodicVestingAccount_messageType struct{}

func (x fastReflection_PeriodicVestingAccount_messageType) Zero() protoreflect.Message {
	return (*fastReflection_PeriodicVestingAccount)(nil)
}
func (x fastReflection_PeriodicVestingAccount_messageType) New() protoreflect.Message {
	return new(fastReflection_PeriodicVestingAccount)
}
func (x fastReflection_PeriodicVestingAccount_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_PeriodicVestingAccount
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_PeriodicVestingAccount) Descriptor() protoreflect.MessageDescriptor {
	return md_PeriodicVestingAccount
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_PeriodicVestingAccount) Type() protoreflect.MessageType {
	return _fastReflection_PeriodicVestingAccount_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_PeriodicVestingAccount) New() protoreflect.Message {
	return new(fastReflection_PeriodicVestingAccount)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_PeriodicVestingAccount) Interface() protoreflect.ProtoMessage {
	return (*PeriodicVestingAccount)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_PeriodicVestingAccount) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.BaseVestingAccount != nil {
		value := protoreflect.ValueOfMessage(x.BaseVestingAccount.ProtoReflect())
		if !f(fd_PeriodicVestingAccount_base_vesting_account, value) {
			return
		}
	}
	if x.StartTime != int64(0) {
		value := protoreflect.ValueOfInt64(x.StartTime)
		if !f(fd_PeriodicVestingAccount_start_time, value) {
			return
		}
	}
	if len(x.VestingPeriods) != 0 {
		value := protoreflect.ValueOfList(&_PeriodicVestingAccount_3_list{list: &x.VestingPeriods})
		if !f(fd_PeriodicVestingAccount_vesting_periods, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_PeriodicVestingAccount) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.vesting.v1beta1.PeriodicVestingAccount.base_vesting_account":
		return x.BaseVestingAccount != nil
	case "cosmos.vesting.v1beta1.PeriodicVestingAccount.start_time":
		return x.StartTime != int64(0)
	case "cosmos.vesting.v1beta1.PeriodicVestingAccount.vesting_periods":
		return len(x.VestingPeriods) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.vesting.v1beta1.PeriodicVestingAccount"))
		}
		panic(fmt.Errorf("message cosmos.vesting.v1beta1.PeriodicVestingAccount does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_PeriodicVestingAccount) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.vesting.v1beta1.PeriodicVestingAccount.base_vesting_account":
		x.BaseVestingAccount = nil
	case "cosmos.vesting.v1beta1.PeriodicVestingAccount.start_time":
		x.StartTime = int64(0)
	case "cosmos.vesting.v1beta1.PeriodicVestingAccount.vesting_periods":
		x.VestingPeriods = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.vesting.v1beta1.PeriodicVestingAccount"))
		}
		panic(fmt.Errorf("message cosmos.vesting.v1beta1.PeriodicVestingAccount does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_PeriodicVestingAccount) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.vesting.v1beta1.PeriodicVestingAccount.base_vesting_account":
		value := x.BaseVestingAccount
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "cosmos.vesting.v1beta1.PeriodicVestingAccount.start_time":
		value := x.StartTime
		return protoreflect.ValueOfInt64(value)
	case "cosmos.vesting.v1beta1.PeriodicVestingAccount.vesting_periods":
		if len(x.VestingPeriods) == 0 {
			return protoreflect.ValueOfList(&_PeriodicVestingAccount_3_list{})
		}
		listValue := &_PeriodicVestingAccount_3_list{list: &x.VestingPeriods}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.vesting.v1beta1.PeriodicVestingAccount"))
		}
		panic(fmt.Errorf("message cosmos.vesting.v1beta1.PeriodicVestingAccount does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_PeriodicVestingAccount) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.vesting.v1beta1.PeriodicVestingAccount.base_vesting_account":
		x.BaseVestingAccount = value.Message().Interface().(*BaseVestingAccount)
	case "cosmos.vesting.v1beta1.PeriodicVestingAccount.start_time":
		x.StartTime = value.Int()
	case "cosmos.vesting.v1beta1.PeriodicVestingAccount.vesting_periods":
		lv := value.List()
		clv := lv.(*_PeriodicVestingAccount_3_list)
		x.VestingPeriods = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.vesting.v1beta1.PeriodicVestingAccount"))
		}
		panic(fmt.Errorf("message cosmos.vesting.v1beta1.PeriodicVestingAccount does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_PeriodicVestingAccount) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.vesting.v1beta1.PeriodicVestingAccount.base_vesting_account":
		if x.BaseVestingAccount == nil {
			x.BaseVestingAccount = new(BaseVestingAccount)
		}
		return protoreflect.ValueOfMessage(x.BaseVestingAccount.ProtoReflect())
	case "cosmos.vesting.v1beta1.PeriodicVestingAccount.vesting_periods":
		if x.VestingPeriods == nil {
			x.VestingPeriods = []*Period{}
		}
		value := &_PeriodicVestingAccount_3_list{list: &x.VestingPeriods}
		return protoreflect.ValueOfList(value)
	case "cosmos.vesting.v1beta1.PeriodicVestingAccount.start_time":
		panic(fmt.Errorf("field start_time of message cosmos.vesting.v1beta1.PeriodicVestingAccount is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.vesting.v1beta1.PeriodicVestingAccount"))
		}
		panic(fmt.Errorf("message cosmos.vesting.v1beta1.PeriodicVestingAccount does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_PeriodicVestingAccount) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.vesting.v1beta1.PeriodicVestingAccount.base_vesting_account":
		m := new(BaseVestingAccount)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "cosmos.vesting.v1beta1.PeriodicVestingAccount.start_time":
		return protoreflect.ValueOfInt64(int64(0))
	case "cosmos.vesting.v1beta1.PeriodicVestingAccount.vesting_periods":
		list := []*Period{}
		return protoreflect.ValueOfList(&_PeriodicVestingAccount_3_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.vesting.v1beta1.PeriodicVestingAccount"))
		}
		panic(fmt.Errorf("message cosmos.vesting.v1beta1.PeriodicVestingAccount does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_PeriodicVestingAccount) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.vesting.v1beta1.PeriodicVestingAccount", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_PeriodicVestingAccount) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_PeriodicVestingAccount) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_PeriodicVestingAccount) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_PeriodicVestingAccount) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*PeriodicVestingAccount)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.BaseVestingAccount != nil {
			l = options.Size(x.BaseVestingAccount)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.StartTime != 0 {
			n += 1 + runtime.Sov(uint64(x.StartTime))
		}
		if len(x.VestingPeriods) > 0 {
			for _, e := range x.VestingPeriods {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*PeriodicVestingAccount)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.VestingPeriods) > 0 {
			for iNdEx := len(x.VestingPeriods) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.VestingPeriods[iNdEx])
				if err != nil {
					return protoiface.MarshalOutput{
						NoUnkeyedLiterals: input.NoUnkeyedLiterals,
						Buf:               input.Buf,
					}, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
				i--
				dAtA[i] = 0x1a
			}
		}
		if x.StartTime != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.StartTime))
			i--
			dAtA[i] = 0x10
		}
		if x.BaseVestingAccount != nil {
			encoded, err := options.Marshal(x.BaseVestingAccount)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*PeriodicVestingAccount)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: PeriodicVestingAccount: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: PeriodicVestingAccount: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field BaseVestingAccount", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.BaseVestingAccount == nil {
					x.BaseVestingAccount = &BaseVestingAccount{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.BaseVestingAccount); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 2:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field StartTime", wireType)
				}
				x.StartTime = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.StartTime |= int64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field VestingPeriods", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.VestingPeriods = append(x.VestingPeriods, &Period{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.VestingPeriods[len(x.VestingPeriods)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_PermanentLockedAccount                      protoreflect.MessageDescriptor
	fd_PermanentLockedAccount_base_vesting_account protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_vesting_v1beta1_vesting_proto_init()
	md_PermanentLockedAccount = File_cosmos_vesting_v1beta1_vesting_proto.Messages().ByName("PermanentLockedAccount")
	fd_PermanentLockedAccount_base_vesting_account = md_PermanentLockedAccount.Fields().ByName("base_vesting_account")
}

var _ protoreflect.Message = (*fastReflection_PermanentLockedAccount)(nil)

type fastReflection_PermanentLockedAccount PermanentLockedAccount

func (x *PermanentLockedAccount) ProtoReflect() protoreflect.Message {
	return (*fastReflection_PermanentLockedAccount)(x)
}

func (x *PermanentLockedAccount) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_vesting_v1beta1_vesting_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_PermanentLockedAccount_messageType fastReflection_PermanentLockedAccount_messageType
var _ protoreflect.MessageType = fastReflection_PermanentLockedAccount_messageType{}

type fastReflection_PermanentLockedAccount_messageType struct{}

func (x fastReflection_PermanentLockedAccount_messageType) Zero() protoreflect.Message {
	return (*fastReflection_PermanentLockedAccount)(nil)
}
func (x fastReflection_PermanentLockedAccount_messageType) New() protoreflect.Message {
	return new(fastReflection_PermanentLockedAccount)
}
func (x fastReflection_PermanentLockedAccount_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_PermanentLockedAccount
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_PermanentLockedAccount) Descriptor() protoreflect.MessageDescriptor {
	return md_PermanentLockedAccount
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_PermanentLockedAccount) Type() protoreflect.MessageType {
	return _fastReflection_PermanentLockedAccount_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_PermanentLockedAccount) New() protoreflect.Message {
	return new(fastReflection_PermanentLockedAccount)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_PermanentLockedAccount) Interface() protoreflect.ProtoMessage {
	return (*PermanentLockedAccount)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_PermanentLockedAccount) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.BaseVestingAccount != nil {
		value := protoreflect.ValueOfMessage(x.BaseVestingAccount.ProtoReflect())
		if !f(fd_PermanentLockedAccount_base_vesting_account, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_PermanentLockedAccount) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.vesting.v1beta1.PermanentLockedAccount.base_vesting_account":
		return x.BaseVestingAccount != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.vesting.v1beta1.PermanentLockedAccount"))
		}
		panic(fmt.Errorf("message cosmos.vesting.v1beta1.PermanentLockedAccount does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_PermanentLockedAccount) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.vesting.v1beta1.PermanentLockedAccount.base_vesting_account":
		x.BaseVestingAccount = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.vesting.v1beta1.PermanentLockedAccount"))
		}
		panic(fmt.Errorf("message cosmos.vesting.v1beta1.PermanentLockedAccount does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_PermanentLockedAccount) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.vesting.v1beta1.PermanentLockedAccount.base_vesting_account":
		value := x.BaseVestingAccount
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.vesting.v1beta1.PermanentLockedAccount"))
		}
		panic(fmt.Errorf("message cosmos.vesting.v1beta1.PermanentLockedAccount does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_PermanentLockedAccount) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.vesting.v1beta1.PermanentLockedAccount.base_vesting_account":
		x.BaseVestingAccount = value.Message().Interface().(*BaseVestingAccount)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.vesting.v1beta1.PermanentLockedAccount"))
		}
		panic(fmt.Errorf("message cosmos.vesting.v1beta1.PermanentLockedAccount does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_PermanentLockedAccount) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.vesting.v1beta1.PermanentLockedAccount.base_vesting_account":
		if x.BaseVestingAccount == nil {
			x.BaseVestingAccount = new(BaseVestingAccount)
		}
		return protoreflect.ValueOfMessage(x.BaseVestingAccount.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.vesting.v1beta1.PermanentLockedAccount"))
		}
		panic(fmt.Errorf("message cosmos.vesting.v1beta1.PermanentLockedAccount does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_PermanentLockedAccount) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.vesting.v1beta1.PermanentLockedAccount.base_vesting_account":
		m := new(BaseVestingAccount)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.vesting.v1beta1.PermanentLockedAccount"))
		}
		panic(fmt.Errorf("message cosmos.vesting.v1beta1.PermanentLockedAccount does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_PermanentLockedAccount) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.vesting.v1beta1.PermanentLockedAccount", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_PermanentLockedAccount) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_PermanentLockedAccount) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_PermanentLockedAccount) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_PermanentLockedAccount) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*PermanentLockedAccount)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.BaseVestingAccount != nil {
			l = options.Size(x.BaseVestingAccount)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*PermanentLockedAccount)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.BaseVestingAccount != nil {
			encoded, err := options.Marshal(x.BaseVestingAccount)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*PermanentLockedAccount)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: PermanentLockedAccount: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: PermanentLockedAccount: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field BaseVestingAccount", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.BaseVestingAccount == nil {
					x.BaseVestingAccount = &BaseVestingAccount{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.BaseVestingAccount); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: cosmos/vesting/v1beta1/vesting.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// BaseVestingAccount implements the VestingAccount interface. It contains all
// the necessary fields needed for any vesting account implementation.
type BaseVestingAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseAccount      *v1beta11.BaseAccount `protobuf:"bytes,1,opt,name=base_account,json=baseAccount,proto3" json:"base_account,omitempty"`
	OriginalVesting  []*v1beta1.Coin       `protobuf:"bytes,2,rep,name=original_vesting,json=originalVesting,proto3" json:"original_vesting,omitempty"`
	DelegatedFree    []*v1beta1.Coin       `protobuf:"bytes,3,rep,name=delegated_free,json=delegatedFree,proto3" json:"delegated_free,omitempty"`
	DelegatedVesting []*v1beta1.Coin       `protobuf:"bytes,4,rep,name=delegated_vesting,json=delegatedVesting,proto3" json:"delegated_vesting,omitempty"`
	EndTime          int64                 `protobuf:"varint,5,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
}

func (x *BaseVestingAccount) Reset() {
	*x = BaseVestingAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_vesting_v1beta1_vesting_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseVestingAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseVestingAccount) ProtoMessage() {}

// Deprecated: Use BaseVestingAccount.ProtoReflect.Descriptor instead.
func (*BaseVestingAccount) Descriptor() ([]byte, []int) {
	return file_cosmos_vesting_v1beta1_vesting_proto_rawDescGZIP(), []int{0}
}

func (x *BaseVestingAccount) GetBaseAccount() *v1beta11.BaseAccount {
	if x != nil {
		return x.BaseAccount
	}
	return nil
}

func (x *BaseVestingAccount) GetOriginalVesting() []*v1beta1.Coin {
	if x != nil {
		return x.OriginalVesting
	}
	return nil
}

func (x *BaseVestingAccount) GetDelegatedFree() []*v1beta1.Coin {
	if x != nil {
		return x.DelegatedFree
	}
	return nil
}

func (x *BaseVestingAccount) GetDelegatedVesting() []*v1beta1.Coin {
	if x != nil {
		return x.DelegatedVesting
	}
	return nil
}

func (x *BaseVestingAccount) GetEndTime() int64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

// ContinuousVestingAccount implements the VestingAccount interface. It
// continuously vests by unlocking coins linearly with respect to time.
type ContinuousVestingAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseVestingAccount *BaseVestingAccount `protobuf:"bytes,1,opt,name=base_vesting_account,json=baseVestingAccount,proto3" json:"base_vesting_account,omitempty"`
	StartTime          int64               `protobuf:"varint,2,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
}

func (x *ContinuousVestingAccount) Reset() {
	*x = ContinuousVestingAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_vesting_v1beta1_vesting_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContinuousVestingAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContinuousVestingAccount) ProtoMessage() {}

// Deprecated: Use ContinuousVestingAccount.ProtoReflect.Descriptor instead.
func (*ContinuousVestingAccount) Descriptor() ([]byte, []int) {
	return file_cosmos_vesting_v1beta1_vesting_proto_rawDescGZIP(), []int{1}
}

func (x *ContinuousVestingAccount) GetBaseVestingAccount() *BaseVestingAccount {
	if x != nil {
		return x.BaseVestingAccount
	}
	return nil
}

func (x *ContinuousVestingAccount) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

// DelayedVestingAccount implements the VestingAccount interface. It vests all
// coins after a specific time, but non prior. In other words, it keeps them
// locked until a specified time.
type DelayedVestingAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseVestingAccount *BaseVestingAccount `protobuf:"bytes,1,opt,name=base_vesting_account,json=baseVestingAccount,proto3" json:"base_vesting_account,omitempty"`
}

func (x *DelayedVestingAccount) Reset() {
	*x = DelayedVestingAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_vesting_v1beta1_vesting_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelayedVestingAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelayedVestingAccount) ProtoMessage() {}

// Deprecated: Use DelayedVestingAccount.ProtoReflect.Descriptor instead.
func (*DelayedVestingAccount) Descriptor() ([]byte, []int) {
	return file_cosmos_vesting_v1beta1_vesting_proto_rawDescGZIP(), []int{2}
}

func (x *DelayedVestingAccount) GetBaseVestingAccount() *BaseVestingAccount {
	if x != nil {
		return x.BaseVestingAccount
	}
	return nil
}

// Period defines a length of time and amount of coins that will vest.
type Period struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Length int64           `protobuf:"varint,1,opt,name=length,proto3" json:"length,omitempty"`
	Amount []*v1beta1.Coin `protobuf:"bytes,2,rep,name=amount,proto3" json:"amount,omitempty"`
}

func (x *Period) Reset() {
	*x = Period{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_vesting_v1beta1_vesting_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Period) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Period) ProtoMessage() {}

// Deprecated: Use Period.ProtoReflect.Descriptor instead.
func (*Period) Descriptor() ([]byte, []int) {
	return file_cosmos_vesting_v1beta1_vesting_proto_rawDescGZIP(), []int{3}
}

func (x *Period) GetLength() int64 {
	if x != nil {
		return x.Length
	}
	return 0
}

func (x *Period) GetAmount() []*v1beta1.Coin {
	if x != nil {
		return x.Amount
	}
	return nil
}

// PeriodicVestingAccount implements the VestingAccount interface. It
// periodically vests by unlocking coins during each specified period.
type PeriodicVestingAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseVestingAccount *BaseVestingAccount `protobuf:"bytes,1,opt,name=base_vesting_account,json=baseVestingAccount,proto3" json:"base_vesting_account,omitempty"`
	StartTime          int64               `protobuf:"varint,2,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	VestingPeriods     []*Period           `protobuf:"bytes,3,rep,name=vesting_periods,json=vestingPeriods,proto3" json:"vesting_periods,omitempty"`
}

func (x *PeriodicVestingAccount) Reset() {
	*x = PeriodicVestingAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_vesting_v1beta1_vesting_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeriodicVestingAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeriodicVestingAccount) ProtoMessage() {}

// Deprecated: Use PeriodicVestingAccount.ProtoReflect.Descriptor instead.
func (*PeriodicVestingAccount) Descriptor() ([]byte, []int) {
	return file_cosmos_vesting_v1beta1_vesting_proto_rawDescGZIP(), []int{4}
}

func (x *PeriodicVestingAccount) GetBaseVestingAccount() *BaseVestingAccount {
	if x != nil {
		return x.BaseVestingAccount
	}
	return nil
}

func (x *PeriodicVestingAccount) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *PeriodicVestingAccount) GetVestingPeriods() []*Period {
	if x != nil {
		return x.VestingPeriods
	}
	return nil
}

// PermanentLockedAccount implements the VestingAccount interface. It does
// not ever release coins, locking them indefinitely. Coins in this account can
// still be used for delegating and for governance votes even while locked.
//
// Since: cosmos-sdk 0.43
type PermanentLockedAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseVestingAccount *BaseVestingAccount `protobuf:"bytes,1,opt,name=base_vesting_account,json=baseVestingAccount,proto3" json:"base_vesting_account,omitempty"`
}

func (x *PermanentLockedAccount) Reset() {
	*x = PermanentLockedAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_vesting_v1beta1_vesting_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PermanentLockedAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PermanentLockedAccount) ProtoMessage() {}

// Deprecated: Use PermanentLockedAccount.ProtoReflect.Descriptor instead.
func (*PermanentLockedAccount) Descriptor() ([]byte, []int) {
	return file_cosmos_vesting_v1beta1_vesting_proto_rawDescGZIP(), []int{5}
}

func (x *PermanentLockedAccount) GetBaseVestingAccount() *BaseVestingAccount {
	if x != nil {
		return x.BaseVestingAccount
	}
	return nil
}

var File_cosmos_vesting_v1beta1_vesting_proto protoreflect.FileDescriptor

var file_cosmos_vesting_v1beta1_vesting_proto_rawDesc = []byte{
	0x0a, 0x24, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x76, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67,
	0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x76, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x76,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x14,
	0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x62, 0x61, 0x73,
	0x65, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x61, 0x75, 0x74,
	0x68, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xea, 0x03, 0x0a, 0x12, 0x42, 0x61, 0x73, 0x65, 0x56, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x49, 0x0a, 0x0c, 0x62,
	0x61, 0x73, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x20, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x42, 0x04, 0xd0, 0xde, 0x1f, 0x01, 0x52, 0x0b, 0x62, 0x61, 0x73, 0x65, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x76, 0x0a, 0x10, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e,
	0x61, 0x6c, 0x5f, 0x76, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x42, 0x30, 0xc8, 0xde, 0x1f,
	0x00, 0xaa, 0xdf, 0x1f, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64,
	0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x73, 0x52, 0x0f, 0x6f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x56, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x72,
	0x0a, 0x0e, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x66, 0x72, 0x65, 0x65,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x69,
	0x6e, 0x42, 0x30, 0xc8, 0xde, 0x1f, 0x00, 0xaa, 0xdf, 0x1f, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x43, 0x6f,
	0x69, 0x6e, 0x73, 0x52, 0x0d, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x46, 0x72,
	0x65, 0x65, 0x12, 0x78, 0x0a, 0x11, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x76, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x42, 0x30, 0xc8, 0xde, 0x1f, 0x00, 0xaa, 0xdf,
	0x1f, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x73, 0x52, 0x10, 0x64, 0x65, 0x6c, 0x65,
	0x67, 0x61, 0x74, 0x65, 0x64, 0x56, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x19, 0x0a, 0x08,
	0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x3a, 0x08, 0x88, 0xa0, 0x1f, 0x00, 0x98, 0xa0, 0x1f,
	0x00, 0x22, 0xa7, 0x01, 0x0a, 0x18, 0x43, 0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x75, 0x6f, 0x75, 0x73,
	0x56, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x62,
	0x0a, 0x14, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x76, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x63,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x76, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x56, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x67, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x04, 0xd0, 0xde, 0x1f, 0x01, 0x52, 0x12,
	0x62, 0x61, 0x73, 0x65, 0x56, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d,
	0x65, 0x3a, 0x08, 0x88, 0xa0, 0x1f, 0x00, 0x98, 0xa0, 0x1f, 0x00, 0x22, 0x85, 0x01, 0x0a, 0x15,
	0x44, 0x65, 0x6c, 0x61, 0x79, 0x65, 0x64, 0x56, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x62, 0x0a, 0x14, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x76, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x76, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x42, 0x61, 0x73,
	0x65, 0x56, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42,
	0x04, 0xd0, 0xde, 0x1f, 0x01, 0x52, 0x12, 0x62, 0x61, 0x73, 0x65, 0x56, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x67, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x3a, 0x08, 0x88, 0xa0, 0x1f, 0x00, 0x98,
	0xa0, 0x1f, 0x00, 0x22, 0x8b, 0x01, 0x0a, 0x06, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x63, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x69,
	0x6e, 0x42, 0x30, 0xc8, 0xde, 0x1f, 0x00, 0xaa, 0xdf, 0x1f, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x43, 0x6f,
	0x69, 0x6e, 0x73, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x3a, 0x04, 0x98, 0xa0, 0x1f,
	0x00, 0x22, 0xf4, 0x01, 0x0a, 0x16, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x69, 0x63, 0x56, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x62, 0x0a, 0x14,
	0x62, 0x61, 0x73, 0x65, 0x5f, 0x76, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x2e, 0x76, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x56, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x04, 0xd0, 0xde, 0x1f, 0x01, 0x52, 0x12, 0x62, 0x61,
	0x73, 0x65, 0x56, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x4d, 0x0a, 0x0f, 0x76, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f,
	0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x2e, 0x76, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x0e,
	0x76, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x73, 0x3a, 0x08,
	0x88, 0xa0, 0x1f, 0x00, 0x98, 0xa0, 0x1f, 0x00, 0x22, 0x86, 0x01, 0x0a, 0x16, 0x50, 0x65, 0x72,
	0x6d, 0x61, 0x6e, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x62, 0x0a, 0x14, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x76, 0x65, 0x73, 0x74,
	0x69, 0x6e, 0x67, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2a, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x76, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x56,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x04, 0xd0,
	0xde, 0x1f, 0x01, 0x52, 0x12, 0x62, 0x61, 0x73, 0x65, 0x56, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x3a, 0x08, 0x88, 0xa0, 0x1f, 0x00, 0x98, 0xa0, 0x1f,
	0x00, 0x42, 0xec, 0x01, 0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2e, 0x76, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x42, 0x0c, 0x56, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x76, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x67, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x76, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x67, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x56, 0x58, 0xaa, 0x02,
	0x16, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x56, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e,
	0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xca, 0x02, 0x16, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x5c, 0x56, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0xe2, 0x02, 0x22, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5c, 0x56, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x67, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x18, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x3a, 0x3a,
	0x56, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cosmos_vesting_v1beta1_vesting_proto_rawDescOnce sync.Once
	file_cosmos_vesting_v1beta1_vesting_proto_rawDescData = file_cosmos_vesting_v1beta1_vesting_proto_rawDesc
)

func file_cosmos_vesting_v1beta1_vesting_proto_rawDescGZIP() []byte {
	file_cosmos_vesting_v1beta1_vesting_proto_rawDescOnce.Do(func() {
		file_cosmos_vesting_v1beta1_vesting_proto_rawDescData = protoimpl.X.CompressGZIP(file_cosmos_vesting_v1beta1_vesting_proto_rawDescData)
	})
	return file_cosmos_vesting_v1beta1_vesting_proto_rawDescData
}

var file_cosmos_vesting_v1beta1_vesting_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_cosmos_vesting_v1beta1_vesting_proto_goTypes = []interface{}{
	(*BaseVestingAccount)(nil),       // 0: cosmos.vesting.v1beta1.BaseVestingAccount
	(*ContinuousVestingAccount)(nil), // 1: cosmos.vesting.v1beta1.ContinuousVestingAccount
	(*DelayedVestingAccount)(nil),    // 2: cosmos.vesting.v1beta1.DelayedVestingAccount
	(*Period)(nil),                   // 3: cosmos.vesting.v1beta1.Period
	(*PeriodicVestingAccount)(nil),   // 4: cosmos.vesting.v1beta1.PeriodicVestingAccount
	(*PermanentLockedAccount)(nil),   // 5: cosmos.vesting.v1beta1.PermanentLockedAccount
	(*v1beta11.BaseAccount)(nil),     // 6: cosmos.auth.v1beta1.BaseAccount
	(*v1beta1.Coin)(nil),             // 7: cosmos.base.v1beta1.Coin
}
var file_cosmos_vesting_v1beta1_vesting_proto_depIdxs = []int32{
	6,  // 0: cosmos.vesting.v1beta1.BaseVestingAccount.base_account:type_name -> cosmos.auth.v1beta1.BaseAccount
	7,  // 1: cosmos.vesting.v1beta1.BaseVestingAccount.original_vesting:type_name -> cosmos.base.v1beta1.Coin
	7,  // 2: cosmos.vesting.v1beta1.BaseVestingAccount.delegated_free:type_name -> cosmos.base.v1beta1.Coin
	7,  // 3: cosmos.vesting.v1beta1.BaseVestingAccount.delegated_vesting:type_name -> cosmos.base.v1beta1.Coin
	0,  // 4: cosmos.vesting.v1beta1.ContinuousVestingAccount.base_vesting_account:type_name -> cosmos.vesting.v1beta1.BaseVestingAccount
	0,  // 5: cosmos.vesting.v1beta1.DelayedVestingAccount.base_vesting_account:type_name -> cosmos.vesting.v1beta1.BaseVestingAccount
	7,  // 6: cosmos.vesting.v1beta1.Period.amount:type_name -> cosmos.base.v1beta1.Coin
	0,  // 7: cosmos.vesting.v1beta1.PeriodicVestingAccount.base_vesting_account:type_name -> cosmos.vesting.v1beta1.BaseVestingAccount
	3,  // 8: cosmos.vesting.v1beta1.PeriodicVestingAccount.vesting_periods:type_name -> cosmos.vesting.v1beta1.Period
	0,  // 9: cosmos.vesting.v1beta1.PermanentLockedAccount.base_vesting_account:type_name -> cosmos.vesting.v1beta1.BaseVestingAccount
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_cosmos_vesting_v1beta1_vesting_proto_init() }
func file_cosmos_vesting_v1beta1_vesting_proto_init() {
	if File_cosmos_vesting_v1beta1_vesting_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cosmos_vesting_v1beta1_vesting_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseVestingAccount); i {
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
		file_cosmos_vesting_v1beta1_vesting_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContinuousVestingAccount); i {
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
		file_cosmos_vesting_v1beta1_vesting_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelayedVestingAccount); i {
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
		file_cosmos_vesting_v1beta1_vesting_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Period); i {
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
		file_cosmos_vesting_v1beta1_vesting_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeriodicVestingAccount); i {
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
		file_cosmos_vesting_v1beta1_vesting_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PermanentLockedAccount); i {
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
			RawDescriptor: file_cosmos_vesting_v1beta1_vesting_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cosmos_vesting_v1beta1_vesting_proto_goTypes,
		DependencyIndexes: file_cosmos_vesting_v1beta1_vesting_proto_depIdxs,
		MessageInfos:      file_cosmos_vesting_v1beta1_vesting_proto_msgTypes,
	}.Build()
	File_cosmos_vesting_v1beta1_vesting_proto = out.File
	file_cosmos_vesting_v1beta1_vesting_proto_rawDesc = nil
	file_cosmos_vesting_v1beta1_vesting_proto_goTypes = nil
	file_cosmos_vesting_v1beta1_vesting_proto_depIdxs = nil
}
