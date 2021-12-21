package core

import (
	"encoding/json"
	"fmt"
)

// read data in file "pkg/core/vietnam_provinces.json"
func ReadVietnamProvincesData(rawData []byte) (
	provinces []Province, districts []District, communes []Commune, err error) {
	type VietnamProvincesData []struct {
		Id        string
		Code      string
		Name      string
		Districts []struct {
			Id    string
			Name  string
			Wards []struct {
				Id     string
				Name   string
				Prefix string // "Xã" / "Phường"
			}
			Streets  []struct{} // don't care
			Projects []struct{} // don't care
		}
	}
	var dataObj VietnamProvincesData
	err = json.Unmarshal(rawData, &dataObj)
	if err != nil {
		return nil, nil, nil, err
	}
	for i, rawProvince := range dataObj {
		province := Province{
			Id:   fmt.Sprintf("%02d", i+1),
			Name: rawProvince.Name,
		}
		provinces = append(provinces, province)
		for k, rawDistrict := range rawProvince.Districts {
			district := District{
				Id:         province.Id + fmt.Sprintf("%02d", k+1),
				ProvinceId: province.Id,
				Name:       rawDistrict.Name,
			}
			districts = append(districts, district)
			for m, rawWard := range rawDistrict.Wards {
				commune := Commune{
					Id:         district.Id + fmt.Sprintf("%02d", m+1),
					DistrictId: district.Id,
					Name:       rawWard.Name,
				}
				communes = append(communes, commune)
			}
		}
	}
	return provinces, districts, communes, nil
}
