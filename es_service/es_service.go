package es_service

import (
	"context"
	"fmt"
	"github.com/Byfengfeng/es/utils"
	"github.com/olivere/elastic/v7"
)

type RangeTime struct {
	MinTime int64
	MaxTime int64
}

type EsService struct {
	*elastic.Client
	Index,EsType string
}

//存储
func (e *EsService) Save(data interface{}) {
	//使用结构体
	_,err := e.Client.Index().
		Index(e.Index).
		Type(e.EsType).
		BodyJson(data).
		Do(context.Background())
	if err != nil {
		e.Save(data)
	}
}

//存储多个
func (e *EsService) SaveAll(datas ...interface{}) {
	bulk := e.Client.Bulk()
	if len(datas) > 0{
		for _,v := range datas {
			req := elastic.NewBulkIndexRequest().
				Index(e.Index).
				Type(e.EsType).
				Doc(v)
			bulk.Add(req)
			if bulk.NumberOfActions() > 100000 {
				_, err := bulk.Do(context.Background())
				if err != nil {
					fmt.Println(err)
					e.SaveAll(datas...)
				}
			}
		}
		_, err := bulk.Do(context.Background())
		if err != nil {
			e.SaveAll(datas...)
		}
	}
}

//删除
func (e *EsService) Remove(id string) {
	res,err := e.Client.Delete().
		Index(e.Index).
		Type(e.EsType).
		Id(id).
		Do(context.Background())
	if err != nil {
		println(err.Error())
		return
	}
	fmt.Printf("delete result %s\n", res.Result)
}

//查询所有数据
func (e *EsService) QueryAll(data interface{}, queryNum int) []interface{} {
	var res *elastic.SearchResult
	var err error
	res, err = e.Client.Search(e.Index).Type(e.EsType).Size(queryNum).Do(context.Background())
	if err != nil {
		println(err.Error())
	}
	return utils.GetDataList(res, err, data)
}

//根据时间范围查询数据
func (e *EsService) QueryRange(data interface{}, queryNum int,rangeTimeKey string, rangeTimeValue *RangeTime) []interface{} {
	var res *elastic.SearchResult
	var err error

	boolSearch := elastic.NewBoolQuery().
		Filter(elastic.NewRangeQuery(rangeTimeKey).Gte(rangeTimeValue.MinTime).Lte(rangeTimeValue.MaxTime))
	res, err = e.Client.Search(e.Index).Type(e.EsType).Query(boolSearch).Size(queryNum).
		Do(context.Background())

	if err != nil {
		println(err.Error())
	}

	return utils.GetDataList(res, err, data)
}

//查询单条数据
func (e *EsService) QueryOne(data interface{}, key string, value string,size int) interface{} {
	var res *elastic.SearchResult
	var err error

	boolSearch := elastic.NewBoolQuery().
		Filter(elastic.NewTermsQuery(key, value))
	res, err = e.Client.Search(e.Index).Type(e.EsType).Query(boolSearch).Size(size).
		Do(context.Background())

	if err != nil {
		println(err.Error())
	}

	return utils.GetDataOne(res, err, data)
}

//根据索引查询相匹配的数据条
func (e *EsService) QueryOneList(data interface{}, key string, value string,size int) []interface{} {
	var res *elastic.SearchResult
	var err error

	boolSearch := elastic.NewBoolQuery().
		Filter(elastic.NewTermsQuery(key, value))
	res, err = e.Client.Search(e.Index).Type(e.EsType).Query(boolSearch).Size(size).
		Do(context.Background())

	if err != nil {
		println(err.Error())
	}

	return utils.GetDataList(res, err, data)
}

//根据时间范围查询数据 TODO
func (e *EsService) QueryLog(data interface{}, queryNum int,
	kv map[string]string, rangeTimeKey string, rangeTimeValue *RangeTime) []interface{} {
	var res *elastic.SearchResult
	var err error

	boolSearch := elastic.NewBoolQuery().
		Filter(elastic.NewRangeQuery(rangeTimeKey).Gte(rangeTimeValue.MinTime).Lte(rangeTimeValue.MaxTime))

	if len(kv) > 0{
		condition := utils.AppendCondition(kv)
		for _,v := range condition {
			boolSearch.Filter(v)
		}
	}

	res, err = e.Client.Search(e.Index).Type(e.EsType).Query(boolSearch).Size(queryNum).
		Do(context.Background())

	if err != nil {
		println(err.Error())
	}

	return utils.GetDataList(res, err, data)
}

//根据时间范围及服务名查询数据
func (e *EsService) QueryTimeLog(data interface{}, queryNum int,
	kv map[string]string, rangeTimeKey string, rangeTimeValue *RangeTime) []interface{} {
	var res *elastic.SearchResult
	var err error
	boolSearch := elastic.NewBoolQuery().
		Filter(elastic.NewRangeQuery(rangeTimeKey).Gte(rangeTimeValue.MinTime).Lte(rangeTimeValue.MaxTime))
	if len(kv) > 0{
		condition := utils.AppendCondition(kv)
		for _,v := range condition {
			boolSearch.Filter(v)
		}
	}

	res, err = e.Client.Search(e.Index).Type(e.EsType).Query(boolSearch).Size(queryNum).
		Do(context.Background())

	if err != nil {
		println(err.Error())
	}

	return utils.GetDataList(res, err, data)
}

//查询es上面所有索引库
func (e *EsService) GetIndexNames()[]string  {
	names, err := e.Client.IndexNames()
	if err != nil {
		println(err.Error())
	}

	return names
}
