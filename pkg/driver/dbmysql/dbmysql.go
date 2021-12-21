package dbmysql

import (
	"github.com/duythuong2308/web_back_end/pkg/core"
	"gorm.io/gorm"
)

type Repo struct {
	DB *gorm.DB
}

func (r Repo) ReadProvince(provinceId string) (core.Province, error) {
	var ret core.Province
	err := r.DB.Debug().
		Where(core.Province{Id: provinceId}).
		First(&ret).Error
	return ret, err
}

func (r Repo) ReadDistrict(districtId string) (core.District, error) {
	var ret core.District
	err := r.DB.Debug() .
		Where(core.District{Id:districtId}).
		First(&ret).Error
	return ret, err
}

func (r Repo) ReadCommnue(communeId string) (core.Commune, error) {
	var ret core.Commune
	err := r.DB.Debug().
		Where(core.Commune{Id: communeId}).
		First(&ret).Error
	return ret, err
}

func (r Repo) UpsertProvince(province core.Province) error {
	return r.DB.Debug().Save(&province).Error
}

func (r Repo) UpsertDistrict(district core.District) error {
	return r.DB.Debug().Save(&district).Error
}

func (r Repo) UpsertCommune(commune core.Commune) error {
	return r.DB.Debug().Save(&commune).Error
}
