package esUtils

import (
	"errors"
	"github.com/olivere/elastic/v7"
	"reflect"
	"sync"
	"time"
)

//get es data list
func GetDataList(res *elastic.SearchResult, err error, dataType interface{}) []interface{} {
	if err != nil {
		print(err.Error())
		return nil
	}
	list := make([]interface{}, 0)
	for _, item := range res.Each(reflect.TypeOf(dataType)) {
		t := item.(interface{})
		list = append(list, t)
	}
	return list
}

//get es one data
func GetDataOne(res *elastic.SearchResult, err error, dataType interface{}) interface{} {
	if err != nil {
		print(err.Error())
		return nil
	}
	for _, item := range res.Each(reflect.TypeOf(dataType)) {
		dataType = item.(interface{})
	}
	return dataType
}

//append query condition
func AppendCondition(kv map[string]string)[]*elastic.TermsQuery  {
	termsQuery := make([]*elastic.TermsQuery,0)
	for k, v := range kv {
		termsQuery = append(termsQuery,elastic.NewTermsQuery(k, v))
	}
	return termsQuery
}

func GetId()(id int64)  {
	id = node.GetId()
	return
}

var node, _ = NewWorker(1)

const (
	workerBits  uint8 = 10
	numberBits  uint8 = 12
	workerMax   int64 = -1 ^ (-1 << workerBits)
	numberMax   int64 = -1 ^ (-1 << numberBits)
	timeShift   uint8 = workerBits + numberBits
	workerShift uint8 = numberBits
	startTime   int64 = 1525705533000 // 如果在程序跑了一段时间修改了epoch这个值 可能会导致生成相同的ID
)

type Worker struct {
	mu        sync.Mutex
	timestamp int64
	workerId  int64
	number    int64
}

func NewWorker(workerId int64) (*Worker, error) {
	if workerId < 0 || workerId > workerMax {
		return nil, errors.New("Worker ID excess of quantity")
	}
	// 生成一个新节点
	return &Worker{
		timestamp: 0,
		workerId:  workerId,
		number:    0,
	}, nil
}

func (w *Worker) GetId() int64 {
	w.mu.Lock()
	defer w.mu.Unlock()
	now := time.Now().UnixNano() / 1e6
	if w.timestamp == now {
		w.number++
		if w.number > numberMax {
			for now <= w.timestamp {
				now = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		w.number = 0
		w.timestamp = now
	}
	ID := int64((now-startTime)<<timeShift | (w.workerId << workerShift) | (w.number))
	return ID
}
