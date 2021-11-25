// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

//nolint:gocritic
package erc20

import "github.com/iotaledger/wasp/packages/vm/wasmlib/go/wasmlib"

type Erc20Events struct{}

func (e Erc20Events) Transfer(amount int64, from wasmlib.ScAgentID, to wasmlib.ScAgentID) {
	wasmlib.NewEventEncoder("erc20.transfer").
		Int64(amount).
		AgentID(from).
		AgentID(to).
		Emit()
}
