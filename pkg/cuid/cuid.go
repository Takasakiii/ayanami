package cuid

import (
	"fmt"
	"github.com/nrednav/cuid2"
)

type Cuid struct {
	cuid func() string
}

func NewCuid() (*Cuid, error) {
	cuid, err := cuid2.Init()
	if err != nil {
		return nil, fmt.Errorf("cuid init failed: %v", err)
	}

	return &Cuid{
		cuid: cuid,
	}, nil
}
