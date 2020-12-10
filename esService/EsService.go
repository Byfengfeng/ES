package esService

import (
	"context"
	"fmt"
	"github.com/Byfengfeng/es/esUtils"
	"github.com/olivere/elastic/v7"
	"strconv"
)

type RangeTime struct {
	MinTime int64
	MaxTime int64
}


//存储
func Save(client *elastic.Client, data interface{}, indexDB string,esType string, id int64) {
	//使用结构体
	_, err := client.Index().
		Index(indexDB).
		Type(esType).
		Id(strconv.FormatInt(id, 10)).
		BodyJson(data).
		Do(context.Background())
	if err != nil {
		Save(client,data,indexDB,esType,id)
	}
}

//存储多个
func SaveAll(client *elastic.Client, indexDB string,esType string, id int64,datas ...interface{},) {
	bulk := client.Bulk()
	if len(datas) > 0{
		for _,v := range datas {
			 data := elastic.NewBulkIndexRequest().
				Index(indexDB).
				Type(esType).
				Id(strconv.FormatInt(id, 10)).
				Doc(v)
			bulk.Add(data)
		}
		_, err := bulk.Do(context.Background())
		if err != nil {
			SaveAll(client,indexDB,esType,id,datas)
		}
	}

}



//删除
func Remove(client *elastic.Client, indexDB string,esType string, id string) {
	res, err := client.Delete().Index(indexDB).
		Id(id).
		Type(esType).
		Do(context.Background())
	if err != nil {
		println(err.Error())
		return
	}
	fmt.Printf("delete result %s\n", res.Result)
}

//查询所有数据
func QueryAll(client *elastic.Client, data interface{}, queryNum int, indexDB string,esType string) []interface{} {
	var res *elastic.SearchResult
	var err error
	res, err = client.Search(indexDB).Type(esType).Size(queryNum).Do(context.Background())
	if err != nil {
		println(err.Error())
	}
	return esUtils.GetDataList(res, err, data)
}

//根据时间范围查询数据
func QueryRange(client *elastic.Client, data interface{}, queryNum int,
	indexDB string,esType string, rangeTimeKey string, rangeTimeValue *RangeTime) []interface{} {
	var res *elastic.SearchResult
	var err error

	boolSearch := elastic.NewBoolQuery().
		Filter(elastic.NewRangeQuery(rangeTimeKey).Gte(rangeTimeValue.MinTime).Lte(rangeTimeValue.MaxTime))
	res, err = client.Search(indexDB).Type(esType).Query(boolSearch).Size(queryNum).
		Do(context.Background())

	if err != nil {
		println(err.Error())
	}

	return esUtils.GetDataList(res, err, data)
}

//查询单条数据

func QueryOne(client *elastic.Client, data interface{}, indexDB string,esType string, key string, value string,size int) interface{} {
	var res *elastic.SearchResult
	var err error

	boolSearch := elastic.NewBoolQuery().
		Filter(elastic.NewTermsQuery(key, value))
	res, err = client.Search(indexDB).Type(esType).Query(boolSearch).Size(size).
		Do(context.Background())

	if err != nil {
		println(err.Error())
	}

	return esUtils.GetDataOne(res, err, data)
}

//根据索引查询相匹配的数据条
func QueryOneList(client *elastic.Client, data interface{}, indexDB string,esType string, key string, value string,size int) []interface{} {
	var res *elastic.SearchResult
	var err error

	boolSearch := elastic.NewBoolQuery().
		Filter(elastic.NewTermsQuery(key, value))
	res, err = client.Search(indexDB).Type(esType).Query(boolSearch).Size(size).
		Do(context.Background())

	if err != nil {
		println(err.Error())
	}

	return esUtils.GetDataList(res, err, data)
}

//根据时间范围查询数据 TODO
func QueryLog(client *elastic.Client, data interface{}, queryNum int,
	indexDB string,esType string, kv map[string]string, rangeTimeKey string, rangeTimeValue *RangeTime) []interface{} {
	var res *elastic.SearchResult
	var err error

	boolSearch := elastic.NewBoolQuery().
		Filter(elastic.NewRangeQuery(rangeTimeKey).Gte(rangeTimeValue.MinTime).Lte(rangeTimeValue.MaxTime))

	if len(kv) > 0{
		condition := esUtils.AppendCondition(kv)
		for _,v := range condition {
			boolSearch.Filter(v)
		}
	}

	res, err = client.Search(indexDB).Type(esType).Query(boolSearch).Size(queryNum).
		Do(context.Background())

	if err != nil {
		println(err.Error())
	}

	return esUtils.GetDataList(res, err, data)
}

//根据时间范围及服务名查询数据
func QueryTimeLog(client *elastic.Client, data interface{}, queryNum int,
	indexDB string,esType string,kv map[string]string, rangeTimeKey string, rangeTimeValue *RangeTime) []interface{} {
	var res *elastic.SearchResult
	var err error
	boolSearch := elastic.NewBoolQuery().
		Filter(elastic.NewRangeQuery(rangeTimeKey).Gte(rangeTimeValue.MinTime).Lte(rangeTimeValue.MaxTime))
	if len(kv) > 0{
		condition := esUtils.AppendCondition(kv)
		for _,v := range condition {
			boolSearch.Filter(v)
		}
	}

	res, err = client.Search(indexDB).Type(esType).Query(boolSearch).Size(queryNum).
		Do(context.Background())

	if err != nil {
		println(err.Error())
	}

	return esUtils.GetDataList(res, err, data)
}

//查询es上面所有索引库
func GetIndexNames(client *elastic.Client)[]string  {
	names, err := client.IndexNames()
	if err != nil {
		println(err.Error())
	}
	return names
}
