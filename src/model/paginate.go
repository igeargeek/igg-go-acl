package model

import (
	"context"
	"fmt"
	"sync"

	"github.com/igeargeek/igg-go-acl/config"

	"github.com/igeargeek/igg-golang-api-response/response"
	"github.com/qiniu/qmgo"
)

func getDataPaginate(results interface{}, collection *qmgo.Collection, limit int64, page int64, filter interface{}, sort string) {
	if page < 1 {
		page = 1
	}
	// options.SetSkip(limit * (page - 1))
	if err := collection.Find(context.TODO(), filter).Sort(sort).Limit(limit).Skip(limit * (page - 1)).All(results); err != nil {
		fmt.Printf("Paginate data error %s", err.Error())
	}
}

func getPaginateTotal(total *int64, collection *qmgo.Collection, filter interface{}) {
	count, err := collection.Find(context.TODO(), filter).Count()
	if err != nil {
		fmt.Printf("Paginate count error %s", err.Error())
	}
	*total = count
}

func getPaginate(
	results interface{},
	collectionName string,
	limit int64,
	page int64,
	filter interface{},
	sort string) (response.Pagination, error) {
	resultPaginate := response.Pagination{}
	collection := config.GetDBClient().Collection(collectionName)
	var wg sync.WaitGroup
	wg.Add(2)
	var total int64
	go func() {
		getDataPaginate(results, collection, limit, page, filter, sort)
		wg.Done()
	}()
	go func() {
		getPaginateTotal(&total, collection, filter)
		wg.Done()
	}()
	wg.Wait()

	finalResult := []interface{}{results}
	resultPaginate = response.Pagination{
		Data:    finalResult,
		Total:   total,
		PerPage: limit,
		Page:    page,
	}
	return resultPaginate, nil
}
