package main

import (
	"fmt"
	EsConfig "github.com/Byfengfeng/es/esConfig"
	"github.com/Byfengfeng/es/esService"
)

func main() {
	data := EsConfig.EsData{
		"http://127.0.0.1:9200/",
		"user",
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
	//}
	indexNames := esService.GetIndexNames(esClient)
	for _,indexName := range indexNames {
		fmt.Println(indexName)
	}
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
