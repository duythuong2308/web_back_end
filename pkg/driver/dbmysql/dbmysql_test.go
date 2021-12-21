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
	2  Bà Rịa - Vũng Tàu
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

func TestRepo_UpsertDistrict (t *testing.T) {
	s := `101  Thành phố Châu Đốc
	102  Thành phố Long Xuyên
	103  Thị xã Tân Châu
	104  Huyện An Phú
	105  Huyện Châu Phú
	106  Huyện Châu Thành
	107  Huyện Chợ Mới
	108  Huyện Phú Tân
	109  Huyện Thoại Sơn
	110  Huyện Tịnh Biên
	111  Huyện Tri Tôn
	201  Thành phố Bà Rịa
	202  Thành phố Vũng Tàu
	203  Thị xã Phú Mỹ
	204  Huyện Châu Đức
	205  Huyện Côn Đảo
	206  Huyện Đất Đỏ
	207  Huyện Long Điền
	208  Huyện Xuyên Mộc
	301  Thành phố Bắc Giang
	302  Huyện Hiệp Hoà
	303  Huyện Lạng Giang
	304	 Huyện Lục Nam
	305	 Huyện Lục Ngạn
	306	 Huyện Sơn Động
	307	 Huyện Tân Yên
	308	 Huyện Việt Yên
	309	 Huyện Yên Dũng
	310	 Huyện Yên Thế
	401	Thành phố Bắc Kạn
	402	Huyện Ba Bể
	403	Huyện Bạch Thông
	404	Huyện Chợ Đồn
	405	Huyện Chợ Mới
	406	Huyện Na Rì
	407	Huyện Ngân Sơn
	408	Huyện Pác Nặm
	501  Thành phố Bạc Liêu
	502  Thị xã Giá Rai
	503  Huyện Đông Hải
	504  Huyện Hòa Bình
	505  Huyện Hồng Dân
	506  Huyện Phước Long
	507  Huyện Vĩnh Lợi
	601  Thành phố Bắc Ninh
	602  Thị xã Từ Sơn
	603  Huyện Gia Bình
	604  Huyện Lương Tài
	605  Huyện Quế Võ
	606  Huyện Thuận Thành
	607  Huyện Tiên Du
	608  Huyện Yên Phong
	701	 Thành phố Bến Tre
	702	 Huyện Ba Tri
	703	 Huyện Bình Đại
	704	 Huyện Châu Thành
	705	 Huyện Chợ Lách
	706	 Huyện Giồng Trôm
	707	 Huyện Mỏ Cày Bắc
	708	 Huyện Mỏ Cày Nam
	709	 Huyện Thạnh Phú
	801  Thành phố Quy Nhơn
	802  Thị xã An Nhơn
	803  Thị xã Hoài Nhơn
	804  Huyện An Lão
	805  Huyện Hoài Ân
	806  Huyện Phù Cát
	807  Huyện Phù Mỹ
	808  Huyện Tây Sơn
	809  Huyện Tuy Phước
	810  Huyện Vân Canh
	811  Huyện Vĩnh Thạnh
	901  Thành phố Dĩ An
	902  Thành phố Thủ Dầu Một
	903  Thành phố Thuận An
	904  Thị xã Bến Cát
	905  Thị xã Tân Uyên
	906  Huyện Bắc Tân Uyên
	907  Huyện Bàu Bàng
	908  Huyện Dầu Tiếng
	909  Huyện Phú Giáo
	1001  Thành phố Đồng Xoài
	1002  Thị xã Bình Long
	1003  Thị xã Phước Long
	1004  Huyện Bù Đăng
	1005  Huyện Bù Đốp
	1006  Huyện Bù Gia Mập
	1007  Huyện Chơn Thành
	1008  Huyện Đồng Phú
	1009  Huyện Hớn Quản
	1010  Huyện Lộc Ninh
	1011  Huyện Phú Riềng
	1101  Thành phố Phan Thiết
	1102  Thị xã La Gi
	1103  Huyện Bắc Bình
	1104  Huyện Đức Linh
	1105  Huyện Hàm Tân
	1106  Huyện Hàm Thuận Bắc
	1107  Huyện Hàm Thuận Nam
	1108  Huyện Phú Quý
	1109  Huyện Tánh Linh
	1110  Huyện Tuy Phong
	1201  Thành phố Cà Mau
	1202  Huyện Cái Nước
	1203  Huyện Đầm Dơi
	1204  Huyện Năm Căn
	1205  Huyện Ngọc Hiển
	1206  Huyện Phú Tân
	1207  Huyện Thới Bình
	1208  Huyện Trần Văn Thời
	1209  Huyện U Minh
	1301  Quận Bình Thủy
	1302  Quận Cái Răng
	1303  Quận Ninh Kiều
	1304  Quận Ô Môn
	1305  Quận Thốt Nốt
	1306  Huyện Cờ Đỏ
	1307  Huyện Phong Điền
	1308  Huyện Thới Lai
	1309  Huyện Vĩnh Thạnh
	1401  Thành phố Cao Bằng
	1402  Huyện Bảo Lạc
	1403  Huyện Bảo Lâm
	1404  Huyện Hạ Lang
	1405  Huyện Hà Quảng
	1406  Huyện Hòa An
	1407  Huyện Nguyên Bình
	1408  Huyện Quảng Hòa
	1409  Huyện Thạch An
	1410  Huyện Trùng Khánh
	1501  Quận Cẩm Lệ
	1502  Quận Hải Châu
	1503  Quận Liên Chiểu
	1504  Quận Ngũ Hành Sơn
	1505  Quận Sơn Trà
	1506  Quận Thanh Khê
	1507  Huyện Hòa Vang
	1508  Huyện Hoàng Sa
	1601  Thành phố Buôn Ma Thuột
	1602  Thị xã Buôn Hồ
	1603  Huyện Buôn Đôn
	1604  Huyện Cư Kuin
	1605  Huyện Cư M’gar
	1606  Huyện Ea H’leo
	1607  Huyện Ea Kar
	1608  Huyện Ea Súp
	1609  Huyện Krông Ana
	1610  Huyện Krông Bông
	1611  Huyện Krông Búk
	1612  Huyện Krông Năng
	1613  Huyện Krông Pắk
	1614  Huyện Lắk
	1615  Huyện M’Đrắk
	1701  Thành phố Gia Nghĩa
	1702  Huyện Cư Jút
	1703  Huyện Đắk Glong
	1704  Huyện Đắk Mil
	1705  Huyện Đắk R’lấp
	1706  Huyện Đắk Song
	1707  Huyện Krông Nô
	1708  Huyện Tuy Đức
	1801  Thành phố Điện Biên Phủ
	1802  Thị xã Mường Lay
	1803  Huyện Điện Biên
	1804  Huyện Điện Biên Đông
	1805  Huyện Mường Ảng
	1806  Huyện Mường Chà
	1807  Huyện Mường Nhé
	1808  Huyện Nậm Pồ
	1809  Huyện Tủa Chùa
	1810  Huyện Tuần Giáo
	1901  Thành phố Biên Hòa
	1902  Thành phố Long Khánh
	1903  Huyện Cẩm Mỹ
	1904  Huyện Định Quán
	1905  Huyện Long Thành
	1906  Huyện Nhơn Trạch
	1907  Huyện Tân Phú
	1908  Huyện Thống Nhất
	1909  Huyện Trảng Bom
	1910  Huyện Vĩnh Cửu
	1911  Huyện Xuân Lộc
	2001  Thành phố Cao Lãnh
	2002  Thành phố Sa Đéc
	2003  Thị xã Hồng Ngự
	2004  Huyện Cao Lãnh
	2005  Huyện Châu Thành
	2006  Huyện Hồng Ngự
	2007  Huyện Lai Vung
	2008  Huyện Lấp Vò
	2009  Huyện Tam Nông
	2010  Huyện Tân Hồng
	2011  Huyện Thanh Bình
	2012  Huyện Tháp Mười
	2101  Thành phố Pleiku
	2102  Thị xã An Khê
	2103  Thị xã Ayun Pa
	2104  Huyện Chư Păh
	2105  Huyện Chư Prông
	2106  Huyện Chư Pưh
	2107  Huyện Chư Sê
	2108  Huyện Đắk Đoa
	2109  Huyện Đak Pơ
	2110  Huyện Đức Cơ
	2111  Huyện Ia Grai
	2112  Huyện Ia Pa
	2113  Huyện K’Bang
	2114  Huyện Kông Chro
	2115  Huyện Krông Pa
	2116  Huyện Mang Yang
	2117  Huyện Phú Thiện
	2201  Thành phố Hà Giang
	2202  Huyện Bắc Mê
	2203  Huyện Bắc Quang
	2204  Huyện Đồng Văn
	2205  Huyện Hoàng Su Phì
	2206  Huyện Mèo Vạc
	2207  Huyện Quản Bạ
	2208  Huyện Quang Bình
	2209  Huyện Vị Xuyên
	2210  Huyện Xín Mần
	2211  Huyện Yên Minh
	2301  Thành phố Phủ Lý
	2302  Thị xã Duy Tiên
	2303  Huyện Bình Lục
	2304  Huyện Kim Bảng
	2305  Huyện Lý Nhân
	2306  Huyện Thanh Liêm
	2401  Quận Ba Đình
	2402  Quận Bắc Từ Liêm
	2403  Quận Cầu Giấy
	2404  Quận Đống Đa
	2405  Quận Hà Đông
	2406  Quận Hai Bà Trưng
	2407  Quận Hoàn Kiếm
	2408  Quận Hoàng Mai
	2409  Quận Long Biên
	2410  Quận Nam Từ Liêm
	2411  Quận Tây Hồ
	2412  Quận Thanh Xuân
	2413  Thị xã Sơn Tây
	2414  Huyện Ba Vì
	2415  Huyện Chương Mỹ
	2416  Huyện Đan Phượng
	2417  Huyện Đông Anh
	2418  Huyện Gia Lâm
	2419  Huyện Hoài Đức
	2420  Huyện Mê Linh
	2421  Huyện Mỹ Đức
	2422  Huyện Phú Xuyên
	2423  Huyện Phúc Thọ
	2424  Huyện Quốc Oai
	2425  Huyện Sóc Sơn
	2426  Huyện Thạch Thất
	2427  Huyện Thanh Oai
	2428  Huyện Thanh Trì
	2429  Huyện Thường Tín
	2430  Huyện Ứng Hoà
	2501  Thành phố Hà Tĩnh
	2502  Thị xã Hồng Lĩnh
	2503  Thị xã Kỳ Anh
	2504  Huyện Can Lộc
	2505  Huyện Cẩm Xuyên
	2506  Huyện Đức Thọ
	2507  Huyện Hương Khê
	2508  Huyện Hương Sơn
	2509  Huyện Kỳ Anh
	2510  Huyện Lộc Hà
	2511  Huyện Nghi Xuân
	2512  Huyện Thạch Hà
	2513  Huyện Vũ Quang
	2601  Thành phố Hải Dương
	2602  Thành phố Chí Linh
	2603  Thị xã Kinh Môn
	2604  Huyện Bình Giang
	2605  Huyện Cẩm Giàng
	2606  Huyện Gia Lộc
	2607  Huyện Kim Thành
	2608  Huyện Nam Sách
	2609  Huyện Ninh Giang
	2610  Huyện Thanh Hà
	2611  Huyện Thanh Miện
	2612  Huyện Tứ Kỳ
	2701  Quận Đồ Sơn
	2702  Quận Dương Kinh
	2703  Quận Hải An
	2704  Quận Hồng Bàng
	2705  Quận Kiến An
	2706  Quận Lê Chân
	2707  Quận Ngô Quyền
	2708  Huyện An Dương
	2709  Huyện An Lão
	2710  Huyện Bạch Long Vĩ
	2711  Huyện Cát Hải
	2712  Huyện Kiến Thuỵ
	2713  Huyện Thủy Nguyên
	2714  Huyện Tiên Lãng
	2715  Huyện Vĩnh Bảo
	2801  Thành phố Vị Thanh
	2802  Thị xã Long Mỹ
	2803  Thị xã Ngã Bảy
	2804  Huyện Châu Thành
	2805  Huyện Châu Thành A
	2806  Huyện Long Mỹ
	2807  Huyện Phụng Hiệp
	2808  Huyện Vị Thuỷ`

	lines := strings.Split(s, "\n")
	for i, line := range lines {
		firstSpace := strings.Index(line, " ")
		if firstSpace != -1 {
			line = strings.TrimSpace(line[firstSpace:])
			district := core.District{
				Id:		fmt.Sprintf("%04d", i+1),
				Name:	line,
			}
			err := repo0.UpsertDistrict(district)
			if err != nil {
				t.Errorf("error District: %v", err)
			}
		}
	}
}

func TestRepo_ReadCommune(t *testing.T) {
	s :=``

	lines := strings.Split(s, "\n")
	for i, line := range lines {
		firstSpace := strings.Index(line, " ")
		line = strings.TrimSpace(line[firstSpace:])
		commune := core.Commune{
			Id:		fmt.Sprintf("%06d", i+1),
			Name: 	line,
		}
		err := repo0.UpsertCommune(commune)
		if err != nil {
			t.Errorf("error Commune: %v", err)
		}
	}
}



func UpdatePopulation( )  {

}