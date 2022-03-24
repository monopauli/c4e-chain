package keeper_test

import (
	"testing"

	testkeeper "github.com/chain4energy/c4e-chain/testutil/keeper"
	testutils "github.com/chain4energy/c4e-chain/testutil/module/cfevesting"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	commontestutils "github.com/chain4energy/c4e-chain/testutil/common"

)

func TestVesting(t *testing.T) {
	height := int64(0)
	keeper, ctx := testkeeper.CfevestingKeeperWithBlockHeight(t, height)
	wctx := sdk.WrapSDKContext(ctx)

	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)
	addr := acountsAddresses[0].String()

	accountVestings := testutils.GenerateOneAccountVestingsWithAddressWith10BasedVestings(1, 1, 1)
	accountVestings.Address = addr
	accountVestings.DelegableAddress = ""

	keeper.SetAccountVestings(ctx, accountVestings)

	response, err := keeper.Vesting(wctx, &types.QueryVestingRequest{Address: addr})
	require.NoError(t, err)
	verifyVestingResponse(t, response, accountVestings, height, true)
}

func TestVestingWithDelegableAddress(t *testing.T) {
	height := int64(0)
	keeper, ctx := testkeeper.CfevestingKeeperWithBlockHeight(t, height)
	wctx := sdk.WrapSDKContext(ctx)
	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)
	addr := acountsAddresses[0].String()

	accountVestings := testutils.GenerateOneAccountVestingsWithAddressWith10BasedVestings(1, 1, 1)
	accountVestings.Address = addr

	keeper.SetAccountVestings(ctx, accountVestings)

	response, err := keeper.Vesting(wctx, &types.QueryVestingRequest{Address: addr})
	require.NoError(t, err)
	verifyVestingResponse(t, response, accountVestings, height, true)

}

func TestVestingSomeToWithdraw(t *testing.T) {
	height := int64(10100)
	keeper, ctx := testkeeper.CfevestingKeeperWithBlockHeight(t, height)
	wctx := sdk.WrapSDKContext(ctx)
	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)
	addr := acountsAddresses[0].String()

	accountVestings := testutils.GenerateOneAccountVestingsWithAddressWith10BasedVestings(1, 1, 1)
	accountVestings.Address = addr

	keeper.SetAccountVestings(ctx, accountVestings)

	response, err := keeper.Vesting(wctx, &types.QueryVestingRequest{Address: addr})
	require.NoError(t, err)

	verifyVestingResponse(t, response, accountVestings, height, true)

}

func TestVestingSomeToWithdrawAndSomeWithdrawn(t *testing.T) {
	height := int64(10100)
	keeper, ctx := testkeeper.CfevestingKeeperWithBlockHeight(t, height)
	wctx := sdk.WrapSDKContext(ctx)
	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)
	addr := acountsAddresses[0].String()

	accountVestings := testutils.GenerateOneAccountVestingsWithAddressWith10BasedVestings(1, 1, 1)
	accountVestings.Address = addr
	accountVestings.Vestings[0].Withdrawn = sdk.NewInt(500)
	accountVestings.Vestings[0].LastModificationWithdrawn = sdk.NewInt(500)

	keeper.SetAccountVestings(ctx, accountVestings)

	response, err := keeper.Vesting(wctx, &types.QueryVestingRequest{Address: addr})
	require.NoError(t, err)
	verifyVestingResponse(t, response, accountVestings, height, true)

}

func TestVestingSentAfterLockEndReceivingSide(t *testing.T) {
	height := int64(10100)
	keeper, ctx := testkeeper.CfevestingKeeperWithBlockHeight(t, height)
	wctx := sdk.WrapSDKContext(ctx)
	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)
	addr := acountsAddresses[0].String()

	accountVestings := testutils.GenerateOneAccountVestingsWithAddressWith10BasedVestings(1, 1, 1)
	accountVestings.Address = addr
	accountVestings.Vestings[0].VestingStartBlock = accountVestings.Vestings[0].LockEndBlock
	accountVestings.Vestings[0].LastModificationBlock = accountVestings.Vestings[0].LockEndBlock

	accountVestings.Vestings[0].LockEndBlock -= 100

	keeper.SetAccountVestings(ctx, accountVestings)

	response, err := keeper.Vesting(wctx, &types.QueryVestingRequest{Address: addr})
	require.NoError(t, err)

	verifyVestingResponse(t, response, accountVestings, height, true)

}

func TestVestingSentAfterLockEndSendingSide(t *testing.T) {
	height := int64(10100)
	keeper, ctx := testkeeper.CfevestingKeeperWithBlockHeight(t, height)
	wctx := sdk.WrapSDKContext(ctx)
	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)
	addr := acountsAddresses[0].String()

	accountVestings := testutils.GenerateOneAccountVestingsWithAddressWith10BasedVestings(1, 1, 1)
	accountVestings.Address = addr

	accountVestings.Vestings[0].LastModificationBlock = accountVestings.Vestings[0].LockEndBlock
	accountVestings.Vestings[0].Sent = sdk.NewInt(100000)
	accountVestings.Vestings[0].LastModificationVested = accountVestings.Vestings[0].LastModificationVested.Sub(sdk.NewInt(100000))

	accountVestings.Vestings[0].LockEndBlock -= 100

	keeper.SetAccountVestings(ctx, accountVestings)

	response, err := keeper.Vesting(wctx, &types.QueryVestingRequest{Address: addr})
	require.NoError(t, err)

	verifyVestingResponse(t, response, accountVestings, height, true)

}

func TestVestingSentAfterLockEndSendingSideAndWithdrawn(t *testing.T) {
	height := int64(10100)
	keeper, ctx := testkeeper.CfevestingKeeperWithBlockHeight(t, height)
	wctx := sdk.WrapSDKContext(ctx)
	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)
	addr := acountsAddresses[0].String()

	accountVestings := testutils.GenerateOneAccountVestingsWithAddressWith10BasedVestings(1, 1, 1)
	accountVestings.Address = addr

	accountVestings.Vestings[0].LastModificationBlock = accountVestings.Vestings[0].LockEndBlock
	accountVestings.Vestings[0].Sent = sdk.NewInt(100000)
	accountVestings.Vestings[0].LastModificationVested = accountVestings.Vestings[0].LastModificationVested.Sub(sdk.NewInt(100000))
	accountVestings.Vestings[0].LastModificationWithdrawn = sdk.NewInt(400)

	accountVestings.Vestings[0].LockEndBlock -= 100

	keeper.SetAccountVestings(ctx, accountVestings)

	response, err := keeper.Vesting(wctx, &types.QueryVestingRequest{Address: addr})
	require.NoError(t, err)

	verifyVestingResponse(t, response, accountVestings, height, true)

}

func TestVestingManyVestings(t *testing.T) {
	height := int64(0)
	keeper, ctx := testkeeper.CfevestingKeeperWithBlockHeight(t, height)
	wctx := sdk.WrapSDKContext(ctx)
	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)
	addr := acountsAddresses[0].String()

	accountVestings := testutils.GenerateOneAccountVestingsWithAddressWith10BasedVestings(3, 1, 1)
	accountVestings.Address = addr

	keeper.SetAccountVestings(ctx, accountVestings)

	response, err := keeper.Vesting(wctx, &types.QueryVestingRequest{Address: addr})
	require.NoError(t, err)

	verifyVestingResponse(t, response, accountVestings, height, true)

}


