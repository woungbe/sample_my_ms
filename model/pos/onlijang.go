package pos

import (
	"database/sql"
	"fmt"
	"sample_my_ms/configs"
)

/*
type PosInterface interface {
	Connection(options interface{}) *sql.DB
	GetProduct(types string) (interface{}, error) // 타입은 정하기 나름 !!
	GetEvent(types string) (interface{}, error)   // 이벤트 정하기 나름
}
*/

type ShopList struct {
	ID string `db:"id"` 
	Pcode string `db:"pcode"`
}

type ErrorInsert struct {
	Divtype string `db:"divtype"`
	Params string `db:"params"`
	Explan string `db:"explan"`
}

// 포스 공통함수
type PosConnect struct {
	Ip         string `db:"ip"`
	Port       string `db:"port"`
	DbName     string `db:"dbName"`
	LoginID    string `db:"loginID"`
	Password   string `db:"password"`
	Encrypt    string `db:"encrypt"`
	Lastupdate string `db:"lastupdate"`
}

type MartList struct {
	PosConnect []PosConnect
}


func GetShopProduct(martID string) ([]ShopList, error){
	db := configs.GetMysqlDB()
	query := fmt.Sprintf(`SELECT id,pcode FROM shop_product
	WHERE admin ='%s' AND pcode != '' and is_delete=0 group by pcode limit 0,10000`, martID)
	slist := make([]ShopList, 0)
	err := db.Select(slist, query)
	if err != nil {
		fmt.Println(err)
	}

	return slist, err
}

func SPUpdatePriceFormPos(martID, bcode, taxFree string , price int64) (sql.Result, error) {
	db := configs.GetMysqlDB()
	query := fmt.Sprintf(
		`call SP_UpdatePriceFormPos( '%s', '%s', %d, '%s')`,
		martID, bcode, price, taxFree)
	result, err := db.Exec(query);
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return result, err
}

func InsertErrorLogPos(attr []ErrorInsert){
	db := configs.GetMysqlDB()
	cityState := `insert into pos_errorlog(divtype, params, explan)
	values (?, ?, ?)`

	for _, v:= range attr {
		db.MustExec(cityState, v.Divtype, v.Params, v.Explan)
	}
}

// 
func SaveSystemLog(types, message string) (sql.Result, error){
	db := configs.GetMysqlDB()
    query := fmt.Sprintf(`insert into system_log(type, message)
    values ('%s', '%s'), type, message`, types, message)	
	result,err := db.Exec(query)
	if err != nil {
		fmt.Println("err : ", err.Error())
		return nil, err
	}

	return result,err
}

// update_endtime 
func UpdateEndTime(martID string) (sql.Result, error) {	 
	db := configs.GetMysqlDB()
	query := fmt.Sprintf(
		`update posconfig set lastupdate=now() where mart_id='%s'`, martID)

	result,err := db.Exec(query)
	if err != nil {
		fmt.Println("err : ", err.Error())
		return nil, err
	}

	return result,err
}

// enterprise_devdb
// enterprisedev_db









