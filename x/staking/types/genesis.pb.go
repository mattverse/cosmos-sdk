// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/staking/v1beta1/genesis.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// GenesisState defines the staking module's genesis state.
type GenesisState struct {
	// params defines all the paramaters of related to deposit.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	// last_total_power tracks the total amounts of bonded tokens recorded during
	// the previous end block.
	LastTotalPower github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=last_total_power,json=lastTotalPower,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"last_total_power" yaml:"last_total_power"`
	// last_validator_powers is a special index that provides a historical list
	// of the last-block's bonded validators.
	LastValidatorPowers []LastValidatorPower `protobuf:"bytes,3,rep,name=last_validator_powers,json=lastValidatorPowers,proto3" json:"last_validator_powers" yaml:"last_validator_powers"`
	// delegations defines the validator set at genesis.
	Validators []Validator `protobuf:"bytes,4,rep,name=validators,proto3" json:"validators"`
	// delegations defines the delegations active at genesis.
	Delegations []Delegation `protobuf:"bytes,5,rep,name=delegations,proto3" json:"delegations"`
	// unbonding_delegations defines the unbonding delegations active at genesis.
	UnbondingDelegations []UnbondingDelegation `protobuf:"bytes,6,rep,name=unbonding_delegations,json=unbondingDelegations,proto3" json:"unbonding_delegations" yaml:"unbonding_delegations"`
	// redelegations defines the redelegations active at genesis.
	Redelegations []Redelegation `protobuf:"bytes,7,rep,name=redelegations,proto3" json:"redelegations"`
	Exported      bool           `protobuf:"varint,8,opt,name=exported,proto3" json:"exported,omitempty"`
	// added newly added queues
	BufferedEditValidators []*MsgEditValidator   `protobuf:"bytes,9,rep,name=buffered_edit_validators,json=bufferedEditValidators,proto3" json:"buffered_edit_validators,omitempty"`
	BufferedDelegations    []*MsgDelegate        `protobuf:"bytes,10,rep,name=buffered_delegations,json=bufferedDelegations,proto3" json:"buffered_delegations,omitempty"`
	BufferedRedelegations  []*MsgBeginRedelegate `protobuf:"bytes,11,rep,name=buffered_redelegations,json=bufferedRedelegations,proto3" json:"buffered_redelegations,omitempty"`
	BufferedUndelegations  []*MsgUndelegate      `protobuf:"bytes,12,rep,name=buffered_undelegations,json=bufferedUndelegations,proto3" json:"buffered_undelegations,omitempty"`
	// epoch number at the time of genesis
	EpochNumber int64 `protobuf:"varint,13,opt,name=epoch_number,json=epochNumber,proto3" json:"epoch_number,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b3dec8894f2831b, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetLastValidatorPowers() []LastValidatorPower {
	if m != nil {
		return m.LastValidatorPowers
	}
	return nil
}

func (m *GenesisState) GetValidators() []Validator {
	if m != nil {
		return m.Validators
	}
	return nil
}

func (m *GenesisState) GetDelegations() []Delegation {
	if m != nil {
		return m.Delegations
	}
	return nil
}

func (m *GenesisState) GetUnbondingDelegations() []UnbondingDelegation {
	if m != nil {
		return m.UnbondingDelegations
	}
	return nil
}

func (m *GenesisState) GetRedelegations() []Redelegation {
	if m != nil {
		return m.Redelegations
	}
	return nil
}

func (m *GenesisState) GetExported() bool {
	if m != nil {
		return m.Exported
	}
	return false
}

func (m *GenesisState) GetBufferedEditValidators() []*MsgEditValidator {
	if m != nil {
		return m.BufferedEditValidators
	}
	return nil
}

func (m *GenesisState) GetBufferedDelegations() []*MsgDelegate {
	if m != nil {
		return m.BufferedDelegations
	}
	return nil
}

func (m *GenesisState) GetBufferedRedelegations() []*MsgBeginRedelegate {
	if m != nil {
		return m.BufferedRedelegations
	}
	return nil
}

func (m *GenesisState) GetBufferedUndelegations() []*MsgUndelegate {
	if m != nil {
		return m.BufferedUndelegations
	}
	return nil
}

func (m *GenesisState) GetEpochNumber() int64 {
	if m != nil {
		return m.EpochNumber
	}
	return 0
}

// LastValidatorPower required for validator set update logic.
type LastValidatorPower struct {
	// address is the address of the validator.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// power defines the power of the validator.
	Power int64 `protobuf:"varint,2,opt,name=power,proto3" json:"power,omitempty"`
}

