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
	s.AddHandler("GET", "/village", s.getVillages)
	s.AddHandler("GET", "/village/:id", s.getVillage)
	s.AddHandler("POST", "/village", s.postVillage)
	//s.AddHandler("DELETE", "/village", s.deleteVillage)

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

/* func (s Server) deleteVillage(w http.ResponseWriter, r *http.Request) {
	var village core.Village
	err := s.ReadJson(r, &village)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		s.WriteJson(w, r, Response{Error: err.Error()})
	}
	err = s.Database.DeleteVillage(village)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		s.WriteJson(w, r, Response{Error: err.Error()})
	}
} */
