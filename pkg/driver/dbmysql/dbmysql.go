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

func (r Repo) ReadProvinces() ([]core.Province, error) {
	var ret []core.Province
	err := r.DB.Debug().Find(&ret).Error
	return ret, err
}

func (r Repo) UpsertProvince(province core.Province) error {
	return r.DB.Debug().Save(&province).Error
}

func (r Repo) DeleteProvince(provinceId string) error {
	return r.DB.Delete(&core.Province{}, provinceId).Error
}

func (r Repo) ReadDistrict(districtId string) (core.District, error) {
	var ret core.District
	err := r.DB.Debug().
		Where(core.District{Id: districtId}).
		Preload("Province").
		First(&ret).Error
	return ret, err
}

func (r Repo) ReadDistricts(provinceId string) ([]core.District, error) {
	var ret []core.District
	err := r.DB.Debug().
		Where(core.District{ProvinceId: provinceId}).
		Find(&ret).Error
	return ret, err
}

func (r Repo) UpsertDistrict(district core.District) error {
	return r.DB.Debug().Save(&district).Error
}

func (r Repo) DeleteDistrict(districtId string) error {
	return r.DB.Delete(&core.District{}, districtId).Error
}

func (r Repo) ReadCommnue(communeId string) (core.Commune, error) {
	var ret core.Commune
	err := r.DB.Debug().
		Where(core.Commune{Id: communeId}).
		Preload("District").
		Preload("District.Province").
		First(&ret).Error
	return ret, err
}

func (r Repo) ReadCommunes(districtId string) ([]core.Commune, error) {
	var ret []core.Commune
	err := r.DB.Debug().
		Where(core.Commune{DistrictId: districtId}).
		Find(&ret).Error
	return ret, err
}

func (r Repo) DeleteCommune(communeId string) error {
	return r.DB.Delete(&core.Commune{}, communeId).Error
}

func (r Repo) UpsertCommune(commune core.Commune) error {
	return r.DB.Debug().Save(&commune).Error
}

func (r Repo) UpsertVillage(village core.Village) error {
	return r.DB.Debug().Save(&village).Error
}

func (r Repo) ReadVillages(communeId string) ([]core.Village, error) {
	var rows []core.Village
	err := r.DB.Where(core.Village{CommuneId: communeId}).Find(&rows).Error
	return rows, err
}

func (r Repo) ReadVillage(villageId string) (core.Village, error) {
	var ret core.Village
	err := r.DB.Debug().
		Model(&core.Village{}).
		Where(core.Village{Id: villageId}).
		Preload("Commune").
		Preload("Commune.District").
		Preload("Commune.District.Province").
		First(&ret).Error
	return ret, err
}

func (r Repo) DeleteVillage(villageId string) error {
	return r.DB.Delete(&core.Village{}, villageId).Error
}

func (r Repo) ReadUser(username string) (core.User, error) {
	ret := core.User{Username: username}
	err := r.DB.First(&ret).Error
	return ret, err
}

func (r Repo) UpsertUser(user core.User) error {
	return r.DB.Save(&user).Error
}

func (r Repo) ReadCitizens(villageId string) ([]core.Citizen, error) {
	var rows []core.Citizen
	err := r.DB.Where(core.Citizen{VillageId: villageId}).Find(&rows).Error
	return rows, err
}

func (r Repo) ReadCitizen(citizenId string) (core.Citizen, error) {
	var ret core.Citizen
	err := r.DB.Debug().
		Model(&core.Citizen{}).
		Where(core.Citizen{Id: citizenId}).
		Preload("Village").
		Preload("Village.Commune").
		Preload("Village.Commune.District").
		Preload("Village.Commune.District.Province").
		First(&ret).Error
	return ret, err
}

func (r Repo) UpsertCitizen(citizen core.Citizen) error {
	return r.DB.Debug().Save(&citizen).Error
}

func (r Repo) DeleteCitizen(citizenId string) error {
	return r.DB.Delete(&core.Citizen{}, citizenId).Error
}
