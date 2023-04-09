// Code generated by schema tool; DO NOT EDIT.

// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

#![allow(dead_code)]

use crate::*;
use crate::coreblocklog::*;

pub struct ControlAddressesCall<'a> {
    pub func:    ScView<'a>,
    pub results: ImmutableControlAddressesResults,
}

pub struct GetBlockInfoCall<'a> {
    pub func:    ScView<'a>,
    pub params:  MutableGetBlockInfoParams,
    pub results: ImmutableGetBlockInfoResults,
}

pub struct GetEventsForBlockCall<'a> {
    pub func:    ScView<'a>,
    pub params:  MutableGetEventsForBlockParams,
    pub results: ImmutableGetEventsForBlockResults,
}

pub struct GetEventsForContractCall<'a> {
    pub func:    ScView<'a>,
    pub params:  MutableGetEventsForContractParams,
    pub results: ImmutableGetEventsForContractResults,
}

pub struct GetEventsForRequestCall<'a> {
    pub func:    ScView<'a>,
    pub params:  MutableGetEventsForRequestParams,
    pub results: ImmutableGetEventsForRequestResults,
}

pub struct GetRequestIDsForBlockCall<'a> {
    pub func:    ScView<'a>,
    pub params:  MutableGetRequestIDsForBlockParams,
    pub results: ImmutableGetRequestIDsForBlockResults,
}

pub struct GetRequestReceiptCall<'a> {
    pub func:    ScView<'a>,
    pub params:  MutableGetRequestReceiptParams,
    pub results: ImmutableGetRequestReceiptResults,
}

pub struct GetRequestReceiptsForBlockCall<'a> {
    pub func:    ScView<'a>,
    pub params:  MutableGetRequestReceiptsForBlockParams,
    pub results: ImmutableGetRequestReceiptsForBlockResults,
}

pub struct IsRequestProcessedCall<'a> {
    pub func:    ScView<'a>,
    pub params:  MutableIsRequestProcessedParams,
    pub results: ImmutableIsRequestProcessedResults,
}

pub struct ScFuncs {
}

impl ScFuncs {
    // Returns the current state controller and governing addresses and at what block index they were set.
    pub fn control_addresses(ctx: &impl ScViewCallContext) -> ControlAddressesCall {
        let mut f = ControlAddressesCall {
            func:    ScView::new(ctx, HSC_NAME, HVIEW_CONTROL_ADDRESSES),
            results: ImmutableControlAddressesResults { proxy: Proxy::nil() },
        };
        ScView::link_results(&mut f.results.proxy, &f.func);
        f
    }

    // Returns information about the given block.
    pub fn get_block_info(ctx: &impl ScViewCallContext) -> GetBlockInfoCall {
        let mut f = GetBlockInfoCall {
            func:    ScView::new(ctx, HSC_NAME, HVIEW_GET_BLOCK_INFO),
            params:  MutableGetBlockInfoParams { proxy: Proxy::nil() },
            results: ImmutableGetBlockInfoResults { proxy: Proxy::nil() },
        };
        ScView::link_params(&mut f.params.proxy, &f.func);
        ScView::link_results(&mut f.results.proxy, &f.func);
        f
    }

    // Returns the list of events triggered during the execution of the given block.
    pub fn get_events_for_block(ctx: &impl ScViewCallContext) -> GetEventsForBlockCall {
        let mut f = GetEventsForBlockCall {
            func:    ScView::new(ctx, HSC_NAME, HVIEW_GET_EVENTS_FOR_BLOCK),
            params:  MutableGetEventsForBlockParams { proxy: Proxy::nil() },
            results: ImmutableGetEventsForBlockResults { proxy: Proxy::nil() },
        };
        ScView::link_params(&mut f.params.proxy, &f.func);
        ScView::link_results(&mut f.results.proxy, &f.func);
        f
    }

