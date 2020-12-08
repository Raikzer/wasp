package chainclient

import (
	"github.com/iotaledger/wasp/packages/coretypes"
	"github.com/iotaledger/wasp/packages/kv/dict"

	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/address"
	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/address/signaturescheme"
	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/balance"
	"github.com/iotaledger/wasp/client"
	"github.com/iotaledger/wasp/packages/apilib"
	"github.com/iotaledger/wasp/packages/nodeclient"
	"github.com/iotaledger/wasp/packages/sctransaction"
)

type Client struct {
	NodeClient nodeclient.NodeClient
	WaspClient *client.WaspClient
	ChainID    coretypes.ChainID
	SigScheme  signaturescheme.SignatureScheme
}

func New(
	nodeClient nodeclient.NodeClient,
	waspClient *client.WaspClient,
	chainID coretypes.ChainID,
	sigScheme signaturescheme.SignatureScheme,
) *Client {
	return &Client{
		NodeClient: nodeClient,
		WaspClient: waspClient,
		ChainID:    chainID,
		SigScheme:  sigScheme,
	}
}

type PostRequestParams struct {
	Mint     map[address.Address]int64 // TODO
	Transfer map[balance.Color]int64
	Args     dict.Dict
}

func (c *Client) PostRequest(
	contractHname coretypes.Hname,
	entryPoint coretypes.Hname,
	params ...PostRequestParams,
) (*sctransaction.Transaction, error) {
	par := PostRequestParams{}
	if len(params) > 0 {
		par = params[0]
	}

	return apilib.CreateRequestTransaction(apilib.CreateRequestTransactionParams{
		NodeClient:      c.NodeClient,
		SenderSigScheme: c.SigScheme,
		Mint:            par.Mint,
		RequestSectionParams: []apilib.RequestSectionParams{{
			TargetContractID: coretypes.NewContractID(c.ChainID, contractHname),
			EntryPointCode:   entryPoint,
			Transfer:         par.Transfer,
			Vars:             par.Args,
		}},
		Post: true,
	})
}