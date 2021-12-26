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

		districts, err := database.ReadDistricts(province.Id)
		if err != nil {
			log.Fatalf("error ReadDistricts: %v", err)
		}

		for _, district := range districts {
			err := database.UpsertUser(core.User{
				Username:     district.Id,
				Password:     "123qwe",
				Role:         core.RoleA3,
				BeginDeclare: time.Unix(0, 0),
				EndDeclare:   time.Unix(0, 0),
			})
			if err != nil {
				log.Printf("error CreateUser: %v", err)
			}

			communes, err := database.ReadCommunes(district.Id)
			if err != nil {
				log.Fatalf("error ReadCommunes: %v", err)
			}

			for _, commune := range communes {
				err := database.UpsertUser(core.User{
					Username:     commune.Id,
					Password:     "123qwe",
					Role:         core.RoleB1,
					BeginDeclare: time.Unix(0, 0),
					EndDeclare:   time.Unix(0, 0),
				})
				if err != nil {
					log.Printf("error CreateUser: %v", err)
				}

				villages, err := database.ReadVillages(commune.Id)
				if err != nil {
					log.Fatalf("error ReadVillages: %v", err)
				}

				for _, village := range villages {
					err := database.UpsertUser(core.User{
						Username:     village.Id,
						Password:     "123qwe",
						Role:         core.RoleB2,
						BeginDeclare: time.Unix(0, 0),
						EndDeclare:   time.Unix(0, 0),
					})
					if err != nil {
						log.Printf("error CreateUser: %v", err)
					}
				}
			}
		}
	}
	log.Printf("done")
}
