package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateVestingPool{}, "cfevesting/CreateVestingPool", nil)
	cdc.RegisterConcrete(&MsgWithdrawAllAvailable{}, "cfevesting/WithdrawAllAvailable", nil)
	cdc.RegisterConcrete(&MsgCreateVestingAccount{}, "cfevesting/CreateVestingAccount", nil)
	cdc.RegisterConcrete(&MsgSendToVestingAccount{}, "cfevesting/SendToVestingAccount", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateVestingPool{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgWithdrawAllAvailable{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateVestingAccount{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSendToVestingAccount{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(Amino)
)

func init() {
	RegisterCodec(Amino)
	Amino.Seal()
}