    // Returns the list of events triggered by the given contract
    // during the execution of the given block range.
    pub fn get_events_for_contract(ctx: &impl ScViewCallContext) -> GetEventsForContractCall {
        let mut f = GetEventsForContractCall {
            func:    ScView::new(ctx, HSC_NAME, HVIEW_GET_EVENTS_FOR_CONTRACT),
            params:  MutableGetEventsForContractParams { proxy: Proxy::nil() },
            results: ImmutableGetEventsForContractResults { proxy: Proxy::nil() },
        };
        ScView::link_params(&mut f.params.proxy, &f.func);
        ScView::link_results(&mut f.results.proxy, &f.func);
        f
    }

    // Returns the list of events triggered during the execution of the given request.
    pub fn get_events_for_request(ctx: &impl ScViewCallContext) -> GetEventsForRequestCall {
        let mut f = GetEventsForRequestCall {
            func:    ScView::new(ctx, HSC_NAME, HVIEW_GET_EVENTS_FOR_REQUEST),
            params:  MutableGetEventsForRequestParams { proxy: Proxy::nil() },
            results: ImmutableGetEventsForRequestResults { proxy: Proxy::nil() },
        };
        ScView::link_params(&mut f.params.proxy, &f.func);
        ScView::link_results(&mut f.results.proxy, &f.func);
        f
    }

    // Returns a list with all request IDs in the given block.
    pub fn get_request_i_ds_for_block(ctx: &impl ScViewCallContext) -> GetRequestIDsForBlockCall {
        let mut f = GetRequestIDsForBlockCall {
            func:    ScView::new(ctx, HSC_NAME, HVIEW_GET_REQUEST_I_DS_FOR_BLOCK),
            params:  MutableGetRequestIDsForBlockParams { proxy: Proxy::nil() },
            results: ImmutableGetRequestIDsForBlockResults { proxy: Proxy::nil() },
        };
        ScView::link_params(&mut f.params.proxy, &f.func);
        ScView::link_results(&mut f.results.proxy, &f.func);
        f
    }

    // Returns the receipt for the request with the given ID.
    pub fn get_request_receipt(ctx: &impl ScViewCallContext) -> GetRequestReceiptCall {
        let mut f = GetRequestReceiptCall {
            func:    ScView::new(ctx, HSC_NAME, HVIEW_GET_REQUEST_RECEIPT),
            params:  MutableGetRequestReceiptParams { proxy: Proxy::nil() },
            results: ImmutableGetRequestReceiptResults { proxy: Proxy::nil() },
        };
        ScView::link_params(&mut f.params.proxy, &f.func);
        ScView::link_results(&mut f.results.proxy, &f.func);
        f
    }

    // Returns all request receipts in the given block.
    pub fn get_request_receipts_for_block(ctx: &impl ScViewCallContext) -> GetRequestReceiptsForBlockCall {
        let mut f = GetRequestReceiptsForBlockCall {
            func:    ScView::new(ctx, HSC_NAME, HVIEW_GET_REQUEST_RECEIPTS_FOR_BLOCK),
            params:  MutableGetRequestReceiptsForBlockParams { proxy: Proxy::nil() },
            results: ImmutableGetRequestReceiptsForBlockResults { proxy: Proxy::nil() },
        };
        ScView::link_params(&mut f.params.proxy, &f.func);
        ScView::link_results(&mut f.results.proxy, &f.func);
        f
    }

    // Returns whether the request with ID u has been processed.
    pub fn is_request_processed(ctx: &impl ScViewCallContext) -> IsRequestProcessedCall {
        let mut f = IsRequestProcessedCall {
            func:    ScView::new(ctx, HSC_NAME, HVIEW_IS_REQUEST_PROCESSED),
            params:  MutableIsRequestProcessedParams { proxy: Proxy::nil() },
            results: ImmutableIsRequestProcessedResults { proxy: Proxy::nil() },
        };
        ScView::link_params(&mut f.params.proxy, &f.func);
        ScView::link_results(&mut f.results.proxy, &f.func);
        f
    }
}
