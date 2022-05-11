// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package coreblocklog

import "github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib/wasmtypes"

type ImmutableControlAddressesResults struct {
	proxy wasmtypes.Proxy
}

// the addresses have been set as state controller address or governing address since the following block index
func (s ImmutableControlAddressesResults) BlockIndex() wasmtypes.ScImmutableUint32 {
	return wasmtypes.NewScImmutableUint32(s.proxy.Root(ResultBlockIndex))
}

func (s ImmutableControlAddressesResults) GoverningAddress() wasmtypes.ScImmutableAddress {
	return wasmtypes.NewScImmutableAddress(s.proxy.Root(ResultGoverningAddress))
}

func (s ImmutableControlAddressesResults) StateControllerAddress() wasmtypes.ScImmutableAddress {
	return wasmtypes.NewScImmutableAddress(s.proxy.Root(ResultStateControllerAddress))
}

type MutableControlAddressesResults struct {
	proxy wasmtypes.Proxy
}

// the addresses have been set as state controller address or governing address since the following block index
func (s MutableControlAddressesResults) BlockIndex() wasmtypes.ScMutableUint32 {
	return wasmtypes.NewScMutableUint32(s.proxy.Root(ResultBlockIndex))
}

func (s MutableControlAddressesResults) GoverningAddress() wasmtypes.ScMutableAddress {
	return wasmtypes.NewScMutableAddress(s.proxy.Root(ResultGoverningAddress))
}

func (s MutableControlAddressesResults) StateControllerAddress() wasmtypes.ScMutableAddress {
	return wasmtypes.NewScMutableAddress(s.proxy.Root(ResultStateControllerAddress))
}

type ImmutableGetBlockInfoResults struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableGetBlockInfoResults) BlockInfo() wasmtypes.ScImmutableBytes {
	return wasmtypes.NewScImmutableBytes(s.proxy.Root(ResultBlockInfo))
}

type MutableGetBlockInfoResults struct {
	proxy wasmtypes.Proxy
}

func (s MutableGetBlockInfoResults) BlockInfo() wasmtypes.ScMutableBytes {
	return wasmtypes.NewScMutableBytes(s.proxy.Root(ResultBlockInfo))
}

type ArrayOfImmutableBytes struct {
	proxy wasmtypes.Proxy
}

func (a ArrayOfImmutableBytes) Length() uint32 {
	return a.proxy.Length()
}

func (a ArrayOfImmutableBytes) GetBytes(index uint32) wasmtypes.ScImmutableBytes {
	return wasmtypes.NewScImmutableBytes(a.proxy.Index(index))
}

type ImmutableGetEventsForBlockResults struct {
	proxy wasmtypes.Proxy
}

// native contract, so this is an Array16
func (s ImmutableGetEventsForBlockResults) Event() ArrayOfImmutableBytes {
	return ArrayOfImmutableBytes{proxy: s.proxy.Root(ResultEvent)}
}

type ArrayOfMutableBytes struct {
	proxy wasmtypes.Proxy
}

func (a ArrayOfMutableBytes) AppendBytes() wasmtypes.ScMutableBytes {
	return wasmtypes.NewScMutableBytes(a.proxy.Append())
}

func (a ArrayOfMutableBytes) Clear() {
	a.proxy.ClearArray()
}

func (a ArrayOfMutableBytes) Length() uint32 {
	return a.proxy.Length()
}

func (a ArrayOfMutableBytes) GetBytes(index uint32) wasmtypes.ScMutableBytes {
	return wasmtypes.NewScMutableBytes(a.proxy.Index(index))
}

type MutableGetEventsForBlockResults struct {
	proxy wasmtypes.Proxy
}

// native contract, so this is an Array16
func (s MutableGetEventsForBlockResults) Event() ArrayOfMutableBytes {
	return ArrayOfMutableBytes{proxy: s.proxy.Root(ResultEvent)}
}

type ImmutableGetEventsForContractResults struct {
	proxy wasmtypes.Proxy
}

// native contract, so this is an Array16
func (s ImmutableGetEventsForContractResults) Event() ArrayOfImmutableBytes {
	return ArrayOfImmutableBytes{proxy: s.proxy.Root(ResultEvent)}
}

type MutableGetEventsForContractResults struct {
	proxy wasmtypes.Proxy
}

// native contract, so this is an Array16
func (s MutableGetEventsForContractResults) Event() ArrayOfMutableBytes {
	return ArrayOfMutableBytes{proxy: s.proxy.Root(ResultEvent)}
}

type ImmutableGetEventsForRequestResults struct {
	proxy wasmtypes.Proxy
}

// native contract, so this is an Array16
func (s ImmutableGetEventsForRequestResults) Event() ArrayOfImmutableBytes {
	return ArrayOfImmutableBytes{proxy: s.proxy.Root(ResultEvent)}
}

type MutableGetEventsForRequestResults struct {
	proxy wasmtypes.Proxy
}

// native contract, so this is an Array16
func (s MutableGetEventsForRequestResults) Event() ArrayOfMutableBytes {
	return ArrayOfMutableBytes{proxy: s.proxy.Root(ResultEvent)}
}

