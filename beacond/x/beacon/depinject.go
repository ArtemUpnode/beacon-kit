// SPDX-License-Identifier: MIT
//
// Copyright (c) 2024 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person
// obtaining a copy of this software and associated documentation
// files (the "Software"), to deal in the Software without
// restriction, including without limitation the rights to use,
// copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following
// conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.

package beacon

import (
	"context"
	"os"
	"time"

	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/depinject"
	"cosmossdk.io/depinject/appconfig"
	modulev1alpha1 "github.com/berachain/beacon-kit/beacond/x/beacon/api/module/v1alpha1"
	"github.com/berachain/beacon-kit/beacond/x/beacon/keeper"
	"github.com/berachain/beacon-kit/mod/primitives"
	depositdb "github.com/berachain/beacon-kit/mod/storage/pkg/deposit"
	filedb "github.com/berachain/beacon-kit/mod/storage/pkg/filedb"
	"github.com/cosmos/cosmos-sdk/client/flags"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/spf13/cast"
)

//nolint:gochecknoinits // required by sdk.
func init() {
	appconfig.RegisterModule(&modulev1alpha1.Module{},
		appconfig.Provide(ProvideModule),
	)
}

// DepInjectInput is the input for the dep inject framework.
type DepInjectInput struct {
	depinject.In

	AppOpts      servertypes.AppOptions
	Environment  appmodule.Environment
	ChainSpec    primitives.ChainSpec
	DepositStore *depositdb.KVStore
}

// DepInjectOutput is the output for the dep inject framework.
type DepInjectOutput struct {
	depinject.Out

	Keeper *keeper.Keeper
	Module appmodule.AppModule
}

// ProvideModule is a function that provides the module to the application.
func ProvideModule(in DepInjectInput) DepInjectOutput {

	k := keeper.NewKeeper(
		context.Background(),
		filedb.NewDB(
			filedb.WithRootDirectory(
				cast.ToString(in.AppOpts.Get(flags.FlagHome))+"/data/blobs"),
			filedb.WithFileExtension("ssz"),
			filedb.WithDirectoryPermissions(os.ModePerm),
			filedb.WithLogger(in.Environment.Logger),
		),
		in.Environment.Logger,
		// TODO: Make this configurable.
		//nolint:mnd // will fix with config.
		10*time.Minute,
		in.Environment,
		in.ChainSpec,
		in.DepositStore,
	)

	return DepInjectOutput{
		Keeper: k,
		Module: NewAppModule(k),
	}
}
