package dbmysql

import (
	"log"
	"os"
	"testing"

	"github.com/duythuong2308/web_back_end/pkg/core"
	"github.com/mywrap/mysql"
)

var repo0 *Repo

func TestMain(m *testing.M) {
	cfg := mysql.Config{
		Host:     "127.0.0.1",
		Port:     "3306",
		Username: "root",
		Password: "123qwe",
		Database: "duythuong",
	}
	// cfg = mysql.LoadEnvConfig()
	cli, err := mysql.ConnectViaGORM(cfg)
	if err != nil {
		log.Fatalf("error connect mysql: %v, config: %#v", err, cfg)
	}
	repo0 = &Repo{DB: cli}
	os.Exit(m.Run())
}

func TestRepo_UpsertProvince(t *testing.T) {
	err := repo0.UpsertProvince(core.Province{
		Id:   "01",
		Name: "Province01",
	})
	if err != nil {
		t.Errorf("error UpsertProvince: %v", err)
	}

}

func TestRepo_ReadProvince(t *testing.T) {
	read, err := repo0.ReadProvince("01")
	if err != nil {
		t.Errorf("error ReadProvince: %v", err)
	}
	t.Logf("read: %v", read)

	//read.Name = "Bac Giaaaang"
	//repo0.UpsertProvince(read)
}

func TestRepo_DeleteProvince(t *testing.T) {
	err := repo0.DeleteProvince("01")
	if err != nil {
		t.Errorf("error: %v", err)
	}
}

func TestRepo_UpsertDistrict(t *testing.T) {
	err := repo0.UpsertDistrict(core.District{
		Id: "1010", ProvinceId: "10", Name: "Huyện abc"})
	if err != nil {
		t.Errorf("error District: %v", err)
	}
}

func TestRepo_UpsertCommune(t *testing.T) {
	err := repo0.UpsertCommune(core.Commune{
		Id: "101010", DistrictId: "1010", Name: "Huyện abc"})
	if err != nil {
		t.Errorf("error District: %v", err)
	}
}

func TestRepo_UpsertVillage(t *testing.T) {
	err := repo0.UpsertVillage(core.Village{
		Id: "01010101", CommuneId: "101010", Population: 123})
	if err != nil {
		t.Fatalf("error UpsertVillage: %v", err)
	}
	t.Logf("ok UpsertVillage")
}
