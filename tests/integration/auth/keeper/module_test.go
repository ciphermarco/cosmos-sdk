package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"cosmossdk.io/depinject"
	"cosmossdk.io/log"
	"cosmossdk.io/x/auth/keeper"
	"cosmossdk.io/x/auth/testutil"
	"cosmossdk.io/x/auth/types"

	simtestutil "github.com/cosmos/cosmos-sdk/x/simulation/helper"
)

// Can move to tests/integration
func TestItCreatesModuleAccountOnInitBlock(t *testing.T) {
	var accountKeeper keeper.AccountKeeper
	app, err := simtestutil.SetupAtGenesis(
		depinject.Configs(
			testutil.AppConfig,
			depinject.Supply(log.NewNopLogger()),
		),
		&accountKeeper)
	require.NoError(t, err)

	ctx := app.BaseApp.NewContext(false)
	acc := accountKeeper.GetAccount(ctx, types.NewModuleAddress(types.FeeCollectorName))
	require.NotNil(t, acc)
}