type ImmutableGetLatestBlockInfoResults struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableGetLatestBlockInfoResults) BlockIndex() wasmtypes.ScImmutableUint32 {
	return wasmtypes.NewScImmutableUint32(s.proxy.Root(ResultBlockIndex))
}

func (s ImmutableGetLatestBlockInfoResults) BlockInfo() wasmtypes.ScImmutableBytes {
	return wasmtypes.NewScImmutableBytes(s.proxy.Root(ResultBlockInfo))
}

type MutableGetLatestBlockInfoResults struct {
	proxy wasmtypes.Proxy
}

func (s MutableGetLatestBlockInfoResults) BlockIndex() wasmtypes.ScMutableUint32 {
	return wasmtypes.NewScMutableUint32(s.proxy.Root(ResultBlockIndex))
}

func (s MutableGetLatestBlockInfoResults) BlockInfo() wasmtypes.ScMutableBytes {
	return wasmtypes.NewScMutableBytes(s.proxy.Root(ResultBlockInfo))
}

type ArrayOfImmutableRequestID struct {
	proxy wasmtypes.Proxy
}

func (a ArrayOfImmutableRequestID) Length() uint32 {
	return a.proxy.Length()
}

func (a ArrayOfImmutableRequestID) GetRequestID(index uint32) wasmtypes.ScImmutableRequestID {
	return wasmtypes.NewScImmutableRequestID(a.proxy.Index(index))
}

type ImmutableGetRequestIDsForBlockResults struct {
	proxy wasmtypes.Proxy
}

// native contract, so this is an Array16
func (s ImmutableGetRequestIDsForBlockResults) RequestID() ArrayOfImmutableRequestID {
	return ArrayOfImmutableRequestID{proxy: s.proxy.Root(ResultRequestID)}
}

type ArrayOfMutableRequestID struct {
	proxy wasmtypes.Proxy
}

func (a ArrayOfMutableRequestID) AppendRequestID() wasmtypes.ScMutableRequestID {
	return wasmtypes.NewScMutableRequestID(a.proxy.Append())
}

func (a ArrayOfMutableRequestID) Clear() {
	a.proxy.ClearArray()
}

func (a ArrayOfMutableRequestID) Length() uint32 {
	return a.proxy.Length()
}

func (a ArrayOfMutableRequestID) GetRequestID(index uint32) wasmtypes.ScMutableRequestID {
	return wasmtypes.NewScMutableRequestID(a.proxy.Index(index))
}

type MutableGetRequestIDsForBlockResults struct {
	proxy wasmtypes.Proxy
}

// native contract, so this is an Array16
func (s MutableGetRequestIDsForBlockResults) RequestID() ArrayOfMutableRequestID {
	return ArrayOfMutableRequestID{proxy: s.proxy.Root(ResultRequestID)}
}

type ImmutableGetRequestReceiptResults struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableGetRequestReceiptResults) BlockIndex() wasmtypes.ScImmutableUint32 {
	return wasmtypes.NewScImmutableUint32(s.proxy.Root(ResultBlockIndex))
}

func (s ImmutableGetRequestReceiptResults) RequestIndex() wasmtypes.ScImmutableUint16 {
	return wasmtypes.NewScImmutableUint16(s.proxy.Root(ResultRequestIndex))
}

func (s ImmutableGetRequestReceiptResults) RequestRecord() wasmtypes.ScImmutableBytes {
	return wasmtypes.NewScImmutableBytes(s.proxy.Root(ResultRequestRecord))
}

type MutableGetRequestReceiptResults struct {
	proxy wasmtypes.Proxy
}

func (s MutableGetRequestReceiptResults) BlockIndex() wasmtypes.ScMutableUint32 {
	return wasmtypes.NewScMutableUint32(s.proxy.Root(ResultBlockIndex))
}

func (s MutableGetRequestReceiptResults) RequestIndex() wasmtypes.ScMutableUint16 {
	return wasmtypes.NewScMutableUint16(s.proxy.Root(ResultRequestIndex))
}

func (s MutableGetRequestReceiptResults) RequestRecord() wasmtypes.ScMutableBytes {
	return wasmtypes.NewScMutableBytes(s.proxy.Root(ResultRequestRecord))
}

type ImmutableGetRequestReceiptsForBlockResults struct {
	proxy wasmtypes.Proxy
}

// native contract, so this is an Array16
func (s ImmutableGetRequestReceiptsForBlockResults) RequestRecord() ArrayOfImmutableBytes {
	return ArrayOfImmutableBytes{proxy: s.proxy.Root(ResultRequestRecord)}
}

type MutableGetRequestReceiptsForBlockResults struct {
	proxy wasmtypes.Proxy
}

// native contract, so this is an Array16
func (s MutableGetRequestReceiptsForBlockResults) RequestRecord() ArrayOfMutableBytes {
	return ArrayOfMutableBytes{proxy: s.proxy.Root(ResultRequestRecord)}
}

type ImmutableIsRequestProcessedResults struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableIsRequestProcessedResults) RequestProcessed() wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(s.proxy.Root(ResultRequestProcessed))
}

type MutableIsRequestProcessedResults struct {
	proxy wasmtypes.Proxy
}

func (s MutableIsRequestProcessedResults) RequestProcessed() wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(s.proxy.Root(ResultRequestProcessed))
}
