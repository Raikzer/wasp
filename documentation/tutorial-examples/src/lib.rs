// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

#![allow(dead_code)]
#![allow(unused_imports)]

use solotutorial::*;
use wasmlib::*;

use crate::consts::*;
use crate::params::*;
use crate::results::*;
use crate::state::*;

mod consts;
mod contract;
mod params;
mod results;
mod state;

mod solotutorial;

const EXPORT_MAP: ScExportMap = ScExportMap {
    names: &[
    	FUNC_STORE_STRING,
    	VIEW_GET_STRING,
	],
    funcs: &[
    	func_store_string_thunk,
	],
    views: &[
    	view_get_string_thunk,
	],
};

#[no_mangle]
fn on_call(index: i32) {
	EXPORT_MAP.call(index);
}

#[no_mangle]
fn on_load() {
	EXPORT_MAP.export();
}

pub struct StoreStringContext {
	params: ImmutableStoreStringParams,
	state: MutableSoloTutorialState,
}

fn func_store_string_thunk(ctx: &ScFuncContext) {
	ctx.log("solotutorial.funcStoreString");
	let f = StoreStringContext {
		params: ImmutableStoreStringParams { proxy: params_proxy() },
		state: MutableSoloTutorialState { proxy: state_proxy() },
	};
	ctx.require(f.params.str().exists(), "missing mandatory str");
	func_store_string(ctx, &f);
	ctx.log("solotutorial.funcStoreString ok");
}

pub struct GetStringContext {
	results: MutableGetStringResults,
	state: ImmutableSoloTutorialState,
}

fn view_get_string_thunk(ctx: &ScViewContext) {
	ctx.log("solotutorial.viewGetString");
	let f = GetStringContext {
		results: MutableGetStringResults { proxy: results_proxy() },
		state: ImmutableSoloTutorialState { proxy: state_proxy() },
	};
	view_get_string(ctx, &f);
	ctx.results(&f.results.proxy.kv_store);
	ctx.log("solotutorial.viewGetString ok");
}
