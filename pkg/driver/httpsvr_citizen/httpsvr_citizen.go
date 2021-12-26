package httpsvr_citizen

import (
	"net/http"

	"strings"

	"fmt"

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

	s.AddHandler("POST", "/api/login", s.postLogin)

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

type LoginRequest struct {
	Username string
	Password string
}

type LoginResponse struct {
	User     core.User
	Location interface{} // Province or District or Commune or Village
	Managers []core.User // các cấp trên
	Error    string
}

func (s Server) postLogin(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	err := s.ReadJson(r, &req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		s.WriteJson(w, r, LoginResponse{Error: err.Error()})
		return
	}
	ret, status := s.readLoginInfo(req)
	w.WriteHeader(status)
	s.WriteJson(w, r, ret)
}

func (s Server) readLoginInfo(req LoginRequest) (ret LoginResponse, status int) {
	user, err := s.Database.ReadUser(req.Username)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			ret.Error = "username does not exist"
			return ret, http.StatusUnauthorized
		} else {
			ret.Error = err.Error()
			return ret, http.StatusInternalServerError
		}
	}
	//log.Debugf("user: %+v, req: %+v", user, req)
	if user.Password != req.Password {
		ret.Error = "wrong password"
		return ret, http.StatusUnauthorized
	}
	ret.User = user

	switch user.Role {
	case core.RoleA1:
	case core.RoleA2:
		province, err := s.Database.ReadProvince(user.Username)
		if err != nil {
			ret.Error = fmt.Sprintf("fail link A2 user to province: %v", err)
			return ret, http.StatusInternalServerError
		}
		ret.Location = province
	case core.RoleA3:
		district, err := s.Database.ReadDistrict(user.Username)
		if err != nil {
			ret.Error = fmt.Sprintf("fail link A3 user to district: %v", err)
			return ret, http.StatusInternalServerError
		}
		ret.Location = district
		userProvince, _ := s.Database.ReadUser(district.ProvinceId)
		ret.Managers = append(ret.Managers, userProvince)
	case core.RoleB1:
		commune, err := s.Database.ReadCommnue(user.Username)
		if err != nil {
			ret.Error = fmt.Sprintf("fail link B1 user to commune: %v", err)
			return ret, http.StatusInternalServerError
		}
		ret.Location = commune
		userDistrict, _ := s.Database.ReadUser(commune.DistrictId)
		ret.Managers = append(ret.Managers, userDistrict)
		if commune.District != nil { // sure
			userProvince, _ := s.Database.ReadUser(commune.District.ProvinceId)
			ret.Managers = append(ret.Managers, userProvince)
		}
	case core.RoleB2:
		village, err := s.Database.ReadVillage(user.Username)
		if err != nil {
			ret.Error = fmt.Sprintf("fail link B2 user to village: %v", err)
			return ret, http.StatusInternalServerError
		}
		ret.Location = village
		userCommune, _ := s.Database.ReadUser(village.CommuneId)
		ret.Managers = append(ret.Managers, userCommune)
		if village.Commune != nil {
			userDistrict, _ := s.Database.ReadUser(village.Commune.DistrictId)
			ret.Managers = append(ret.Managers, userDistrict)
			if village.Commune.District != nil { // sure
				userProvince, _ := s.Database.ReadUser(village.Commune.District.ProvinceId)
				ret.Managers = append(ret.Managers, userProvince)
			}
		}
	default:
		ret.Error = "unexpected user role: " + string(user.Role)
		return ret, http.StatusInternalServerError
	}
	return ret, http.StatusOK
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
