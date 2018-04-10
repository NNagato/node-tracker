package storage

import (
	"bytes"
	"log"
	"encoding/binary"
	"encoding/json"

	"github.com/boltdb/bolt"
	"github.com/Gin/node-tracker/common"
)

const (
	PATH string = "/home/ubuntu/go/src/github.com/Gin/node-tracker/storage/gin.db"

	NET_VERSION string = "net_version"
	NET_PEERCOUNT string = "net_peerCount"
	NET_LISTENING string = "net_listening"
	ETH_PROTOCOLVERSION string = "eth_protocolVersion"
	GAS_PRICE string = "eth_gasPrice"
	BLOCK_NUM string = "eth_blockNumber"
	GET_BALANCE string = "eth_getBalance"
	GET_STORAGE_AT string = "eth_getStorageAt"
	GET_BLOCKTXCOUNT_BY_NUM string = "eth_getBlockTransactionCountByNumber"
	GET_TX_COUNT string = "eth_getTransactionCount"
	GET_CODE string = "eth_getCode"
	SIGN string = "eth_sign"
	SEND_TX string = "eth_sendTransaction"
	SEND_RAW_TX string = "eth_sendRawTransaction"
	ETH_CALL string = "eth_call"
	ESTIMAT_GAS string = "eth_estimateGas"
	GET_BLOCK_BY_HASH string = "eth_getBlockByHash"
	GET_BLOCK_BY_NUM string = "eth_getBlockByNumber"
	GET_TX_BY_HASH string = "eth_getTransactionByHash"
	GET_TX_BY_BLOCK_HASH_AND_INDEX string = "eth_getTransactionByBlockHashAndIndex"
	GET_TX_BY_BLOCK_NUM_AND_INDEX string = "eth_getTransactionByBlockNumberAndIndex"
	GET_TX_RECEIPT string = "eth_getTransactionReceipt"
	GET_UNCLE_BY_BLOCK_HASH_AND_INDEX string = "eth_getUncleByBlockHashAndIndex"
	GET_UNCLE_BY_BLOCK_NUM_AND_INDEX string = "eth_getUncleByBlockNumberAndIndex"
	COMPILE_SOLIDITY string = "eth_compileSolidity"
	GET_LOG string = "eth_getLogs"
	ONE string = "1"
)

var ListBucket = []string{
	NET_VERSION,
	NET_PEERCOUNT,
	NET_LISTENING,
	ETH_PROTOCOLVERSION,
	GAS_PRICE,
	BLOCK_NUM,
	GET_BALANCE,
	GET_STORAGE_AT,
	GET_BLOCKTXCOUNT_BY_NUM,
	GET_TX_COUNT,
	GET_CODE,
	SIGN,
	SEND_TX,
	SEND_RAW_TX,
	ETH_CALL,
	ESTIMAT_GAS,
	GET_BLOCK_BY_HASH,
	GET_BLOCK_BY_NUM,
	GET_TX_BY_HASH,
	GET_TX_BY_BLOCK_HASH_AND_INDEX,
	GET_TX_BY_BLOCK_NUM_AND_INDEX,
	GET_TX_RECEIPT,
	GET_UNCLE_BY_BLOCK_HASH_AND_INDEX,
	GET_UNCLE_BY_BLOCK_NUM_AND_INDEX,
	COMPILE_SOLIDITY,
	GET_LOG,
	ONE,
}

type BoltStorage struct {
	db *bolt.DB
}

