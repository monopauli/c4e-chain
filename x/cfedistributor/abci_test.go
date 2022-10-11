package cfedistributor_test

import (
	"testing"

	testapp "github.com/chain4energy/c4e-chain/testutil/app"
	// "github.com/chain4energy/c4e-chain/app"

	"github.com/chain4energy/c4e-chain/x/cfedistributor/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	// tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	// banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	commontestutils "github.com/chain4energy/c4e-chain/testutil/common"

)

type DestinationType int64

const (
	MainCollector DestinationType = iota
	ModuleAccount
	InternalAccount
	BaseAccount
)

const c4eDistributorCollectorName = types.GreenEnergyBoosterCollector
const noValidatorsCollectorName = types.GovernanceBoosterCollector

// const c4eDistributorCollectorName = "c4e_distributor"
// const noValidatorsCollectorName = "no_validators"

func prepareBurningDistributor(destinationType DestinationType) types.SubDistributor {
	var address string
	if destinationType == BaseAccount {
		address = "cosmos13zg4u07ymq83uq73t2cq3dj54jj37zzgr3hlck"
	} else {
		address = c4eDistributorCollectorName
	}

	var destAccount = types.Account{}
	destAccount.Id = address

	if destinationType == ModuleAccount {
		destAccount.Type = "MODULE_ACCOUNT"
	} else if destinationType == InternalAccount {
		destAccount.Type = "INTERNAL_ACCOUNT"
	} else {
		destAccount.Type = "BASE_ACCOUNT"
	}

	if destinationType == MainCollector {
		destAccount.Type = "MAIN"
	}

	burnShare := types.BurnShare{
		Percent: sdk.MustNewDecFromStr("51"),
	}

	destination := types.Destination{
		Account:   destAccount,
		Share:     nil,
		BurnShare: &burnShare,
	}

	distributor1 := types.SubDistributor{
		Name:        "tx_fee_distributor",
		Sources:     []*types.Account{{Id: authtypes.FeeCollectorName, Type: "MODULE_ACCOUNT"}},
		Destination: destination,
	}

	return distributor1
}

func prepareInflationToPassAcoutSubDistr(passThroughAccoutType DestinationType) types.SubDistributor {
	source := types.Account{
		Id:   "c4e",
		Type: "MAIN",
	}

	var address string
	if passThroughAccoutType == BaseAccount {
		address = "cosmos13zg4u07ymq83uq73t2cq3dj54jj37zzgr3hlck"
	} else {
		address = c4eDistributorCollectorName
	}

	var destAccount = types.Account{
		Id: address,
	}

	if passThroughAccoutType == ModuleAccount {
		destAccount.Type = "MODULE_ACCOUNT"
	} else if passThroughAccoutType == InternalAccount {
		destAccount.Type = "INTERNAL_ACCOUNT"
	} else {
		destAccount.Type = "BASE_ACCOUNT"
	}

	if passThroughAccoutType == MainCollector {
		destAccount.Type = "MAIN"
	}

	burnShare := types.BurnShare{
		Percent: sdk.MustNewDecFromStr("0"),
	}

	destination := types.Destination{
		Account:   destAccount,
		Share:     nil,
		BurnShare: &burnShare,
	}
	return types.SubDistributor{
		Name:        "pass_distributor",
		Sources:     []*types.Account{&source},
		Destination: destination,
	}
}

