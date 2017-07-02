package legacy

import (
	"fmt"
	"io"

	"chain/encoding/blockchain"
	"chain/errors"
	"chain/protocol/bc"
)

type (
	TxInput struct {
		AssetVersion  uint64
		ReferenceData []byte
		TypedInput

		// Unconsumed suffixes of the commitment and witness extensible
		// strings.
		CommitmentSuffix []byte
		WitnessSuffix    []byte
	}

	TypedInput interface {
		IsIssuance() bool
	}
)

var errBadAssetID = errors.New("asset ID does not match other issuance parameters")

func (t *TxInput) AssetAmount() bc.AssetAmount {
	if ii, ok := t.TypedInput.(*IssuanceInput); ok {
		assetID := ii.AssetID()
		return bc.AssetAmount{
			AssetId: &assetID,
			Amount:  ii.Amount,
		}
	}
	si := t.TypedInput.(*SpendInput)
	return si.AssetAmount
}

func (t *TxInput) AssetID() bc.AssetID {
	if ii, ok := t.TypedInput.(*IssuanceInput); ok {
		return ii.AssetID()
	}
	si := t.TypedInput.(*SpendInput)
	return *si.AssetId
}

func (t *TxInput) Amount() uint64 {
	if ii, ok := t.TypedInput.(*IssuanceInput); ok {
		return ii.Amount
	}
	si := t.TypedInput.(*SpendInput)
	return si.Amount
}

func (t *TxInput) ControlProgram() []byte {
	if si, ok := t.TypedInput.(*SpendInput); ok {
		return si.ControlProgram
	}
	return nil
}

func (t *TxInput) IssuanceProgram() []byte {
	if ii, ok := t.TypedInput.(*IssuanceInput); ok {
		return ii.IssuanceProgram
	}
	return nil
}

func (t *TxInput) Arguments() [][]byte {
	switch inp := t.TypedInput.(type) {
	case *IssuanceInput:
		return inp.Arguments
	case *SpendInput:
		return inp.Arguments
	}
	return nil
}

func (t *TxInput) SetArguments(args [][]byte) {
	switch inp := t.TypedInput.(type) {
	case *IssuanceInput:
		inp.Arguments = args
	case *SpendInput:
		inp.Arguments = args
	}
}

func (t *TxInput) readFrom(r *blockchain.Reader) (err error) {
	t.AssetVersion, err = blockchain.ReadVarint63(r)
	if err != nil {
		return err
	}

	var (
		ii      *IssuanceInput
		si      *SpendInput
		assetID bc.AssetID
	)

	t.CommitmentSuffix, err = blockchain.ReadExtensibleString(r, func(r *blockchain.Reader) error {
		if t.AssetVersion != 1 {
			return nil
		}
		var icType [1]byte
		_, err = io.ReadFull(r, icType[:])
		if err != nil {
			return errors.Wrap(err, "reading input commitment type")
		}
		switch icType[0] {
		case 0:
			ii = new(IssuanceInput)

			ii.Nonce, err = blockchain.ReadVarstr31(r)
			if err != nil {
				return err
			}
			_, err = assetID.ReadFrom(r)
			if err != nil {
				return err
			}
			ii.Amount, err = blockchain.ReadVarint63(r)
			if err != nil {
				return err
			} 
