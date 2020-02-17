package loopdb

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"time"

	"github.com/lightningnetwork/lnd/routing/route"
)

// LoopInContract contains the data that is serialized to persistent storage for
// pending loop in swaps.
type LoopInContract struct {
	SwapContract

	// SweepConfTarget specifies the targeted confirmation target for the
	// client sweep tx.
	HtlcConfTarget int32

	// LastHop is the last hop to use for the loop in swap (optional).
	LastHop *route.Vertex

	// ExternalHtlc specifies whether the htlc is published by an external
	// source.
	ExternalHtlc bool
}

// LoopIn is a combination of the contract and the updates.
type LoopIn struct {
	Loop

	Contract *LoopInContract
}

// LastUpdateTime returns the last update time of this swap.
func (s *LoopIn) LastUpdateTime() time.Time {
	lastUpdate := s.LastUpdate()
	if lastUpdate == nil {
		return s.Contract.InitiationTime
	}

	return lastUpdate.Time
}

// serializeLoopInContract serialize the loop in contract into a byte slice.
func serializeLoopInContract(swap *LoopInContract) (
	[]byte, error) {

	var b bytes.Buffer

	if err := binary.Write(&b, byteOrder, swap.InitiationTime.UnixNano()); err != nil {
		return nil, err
	}

	if err := binary.Write(&b, byteOrder, swap.Preimage); err != nil {
		return nil, err
	}

	if err := binary.Write(&b, byteOrder, swap.AmountRequested); err != nil {
		return nil, err
	}

	n, err := b.Write(swap.SenderKey[:])
	if err != nil {
		return nil, err
	}
	if n != keyLength {
		return nil, fmt.Errorf("sender key has invalid length")
	}

	n, err = b.Write(swap.ReceiverKey[:])
	if err != nil {
		return nil, err
	}
	if n != keyLength {
		return nil, fmt.Errorf("receiver key has invalid length")
	}

	if err := binary.Write(&b, byteOrder, swap.CltvExpiry); err != nil {
		return nil, err
	}

	if err := binary.Write(&b, byteOrder, swap.MaxMinerFee); err != nil {
		return nil, err
	}

	if err := binary.Write(&b, byteOrder, swap.MaxSwapFee); err != nil {
		return nil, err
	}

	if err := binary.Write(&b, byteOrder, swap.InitiationHeight); err != nil {
		return nil, err
	}

	if err := binary.Write(&b, byteOrder, swap.HtlcConfTarget); err != nil {
		return nil, err
	}

	var lastHop route.Vertex
	if swap.LastHop != nil {
		lastHop = *swap.LastHop
	}
	if err := binary.Write(&b, byteOrder, lastHop[:]); err != nil {
		return nil, err
	}

	if err := binary.Write(&b, byteOrder, swap.ExternalHtlc); err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

// deserializeLoopInContract deserializes the loop in contract from a byte slice.
func deserializeLoopInContract(value []byte) (*LoopInContract, error) {
	r := bytes.NewReader(value)

	contract := LoopInContract{}
	var err error
	var unixNano int64
	if err := binary.Read(r, byteOrder, &unixNano); err != nil {
		return nil, err
	}
	contract.InitiationTime = time.Unix(0, unixNano)

	if err := binary.Read(r, byteOrder, &contract.Preimage); err != nil {
		return nil, err
	}

	err = binary.Read(r, byteOrder, &contract.AmountRequested)
	if err != nil {
		return nil, err
	}

	n, err := r.Read(contract.SenderKey[:])
	if err != nil {
		return nil, err
	}
	if n != keyLength {
		return nil, fmt.Errorf("sender key has invalid length")
	}

	n, err = r.Read(contract.ReceiverKey[:])
	if err != nil {
		return nil, err
	}
	if n != keyLength {
		return nil, fmt.Errorf("receiver key has invalid length")
	}

	if err := binary.Read(r, byteOrder, &contract.CltvExpiry); err != nil {
		return nil, err
	}
	if err := binary.Read(r, byteOrder, &contract.MaxMinerFee); err != nil {
		return nil, err
	}

	if err := binary.Read(r, byteOrder, &contract.MaxSwapFee); err != nil {
		return nil, err
	}

	if err := binary.Read(r, byteOrder, &contract.InitiationHeight); err != nil {
		return nil, err
	}

	if err := binary.Read(r, byteOrder, &contract.HtlcConfTarget); err != nil {
		return nil, err
	}

	var lastHop route.Vertex
	if err := binary.Read(r, byteOrder, lastHop[:]); err != nil {
		return nil, err
	}
	var noLastHop route.Vertex
	if lastHop != noLastHop {
		contract.LastHop = &lastHop
	}

	if err := binary.Read(r, byteOrder, &contract.ExternalHtlc); err != nil {
		return nil, err
	}

	return &contract, nil
}