func prepareInflationSubDistributor(sourceAccoutType DestinationType, toValidators bool) types.SubDistributor {

	var address string
	if sourceAccoutType == BaseAccount {
		address = "cosmos13zg4u07ymq83uq73t2cq3dj54jj37zzgr3hlck"
	} else {
		address = c4eDistributorCollectorName
	}

	var source = types.Account{
		Id: address,
	}

	if sourceAccoutType == ModuleAccount {
		source.Type = "MODULE_ACCOUNT"
	} else if sourceAccoutType == InternalAccount {
		source.Type = "INTERNAL_ACCOUNT"
	} else {
		source.Type = "BASE_ACCOUNT"
	}

	if sourceAccoutType == MainCollector {
		source.Type = "MAIN"
	}

	// source := types.Account{IsMainCollector: true, IsModuleAccount: false, IsInternalAccount: false}

	burnShare := types.BurnShare{
		Percent: sdk.MustNewDecFromStr("0"),
	}

	var destName string
	if toValidators {
		destName = types.ValidatorsRewardsCollector
	} else {
		destName = noValidatorsCollectorName
	}

	var destAccount = types.Account{
		Id:   destName,
		Type: "MODULE_ACCOUNT",
	}

	var shareDevelopmentFundAccount = types.Account{
		Id:   "cosmos1p20lmfzp4g9vywl2jxwexwh6akvkxzpa6hdrag",
		Type: "BASE_ACCOUNT",
	}

	shareDevelopmentFund := types.Share{
		Name:    "development_fund",
		Percent: sdk.MustNewDecFromStr("10.345"),
		Account: shareDevelopmentFundAccount,
	}

	destination := types.Destination{
		Account:   destAccount,
		Share:     []*types.Share{&shareDevelopmentFund},
		BurnShare: &burnShare,
	}

	return types.SubDistributor{
		Name:        "tx_fee_distributor",
		Sources:     []*types.Account{&source},
		Destination: destination,
	}
}

func TestBurningDistributorMainCollectorDes(t *testing.T) {
	BurningDistributorTest(t, MainCollector)
}

func TestBurningDistributorModuleAccountDest(t *testing.T) {
	BurningDistributorTest(t, ModuleAccount)
}

func TestBurningDistributorInternalAccountDest(t *testing.T) {
	BurningDistributorTest(t, InternalAccount)
}

func TestBurningDistributorBaseAccountDest(t *testing.T) {
	BurningDistributorTest(t, BaseAccount)
}

