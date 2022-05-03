package transaction

import (
	"math/big"

	iotago "github.com/iotaledger/iota.go/v3"
	"github.com/iotaledger/wasp/packages/cryptolib"
	"github.com/iotaledger/wasp/packages/iscp"
	"github.com/iotaledger/wasp/packages/parameters"
	"github.com/iotaledger/wasp/packages/util"
	"golang.org/x/xerrors"
)

type NewRequestTransactionParams struct {
	SenderKeyPair                *cryptolib.KeyPair
	SenderAddress                iotago.Address // might be different from the senderKP address (when sending as NFT or alias)
	UnspentOutputs               iotago.OutputSet
	UnspentOutputIDs             iotago.OutputIDs
	Request                      *iscp.RequestParameters
	NFT                          *iscp.NFT
	L1                           *parameters.L1
	DisableAutoAdjustDustDeposit bool // if true, the minimal dust deposit won't be adjusted automatically
}

// NewRequestTransaction creates a transaction including one or more requests to a chain.
// Empty assets in the request data defaults to 1 iota, which later is adjusted to the dust minimum
// Assumes all UnspentOutputs and corresponding UnspentOutputIDs can be used as inputs, i.e. are
// unlockable for the sender address
func NewRequestTransaction(par NewRequestTransactionParams) (*iotago.Transaction, error) {
	outputs := iotago.Outputs{}
	sumIotasOut := uint64(0)
	sumTokensOut := make(map[iotago.NativeTokenID]*big.Int)
	sumNFTsOut := make(map[iotago.NFTID]bool)

	req := par.Request

	// create outputs, sum totals needed
	assets := req.FungibleTokens
	if assets == nil {
		// if assets not specified, the minimum dust deposit will be adjusted by vmtxbuilder.MakeBasicOutput
		assets = &iscp.FungibleTokens{}
	}
	var out iotago.Output
	// will adjust to minimum dust deposit
	out = MakeBasicOutput(
		req.TargetAddress,
		par.SenderAddress,
		assets,
		&iscp.RequestMetadata{
			SenderContract: 0,
			TargetContract: req.Metadata.TargetContract,
			EntryPoint:     req.Metadata.EntryPoint,
			Params:         req.Metadata.Params,
			Allowance:      req.Metadata.Allowance,
			GasBudget:      req.Metadata.GasBudget,
		},
		req.Options,
		par.L1.RentStructure(),
		par.DisableAutoAdjustDustDeposit,
	)
	if par.NFT != nil {
		out = NftOutputFromBasicOutput(out.(*iotago.BasicOutput), par.NFT)
	}

	requiredDustDeposit := out.VBytes(par.L1.RentStructure(), nil)
	if out.Deposit() < requiredDustDeposit {
		return nil, xerrors.Errorf("%v: available %d < required %d iotas",
			ErrNotEnoughIotasForDustDeposit, out.Deposit(), requiredDustDeposit)
	}
	outputs = append(outputs, out)
	sumIotasOut += out.Deposit()
	sumTokensOut = addNativeTokens(sumTokensOut, out)
	if par.NFT != nil {
		sumNFTsOut[par.NFT.ID] = true
	}

	outputs, sumIotasOut, sumTokensOut, sumNFTsOut = updateOutputsWhenSendingOnBehalfOf(par, outputs, sumIotasOut, sumTokensOut, sumNFTsOut)

	inputIDs, remainder, err := computeInputsAndRemainder(par.SenderKeyPair.Address(), sumIotasOut, sumTokensOut, sumNFTsOut, par.UnspentOutputs, par.UnspentOutputIDs, par.L1.RentStructure())
	if err != nil {
		return nil, err
	}

	if remainder != nil {
		outputs = append(outputs, remainder)
	}

	inputsCommitment := inputIDs.OrderedSet(par.UnspentOutputs).MustCommitment()
	return CreateAndSignTx(inputIDs, inputsCommitment, outputs, par.SenderKeyPair, par.L1.Protocol.NetworkID())
}

func outputMatchesSendAsAddress(output iotago.Output, oID iotago.OutputID, address iotago.Address) bool {
	switch o := output.(type) {
	case *iotago.NFTOutput:
		if address.Equal(util.NFTIDFromNFTOutput(o, oID).ToAddress()) {
			return true
		}
	case *iotago.AliasOutput:
		if address.Equal(o.AliasID.ToAddress()) {
			return true
		}
	}
	return false
}

func addNativeTokens(sumTokensOut map[iotago.NativeTokenID]*big.Int, out iotago.Output) map[iotago.NativeTokenID]*big.Int {
	for _, nt := range out.NativeTokenSet() {
		s, ok := sumTokensOut[nt.ID]
		if !ok {
			s = new(big.Int)
		}
		s.Add(s, nt.Amount)
		sumTokensOut[nt.ID] = s
	}
	return sumTokensOut
}

func updateOutputsWhenSendingOnBehalfOf(
	par NewRequestTransactionParams,
	outputs iotago.Outputs,
	sumIotasOut uint64,
	sumTokensOut map[iotago.NativeTokenID]*big.Int,
	sumNFTsOut map[iotago.NFTID]bool,
) (
	iotago.Outputs,
	uint64,
	map[iotago.NativeTokenID]*big.Int,
	map[iotago.NFTID]bool,
) {
	if par.SenderAddress.Equal(par.SenderKeyPair.Address()) {
		return outputs, sumIotasOut, sumTokensOut, sumNFTsOut
	}
	// sending request "on behalf of" (need NFT or alias output as input/output)

	for _, output := range outputs {
		if outputMatchesSendAsAddress(output, iotago.OutputID{}, par.SenderAddress) {
			// if already present in the outputs, no need to do anything
			return outputs, sumIotasOut, sumTokensOut, sumNFTsOut
		}
	}
	for outID, out := range par.UnspentOutputs {
		// find the output that matches the "send as" address
		if !outputMatchesSendAsAddress(out, outID, par.SenderAddress) {
			continue
		}
		if nftOut, ok := out.(*iotago.NFTOutput); ok {
			if nftOut.NFTID.Empty() {
				// if this is the first time the NFT output transitions, we need to fill the correct NFTID
				nftOut.NFTID = iotago.NFTIDFromOutputID(outID)
			}
			sumNFTsOut[nftOut.NFTID] = true
		}
		// found the needed output
		outputs = append(outputs, out)
		sumIotasOut += out.Deposit()
		sumTokensOut = addNativeTokens(sumTokensOut, out)
		return outputs, sumIotasOut, sumTokensOut, sumNFTsOut
	}
	panic("unable to build tx, 'sendAs' output not found")
}
