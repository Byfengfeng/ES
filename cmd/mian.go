package main

import (
	"fmt"
	EsConfig "github.com/Byfengfeng/es/es_config"
	//"sync"
)

//var luck sync.Mutex
var n = 0
func main() {
	data := EsConfig.EsData{
		"http://127.0.0.1:9200/",
		"esuser12",
		"123456111",
		"s",
		"1233",
	}
	fmt.Println(data)
	per := struct{name string}{name: "123"}
	//rangeTime := es_service.RangeTime{1605593790, 1605593797}
	esClient := EsConfig.NewEsClient(&data,true)
	esClient.SaveAll(&per)
	//bulk := esClient.Bulk()
	//aulk := esClient.Bulk()
	//for i:=0; i<= 20000;i++ {
	//	per := struct{name string}{name: "123"}
	//	req := elastic.NewBulkIndexRequest().
	//		Index("1").
	//		Type("1").
	//		Id(strconv.FormatInt(utils.GetId(), 10)).
	//		Doc(per)
	//	if n > 0{
	//		aulk.Add(req)
	//	}

		//luck.Lock()
		//n = 1
		//bulk.Add(req)
		//if bulk.NumberOfActions() == 1000 {
		//	bulk.Do(context.Background())
		//}
		//if aulk.NumberOfActions() > 0{
		//	aulk.Do(context.Background())
		//}
		//n = 0
		//luck.Unlock()
	//}

	//bulk.Do(context.Background())
	//bulk.Reset()
	//fmt.Println(bulk)
	//request := elastic.BulkableRequest()


	//ke := make(map[string]string,0)
	//ke["tid"] = "1"
	//ke["uid"] = "456"
	//queryAll := es_service.QueryTimeLog(esClient, Log{}, 15, data.IndexDBName,"log",ke,"time",&rangeTime)
	//for i := range queryAll {
	//	fmt.Println(queryAll[i])
	//}indexDB string,esType string, id int64
	//wg := sync.WaitGroup{}
	//var list map[string]string
	//count := 10000
	//wg.Add(count)
	//lock := sync.Mutex{}
	////startTime := time.Now().UnixNano() / 1e6
	//for i:=0; i < count; i++{
	//	go func() {
	//		lock.Lock()
	//		a := Ads{
	//			name: fmt.Sprint(utils.GetId()),
	//		}
	//		marshal, _ := json.Marshal(a)
	//		list[a.name] = string(marshal)
	//		lock.Unlock()
	//		wg.Done()
	//	}()
	//}
	//wg.Wait()
	//es_service.SaveAll(esClient,"1","1",utils.GetId(),&list)
	//fmt.Println("sleep")
	//fmt.Println(time.Now().UnixNano() / 1e6 - startTime)
	//time.Sleep(time.Second*20000)
	//fmt.Println(es_service.Num)
	//indexNames := es_service.GetIndexNames(esClient)
	//for _,indexName := range indexNames {
	//	fmt.Println(indexName)
	//}
	//fmt.Println(zap.DebugLevel)
}

type Ads struct{
	name string
}

type ByTime []Log

//排序
func (a ByTime) Len() int           { return len(a) }
func (a ByTime) Less(i, j int) bool { return a[i].Time < a[j].Time }
func (a ByTime) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type Log struct {
	Time       int64  `json:"time"`
	LogLevel   string `json:"log_level"`
	Src        string `json:"src"`
	ServerName string `json:"server_name"`
	Nick       string `json:"nick"`
	Age        int8   `json:"age"`
	Uid        int64  `json:"uid"`
	Tid        int64
}

func bules()  {

}
