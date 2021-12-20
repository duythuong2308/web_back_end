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

type Commune struct {
	Id 			int
	ProvinceId	int
	DistrictId	int
	Name 		string

}

type Village struct {
	Id 			int
	ProvinceId	int
	DistrictId	int
	CommuneId	int
	Name 		string
	Population	int
}

