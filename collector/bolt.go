package collector

import (
	"github.com/Gin/node-tracker/common"
)

type BoltStorage interface {
	StorageTimeResponse(allSaveData map[string][][]float64) error
	GetTimeResponseData(fromTime uint64) ([]common.TimeResponse, error)
	GetLatestVersion() float64
}