package EsConfig

import (
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
	"os"
)

func NewEsClient(esData *EsData,startLog bool) *elastic.Client {
	if esData == nil {
		panic("esConfigData is null")
	}
	var client *elastic.Client
	var err error
	if startLog {
		errorlog := log.New(os.Stdout, "", log.LstdFlags)
		client, err = elastic.NewClient(
			elastic.SetErrorLog(errorlog),
			elastic.SetURL(esData.Host),
			elastic.SetBasicAuth(esData.UserName, esData.PassWord))
	}else{
		client, err = elastic.NewClient(
			elastic.SetURL(esData.Host),
			elastic.SetBasicAuth(esData.UserName, esData.PassWord))
	}
	if err != nil {
		fmt.Println("es连接异常",err)
	}
	return client
}

type EsInterface interface {
	getEsData(...string) EsData
}

type EsData struct {
	Host        string `json:"host"`
	UserName    string `json:"user_name"`
	PassWord    string `json:"pass_word"`
	IndexDBName string `mapstructure:"index-db-name"`
}
