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
	s.AddHandler("GET", "/hello", s.getHello)
	s.AddHandler("GET", "/province", s.getProvinces)
	s.AddHandler("GET", "/district", s.getDistricts)
	s.AddHandler("GET", "/commune", s.getCommunes)
	s.AddHandler("POST", "/province", s.postProvinces)
	s.AddHandler("POST", "/district", s.postDistricts)
	s.AddHandler("POST", "/commune", s.postCommunes)
	s.AddHandler("POST", "/village", s.postVillage)

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

func (s Server) postProvinces(w http.ResponseWriter, r *http.Request) {
	var province core.Province
	err := s.ReadJson(r, &province)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		s.WriteJson(w, r, Response{Error: err.Error()})
	}
	err = s.Database.UpsertCommune(province)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		s.WriteJson(w, r, Response{Error: err.Error()})
	}
	s.WriteJson(w, r, Response{Data: province})
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

func (s Server) postDistricts(w http.ResponseWriter, r *http.Request) {
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

func (s Server) postCommunes(w http.ResponseWriter, r *http.Request) {
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
