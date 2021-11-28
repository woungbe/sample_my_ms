package sqlxDB

import (
	"fmt"
	"log"
	"time"

	_ "github.com/denisenkom/go-mssqldb"

	"github.com/jmoiron/sqlx"
)

type SQLXforMssql struct {
	bInited         bool
	stDBConnectInfo DBConnectInfo //연결정보
	mDBconn         *sqlx.DB
	mMaxOpenConn    int
	mMaxIdelConn    int

	mJobTime time.Time
	idleTime int64
	closeFlg bool //데이터베이스 close됐는가
}

func (ty *SQLXforMssql) InitDB(tyDbInfoStruct DBConnectInfo, nMaxOpenConn int, nMaxIdelConn int, nIdleChkTime int64) {
	ty.stDBConnectInfo = tyDbInfoStruct
	ty.mMaxOpenConn = nMaxOpenConn
	ty.mMaxIdelConn = nMaxIdelConn

	ty.idleTime = nIdleChkTime
	ty.closeFlg = false
	if ty.idleTime < 60 {
		ty.idleTime = 60
	}
}

func (ty *SQLXforMssql) ConnectDB() (dbconn *sqlx.DB, err error) {

	// 접속경로 지정하는 부분 ex => "root:fnakfn100djr@(localhost:3306)/sakila"
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s",
		ty.stDBConnectInfo.StrIP, ty.stDBConnectInfo.StrID, ty.stDBConnectInfo.StrPasswd, ty.stDBConnectInfo.NPort, ty.stDBConnectInfo.StrDBname)

	db, err := sqlx.Connect("sqlserver", connString)
	if err != nil {
		log.Fatalln(err)
	}

	db.SetMaxIdleConns(ty.mMaxIdelConn)
	db.SetMaxOpenConns(ty.mMaxOpenConn)

	ty.mDBconn = db
	ty.bInited = true

	ty.mJobTime = time.Now()
	ty.closeFlg = false
	// go ty.dbIdleUnconnectionChecker() // DB connection 문제 인거 같은데. 이제 해결된듯,,
	return db, nil
}
func (ty *SQLXforMssql) GetDBConn() *sqlx.DB {
	return ty.mDBconn
}
