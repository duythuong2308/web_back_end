package main

import (
	"log"

	"time"

	"github.com/duythuong2308/web_back_end/pkg/core"
	"github.com/duythuong2308/web_back_end/pkg/driver/dbmysql"
	"github.com/mywrap/mysql"
)

func main() {
	cfg := mysql.Config{
		Host:     "127.0.0.1",
		Port:     "3306",
		Username: "root",
		Password: "123qwe",
		Database: "duythuong",
	}
	mysqlCli, err := mysql.ConnectViaGORM(cfg)
	if err != nil {
		log.Fatalf("error connect mysql: %v, config: %#v", err, cfg)
	}
	database := &dbmysql.Repo{DB: mysqlCli}

	provinces, err := database.ReadProvinces()
	if err != nil {
		log.Fatalf("error ReadProvinces: %v", err)
	}

	for _, province := range provinces {
		err := database.UpsertUser(core.User{
			Username:     province.Id,
			Password:     "123qwe",
			Role:         core.RoleA2,
			BeginDeclare: time.Unix(0, 0),
			EndDeclare:   time.Unix(0, 0),
		})
		if err != nil {
			log.Printf("error CreateUser: %v", err)
		}
	}
	log.Printf("done")
}
