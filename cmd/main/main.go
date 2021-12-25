package main

import (
	"net/http"
	"path/filepath"

	"github.com/duythuong2308/web_back_end/pkg/driver/dbmysql"
	"github.com/duythuong2308/web_back_end/pkg/driver/httpsvr_citizen"
	"github.com/mywrap/gofast"
	"github.com/mywrap/log"
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

	server := httpsvr_citizen.NewServer(database)
	projectRootDir, err := gofast.GetProjectRootPath()
	if err != nil {
		log.Fatalf("error projectRootDir: %v, %v", projectRootDir, err)
	}
	server.Router.ServeFiles("/gui/*filepath",
		http.Dir(filepath.Join(projectRootDir, "/gui")))
	server.AddHandler("GET", "/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/gui", http.StatusSeeOther)
	})
	listenPort := ":39539"
	log.Printf("listening on http://127.0.0.1%v", listenPort)
	err = server.ListenAndServe(listenPort)
	if err != nil {
		log.Fatalf("err ListenAndServe: %v", err)
	}
}
