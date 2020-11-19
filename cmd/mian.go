package main

import (
	"github.com/Byfengfeng/es/esService"

	Es "github.com/Byfengfeng/es/esConfig"

	"fmt"
)

func main() {
	data := Es.EsData{
		"http://127.0.0.1:9200/",
		"user",
		"123456",
		"华山",
	}
	esClient := Es.NewEsClient(&data)
	queryAll := esService.QueryAll(esClient, Log{}, 15, data.IndexDBName)
	for i := range queryAll {
		fmt.Println(queryAll[i])
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
