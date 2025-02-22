package migrations

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/sei-protocol/sei-chain/x/dex/types"
)

func PriceSnapshotUpdate(ctx sdk.Context, paramStore paramtypes.Subspace) error {
	migratePriceSnapshotParam(ctx, paramStore)
	return nil
}

func migratePriceSnapshotParam(ctx sdk.Context, paramStore paramtypes.Subspace) error {
	defaultParams := types.Params{
		PriceSnapshotRetention: types.DefaultPriceSnapshotRetention,
	}
	paramStore.SetParamSet(ctx, &defaultParams)
	return nil
}
