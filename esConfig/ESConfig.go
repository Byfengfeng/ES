package EsConfig

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/olivere/elastic/v7"
)

func NewEsClient(esData *EsData) *elastic.Client {
	if esData == nil {
		panic("esConfigData is null")
	}

	errorlog := log.New(os.Stdout, "", log.LstdFlags)
	client, err := elastic.NewClient(elastic.SetErrorLog(errorlog),
		elastic.SetURL(esData.Host),
		elastic.SetBasicAuth(esData.UserName, esData.PassWord))
	if err != nil {
		panic(err)
	}
	info, code, err := client.Ping(esData.Host).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	esversion, err := client.ElasticsearchVersion(esData.Host)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Elasticsearch version %s\n", esversion)
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
