package core

import (
	"io/ioutil"
	"testing"
)

func TestReadVietnamProvincesData(t *testing.T) {
	rawData, err := ioutil.ReadFile("vietnam_provinces.json")
	if err != nil {
		t.Fatal(err)
	}
	provinces, districts, wards, err := ReadVietnamProvincesData(rawData)
	if err != nil {
		t.Fatal(err)
	}
	if len(provinces) == 0 || len(districts) == 0 || len(wards) == 0 {
		t.Fatal("empty provinces or empty districts or empty wards")
	}
	t.Logf("provinces: %v, %v", len(provinces), provinces)
	t.Logf("districts: %v, %v", len(districts), districts[0])
	t.Logf("wards: %v, %v", len(wards), wards[0])
}