func BurningDistributorTest(t *testing.T, destinationType DestinationType) {
	senderCoin := sdk.NewInt(1017)
	
	testHelper, ctx := testapp.SetupTestApp(t)

	testHelper.BankUtils.AddDefaultDenomCoinsToModule(ctx, senderCoin, authtypes.FeeCollectorName)
	
	require.EqualValues(t, testHelper.InitialValidatorsCoin.AddAmount(senderCoin), testHelper.App.BankKeeper.GetSupply(ctx, commontestutils.DefaultTestDenom))
	var subdistributors []types.SubDistributor
	subdistributors = append(subdistributors, prepareBurningDistributor(destinationType))

	testHelper.App.CfedistributorKeeper.SetParams(ctx, types.NewParams(subdistributors))
	ctx = ctx.WithBlockHeight(int64(2))
	testHelper.App.BeginBlocker(ctx, abci.RequestBeginBlock{})

	//coin on "burnState" should be equal 498, remains: 1 and 0.33 on remains
	burnState, _ := testHelper.App.CfedistributorKeeper.GetState(ctx, "burn_state_key")
	ctx.Logger().Error(burnState.String())
	//burnState, _ := app.CfedistributorKeeper.GetAllStates()
	coinRemains := burnState.CoinsStates
	require.EqualValues(t, sdk.MustNewDecFromStr("0.67"), coinRemains.AmountOf(commontestutils.DefaultTestDenom))

	if destinationType == MainCollector {
		mainCollectorCoins :=
		testHelper.App.CfedistributorKeeper.GetAccountCoinsForModuleAccount(ctx, types.DistributorMainAccount)
		require.EqualValues(t, 1, len(testHelper.App.CfedistributorKeeper.GetAllStates(ctx)))
		require.EqualValues(t, sdk.NewInt(499), mainCollectorCoins.AmountOf(commontestutils.DefaultTestDenom))
	} else if destinationType == ModuleAccount {
		mainCollectorCoins :=
		testHelper.App.CfedistributorKeeper.GetAccountCoinsForModuleAccount(ctx, types.DistributorMainAccount)
		c4eModulAccountCoins :=
		testHelper.App.CfedistributorKeeper.GetAccountCoinsForModuleAccount(ctx, c4eDistributorCollectorName)
		require.EqualValues(t, 2, len(testHelper.App.CfedistributorKeeper.GetAllStates(ctx)))
		require.EqualValues(t, sdk.NewInt(498), c4eModulAccountCoins.AmountOf(commontestutils.DefaultTestDenom))
		require.EqualValues(t, sdk.NewInt(1), mainCollectorCoins.AmountOf(commontestutils.DefaultTestDenom))
		c4eDistrState, _ := testHelper.App.CfedistributorKeeper.GetState(ctx, c4eDistributorCollectorName)
		coinRemains := c4eDistrState.CoinsStates
		require.EqualValues(t, sdk.MustNewDecFromStr("0.33"), coinRemains.AmountOf(commontestutils.DefaultTestDenom))

	} else if destinationType == InternalAccount {
		mainCollectorCoins :=
		testHelper.App.CfedistributorKeeper.GetAccountCoinsForModuleAccount(ctx, types.DistributorMainAccount)
		require.EqualValues(t, 2, len(testHelper.App.CfedistributorKeeper.GetAllStates(ctx)))
		require.EqualValues(t, sdk.NewInt(499), mainCollectorCoins.AmountOf(commontestutils.DefaultTestDenom))
		c4eDistrState, _ := testHelper.App.CfedistributorKeeper.GetState(ctx, c4eDistributorCollectorName)
		coinRemains := c4eDistrState.CoinsStates
		require.EqualValues(t, sdk.MustNewDecFromStr("498.33"), coinRemains.AmountOf(commontestutils.DefaultTestDenom))
	} else {
		address, _ := sdk.AccAddressFromBech32("cosmos13zg4u07ymq83uq73t2cq3dj54jj37zzgr3hlck")
		mainCollectorCoins :=
		testHelper.App.CfedistributorKeeper.GetAccountCoinsForModuleAccount(ctx, types.DistributorMainAccount)

		accountCoins :=
		testHelper.App.CfedistributorKeeper.GetAccountCoins(ctx, address)

		require.EqualValues(t, 2, len(testHelper.App.CfedistributorKeeper.GetAllStates(ctx)))
		ctx.Logger().Error(accountCoins.AmountOf(commontestutils.DefaultTestDenom).String())
		println("Amount: " + accountCoins.AmountOf(commontestutils.DefaultTestDenom).String())

		require.EqualValues(t, sdk.NewInt(498), accountCoins.AmountOf(commontestutils.DefaultTestDenom))
		require.EqualValues(t, sdk.NewInt(1), mainCollectorCoins.AmountOf(commontestutils.DefaultTestDenom))

		c4eDistrState, _ := testHelper.App.CfedistributorKeeper.GetState(ctx, "cosmos13zg4u07ymq83uq73t2cq3dj54jj37zzgr3hlck")
		coinRemains := c4eDistrState.CoinsStates
		require.EqualValues(t, sdk.MustNewDecFromStr("0.33"), coinRemains.AmountOf("uc4e"))
	}
	require.EqualValues(t, sdk.NewCoin(commontestutils.DefaultTestDenom, sdk.NewInt(499)).Add(testHelper.InitialValidatorsCoin), testHelper.App.BankKeeper.GetSupply(ctx, commontestutils.DefaultTestDenom))
}

func TestBurningWithInflationDistributorPassThroughMainCollector(t *testing.T) {
	BurningWithInflationDistributorTest(t, MainCollector, true)
}

func TestBurningWithInflationDistributorPassThroughModuleAccount(t *testing.T) {
	BurningWithInflationDistributorTest(t, ModuleAccount, true)
}

func TestBurningWithInflationDistributorPassInternalAccountAccount(t *testing.T) {
	BurningWithInflationDistributorTest(t, InternalAccount, true)
}

func TestBurningWithInflationDistributorPassBaseAccountAccount(t *testing.T) {
	BurningWithInflationDistributorTest(t, BaseAccount, true)
}

func TestBurningWithInflationDistributorPassThroughMainCollectorNoValidators(t *testing.T) {
	BurningWithInflationDistributorTest(t, MainCollector, false)
}

func TestBurningWithInflationDistributorPassThroughModuleAccountNoValidators(t *testing.T) {
	BurningWithInflationDistributorTest(t, ModuleAccount, false)
}

