package types

import (
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	// KeyEpochIdentifier is the key for the epoch identifier
	KeyEpochIdentifier = []byte("EpochIdentifier")
)

const (
	defaultEpochIdentifier = "hour"
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(epcohIdentifier string) Params {
	return Params{
		EpochIdentifier: epcohIdentifier,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(defaultEpochIdentifier)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyEpochIdentifier, &p.EpochIdentifier, func(_ interface{}) error { return nil }),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if p.EpochIdentifier == "" {
		return ErrEmptyEpochIdentifier
	}
	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}
