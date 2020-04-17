package rpc_tool

import (
	"fmt"

	"testing"
)

type DynamicGlobalProperties struct {
	//ID                             types.ObjectID `json:"id"`
	HeadBlockNumber uint32 `json:"head_block_number"`
	HeadBlockID     string `json:"head_block_id"`
	//Time                           types.Time     `json:"time"`
	//CurrentWitness                 types.ObjectID `json:"current_witness"`
	//NextMaintenanceTime            types.Time     `json:"next_maintenance_time"`
	//LastBudgetTime                 types.Time     `json:"last_budget_time"`
	AccountsRegisteredThisInterval int    `json:"accounts_registered_this_interval"`
	DynamicFlags                   int    `json:"dynamic_flags"`
	RecentSlotsFilled              string `json:"recent_slots_filled"`
	LastIrreversibleBlockNum       uint32 `json:"last_irreversible_block_num"`
	CurrentAslot                   int64  `json:"current_aslot"`
	WitnessBudget                  int64  `json:"witness_budget"`
	RecentlyMissedCount            int64  `json:"recently_missed_count"`
	Parameters                     string `json:"parameters"`
}

func Test_Http(t *testing.T) {
	rpc := NewRPcTool("https://testnet.gxchain.org")
	var resp DynamicGlobalProperties
	err := rpc.Call("get_dynamic_global_properties", []interface{}{}, &resp)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

func Test_Websocket(t *testing.T) {
	rpc := NewRPcTool("wss://testnet.gxchain.org")
	var resp interface{}
	err := rpc.Call("get_dynamic_global_properties", []interface{}{}, &resp)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
