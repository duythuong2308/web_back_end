package core

type Province struct {
	Id   string `gorm:"primary_key;type:varchar(191)"`
	Name string `gorm:"type:varchar(191)"`
}

type District struct {
	Id         string    `gorm:"primary_key;type:varchar(191)"`
	ProvinceId string    // this field name follows gorm foreign key convention
	Province   *Province `gorm:"constraint:fk_districts_provinceid,OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Name       string    `gorm:"type:varchar(191)"`
}

type Commune struct {
	Id         string `gorm:"primary_key;type:varchar(191)"`
	DistrictId string
	District   *District `gorm:"constraint:fk_communes_districtid,OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Name       string    `gorm:"type:varchar(191)"`
}

type Village struct {
	Id         string `gorm:"primary_key;type:varchar(191)"`
	CommuneId  string
	Commune    *Commune `gorm:"constraint:fk_villages_communeid,OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Name       string   `gorm:"type:varchar(191)"`
	Population int
}

type Citizen struct {
	Id 					string	`gorm:"type:varchar(191)"`
	VillageId			string
	Village				*Village `gorm:"constraint:fk_citizen_villageid,OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Name 				string	`gorm:"type:varchar(191)"`
	DateOfBirth 		string	`gorm:"type:varchar(191)"`
	Gender				string	`gorm:"type:varchar(191)"`
	PlaceOfBirth		string	`gorm:"type:varchar(191)"`
	PernamentAddress	string	`gorm:"type:varchar(1023)"`
	TemporaryAddress	string	`gorm:"type:varchar(1023)"`
	Religion			string	`gorm:"type:varchar(191)"`
	EducationLevel		string	`gorm:"type:varchar(191)"`
	Job					string	`gorm:"type:varchar(191)"`
}