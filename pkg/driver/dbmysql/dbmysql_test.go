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

func _TestRepo_DeleteProvince(t *testing.T) {
	err := repo0.DeleteProvince("01")
	if err != nil {
		t.Errorf("error: %v", err)
	}
}

func TestRepo_UpsertDistrict(t *testing.T) {
	err := repo0.UpsertDistrict(core.District{
		Id: "0101", ProvinceId: "01", Name: "huyện 0101"})
	if err != nil {
		t.Errorf("error District: %v", err)
	}
	err = repo0.UpsertDistrict(core.District{
		Id: "1010", ProvinceId: "10", Name: "huyện 1010"})
	if err != nil {
		t.Errorf("error District: %v", err)
	}
}

func TestRepo_ReadDistrict(t *testing.T) {
	read, err := repo0.ReadDistrict("0101")
	if err != nil {
		t.Errorf("error ReadDistrict: %v", err)
	}
	t.Logf("read: %v", read)
	if read.Province == nil {
		t.Fatalf("error ReadVillage nil Province")
	}
	t.Logf("province: %+v", read.Province)
}

func _TestRepo_DeleteDistrict(t *testing.T) {
	err := repo0.DeleteDistrict("0101")
	if err != nil {
		t.Errorf("error: %v", err)
	}
}

func TestRepo_UpsertCommune(t *testing.T) {
	err := repo0.UpsertCommune(core.Commune{
		Id: "101010", DistrictId: "1010", Name: "xã abc"})
	if err != nil {
		t.Errorf("error District: %v", err)
	}
}

func TestRepo_ReadCommnue(t *testing.T) {
	read, err := repo0.ReadCommnue("101010")
	if err != nil {
		t.Errorf("error ReadCommune: %v", err)
	}
	t.Logf("read: %v", read)
	if read.District == nil {
		t.Fatalf("error ReadVillage nil District")
	}
	t.Logf("district: %+v", read.District)
	if read.District.Province == nil {
		t.Fatalf("error ReadVillage nil Province")
	}
	t.Logf("province: %+v", read.District.Province)
}

//func TestRepo_DeleteCommune(t *testing.T) {
//	err := repo0.DeleteCommue("101010")
//	if err != nil {
//		t.Errorf("error: %v", err)
//	}
//}

func TestRepo_UpsertVillage(t *testing.T) {
	err := repo0.UpsertVillage(core.Village{
		Id: "01010101", CommuneId: "101010", Population: 123, Name: "thôn 01010101"})
	if err != nil {
		t.Fatalf("error UpsertVillage: %v", err)
	}
	t.Logf("ok UpsertVillage")
}

func TestRepo_ReadVillage(t *testing.T) {
	read, err := repo0.ReadVillage("01010101")
	if err != nil {
		t.Errorf("error ReadVillage: %v", err)
	}
	t.Logf("village: %+v", read)
	if read.Commune == nil {
		t.Fatalf("error ReadVillage nil Commune")
	}
	t.Logf("commune: %+v", read.Commune)
	if read.Commune.District == nil {
		t.Fatalf("error ReadVillage nil District")
	}
	t.Logf("district: %+v", read.Commune.District)
	if read.Commune.District.Province == nil {
		t.Fatalf("error ReadVillage nil Province")
	}
	t.Logf("province: %+v", read.Commune.District.Province)
}

func _TestRepo_DeleteVillage(t *testing.T) {
	err := repo0.DeleteVillage("01010101")
	if err != nil {
		t.Errorf("error: %v", err)
	}
}

func TestRepo_ReadUser(t *testing.T) {
	row, err := repo0.ReadUser("0101")
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("user: %+v", row)
	}
}