func TestBurningWithInflationDistributorPassInternalAccountAccountNoValidators(t *testing.T) {
	BurningWithInflationDistributorTest(t, InternalAccount, false)
}

func TestBurningWithInflationDistributorPassBaseAccountAccountNoValidators(t *testing.T) {
	BurningWithInflationDistributorTest(t, BaseAccount, false)
}

func BurningWithInflationDistributorTest(t *testing.T, passThroughAccoutType DestinationType, toValidators bool) {

	testHelper, ctx := testapp.SetupTestApp(t)
	

	//prepare module account with coin to distribute fee_collector 1017
	cointToMint := sdk.NewInt(1017)

	testHelper.BankUtils.AddDefaultDenomCoinsToModule(ctx, cointToMint, authtypes.FeeCollectorName)

	cointToMintFromInflation := sdk.NewInt(5044)
	testHelper.BankUtils.AddDefaultDenomCoinsToModule(ctx, cointToMintFromInflation, types.DistributorMainAccount)

	initialCoinAmount := testHelper.InitialValidatorsCoin.AddAmount(cointToMint).AddAmount(cointToMintFromInflation)
	require.EqualValues(t, initialCoinAmount, testHelper.App.BankKeeper.GetSupply(ctx, commontestutils.DefaultTestDenom))

	// require.EqualValues(t, sdk.NewCoin(denom, sdk.NewInt(1017+5044)).Add(valCoin), app.BankKeeper.GetSupply(ctx, denom))

	var subDistributors []types.SubDistributor

	subDistributors = append(subDistributors, prepareBurningDistributor(MainCollector))
	if passThroughAccoutType != MainCollector {
		subDistributors = append(subDistributors, prepareInflationToPassAcoutSubDistr(passThroughAccoutType))
	}
	subDistributors = append(subDistributors, prepareInflationSubDistributor(passThroughAccoutType, toValidators))

	testHelper.App.CfedistributorKeeper.SetParams(ctx, types.NewParams(subDistributors))
	ctx = ctx.WithBlockHeight(int64(2))
	testHelper.App.BeginBlocker(ctx, abci.RequestBeginBlock{})

	if passThroughAccoutType == MainCollector {
		require.EqualValues(t, 3, len(testHelper.App.CfedistributorKeeper.GetAllStates(ctx)))
	} else if passThroughAccoutType == ModuleAccount {
		require.EqualValues(t, 4, len(testHelper.App.CfedistributorKeeper.GetAllStates(ctx)))
	} else if passThroughAccoutType == InternalAccount {
		require.EqualValues(t, 4, len(testHelper.App.CfedistributorKeeper.GetAllStates(ctx)))
	} else {
		require.EqualValues(t, 4, len(testHelper.App.CfedistributorKeeper.GetAllStates(ctx)))
	}

	// coins flow:
	// fee 1017*51% = 518.67 to burn, so 518 burned - and burn remains 0.67

	require.EqualValues(t, initialCoinAmount.SubAmount(sdk.NewInt(518)), testHelper.App.BankKeeper.GetSupply(ctx, commontestutils.DefaultTestDenom))

	// require.EqualValues(t, sdk.NewCoin(denom, sdk.NewInt(1017+5044-518)).Add(valCoin), app.BankKeeper.GetSupply(ctx, commontestutils.DefaultTestDenom))
	burnState, _ := testHelper.App.CfedistributorKeeper.GetBurnState(ctx)
	coinRemains := burnState.CoinsStates
	require.EqualValues(t, sdk.MustNewDecFromStr("0.67"), coinRemains.AmountOf("uc4e"))

	// added 499 to main collector
	// main collector state = 499 + 5044 = 5543, but 5543 - 0,67 = 5542.33 to distribute

	if passThroughAccoutType == ModuleAccount || passThroughAccoutType == InternalAccount {
		// 5542.33 moved to c4e_distributor module or internal account
		// and all is distributed further, and 0 in remains
		c4eDIstrCoins := testHelper.App.CfedistributorKeeper.GetAccountCoinsForModuleAccount(ctx, c4eDistributorCollectorName)
		require.EqualValues(t, sdk.MustNewDecFromStr("0"), c4eDIstrCoins.AmountOf(commontestutils.DefaultTestDenom).ToDec())

		remains, _ := testHelper.App.CfedistributorKeeper.GetState(ctx, c4eDistributorCollectorName)
		//require.EqualValues(t, passThroughAccoutType == ModuleAccount, remains.Account.IsModuleAccount)
		//require.EqualValues(t, passThroughAccoutType == InternalAccount, remains.Account.IsInternalAccount)
		//require.EqualValues(t, false, remains.Account.IsMainCollector)

		coinRemainsDevelopmentFund := remains.CoinsStates
		require.EqualValues(t, sdk.MustNewDecFromStr("0"), coinRemainsDevelopmentFund.AmountOf("uc4e"))
	} else if passThroughAccoutType == BaseAccount {
		// 5542.33 moved to cosmos13zg4u07ymq83uq73t2cq3dj54jj37zzgr3hlck account
		// and all is distributed further, and 0 in remains
		address, _ := sdk.AccAddressFromBech32("cosmos13zg4u07ymq83uq73t2cq3dj54jj37zzgr3hlck")

		c4eDIstrCoins := testHelper.App.CfedistributorKeeper.GetAccountCoins(ctx, address)
		require.EqualValues(t, sdk.MustNewDecFromStr("0"), c4eDIstrCoins.AmountOf(commontestutils.DefaultTestDenom).ToDec())

		remains, _ := testHelper.App.CfedistributorKeeper.GetState(ctx, "cosmos13zg4u07ymq83uq73t2cq3dj54jj37zzgr3hlck")
		//require.EqualValues(t, passThroughAccoutType == ModuleAccount, remains.Account.IsModuleAccount)
		//require.EqualValues(t, passThroughAccoutType == InternalAccount, remains.Account.IsInternalAccount)
		//require.EqualValues(t, false, remains.Account.IsMainCollector)

		coinRemainsDevelopmentFund := remains.CoinsStates
		require.EqualValues(t, sdk.MustNewDecFromStr("0"), coinRemainsDevelopmentFund.AmountOf("uc4e"))
	}

	// 5542.33*10.345% = 573.3540385 to cosmos1p20lmfzp4g9vywl2jxwexwh6akvkxzpa6hdrag, so
	// 573 on cosmos1p20lmfzp4g9vywl2jxwexwh6akvkxzpa6hdrag and 0.3540385 on its distributor state

	acc, _ := sdk.AccAddressFromBech32("cosmos1p20lmfzp4g9vywl2jxwexwh6akvkxzpa6hdrag")
	developmentFundAccount := testHelper.App.CfedistributorKeeper.GetAccountCoins(ctx, acc)
	require.EqualValues(t, sdk.MustNewDecFromStr("573"), developmentFundAccount.AmountOf(commontestutils.DefaultTestDenom).ToDec())

	remains, _ := testHelper.App.CfedistributorKeeper.GetState(ctx, "cosmos1p20lmfzp4g9vywl2jxwexwh6akvkxzpa6hdrag")
	coinRemainsDevelopmentFund := remains.CoinsStates
	require.EqualValues(t, sdk.MustNewDecFromStr("0.3540385"), coinRemainsDevelopmentFund.AmountOf("uc4e"))

	// 5542.33 - 573.3540385 = 4968.9759615 to validators_rewards_collector, so
	// 4968 on validators_rewards_collector or no_validators module account and 0.9759615 on its distributor state

	if toValidators {
		// validators_rewards_collector coins sent to vaalidator distribition so amount is 0,

		validatorRewardCollectorAccountCoin := testHelper.App.CfedistributorKeeper.GetAccountCoinsForModuleAccount(ctx, types.ValidatorsRewardsCollector)
		require.EqualValues(t, sdk.MustNewDecFromStr("0"), validatorRewardCollectorAccountCoin.AmountOf(commontestutils.DefaultTestDenom).ToDec())
		// still 0.9759615 on its distributor state remains
		remains, _ = testHelper.App.CfedistributorKeeper.GetState(ctx, types.ValidatorsRewardsCollector)
		coinRemainsValidatorsReward := remains.CoinsStates
		require.EqualValues(t, sdk.MustNewDecFromStr("0.9759615"), coinRemainsValidatorsReward.AmountOf(commontestutils.DefaultTestDenom))
		// and 4968 to validators rewards
		distrCoins := testHelper.App.CfedistributorKeeper.GetAccountCoins(ctx, testHelper.App.DistrKeeper.GetDistributionAccount(ctx).GetAddress())
		require.EqualValues(t, sdk.NewCoins(sdk.NewCoin(commontestutils.DefaultTestDenom, sdk.NewInt(4968))), distrCoins)
	} else {
		// no_validators module account coins amount is 4968,
		// and remains 0.9759615 on its distributor state

		NoValidatorsCoin := testHelper.App.CfedistributorKeeper.GetAccountCoinsForModuleAccount(ctx, noValidatorsCollectorName)
		require.EqualValues(t, sdk.MustNewDecFromStr("4968"), NoValidatorsCoin.AmountOf(commontestutils.DefaultTestDenom).ToDec())

		remains, _ = testHelper.App.CfedistributorKeeper.GetState(ctx, noValidatorsCollectorName)
		coinRemainsValidatorsReward := remains.CoinsStates
		require.EqualValues(t, sdk.MustNewDecFromStr("0.9759615"), coinRemainsValidatorsReward.AmountOf(commontestutils.DefaultTestDenom))
	}

	// 5543 - 573 - 4968 = 2 (its ramains 0,67 + 0.3540385 + 0.9759615 = 2) on main collector
	coinOnDistributorAccount :=
		testHelper.App.CfedistributorKeeper.GetAccountCoinsForModuleAccount(ctx, types.DistributorMainAccount)
	require.EqualValues(t, sdk.MustNewDecFromStr("2"), coinOnDistributorAccount.AmountOf(commontestutils.DefaultTestDenom).ToDec())

}

