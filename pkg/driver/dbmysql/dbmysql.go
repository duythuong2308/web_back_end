package dbmysql

import (
	"github.com/duythuong2308/web_back_end/pkg/core"
	"gorm.io/gorm"
)

type Repo struct {
	DB *gorm.DB
}

func (r Repo) UpsertProvince(province core.Province) error {
	return r.DB.Debug().Save(&province).Error
}
