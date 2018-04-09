package collector

import (
	"strings"
	"encoding/json"
	"strconv"
	"time"
	"log"

	"github.com/hpcloud/tail"
	"github.com/Gin/node-tracker/storage"
	"github.com/Gin/node-tracker/common"
)

const (
	tickDistance float64 = 60
	timeBack int64 = 43200 // 12 hours in second
)

type JSON struct {
	TimeHandled float64
	TimeUTC     string `json:"time_utc"`
	RequestTime string `json:"request_time"`
	RequestBody string `json:"request_body"`
}

type SaveData struct {
	Time         int64
	TimeResponse float64
	Method       string 
}

type Collector struct {
	db BoltStorage
}

func NewCollector() *Collector {
	storage := storage.NewStorage()
	return &Collector{
		db: storage,
	}
}

// const file_log string = "/home/gin/Gin/Testspace/golang/write-and-read-file/file.txt"
const file_log string = "/var/log/nginx/rpc.log"

func (self *Collector) GetLog() {
	latestVersion := self.db.GetLatestVersion()
	t, err := tail.TailFile(file_log, tail.Config{
		Follow: true, 
		ReOpen: true})

	if err != nil {
		log.Println(err)
		panic(err)
	}

	var jsonData JSON
	var text string
	var jsonArray []JSON
	for line := range t.Lines {
		text = strings.Replace(line.Text, "\\x22", "", -1)
		err = json.Unmarshal([]byte(text), &jsonData)
		if err != nil {
			log.Println(text)
			log.Println("error marshal: ", err)
		}
		timeRequest, _ := time.Parse(time.RFC3339, jsonData.TimeUTC)
		jsonData.TimeHandled = float64(timeRequest.Unix())
		if jsonData.RequestBody != "-" && jsonData.TimeHandled >= latestVersion {
			jsonArray = append(jsonArray, jsonData)
			if len(jsonArray) == 2000 {
				log.Println("new pack to save")
				go self.HandleJsonArray(jsonArray)
				jsonArray = []JSON{}
			}
		}
	}
}

func (self *Collector) HandleJsonArray(jsonArray []JSON) {
	allSaveData := make(map[string][][]float64)
	for _, json := range jsonArray {
		bodyArray := strings.Split(json.RequestBody, ",")
		methodArray := strings.Split(bodyArray[2], ":")
		method := methodArray[1]
		// log.Println("method: ", method, len(allSaveData[method]))
		timeResponse, _ := strconv.ParseFloat(json.RequestTime, 64)

		allSaveData[method] = append(allSaveData[method], []float64{json.TimeHandled, timeResponse})		
	}
	trueStoreData := makeTrueData(allSaveData)
	self.db.StorageTimeResponse(trueStoreData)
}

func makeTrueData(allSaveData map[string][][]float64) map[string][][]float64 {
	trueStoreData := make(map[string][][]float64)

	for rpc, saveData := range allSaveData {
		var sum float64
		var tickTime float64
		var countIndex int

		for i, data := range saveData {
			if len(trueStoreData[rpc]) == 0 && i == 0 {
				var firstPackTick float64 = 0
				if (data[0] / tickDistance) - float64(uint64(data[0] / tickDistance)) == 0 {
					firstPackTick = data[0]
				} else {
					firstPackTick = float64(int64((data[0] + tickDistance)/tickDistance)) * tickDistance	
				}
				trueStoreData[rpc] = append(trueStoreData[rpc], []float64{firstPackTick, data[1]})
				tickTime = firstPackTick + tickDistance
				countIndex = 1
				sum = 0
				continue
			} 

			if data[0] <= tickTime {
				sum += data[1]
				if i == len(saveData) - 1 {
					avg := sum / float64(i - countIndex + 1)
					lenArray := len(trueStoreData[rpc]) - 1
					oldVal := trueStoreData[rpc][lenArray][1]
					trueStoreData[rpc][lenArray][1] = (avg + oldVal)/float64(2)
				}
			} else {
				if i == countIndex {
					aloneTick := float64(int64((data[0])/tickDistance)) * tickDistance
					trueStoreData[rpc] = append(trueStoreData[rpc], []float64{aloneTick, data[1]})
					tickTime = aloneTick + tickDistance
					countIndex += 1
				}
				if i > countIndex {
					avg := sum / float64(i - countIndex)
					trueStoreData[rpc] = append(trueStoreData[rpc], []float64{tickTime, avg})
					tickTime += tickDistance
					countIndex = i
					sum = data[1]
				}
			}
		}
	}
	return trueStoreData
}

func (self *Collector) GetData() ([]common.TimeResponse, error) {
	timeNow := time.Now().UTC().Unix()
	fromTime := uint64(timeNow - timeBack)
	// fromTime = uint64(1523082663)
	result, err := self.db.GetTimeResponseData(fromTime)
	if err != nil {
		log.Println(err)
		return []common.TimeResponse{}, err
	}
	return result, nil
}
