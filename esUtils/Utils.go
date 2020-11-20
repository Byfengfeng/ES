package esUtils

import (
	"github.com/olivere/elastic/v7"
	"reflect"
)

//get es data list
func GetDataList(res *elastic.SearchResult, err error, dataType interface{}) []interface{} {
	if err != nil {
		print(err.Error())
		return nil
	}
	list := make([]interface{}, 0)
	for _, item := range res.Each(reflect.TypeOf(dataType)) {
		t := item.(interface{})
		list = append(list, t)
	}
	return list
}

//get es one data
func GetDataOne(res *elastic.SearchResult, err error, dataType interface{}) interface{} {
	if err != nil {
		print(err.Error())
		return nil
	}
	for _, item := range res.Each(reflect.TypeOf(dataType)) {
		dataType = item.(interface{})
	}
	return dataType
}

//append query condition
func AppendCondition(kv map[string]string)[]*elastic.TermsQuery  {
	termsQuery := make([]*elastic.TermsQuery,0)
	for k, v := range kv {
		termsQuery = append(termsQuery,elastic.NewTermsQuery(k, v))
	}
	return termsQuery
}
