package configs

import (
	"encoding/json"
	"io/ioutil"

	"sample_my_ms/sqlxDB"

	"github.com/go-playground/validator"
	"github.com/gorilla/sessions"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	logging "github.com/op/go-logging"
)

var debugmodeflg bool //디버그 모드 플레그 디버그:true,

type ConfigData struct {
	ServerName   string //서버이름
	RunMode      string // = release (서비스용) , debug (디버그용)
	Ssluse       string //= Y,N
	Sslkey       string //= ssl key 파일
	Sslcrt       string //= ssl crt 파일
	MysqlDBConfig sqlxDB.DBConnectInfo
	MssqlDBConfig  sqlxDB.DBConnectInfo
	HttpConfig   httpConfig
	CookieConfig CookieConfig // cookie 관련 정보!1 - 수정하고 싶다면 수정하셈 !
}

type SingleConfig struct {
	cnf    ConfigData
	CnfLog StaticLog
	E      *echo.Echo
}

var ConfigPTR *SingleConfig
var MysqlObj *sqlxDB.SQLXforMysql //메인 db
var MssqlObj *sqlxDB.SQLXforMssql // mssql

func IsDebugmode() bool {
	return debugmodeflg
}

func SetDebugmode(b bool) {
	debugmodeflg = b
}

func GetConfig() *SingleConfig {

	if ConfigPTR == nil {
		ConfigPTR = new(SingleConfig)
	}
	return ConfigPTR
}

func GetConfigData() *ConfigData {
	return &ConfigPTR.cnf
}

func GetConfigLog() *logging.Logger {
	return ConfigPTR.CnfLog.Log
}

func GetMysqlDB() *sqlx.DB {
	return MysqlObj.GetDBConn()
}

func GetMssqlDB() *sqlx.DB {
	return MssqlObj.GetDBConn()
}

// 시스템 세팅 값 저장!!
func (ty *SingleConfig) InitConfig(cfPath string) error {

	er := ty.loadConfig(cfPath)

	if er != nil {
		return er
	}

	// mysql 등록
	MysqlObj = new(sqlxDB.SQLXforMysql) //mysql db
	MysqlObj.InitDB(ty.cnf.MysqlDBConfig, 10, 10, 300)
	_, err := MysqlObj.ConnectDB()
	if err != nil {
		return err
	}

	// mysql 등록
	MssqlObj = new(sqlxDB.SQLXforMssql) //mssql db
	MssqlObj.InitDB(ty.cnf.MssqlDBConfig, 10, 10, 300)
	_, err2 := MssqlObj.ConnectDB()
	if err2 != nil {
		return err2
	}

	// http settting
	ty.E = echo.New() //  echo 초기화

	// 로그 시작 로딩
	// ty.InitHttp()

	return er
}

// config 파일 로드
func (ty *SingleConfig) loadConfig(cfPath string) error {

	b, err := ioutil.ReadFile(cfPath)
	if err != nil {
		log.Warn("Warn", "config file Not found", "mfconfig.json")
		return err
	}

	er := json.Unmarshal(b, &ty.cnf)
	if er != nil {
		log.Error("Error", "설정로드에러", er.Error())
		return er
	}

	if ty.cnf.RunMode == "debug" {
		SetDebugmode(true)
	} else {
		SetDebugmode(false)
	}

	return nil
}

// https 설정할때 쓰면 됨
func (ty *SingleConfig) InitHttp() {

	// ty.E.Use(middleware.CORS()) // 이상한데서 오면 막는다. - api는 이거 풀어놔야지..
	ty.E.Use(session.Middleware(
		sessions.NewCookieStore([]byte("cdms-PDAServer")))) // session store생성함
	ty.E.Use(middleware.Recover()) // 에러 나면 다시 살려주는애인데. .못살릴 수도 있음.

	ty.E.Validator = &CustomValidator{validator: validator.New()} // validator 체크 - 꼭 있어야되는애, tag랑 같이 쓰이는 듯

	ty.E.Static("/", "./public")

}
