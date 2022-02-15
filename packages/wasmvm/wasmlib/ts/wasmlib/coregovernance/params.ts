// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

import * as wasmtypes from "wasmlib/wasmtypes";
import * as sc from "./index";

export class ImmutableAddAllowedStateControllerAddressParams extends wasmtypes.ScProxy {
	chainOwner(): wasmtypes.ScImmutableAgentID {
		return new wasmtypes.ScImmutableAgentID(this.proxy.root(sc.ParamChainOwner));
	}

	feeColor(): wasmtypes.ScImmutableColor {
		return new wasmtypes.ScImmutableColor(this.proxy.root(sc.ParamFeeColor));
	}

	stateControllerAddress(): wasmtypes.ScImmutableAddress {
		return new wasmtypes.ScImmutableAddress(this.proxy.root(sc.ParamStateControllerAddress));
	}
}

export class MutableAddAllowedStateControllerAddressParams extends wasmtypes.ScProxy {
	chainOwner(): wasmtypes.ScMutableAgentID {
		return new wasmtypes.ScMutableAgentID(this.proxy.root(sc.ParamChainOwner));
	}

	feeColor(): wasmtypes.ScMutableColor {
		return new wasmtypes.ScMutableColor(this.proxy.root(sc.ParamFeeColor));
	}

	stateControllerAddress(): wasmtypes.ScMutableAddress {
		return new wasmtypes.ScMutableAddress(this.proxy.root(sc.ParamStateControllerAddress));
	}
}

export class ImmutableDelegateChainOwnershipParams extends wasmtypes.ScProxy {
	chainOwner(): wasmtypes.ScImmutableAgentID {
		return new wasmtypes.ScImmutableAgentID(this.proxy.root(sc.ParamChainOwner));
	}
}

export class MutableDelegateChainOwnershipParams extends wasmtypes.ScProxy {
	chainOwner(): wasmtypes.ScMutableAgentID {
		return new wasmtypes.ScMutableAgentID(this.proxy.root(sc.ParamChainOwner));
	}
}

export class ImmutableRemoveAllowedStateControllerAddressParams extends wasmtypes.ScProxy {
	stateControllerAddress(): wasmtypes.ScImmutableAddress {
		return new wasmtypes.ScImmutableAddress(this.proxy.root(sc.ParamStateControllerAddress));
	}
}

export class MutableRemoveAllowedStateControllerAddressParams extends wasmtypes.ScProxy {
	stateControllerAddress(): wasmtypes.ScMutableAddress {
		return new wasmtypes.ScMutableAddress(this.proxy.root(sc.ParamStateControllerAddress));
	}
}

export class ImmutableRotateStateControllerParams extends wasmtypes.ScProxy {
	stateControllerAddress(): wasmtypes.ScImmutableAddress {
		return new wasmtypes.ScImmutableAddress(this.proxy.root(sc.ParamStateControllerAddress));
	}
}

export class MutableRotateStateControllerParams extends wasmtypes.ScProxy {
	stateControllerAddress(): wasmtypes.ScMutableAddress {
		return new wasmtypes.ScMutableAddress(this.proxy.root(sc.ParamStateControllerAddress));
	}
}

export class ImmutableSetChainInfoParams extends wasmtypes.ScProxy {
	maxBlobSize(): wasmtypes.ScImmutableInt32 {
		return new wasmtypes.ScImmutableInt32(this.proxy.root(sc.ParamMaxBlobSize));
	}

	maxEventSize(): wasmtypes.ScImmutableInt16 {
		return new wasmtypes.ScImmutableInt16(this.proxy.root(sc.ParamMaxEventSize));
	}

	maxEventsPerReq(): wasmtypes.ScImmutableInt16 {
		return new wasmtypes.ScImmutableInt16(this.proxy.root(sc.ParamMaxEventsPerReq));
	}

