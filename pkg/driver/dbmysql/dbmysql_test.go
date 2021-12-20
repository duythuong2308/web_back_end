package dbmysql

import (
	"log"
	"os"
	"strings"
	"testing"

	"fmt"

	"github.com/duythuong2308/web_back_end/pkg/core"
	"github.com/mywrap/mysql"
)

var repo0 *Repo

func TestMain(m *testing.M) {
	cfg := mysql.Config{
		Host:     "127.0.0.1",
		Port:     "3306",
		Username: "root",
		Password: "123qwe",
		Database: "duythuong",
	}
	// cfg = mysql.LoadEnvConfig()
	cli, err := mysql.ConnectViaGORM(cfg)
	if err != nil {
		log.Fatalf("error connect mysql: %v, config: %#v", err, cfg)
	}
	repo0 = &Repo{DB: cli}
	os.Exit(m.Run())
}

func TestRepo_UpsertProvince(t *testing.T) {
	s := `1  An Giang
	2  Bà rịa - Vũng tàu
	3  Bắc Giang
	4  Bắc Kạn
	5  Bạc Liêu
	6  Bắc Ninh
	7  Bến Tre
	8  Bình Định
	9  Bình Dương
	10  Bình Phước
	11  Bình Thuận
	12  Cà Mau
	13  Cần Thơ
	14  Cao Bằng
	15  Đà Nẵng
	16  Đắk Lắk
	17  Đắk Nông
	18  Điện Biên
	19  Đồng Nai
	20  Đồng Tháp
	21  Gia Lai
	22  Hà Giang
	23  Hà Nam
	24  Hà Nội
	25  Hà Tĩnh
	26  Hải Dương
	27  Hải Phòng
	28  Hậu Giang
	29  Hòa Bình
	30  Hưng Yên
	31  Khánh Hòa
	32  Kiên Giang
	33  Kon Tum
	34  Lai Châu
	35  Lâm Đồng
	36  Lạng Sơn
	37  Lào Cai
	38  Long An
	39  Nam Định
	40  Nghệ An
	41  Ninh Bình
	42  Ninh Thuận
	43  Phú Thọ
	44  Phú Yên
	45  Quảng Bình
	46  Quảng Nam
	47  Quảng Ngãi
	48  Quảng Ninh
	49  Quảng Trị
	50  Sóc Trăng
	51  Sơn La
	52  Tây Ninh
	53  Thái Bình
	54  Thái Nguyên
	55  Thanh Hóa
	56  Thừa Thiên Huế
	57  Tiền Giang
	58  Thành phố Hồ Chí Minh
	59  Trà Vinh
	60  Tuyên Quang
	61  Vĩnh Long
	62  Vĩnh Phúc
	63  Yên Bái`
	lines := strings.Split(s, "\n")
	for i, line := range lines {
		firstSpace := strings.Index(line, " ")
		if firstSpace != -1 {
			line = strings.TrimSpace(line[firstSpace:])
			province := core.Province{
				Id:   fmt.Sprintf("%02d", i+1),
				Name: line,
			}
			err := repo0.UpsertProvince(province)
			if err != nil {
				t.Errorf("error UpsertProvince: %v", err)
			}
		}
	}

}

func TestRepo_UpsertProvince2(t *testing.T) {
	read, err := repo0.ReadProvince("03")
	if err != nil {
		t.Errorf("error ReadProvince: %v", err)
	}
	t.Logf("read: %v", read)

	//read.Name = "Bac Giaaaang"
	//repo0.UpsertProvince(read)
}
