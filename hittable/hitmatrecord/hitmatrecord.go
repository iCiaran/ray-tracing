package hitmatrecord

import (
	"github.com/iCiaran/ray-tracing/hittable/hitrecord"
	"github.com/iCiaran/ray-tracing/material"
)

type HitMatRecord struct {
	Rec *hitrecord.HitRecord
	Mat material.Material
}

func New() *HitMatRecord {
	return &HitMatRecord{hitrecord.New(), nil}
}
