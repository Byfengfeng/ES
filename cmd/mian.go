package main

import (
	EsConfig "github.com/Byfengfeng/es/esConfig"
	"github.com/Byfengfeng/es/esService"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	data := EsConfig.EsData{
		"http://122.228.10.99:9200/",
		"esuser",
		"123456",
		"log",
	}
	//rangeTime := esService.RangeTime{1605593790, 1605593797}
	esClient := EsConfig.NewEsClient(&data)
	//ke := make(map[string]string,0)
	//ke["tid"] = "1"
	//ke["uid"] = "456"
	//queryAll := esService.QueryTimeLog(esClient, Log{}, 15, data.IndexDBName,"log",ke,"time",&rangeTime)
	//for i := range queryAll {
	//	fmt.Println(queryAll[i])
	//}indexDB string,esType string, id int64
	per := struct{name string}{name: "123"}
	wg := sync.WaitGroup{}
	count := 3000
	wg.Add(count)
	var i int64 = 0
	for i < 3000 {
		//go func() {
			esService.Save(esClient,per,"1","1",i)
		//	wg.Done()
		//}()
		atomic.AddInt64(&i,1)
	}
	wg.Wait()
	time.Sleep(time.Second*20)
	//indexNames := esService.GetIndexNames(esClient)
	//for _,indexName := range indexNames {
	//	fmt.Println(indexName)
	//}
	//fmt.Println(zap.DebugLevel)
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
