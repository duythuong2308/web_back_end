package main

import (
	"github.com/duythuong2308/web_back_end/pkg/core"
	"github.com/mywrap/log"
	"github.com/mywrap/mysql"
)

func main() {
	mysqlConf := mysql.Config{
		Host:     "127.0.0.1",
		Port:     "3306",
		Username: "root",
		Password: "123qwe",
		Database: "duythuong",
	}
	db, err := mysql.ConnectViaGORM(mysqlConf)
	if err != nil {
		log.Fatalf("error connect MySQL: %v, config: %#v", err, mysqlConf)
	}
	//db = db.Debug()
	//if err != nil {
	//	log.Fatal(err)
	//}

	// CREATE DATABASE duythuong /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;

	err = db.AutoMigrate(&core.Province{})
	if err != nil {
		log.Printf("error create table: %v", err)
	}
	err = db.AutoMigrate(&core.District{})
	if err != nil {
		log.Printf("error create table: %v", err)
	}
	err = db.AutoMigrate(&core.Commune{})
	if err != nil {
		log.Printf("error create table: %v", err)
	}
	err = db.AutoMigrate(&core.Village{})
	if err != nil {
		log.Printf("error create table: %v", err)
	}
	err = db.AutoMigrate(&core.User{})
	if err != nil {
		log.Printf("error create table: %v", err)
	}
	err = db.AutoMigrate(&core.Citizen{})
	if err != nil {
		log.Printf("error create table: %v", err)
	}
	log.Print("done")
}
