package httpsvr_citizen

import (
	"net/http"

	"github.com/duythuong2308/web_back_end/pkg/core"
	"github.com/duythuong2308/web_back_end/pkg/driver/dbmysql"
	"github.com/mywrap/httpsvr"
)

type Server struct {
	*httpsvr.Server
	Database *dbmysql.Repo
}

func NewServer(database *dbmysql.Repo) *Server {
	s := &Server{
		Server:   httpsvr.NewServer(),
		Database: database,
	}
	s.AddHandler("GET", "/api/hello", s.getHello)

	s.AddHandler("GET", "/api/province", s.getProvinces)
	s.AddHandler("POST", "/api/province", s.postProvince)
	s.AddHandler("DELETE", "/api/province", s.deleteProvince)

	s.AddHandler("GET", "/api/district", s.getDistricts)
	s.AddHandler("POST", "/api/district", s.postDistrict)
	s.AddHandler("DELETE", "/api/district", s.deleteDistrict)

	s.AddHandler("GET", "/api/commune", s.getCommunes)
	s.AddHandler("POST", "/api/commune", s.postCommune)
	s.AddHandler("DELETE", "/api/commune", s.deleteCommune)

	s.AddHandler("GET", "/api/village", s.getVillages)
	s.AddHandler("GET", "/api/village/:id", s.getVillage)
	s.AddHandler("POST", "/api/village", s.postVillage)
	s.AddHandler("DELETE", "/api/village", s.deleteVillage)

	return s
}

func (s Server) getHello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`hehe`))
}

// common response format for all requests
type Response struct {
	Data  interface{}
	Error string
}

func (s Server) getProvinces(w http.ResponseWriter, r *http.Request) {
	provinces, err := s.Database.ReadProvinces()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		s.WriteJson(w, r, Response{Error: err.Error()})
		return
	}
	s.WriteJson(w, r, Response{Data: provinces})
}

func (s Server) postProvince(w http.ResponseWriter, r *http.Request) {
	var province core.Province
	err := s.ReadJson(r, &province)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		s.WriteJson(w, r, Response{Error: err.Error()})
	}
	err = s.Database.UpsertProvince(province)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		s.WriteJson(w, r, Response{Error: err.Error()})
	}
	s.WriteJson(w, r, Response{Data: province})
}

func (s Server) deleteProvince(w http.ResponseWriter, r *http.Request) {
	var province core.Province
	err := s.ReadJson(r, &province)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		s.WriteJson(w, r, Response{Error: err.Error()})
	}
	err = s.Database.DeleteProvince(province.Id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		s.WriteJson(w, r, Response{Error: err.Error()})
	}
	s.WriteJson(w, r, Response{Data: "deleted province"})
}

func (s Server) getDistricts(w http.ResponseWriter, r *http.Request) {
	provinceId := r.FormValue("provinceId")
	districts, err := s.Database.ReadDistricts(provinceId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		s.WriteJson(w, r, Response{Error: err.Error()})
		return
	}
	s.WriteJson(w, r, Response{Data: districts})
}

func (s Server) postDistrict(w http.ResponseWriter, r *http.Request) {
	var district core.District
	err := s.ReadJson(r, &district)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		s.WriteJson(w, r, Response{Error: err.Error()})
	}
	err = s.Database.UpsertDistrict(district)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		s.WriteJson(w, r, Response{Error: err.Error()})
	}
	s.WriteJson(w, r, Response{Data: district})
}

func (s Server) deleteDistrict(w http.ResponseWriter, r *http.Request) {
	var row core.District
	err := s.ReadJson(r, &row)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		s.WriteJson(w, r, Response{Error: err.Error()})
	}
	err = s.Database.DeleteDistrict(row.Id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		s.WriteJson(w, r, Response{Error: err.Error()})
	}
	s.WriteJson(w, r, Response{Data: "deleted district"})
}

func (s Server) getCommunes(w http.ResponseWriter, r *http.Request) {
	districtId := r.FormValue("districtId")
	communes, err := s.Database.ReadCommunes(districtId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		s.WriteJson(w, r, Response{Error: err.Error()})
		return
	}
	s.WriteJson(w, r, Response{Data: communes})
}

func (s Server) postCommune(w http.ResponseWriter, r *http.Request) {
	var commune core.Commune
	err := s.ReadJson(r, &commune)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		s.WriteJson(w, r, Response{Error: err.Error()})
	}
	err = s.Database.UpsertCommune(commune)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		s.WriteJson(w, r, Response{Error: err.Error()})
	}
	s.WriteJson(w, r, Response{Data: commune})
}

func (s Server) deleteCommune(w http.ResponseWriter, r *http.Request) {
	var row core.Commune
	err := s.ReadJson(r, &row)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		s.WriteJson(w, r, Response{Error: err.Error()})
	}
	err = s.Database.DeleteCommune(row.Id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		s.WriteJson(w, r, Response{Error: err.Error()})
	}
	s.WriteJson(w, r, Response{Data: "deleted commune"})
}

func (s Server) getVillage(w http.ResponseWriter, r *http.Request) {
	villageId := httpsvr.GetUrlParams(r)["id"]
	village, err := s.Database.ReadVillage(villageId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		s.WriteJson(w, r, Response{Error: err.Error()})
		return
	}
	s.WriteJson(w, r, Response{Data: village})
}

func (s Server) getVillages(w http.ResponseWriter, r *http.Request) {
	communeId := r.FormValue("communeId")
	villages, err := s.Database.ReadVillages(communeId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		s.WriteJson(w, r, Response{Error: err.Error()})
		return
	}
	s.WriteJson(w, r, Response{Data: villages})
}

func (s Server) postVillage(w http.ResponseWriter, r *http.Request) {
	var village core.Village
	err := s.ReadJson(r, &village)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		s.WriteJson(w, r, Response{Error: err.Error()})
	}
	err = s.Database.UpsertVillage(village)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		s.WriteJson(w, r, Response{Error: err.Error()})
	}
	s.WriteJson(w, r, Response{Data: village})
}

func (s Server) deleteVillage(w http.ResponseWriter, r *http.Request) {
	var row core.Village
	err := s.ReadJson(r, &row)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		s.WriteJson(w, r, Response{Error: err.Error()})
	}
	err = s.Database.DeleteVillage(row.Id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		s.WriteJson(w, r, Response{Error: err.Error()})
	}
	s.WriteJson(w, r, Response{Data: "deleted village"})
}
