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

func (r Repo) UpsertProvince(province core.Province) error {
	return r.DB.Debug().Save(&province).Error
}

func (r Repo) UpsertVillage(villageId string, newPopulation int) error {
	var village core.Village
	err := r.DB.Debug().
		Where(&core.Village{Id: villageId}).
		First(&village).Error
	if err != nil {
		return err
	}
	village.Population = newPopulation
	return r.DB.Debug().Save(&village).Error
}
