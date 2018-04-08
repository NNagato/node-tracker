package common

import (
	"sync"
)

type TimeResponse struct {
	RPC    string
	Data   map[uint64]float64
}

type DataTimeResponse struct {
	mu   *sync.RWMutex
	Data []TimeResponse
}

func NewDataTimeResponse() *DataTimeResponse {
	return &DataTimeResponse{
		mu: &sync.RWMutex{},
	}
}

func (self *DataTimeResponse) Add(bucket string, timeData map[uint64]float64) {
	self.mu.Lock()

	timeResponse := TimeResponse{
		RPC: bucket,
		Data: timeData,
	}
	tempData := self.Data
	tempData = append(tempData, timeResponse)
	self.Data = tempData
	self.mu.Unlock()
}

func (self *DataTimeResponse) GetData() []TimeResponse {
	self.mu.Lock()
	defer self.mu.Unlock()
	
	data := self.Data
	return data
}