func NewStorage() *BoltStorage {
	var err error
	var db *bolt.DB
	db, err = bolt.Open(PATH, 0600, nil)
	if err != nil {
		panic(err)
	}

	db.Update(func(tx *bolt.Tx) error {
		tx.CreateBucket([]byte(NET_VERSION))
		tx.CreateBucket([]byte(NET_PEERCOUNT))
		tx.CreateBucket([]byte(NET_LISTENING))
		tx.CreateBucket([]byte(ETH_PROTOCOLVERSION))
		tx.CreateBucket([]byte(GAS_PRICE))
		tx.CreateBucket([]byte(BLOCK_NUM))
		tx.CreateBucket([]byte(GET_BALANCE))
		tx.CreateBucket([]byte(GET_STORAGE_AT))
		tx.CreateBucket([]byte(GET_BLOCKTXCOUNT_BY_NUM))
		tx.CreateBucket([]byte(GET_TX_COUNT))
		tx.CreateBucket([]byte(GET_CODE))
		tx.CreateBucket([]byte(SIGN))
		tx.CreateBucket([]byte(SEND_TX))
		tx.CreateBucket([]byte(SEND_RAW_TX))
		tx.CreateBucket([]byte(ETH_CALL))
		tx.CreateBucket([]byte(ESTIMAT_GAS))
		tx.CreateBucket([]byte(GET_BLOCK_BY_HASH))
		tx.CreateBucket([]byte(GET_BLOCK_BY_NUM))
		tx.CreateBucket([]byte(GET_TX_BY_HASH))
		tx.CreateBucket([]byte(GET_TX_BY_BLOCK_HASH_AND_INDEX))
		tx.CreateBucket([]byte(GET_TX_BY_BLOCK_NUM_AND_INDEX))
		tx.CreateBucket([]byte(GET_TX_RECEIPT))
		tx.CreateBucket([]byte(GET_UNCLE_BY_BLOCK_HASH_AND_INDEX))
		tx.CreateBucket([]byte(GET_UNCLE_BY_BLOCK_NUM_AND_INDEX))
		tx.CreateBucket([]byte(COMPILE_SOLIDITY))
		tx.CreateBucket([]byte(GET_LOG))
		tx.CreateBucket([]byte(ONE))

		return nil
	})
	storage := &BoltStorage{db}
	return storage
}

func uint64ToBytes(u uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, u)
	return b
}

func bytesToUint64(b []byte) uint64 {
	return binary.BigEndian.Uint64(b)
}

func (self *BoltStorage) StorageTimeResponse(allSaveData map[string][][]float64) error {
	var err error
	self.db.Update(func(tx *bolt.Tx) error {
		for rpc, d := range(allSaveData) {
			var dataJson []byte
			b := tx.Bucket([]byte(rpc))

			for _, data := range(d) {
				timeCount := uint64(data[0])
				responseTime := data[1]
				dataJson, err = json.Marshal(responseTime)
				if err != nil {
					log.Println(err)
					return err
				}
				err = b.Put(uint64ToBytes(timeCount), dataJson)
				if err != nil {
					return err
				}
			}
		}
		return err
	})
	return err
}

func (self *BoltStorage) GetTimeResponseData(fromTime uint64) ([]common.TimeResponse, error) {
	data := common.NewDataTimeResponse()
	var err error
	self.db.View(func(tx *bolt.Tx) error {
		for _, bucket := range(ListBucket) {
			queryData(tx, data, bucket, fromTime)
		}
		return nil
	})
	return data.GetData(), err
}

func queryData(tx *bolt.Tx, data *common.DataTimeResponse, bucket string, fromTime uint64) {
	// log.Println("fromTime: ", fromTime)
	result := make(map[uint64]float64)

	b := tx.Bucket([]byte(bucket))
	c := b.Cursor()
	max, _ := c.Last()
	min := uint64ToBytes(fromTime)
	// min, _ = c.First()

	for k, v := c.Seek(min); k != nil && bytes.Compare(k, max) <= 0; k, v = c.Next() {
		tickTime := bytesToUint64(k)
		var responseTime float64
		err := json.Unmarshal(v, &responseTime)
		if err != nil {
			log.Println(err)
		}
		result[tickTime] = responseTime
	}

	if len(result) > 0 {
		data.Add(bucket, result)
	}
}

func (self *BoltStorage) GetLatestVersion() float64 {
	latestTimeByte := uint64ToBytes(0)
	self.db.View(func(tx *bolt.Tx) error {
		for _, bucket := range(ListBucket) {
			b := tx.Bucket([]byte(bucket))
			c := b.Cursor()
			latest, _ := c.Last()
			if bytes.Compare(latestTimeByte, latest) < 0 {
				latestTimeByte = latest
			}
		}
		return nil
	})
	latestTime := bytesToUint64(latestTimeByte)
	return float64(latestTime)
}
