package es_config

import (
	"fmt"
	"github.com/Byfengfeng/es/es_service"
	"github.com/olivere/elastic/v7"
	"log"
	"os"
)

func NewEsClient(esData *EsData,startLog bool) (esService es_service.EsService) {
	if esData == nil {
		panic("esConfigData is null")
	}
	var err error
	if startLog {
		errorlog := log.New(os.Stdout, "", log.LstdFlags)
		esService.Client, err = elastic.NewClient(
			elastic.SetErrorLog(errorlog),
			elastic.SetURL(esData.Host),
			elastic.SetBasicAuth(esData.UserName, esData.PassWord))
	}else{
		esService.Client, err = elastic.NewClient(
			elastic.SetURL(esData.Host),
			elastic.SetBasicAuth(esData.UserName, esData.PassWord))
	}
	esService.Index = esData.IndexDBName
	esService.EsType = esData.EsType
	if err != nil {
		fmt.Println("es连接异常",err)
	}

	return
}

type EsInterface interface {
	getEsData(...string) EsData
}

type EsData struct {
	Host        string `json:"host"`
	UserName    string `json:"user_name"`
	PassWord    string `json:"pass_word"`
	IndexDBName string `mapstructure:"index-db-name"`
	EsType		string `json:"es_type"`
}
