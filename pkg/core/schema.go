package core

type Province struct {
	Id   int
	Name string
}

type District struct {
	Id         int
	ProvinceId int
	Name       string
}
