package reqstatus

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/iotaledger/wasp/packages/chain"
	"github.com/iotaledger/wasp/packages/chains"
	"github.com/iotaledger/wasp/packages/iscp"
	"github.com/iotaledger/wasp/packages/vm/gas"
	"github.com/iotaledger/wasp/packages/webapi/httperrors"
	"github.com/iotaledger/wasp/packages/webapi/model"
	"github.com/iotaledger/wasp/packages/webapi/routes"
	"github.com/labstack/echo/v4"
	"github.com/pangpanglabs/echoswagger/v2"
	"golang.org/x/xerrors"
)

// TODO ch.GetRequestReceipt(reqID) is succeptible to optimistic state read error - it doesn't retry atm,
// should probably have a similar retry strategy to the CallView functionality

type reqstatusWebAPI struct {
	getChain func(chainID *iscp.ChainID) chain.ChainRequests
}

// TODO  add examples for receipt json
func AddEndpoints(server echoswagger.ApiRouter, getChain chains.ChainProvider) {
	r := &reqstatusWebAPI{func(chainID *iscp.ChainID) chain.ChainRequests {
		return getChain(chainID)
	}}

	server.GET(routes.RequestReceipt(":chainID", ":reqID"), r.handleRequestReceipt).
		SetSummary("Get the processing status of a given request in the node").
		AddParamPath("", "chainID", "ChainID (Hex)").
		AddParamPath("", "reqID", "Request ID").
		AddResponse(http.StatusOK, "Request Receipt", model.RequestReceiptResponse{}, nil)

	server.GET(routes.WaitRequestProcessed(":chainID", ":reqID"), r.handleWaitRequestProcessed).
		SetSummary("Wait until the given request has been processed by the node").
		AddParamPath("", "chainID", "ChainID (Hex)").
		AddParamPath("", "reqID", "Request ID").
		AddParamBody(model.WaitRequestProcessedParams{}, "Params", "Optional parameters", false).
		AddResponse(http.StatusOK, "Request Receipt", model.RequestReceiptResponse{}, nil)
}

func (r *reqstatusWebAPI) handleRequestReceipt(c echo.Context) error {
	ch, reqID, err := r.parseParams(c)
	if err != nil {
		return err
	}

	receiptResponse, err := getTranslatedReceipt(ch, reqID)
	if err != nil {
		return httperrors.ServerError(err.Error())
	}
	return c.JSON(http.StatusOK, receiptResponse)
}

const WaitRequestProcessedDefaultTimeout = 30 * time.Second

func (r *reqstatusWebAPI) handleWaitRequestProcessed(c echo.Context) error {
	ch, reqID, err := r.parseParams(c)
	if err != nil {
		return err
	}

	req := model.WaitRequestProcessedParams{
		Timeout: WaitRequestProcessedDefaultTimeout,
	}
	if c.Request().Header.Get("Content-Type") == "application/json" {
		if err := c.Bind(&req); err != nil {
			return httperrors.BadRequest("Invalid request body")
		}
	}

	tryGetReceipt := func() error {
		receiptResponse, err := getTranslatedReceipt(ch, reqID)
		if err != nil {
			return httperrors.ServerError(err.Error())
		}
		if receiptResponse != nil {
			return c.JSON(http.StatusOK, receiptResponse)
		}
		return nil
	}

	ret := tryGetReceipt()
	if ret != nil {
		return ret
	}

	// subscribe to event
	requestProcessed := make(chan bool)
	attachID := ch.AttachToRequestProcessed(func(rid iscp.RequestID) {
		if rid == reqID {
			requestProcessed <- true
		}
	})
	defer ch.DetachFromRequestProcessed(attachID)

	select {
	case <-requestProcessed:
		ret := tryGetReceipt()
		if ret != nil {
			return ret
		}
		return httperrors.ServerError("Unexpected error, receipt not found after request has been processed")
	case <-time.After(req.Timeout):
		// check again, in case event was triggered just before we subscribed
		ret := tryGetReceipt()
		if ret != nil {
			return ret
		}
		return httperrors.Timeout("Timeout while waiting for request to be processed")
	}
}

func (r *reqstatusWebAPI) parseParams(c echo.Context) (chain.ChainRequests, iscp.RequestID, error) {
	chainID, err := iscp.ChainIDFromString(c.Param("chainID"))
	if err != nil {
		return nil, iscp.RequestID{}, httperrors.BadRequest(fmt.Sprintf("Invalid Chain ID %+v: %s", c.Param("chainID"), err.Error()))
	}
	theChain := r.getChain(chainID)
	if theChain == nil {
		return nil, iscp.RequestID{}, httperrors.NotFound(fmt.Sprintf("Chain not found: %s", chainID.String()))
	}
	reqID, err := iscp.RequestIDFromString(c.Param("reqID"))
	if err != nil {
		return nil, iscp.RequestID{}, httperrors.BadRequest(fmt.Sprintf("Invalid request id %+v: %s", c.Param("reqID"), err.Error()))
	}
	return theChain, reqID, nil
}

func getTranslatedReceipt(ch chain.ChainRequests, reqID iscp.RequestID) (*model.RequestReceiptResponse, error) {
	receipt, err := ch.GetRequestReceipt(reqID)
	if err != nil {
		return nil, xerrors.Errorf("error getting request receipt: %s", err)
	}
	if receipt == nil {
		return nil, nil
	}
	iscpReceipt := &iscp.Receipt{
		Request:       receipt.Request,
		Error:         receipt.Error,
		GasBudget:     receipt.GasBudget,
		GasBurned:     receipt.GasBurned,
		GasFeeCharged: receipt.GasFeeCharged,
		BlockIndex:    receipt.BlockIndex,
		RequestIndex:  receipt.RequestIndex,
		// TODO translate receipt error
		TranslatedError: "",
		GasBurnLog:      &gas.BurnLog{},
	}
	receiptJSON, err := json.Marshal(iscpReceipt)
	if err != nil {
		return nil, xerrors.Errorf("error marshaling receipt into JSON: %s", err)
	}
	return &model.RequestReceiptResponse{
		Receipt: string(receiptJSON),
	}, nil
}
