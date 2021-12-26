package core

import (
	"time"
)

type Province struct {
	Id   string `gorm:"primary_key;type:varchar(191)"`
	Name string `gorm:"type:varchar(191)"`
}

type District struct {
	Id         string    `gorm:"primary_key;type:varchar(191)"`
	ProvinceId string    `gorm:"type:varchar(191)"` // this field name follows gorm foreign key convention
	Province   *Province `gorm:"constraint:fk_districts_provinceid,OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Name       string    `gorm:"type:varchar(191)"`
}

type Commune struct {
	Id          string    `gorm:"primary_key;type:varchar(191)"`
	DistrictId  string    `gorm:"type:varchar(191)"`
	District    *District `gorm:"constraint:fk_communes_districtid,OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Name        string    `gorm:"type:varchar(191)"`
	Population  string		`gorm:"type:varchar(191)"`
	IsCompleted bool
}

type Village struct {
	Id         string   `gorm:"primary_key;type:varchar(191)"`
	CommuneId  string   `gorm:"type:varchar(191)"`
	Commune    *Commune `gorm:"constraint:fk_villages_communeid,OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Name       string   `gorm:"type:varchar(191)"`
	Population string	`gorm:"type:varchar(191)"`
}

type User struct {
	Username     string `gorm:"primary_key;type:varchar(191)"`
	Password     string `gorm:"type:varchar(191)" json:"-"`
	Role         Role   `gorm:"type:varchar(191)"`
	BeginDeclare time.Time
	EndDeclare   time.Time
}

type Role string

const (
	RoleA1 Role = "A1" // admin, update provinces
	RoleA2 Role = "A2" // update districts in a province, Id len 2
	RoleA3 Role = "A3" // update communes in a district, Id len 4
	RoleB1 Role = "B1" // update villages in a commune, Id len 6
	RoleB2 Role = "B2" // add citizen in a village, Id len 8
)

type Citizen struct {
	Id               string   `gorm:"type:varchar(191)"`
	VillageId        string   `gorm:"type:varchar(191)"`
	Village          *Village `gorm:"constraint:fk_citizen_villageid,OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Name             string   `gorm:"type:varchar(191)"`
	DateOfBirth      string   `gorm:"type:varchar(191)"`
	Gender           string   `gorm:"type:varchar(191)"`
	PlaceOfBirth     string   `gorm:"type:varchar(191)"`
	PernamentAddress string   `gorm:"type:varchar(1023)"`
	TemporaryAddress string   `gorm:"type:varchar(1023)"`
	Religion         string   `gorm:"type:varchar(191)"`
	EducationLevel   string   `gorm:"type:varchar(191)"`
	Job              string   `gorm:"type:varchar(191)"`
	IdDate           string   `gorm:"type:varchar(191)"`
	IdPlace          string   `gorm:"type:varchar(191)"`
	Note             string   `gorm:"type:varchar(1023)"`
}
