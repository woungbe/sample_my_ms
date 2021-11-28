package pos

import (
	"database/sql"
	"fmt"
	"sample_my_ms/configs"
	"sample_my_ms/defined"
)

/*
요구사항 정리 하기 !!
1.실행 시간 조절이 가능해야됨 !!
 - 마트별 각각 실행시간이 조절되어야됨
  : 실행시간 저장해서 보관 !!

2.정보를 가져와서 정리하는 모델이 필요함 => 자주 변경되는 모델- 심지어 여러개
 - 마트별 각각 정리가 필요함
  : 최근상품리스트 가져오기
  : 데이터 받아서 중간에 믹스 및 후처리
  : 상품등록 리스트로 변경해서 전달
  : 이벤트 저장으로 변경해서 전달
  : 이벤트 상품으로 변경해서 전달

3.해당 데이터를 본서버에 적용시켜야됨 !! => 자주 변경되지 않음, 기능추가시 변경
 - 각 데이터를 가져오기
  : 상품리스트
  : 이벤트리스트
  : 이벤트 상품
 - 각 데이터를 저장하기
  : 상품리스트 저장하기
  : 이벤트 저장하기
  : 이벤트 상품 저장하기
*/

// "github.com/jmoiron/sqlx"
// 상품리스트 등록!!
func GetProductList(glist, lastupdate string) ([]map[string]interface{}, error){
	db := configs.GetMssqlDB()
	whereGlist := ""
	if glist !="" {
		whereGlist = fmt.Sprintf(" WHERE g.goods_bcode in (%s) ", glist)
	}

	wherelastupdate := ""
	if lastupdate !="" {
		wherelastupdate = fmt.Sprintf(" and gl_date >= '%s' ", lastupdate)		
	}

	query := fmt.Sprintf(`select top 1 * from goods as g 
		inner join  
		(select gl_goods, 
		CONVERT(CHAR(23), max(gl_date), 21) as gl_date from goodslog where gl_kind in (1,2)
		%s
		group by gl_goods) as s 
		on g.goods_code = s.gl_goods %s`, wherelastupdate,whereGlist )

	// places := []defined.GoodsList{}
	rows,err := db.Queryx(query)
	if err != nil {
		fmt.Println("err : ", err.Error())
		return nil, err
	}

	var AA []map[string]interface{}
	for rows.Next() {
		results := make(map[string]interface{})
		err = rows.MapScan(results)
		AA = append(AA, results)

		if err !=nil {
			fmt.Println(err)
		}
	}

	return AA, err
}

// 상품 등록하는 것!! 
func SP_UpdatePriceFormPos(attr defined.Attr) (sql.Result, error){
	db := configs.GetMysqlDB()

	query := fmt.Sprintf(`call SP_UpdatePriceFormPos('%s', '%s' ,%d ,%d )`, 
	attr.Mart_id, attr.Bcode, attr.Price, attr.Tax_free)	
	result,err := db.Exec(query)
	if err != nil {
		fmt.Println("err : ", err.Error())
		return nil, err
	}

	return result, err
}