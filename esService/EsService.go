package esService

import (
	"fmt"
	"github.com/Byfengfeng/es/esUtils"
	"strconv"

	"context"
	"github.com/olivere/elastic/v7"
)

type RangeTime struct {
	MinTime int64
	MaxTime int64
}

//存储
func Save(client *elastic.Client, data interface{}, indexDB string, id int64) {
	//使用结构体
	_, err := client.Index().
		Index(indexDB).
		Type("employee").
		Id(strconv.FormatInt(id, 10)).
		BodyJson(data).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
}

//删除
func Remove(client *elastic.Client, indexDB string, id string) {
	res, err := client.Delete().Index(indexDB).
		Id(id).
		Type("employee").
		Do(context.Background())
	if err != nil {
		println(err.Error())
		return
	}
	fmt.Printf("delete result %s\n", res.Result)
}

//查询所有数据
func QueryAll(client *elastic.Client, data interface{}, queryNum int, indexDB string) []interface{} {
	var res *elastic.SearchResult
	var err error
	res, err = client.Search(indexDB).Type("employee").Size(queryNum).Do(context.Background())
	if err != nil {
		println(err.Error())
	}
	return esUtils.GetDataList(res, err, data)
}

//根据时间范围查询数据
func QueryRange(client *elastic.Client, data interface{}, queryNum int,
	indexDB string, rangeTimeKey string, rangeTimeValue *RangeTime) []interface{} {
	var res *elastic.SearchResult
	var err error

	boolSearch := elastic.NewBoolQuery().
		Filter(elastic.NewRangeQuery(rangeTimeKey).Gte(rangeTimeValue.MinTime).Lte(rangeTimeValue.MaxTime))
	res, err = client.Search(indexDB).Type("employee").Query(boolSearch).Size(queryNum).
		Do(context.Background())

	if err != nil {
		println(err.Error())
	}

	return esUtils.GetDataList(res, err, data)
}

//查询单条数据

func QueryOne(client *elastic.Client, data interface{}, indexDB string, key string, value string) interface{} {
	var res *elastic.SearchResult
	var err error

	boolSearch := elastic.NewBoolQuery().
		Filter(elastic.NewTermsQuery(key, value))
	res, err = client.Search(indexDB).Type("employee").Query(boolSearch).Size(1).
		Do(context.Background())

	if err != nil {
		println(err.Error())
	}

	return esUtils.GetDataOne(res, err, data)
}

//根据时间范围查询数据 TODO
func QueryLog(client *elastic.Client, data interface{}, queryNum int,
	indexDB string, kv map[string]string, rangeTimeKey string, rangeTimeValue *RangeTime) []interface{} {
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

	res, err = client.Search(indexDB).Type("employee").Query(boolSearch).Size(queryNum).
		Do(context.Background())

	if err != nil {
		println(err.Error())
	}

	return esUtils.GetDataList(res, err, data)
}

//根据时间范围及服务名查询数据
func QueryTimeLog(client *elastic.Client, data interface{}, queryNum int,
	indexDB string,kv map[string]string, rangeTimeKey string, rangeTimeValue *RangeTime) []interface{} {
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

	res, err = client.Search(indexDB).Type("employee").Query(boolSearch).Size(queryNum).
		Do(context.Background())

	if err != nil {
		println(err.Error())
	}

	return esUtils.GetDataList(res, err, data)
}