func (m *LastValidatorPower) Reset()         { *m = LastValidatorPower{} }
func (m *LastValidatorPower) String() string { return proto.CompactTextString(m) }
func (*LastValidatorPower) ProtoMessage()    {}
func (*LastValidatorPower) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b3dec8894f2831b, []int{1}
}
func (m *LastValidatorPower) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LastValidatorPower) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LastValidatorPower.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LastValidatorPower) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LastValidatorPower.Merge(m, src)
}
func (m *LastValidatorPower) XXX_Size() int {
	return m.Size()
}
func (m *LastValidatorPower) XXX_DiscardUnknown() {
	xxx_messageInfo_LastValidatorPower.DiscardUnknown(m)
}

var xxx_messageInfo_LastValidatorPower proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GenesisState)(nil), "cosmos.staking.v1beta1.GenesisState")
	proto.RegisterType((*LastValidatorPower)(nil), "cosmos.staking.v1beta1.LastValidatorPower")
}

func init() {
	proto.RegisterFile("cosmos/staking/v1beta1/genesis.proto", fileDescriptor_9b3dec8894f2831b)
}

var fileDescriptor_9b3dec8894f2831b = []byte{
	// 626 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x94, 0xcb, 0x4e, 0xdb, 0x40,
	0x14, 0x86, 0xe3, 0x86, 0x4b, 0x98, 0x84, 0xaa, 0x1a, 0x02, 0xb5, 0x50, 0x65, 0x07, 0x97, 0x56,
	0x51, 0x2f, 0xb6, 0xa0, 0x3b, 0xd4, 0x55, 0xd4, 0x16, 0x51, 0x15, 0x84, 0xa6, 0x85, 0x45, 0x55,
	0xc9, 0x1a, 0xe3, 0xc1, 0x58, 0x38, 0x1e, 0xcb, 0x33, 0xa6, 0xb0, 0xaf, 0xaa, 0x2e, 0xfb, 0x08,
	0x3c, 0x0e, 0x4b, 0x96, 0x55, 0x17, 0x51, 0x05, 0x9b, 0xae, 0x51, 0x1f, 0xa0, 0xf2, 0xf8, 0xc2,
	0x24, 0xc1, 0x59, 0x25, 0x73, 0xf4, 0xff, 0xdf, 0x7f, 0xe6, 0x4c, 0x72, 0xc0, 0xea, 0x01, 0x65,
	0x7d, 0xca, 0x2c, 0xc6, 0xf1, 0xb1, 0x1f, 0x7a, 0xd6, 0xc9, 0x9a, 0x43, 0x38, 0x5e, 0xb3, 0x3c,
	0x12, 0x12, 0xe6, 0x33, 0x33, 0x8a, 0x29, 0xa7, 0x70, 0x29, 0x53, 0x99, 0xb9, 0xca, 0xcc, 0x55,
	0xcb, 0x6d, 0x8f, 0x7a, 0x54, 0x48, 0xac, 0xf4, 0x5b, 0xa6, 0x5e, 0xae, 0x62, 0x16, 0xee, 0x4c,
	0xa5, 0x57, 0xa8, 0xf8, 0x69, 0x26, 0x30, 0xfe, 0x35, 0x40, 0x6b, 0x33, 0x6b, 0xe3, 0x23, 0xc7,
	0x9c, 0xc0, 0xd7, 0x60, 0x26, 0xc2, 0x31, 0xee, 0x33, 0x55, 0xe9, 0x28, 0xdd, 0xe6, 0xba, 0x66,
	0xde, 0xdd, 0x96, 0xb9, 0x2b, 0x54, 0xbd, 0xa9, 0x8b, 0x81, 0x5e, 0x43, 0xb9, 0x07, 0x32, 0xf0,
	0x20, 0xc0, 0x8c, 0xdb, 0x9c, 0x72, 0x1c, 0xd8, 0x11, 0xfd, 0x4a, 0x62, 0xf5, 0x5e, 0x47, 0xe9,
	0xb6, 0x7a, 0x5b, 0xa9, 0xee, 0xf7, 0x40, 0x7f, 0xea, 0xf9, 0xfc, 0x28, 0x71, 0xcc, 0x03, 0xda,
	0xb7, 0xf2, 0xe6, 0xb2, 0x8f, 0x97, 0xcc, 0x3d, 0xb6, 0xf8, 0x59, 0x44, 0x98, 0xb9, 0x15, 0xf2,
	0x9b, 0x81, 0xfe, 0xf0, 0x0c, 0xf7, 0x83, 0x0d, 0x63, 0x94, 0x67, 0xa0, 0xfb, 0x69, 0xe9, 0x53,
	0x5a, 0xd9, 0x4d, 0x0b, 0xf0, 0x9b, 0x02, 0x16, 0x85, 0xea, 0x04, 0x07, 0xbe, 0x8b, 0x39, 0x8d,
	0x33, 0x25, 0x53, 0xeb, 0x9d, 0x7a, 0xb7, 0xb9, 0xfe, 0xac, 0xea, 0x0a, 0x1f, 0x30, 0xe3, 0xfb,
	0x85, 0x47, 0xb0, 0x7a, 0xab, 0x69, 0x9b, 0x37, 0x03, 0xfd, 0x91, 0x14, 0x3e, 0x8a, 0x35, 0xd0,
	0x42, 0x30, 0xe6, 0x64, 0x70, 0x13, 0x80, 0x52, 0xc9, 0xd4, 0x29, 0x11, 0xbd, 0x52, 0x15, 0x5d,
	0x9a, 0xf3, 0x01, 0x4a, 0x56, 0xf8, 0x1e, 0x34, 0x5d, 0x12, 0x10, 0x0f, 0x73, 0x9f, 0x86, 0x4c,
	0x9d, 0x16, 0x24, 0xa3, 0x8a, 0xf4, 0xa6, 0x94, 0xe6, 0x28, 0xd9, 0x0c, 0xbf, 0x2b, 0x60, 0x31,
	0x09, 0x1d, 0x1a, 0xba, 0x7e, 0xe8, 0xd9, 0x32, 0x76, 0x46, 0x60, 0x9f, 0x57, 0x61, 0xf7, 0x0a,
	0x93, 0xc4, 0x1f, 0x19, 0xce, 0x9d, 0x5c, 0x03, 0xb5, 0x93, 0x71, 0x2b, 0x83, 0xbb, 0x60, 0x3e,
	0x26, 0x72, 0xfe, 0xac, 0xc8, 0x5f, 0xad, 0xca, 0x47, 0x92, 0x38, 0xbf, 0xd8, 0x30, 0x00, 0x2e,
	0x83, 0x06, 0x39, 0x8d, 0x68, 0xcc, 0x89, 0xab, 0x36, 0x3a, 0x4a, 0xb7, 0x81, 0xca, 0x33, 0x74,
	0x80, 0xea, 0x24, 0x87, 0x87, 0x24, 0x26, 0xae, 0x4d, 0x5c, 0x5f, 0x7a, 0x43, 0xa6, 0xce, 0x89,
	0xe0, 0x6e, 0x55, 0xf0, 0x36, 0xf3, 0xde, 0xba, 0xfe, 0xed, 0xeb, 0xa2, 0xa5, 0x82, 0x34, 0x54,
	0x66, 0x70, 0x1f, 0xb4, 0xcb, 0x0c, 0xf9, 0x62, 0x40, 0xf0, 0x1f, 0x4f, 0xe0, 0xe7, 0x73, 0x21,
	0x68, 0xa1, 0x00, 0xc8, 0x93, 0xc2, 0xa0, 0x4c, 0xb4, 0x87, 0x47, 0xd6, 0x9c, 0xfc, 0x73, 0xde,
	0x66, 0x5e, 0x8f, 0x78, 0x7e, 0x58, 0x8e, 0x8e, 0xa0, 0xc5, 0x82, 0x84, 0x86, 0x46, 0xf7, 0x45,
	0x8a, 0x48, 0x42, 0x39, 0xa2, 0x25, 0x22, 0x9e, 0x4c, 0x88, 0xd8, 0x0b, 0xc7, 0xe9, 0x7b, 0x32,
	0x03, 0xae, 0x80, 0x16, 0x89, 0xe8, 0xc1, 0x91, 0x1d, 0x26, 0x7d, 0x87, 0xc4, 0xea, 0x7c, 0x47,
	0xe9, 0xd6, 0x51, 0x53, 0xd4, 0x76, 0x44, 0xc9, 0xd8, 0x01, 0x70, 0xfc, 0xcf, 0x07, 0x55, 0x30,
	0x8b, 0x5d, 0x37, 0x26, 0x2c, 0x5b, 0x3e, 0x73, 0xa8, 0x38, 0xc2, 0x36, 0x98, 0xbe, 0x5d, 0x26,
	0x75, 0x94, 0x1d, 0x36, 0x1a, 0x3f, 0xce, 0xf5, 0xda, 0xdf, 0x73, 0xbd, 0xd6, 0x7b, 0x77, 0x71,
	0xa5, 0x29, 0x97, 0x57, 0x9a, 0xf2, 0xe7, 0x4a, 0x53, 0x7e, 0x5e, 0x6b, 0xb5, 0xcb, 0x6b, 0xad,
	0xf6, 0xeb, 0x5a, 0xab, 0x7d, 0x7e, 0x31, 0x71, 0xdf, 0x9c, 0x96, 0x9b, 0x51, 0x6c, 0x1e, 0x67,
	0x46, 0x6c, 0xc5, 0x57, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff, 0x23, 0xd9, 0xab, 0x07, 0xb2, 0x05,
	0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.EpochNumber != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.EpochNumber))
		i--
		dAtA[i] = 0x68
	}
	if len(m.BufferedUndelegations) > 0 {
		for iNdEx := len(m.BufferedUndelegations) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.BufferedUndelegations[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x62
		}
	}
	if len(m.BufferedRedelegations) > 0 {
		for iNdEx := len(m.BufferedRedelegations) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.BufferedRedelegations[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x5a
		}
	}
	if len(m.BufferedDelegations) > 0 {
		for iNdEx := len(m.BufferedDelegations) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.BufferedDelegations[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x52
		}
	}
	if len(m.BufferedEditValidators) > 0 {
		for iNdEx := len(m.BufferedEditValidators) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.BufferedEditValidators[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x4a
		}
	}
	if m.Exported {
		i--
		if m.Exported {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x40
	}
	if len(m.Redelegations) > 0 {
		for iNdEx := len(m.Redelegations) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Redelegations[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.UnbondingDelegations) > 0 {
		for iNdEx := len(m.UnbondingDelegations) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.UnbondingDelegations[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.Delegations) > 0 {
		for iNdEx := len(m.Delegations) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Delegations[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.Validators) > 0 {
		for iNdEx := len(m.Validators) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Validators[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.LastValidatorPowers) > 0 {
		for iNdEx := len(m.LastValidatorPowers) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.LastValidatorPowers[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	{
		size := m.LastTotalPower.Size()
		i -= size
		if _, err := m.LastTotalPower.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *LastValidatorPower) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LastValidatorPower) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LastValidatorPower) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Power != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.Power))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.LastTotalPower.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.LastValidatorPowers) > 0 {
		for _, e := range m.LastValidatorPowers {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Validators) > 0 {
		for _, e := range m.Validators {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Delegations) > 0 {
		for _, e := range m.Delegations {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.UnbondingDelegations) > 0 {
		for _, e := range m.UnbondingDelegations {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Redelegations) > 0 {
		for _, e := range m.Redelegations {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.Exported {
		n += 2
	}
	if len(m.BufferedEditValidators) > 0 {
		for _, e := range m.BufferedEditValidators {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.BufferedDelegations) > 0 {
		for _, e := range m.BufferedDelegations {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.BufferedRedelegations) > 0 {
		for _, e := range m.BufferedRedelegations {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.BufferedUndelegations) > 0 {
		for _, e := range m.BufferedUndelegations {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.EpochNumber != 0 {
		n += 1 + sovGenesis(uint64(m.EpochNumber))
	}
	return n
}

func (m *LastValidatorPower) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.Power != 0 {
		n += 1 + sovGenesis(uint64(m.Power))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastTotalPower", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LastTotalPower.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastValidatorPowers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LastValidatorPowers = append(m.LastValidatorPowers, LastValidatorPower{})
			if err := m.LastValidatorPowers[len(m.LastValidatorPowers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validators", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Validators = append(m.Validators, Validator{})
			if err := m.Validators[len(m.Validators)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Delegations", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Delegations = append(m.Delegations, Delegation{})
			if err := m.Delegations[len(m.Delegations)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnbondingDelegations", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UnbondingDelegations = append(m.UnbondingDelegations, UnbondingDelegation{})
			if err := m.UnbondingDelegations[len(m.UnbondingDelegations)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Redelegations", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Redelegations = append(m.Redelegations, Redelegation{})
			if err := m.Redelegations[len(m.Redelegations)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Exported", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Exported = bool(v != 0)
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BufferedEditValidators", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BufferedEditValidators = append(m.BufferedEditValidators, &MsgEditValidator{})
			if err := m.BufferedEditValidators[len(m.BufferedEditValidators)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BufferedDelegations", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BufferedDelegations = append(m.BufferedDelegations, &MsgDelegate{})
			if err := m.BufferedDelegations[len(m.BufferedDelegations)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BufferedRedelegations", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BufferedRedelegations = append(m.BufferedRedelegations, &MsgBeginRedelegate{})
			if err := m.BufferedRedelegations[len(m.BufferedRedelegations)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BufferedUndelegations", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BufferedUndelegations = append(m.BufferedUndelegations, &MsgUndelegate{})
			if err := m.BufferedUndelegations[len(m.BufferedUndelegations)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochNumber", wireType)
			}
			m.EpochNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EpochNumber |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *LastValidatorPower) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
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
			return fmt.Errorf("proto: LastValidatorPower: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LastValidatorPower: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Power", wireType)
			}
			m.Power = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Power |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