	ownerFee(): wasmtypes.ScImmutableInt64 {
		return new wasmtypes.ScImmutableInt64(this.proxy.root(sc.ParamOwnerFee));
	}

	validatorFee(): wasmtypes.ScImmutableInt64 {
		return new wasmtypes.ScImmutableInt64(this.proxy.root(sc.ParamValidatorFee));
	}
}

export class MutableSetChainInfoParams extends wasmtypes.ScProxy {
	maxBlobSize(): wasmtypes.ScMutableInt32 {
		return new wasmtypes.ScMutableInt32(this.proxy.root(sc.ParamMaxBlobSize));
	}

	maxEventSize(): wasmtypes.ScMutableInt16 {
		return new wasmtypes.ScMutableInt16(this.proxy.root(sc.ParamMaxEventSize));
	}

	maxEventsPerReq(): wasmtypes.ScMutableInt16 {
		return new wasmtypes.ScMutableInt16(this.proxy.root(sc.ParamMaxEventsPerReq));
	}

	ownerFee(): wasmtypes.ScMutableInt64 {
		return new wasmtypes.ScMutableInt64(this.proxy.root(sc.ParamOwnerFee));
	}

	validatorFee(): wasmtypes.ScMutableInt64 {
		return new wasmtypes.ScMutableInt64(this.proxy.root(sc.ParamValidatorFee));
	}
}

export class ImmutableSetContractFeeParams extends wasmtypes.ScProxy {
	hname(): wasmtypes.ScImmutableHname {
		return new wasmtypes.ScImmutableHname(this.proxy.root(sc.ParamHname));
	}

	ownerFee(): wasmtypes.ScImmutableInt64 {
		return new wasmtypes.ScImmutableInt64(this.proxy.root(sc.ParamOwnerFee));
	}

	validatorFee(): wasmtypes.ScImmutableInt64 {
		return new wasmtypes.ScImmutableInt64(this.proxy.root(sc.ParamValidatorFee));
	}
}

export class MutableSetContractFeeParams extends wasmtypes.ScProxy {
	hname(): wasmtypes.ScMutableHname {
		return new wasmtypes.ScMutableHname(this.proxy.root(sc.ParamHname));
	}

	ownerFee(): wasmtypes.ScMutableInt64 {
		return new wasmtypes.ScMutableInt64(this.proxy.root(sc.ParamOwnerFee));
	}

	validatorFee(): wasmtypes.ScMutableInt64 {
		return new wasmtypes.ScMutableInt64(this.proxy.root(sc.ParamValidatorFee));
	}
}

export class ImmutableSetDefaultFeeParams extends wasmtypes.ScProxy {
	ownerFee(): wasmtypes.ScImmutableInt64 {
		return new wasmtypes.ScImmutableInt64(this.proxy.root(sc.ParamOwnerFee));
	}

	validatorFee(): wasmtypes.ScImmutableInt64 {
		return new wasmtypes.ScImmutableInt64(this.proxy.root(sc.ParamValidatorFee));
	}
}

export class MutableSetDefaultFeeParams extends wasmtypes.ScProxy {
	ownerFee(): wasmtypes.ScMutableInt64 {
		return new wasmtypes.ScMutableInt64(this.proxy.root(sc.ParamOwnerFee));
	}

	validatorFee(): wasmtypes.ScMutableInt64 {
		return new wasmtypes.ScMutableInt64(this.proxy.root(sc.ParamValidatorFee));
	}
}

export class ImmutableGetFeeInfoParams extends wasmtypes.ScProxy {
	hname(): wasmtypes.ScImmutableHname {
		return new wasmtypes.ScImmutableHname(this.proxy.root(sc.ParamHname));
	}
}

export class MutableGetFeeInfoParams extends wasmtypes.ScProxy {
	hname(): wasmtypes.ScMutableHname {
		return new wasmtypes.ScMutableHname(this.proxy.root(sc.ParamHname));
	}
}