func TestBurningWithInflationDistributorAfter3001Blocks(t *testing.T) {

	testHelper, ctx := testapp.SetupTestApp(t)

	var subdistributors []types.SubDistributor
	subdistributors = append(subdistributors, prepareBurningDistributor(MainCollector))
	subdistributors = append(subdistributors, prepareInflationSubDistributor(MainCollector, true))
	testHelper.App.CfedistributorKeeper.SetParams(ctx, types.NewParams(subdistributors))

	for i := int64(1); i <= 3001; i++ {

		cointToMint := sdk.NewInt(1017)

		testHelper.BankUtils.AddDefaultDenomCoinsToModule(ctx, cointToMint, authtypes.FeeCollectorName)

		cointToMintFromInflation := sdk.NewInt(5044)

		testHelper.BankUtils.AddDefaultDenomCoinsToModule(ctx, cointToMintFromInflation, types.DistributorMainAccount)

		ctx = ctx.WithBlockHeight(int64(i))
		testHelper.App.BeginBlocker(ctx, abci.RequestBeginBlock{})
		testHelper.App.EndBlocker(ctx, abci.RequestEndBlock{})
		burn, _ := sdk.NewDecFromStr("518.67")
		burn = burn.MulInt64(i)
		burn.GT(burn.TruncateDec())
		totalExpected := sdk.NewDec(i * (1017 + 5044)).Sub(burn)

		totalExpectedTruncated := totalExpected.TruncateInt()

		if burn.GT(burn.TruncateDec()) {
			totalExpectedTruncated = totalExpectedTruncated.AddRaw(1)
		}
		require.EqualValues(t, testHelper.InitialValidatorsCoin.AddAmount(totalExpectedTruncated), testHelper.App.BankKeeper.GetSupply(ctx, commontestutils.DefaultTestDenom))
	}

	ctx = ctx.WithBlockHeight(int64(3002))
	testHelper.App.BeginBlocker(ctx, abci.RequestBeginBlock{})
	testHelper.App.EndBlocker(ctx, abci.RequestEndBlock{})

	// coins flow:
	// fee 3001*1017*51% = 1556528.67 to burn, so 1556528 burned - and burn remains 0.67
	// fee 3001*(1017*51%) = 3001*518.67 = 1556528.67 to burn, so 1556528 burned - and burn remains 0.67
	require.EqualValues(t, testHelper.InitialValidatorsCoin.AddAmount(sdk.NewInt(3001*(1017+5044)-1556528)), testHelper.App.BankKeeper.GetSupply(ctx, commontestutils.DefaultTestDenom))

	burnState, _ := testHelper.App.CfedistributorKeeper.GetBurnState(ctx)
	coinRemains := burnState.CoinsStates
	require.EqualValues(t, sdk.MustNewDecFromStr("0.67"), coinRemains.AmountOf("uc4e"))

	// added 3001*1017 - 1556528 = 1495489 to main collector
	// main collector state = 1495489 + 3001*5044 = 16632533, but 16632533 - 0.67 (burning remains) = 16632532.33 to distribute

	// 16632532.33*10.345% = 1720635.4695385 to cosmos1p20lmfzp4g9vywl2jxwexwh6akvkxzpa6hdrag, so
	// 1720635 on cosmos1p20lmfzp4g9vywl2jxwexwh6akvkxzpa6hdrag and 0.4695385 on its distributor state

	acc, _ := sdk.AccAddressFromBech32("cosmos1p20lmfzp4g9vywl2jxwexwh6akvkxzpa6hdrag")
	developmentFundAccount := testHelper.App.CfedistributorKeeper.GetAccountCoins(ctx, acc)
	require.EqualValues(t, sdk.MustNewDecFromStr("1720635"), developmentFundAccount.AmountOf(commontestutils.DefaultTestDenom).ToDec())

	remains, _ := testHelper.App.CfedistributorKeeper.GetState(ctx, "cosmos1p20lmfzp4g9vywl2jxwexwh6akvkxzpa6hdrag")
	coinRemainsDevelopmentFund := remains.CoinsStates
	require.EqualValues(t, sdk.MustNewDecFromStr("0.4695385"), coinRemainsDevelopmentFund.AmountOf("uc4e"))

	// 16632532.33- 1720635.4695385 = 14911896.8604615 to validators_rewards_collector, so
	// 14911896 on validators_rewards_collector or no_validators module account and 0.8604615 on its distributor state

	// validators_rewards_collector coins sent to vaalidator distribition so amount is 0,

	validatorRewardCollectorAccountCoin := testHelper.App.CfedistributorKeeper.GetAccountCoinsForModuleAccount(ctx, types.ValidatorsRewardsCollector)
	require.EqualValues(t, sdk.MustNewDecFromStr("0"), validatorRewardCollectorAccountCoin.AmountOf(commontestutils.DefaultTestDenom).ToDec())
	// still 0.8845 on its distributor state remains
	remains, _ = testHelper.App.CfedistributorKeeper.GetState(ctx, types.ValidatorsRewardsCollector)
	coinRemainsValidatorsReward := remains.CoinsStates
	require.EqualValues(t, sdk.MustNewDecFromStr("0.8604615"), coinRemainsValidatorsReward.AmountOf("uc4e"))
	// and 14906927 to validators rewards
	distrCoins := testHelper.App.CfedistributorKeeper.GetAccountCoins(ctx, testHelper.App.DistrKeeper.GetDistributionAccount(ctx).GetAddress())
	require.EqualValues(t, sdk.NewCoins(sdk.NewCoin("uc4e", sdk.NewInt(14911896))), distrCoins)

	// 16632533 - 1720635 - 14911896 = 1 (its ramains 0.67 + 0.4695385 + 0.8604615 = 2) on main collector
	coinOnDistributorAccount :=
		testHelper.App.CfedistributorKeeper.GetAccountCoinsForModuleAccount(ctx, types.DistributorMainAccount)
	require.EqualValues(t, sdk.MustNewDecFromStr("2"), coinOnDistributorAccount.AmountOf(commontestutils.DefaultTestDenom).ToDec())

}
