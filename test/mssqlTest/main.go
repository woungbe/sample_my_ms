package main

import (
	"fmt"
	"os"
	"sample_my_ms/configs"
	"sample_my_ms/model/pos"
)

func serviceStart() {
	cnf := configs.GetConfig()
	er := cnf.InitConfig("config.json")
	if er != nil {
		fmt.Println("설정정보로드에러")
		os.Exit(1)
	}
}

func main() {
	serviceStart()

	// rows, err := pos.GetProductList("", "2021-11-25 00:00:00")
	// if err != nil {
	// 	fmt.Println(err);
	// }

	// for k1,v1 := range rows {
	// 	rows[k1]["goods_bcode"] = strings.TrimSpace(v1["goods_bcode"].(string))
	// }

	// fmt.Println(rows)
	// SaveSystemLog(types, message string) (sql.Result, error)


	var temp pos.ErrorInsert

	temp.Divtype = "Divtype"  
	temp.Explan = "Explan" 
	temp.Params = "Params" 

	var attr []pos.ErrorInsert
	attr = append(attr, temp)	
	pos.InsertErrorLogPos(attr)



}