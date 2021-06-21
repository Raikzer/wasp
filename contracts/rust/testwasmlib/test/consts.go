// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

package test

import "github.com/iotaledger/wasp/packages/coretypes"

const (
	ScName        = "testwasmlib"
	ScDescription = "Exercise all aspects of WasmLib"
	HScName       = coretypes.Hname(0x89703a45)

	ParamAddress     = "address"
	ParamAgentId     = "agentId"
	ParamBlockIndex  = "blockIndex"
	ParamBytes       = "bytes"
	ParamChainId     = "chainId"
	ParamColor       = "color"
	ParamHash        = "hash"
	ParamHname       = "hname"
	ParamInt16       = "int16"
	ParamInt32       = "int32"
	ParamInt64       = "int64"
	ParamRecordIndex = "recordIndex"
	ParamRequestId   = "requestId"
	ParamString      = "string"

	ResultCount  = "count"
	ResultRecord = "record"

	FuncParamTypes   = "paramTypes"
	ViewBlockRecord  = "blockRecord"
	ViewBlockRecords = "blockRecords"

	HFuncParamTypes   = coretypes.Hname(0x6921c4cd)
	HViewBlockRecord  = coretypes.Hname(0xad13b2f8)
	HViewBlockRecords = coretypes.Hname(0x16e249ea)
)
