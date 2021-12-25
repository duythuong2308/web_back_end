package main

import (
	"io/ioutil"

	"sync"

	"github.com/duythuong2308/web_back_end/pkg/core"
	"github.com/duythuong2308/web_back_end/pkg/driver/dbmysql"
	"github.com/mywrap/log"
	"github.com/mywrap/mysql"
)

func main() {
	mysqlCfg := mysql.Config{
		Host:     "127.0.0.1",
		Port:     "3306",
		Username: "root",
		Password: "123qwe",
		Database: "duythuong",
	}
	mysqlClient, err := mysql.ConnectViaGORM(mysqlCfg)
	database := dbmysql.Repo{DB: mysqlClient}

	rawData, err := ioutil.ReadFile("pkg/core/vietnam_provinces.json")
	if err != nil {
		log.Fatal(err)
	}
	provinces, districts, wards, err := core.ReadVietnamProvincesData(rawData)
	if err != nil {
		log.Fatal(err)
	}

	limitGoroutines := make(chan bool, 100)
	wg := sync.WaitGroup{}

	for _, p := range provinces {
		err := database.UpsertProvince(p)
		if err != nil {
			log.Printf("error UpsertProvince: %v", err)
		}
	}
	wg.Wait()

	for _, d := range districts {
		d := d
		wg.Add(1)
		limitGoroutines <- true
		go func() {
			defer func() {
				wg.Add(-1)
				<-limitGoroutines
			}()
			err := database.UpsertDistrict(d)
			if err != nil {
				log.Printf("error UpsertProvince: %v", err)
			}
		}()
	}
	wg.Wait()

	for _, w := range wards {
		w := w
		wg.Add(1)
		limitGoroutines <- true
		go func() {
			defer func() {
				wg.Add(-1)
				<-limitGoroutines
			}()
			err := database.UpsertCommune(w)
			if err != nil {
				log.Printf("error UpsertProvince: %v", err)
			}
		}()
	}
	wg.Wait()
	log.Printf("done")
}
