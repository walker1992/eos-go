package types

import (
	"fmt"
	"github.com/eosspark/eos-go/common"
	"github.com/eosspark/eos-go/ecc"
	"github.com/eosspark/eos-go/log"
)

type GenesisState struct {
	EosioRootKey     string           `json:"eosio_root_key"`
	InitialTimestamp common.TimePoint `json:"initial_timestamp"`
	InitialKey       ecc.PublicKey    `json:"initial_key"`
}

func (g *GenesisState) NewGenesisState() {
	its, err := common.FromIsoString("2018-09-10T12:00:00")
	if err != nil {
		log.Error("NewGenesisState is error detail:", err)
	}
	//g.EosioRootKey = "EOS6MRyAjQq8ud7hVNYcfnVPJqcVpscN5So8BhtHuGYqET5GDW5CV"
	g.InitialTimestamp = its
	key, err := ecc.NewPublicKey("EOS6MRyAjQq8ud7hVNYcfnVPJqcVpscN5So8BhtHuGYqET5GDW5CV")
	if err != nil {
		log.Error("", err)
	}
	g.InitialKey = key
	g.initial()
}

func (gs *GenesisState) initial() {
	InitialConfiguration := common.Config{
		MaxBlockNetUsage:               common.DefaultConfig.MaxBlockNetUsage,
		TargetBlockNetUsagePct:         common.DefaultConfig.TargetBlockNetUsagePct,
		MaxTransactionNetUsage:         common.DefaultConfig.MaxTransactionNetUsage,
		BasePerTransactionNetUsage:     common.DefaultConfig.BasePerTransactionNetUsage,
		NetUsageLeeway:                 common.DefaultConfig.NetUsageLeeway,
		ContextFreeDiscountNetUsageNum: common.DefaultConfig.ContextFreeDiscountNetUsageNum,
		ContextFreeDiscountNetUsageDen: common.DefaultConfig.ContextFreeDiscountNetUsageDen,

		MaxBlockCpuUsage:       common.DefaultConfig.MaxBlockCpuUsage,
		TargetBlockCpuUsagePct: common.DefaultConfig.TargetBlockCpuUsagePct,
		MaxTransactionCpuUsage: common.DefaultConfig.MaxTransactionCpuUsage,
		MinTransactionCpuUsage: common.DefaultConfig.MinTransactionCpuUsage,

		MaxTrxLifetime:              common.DefaultConfig.MaxTrxLifetime,
		DeferredTrxExpirationWindow: common.DefaultConfig.DeferredTrxExpirationWindow,
		MaxTrxDelay:                 common.DefaultConfig.MaxTrxDelay,
		MaxInlineActionSize:         common.DefaultConfig.MaxInlineActionSize,
		MaxInlineActionDepth:        common.DefaultConfig.MaxInlineActionDepth,
		MaxAuthorityDepth:           common.DefaultConfig.MaxAuthorityDepth,
	}

	fmt.Println(InitialConfiguration)
}