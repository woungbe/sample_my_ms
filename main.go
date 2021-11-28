package main

import (
	"fmt"
	"os"
	"sample_my_ms/configs"
)

/*
 - 간단하게 설계를 하고 리팩토링을 해보자 !!
  1. mysql ,mssql 연결해서 동작하는 거 하나 !!

  정확하게 모르면 탑다운으로 해버려 !!

  pos 에 있는 데이터를 가져와서. mysql DB에 넣는 작업 -
  중간에 조율이 필요함 !!

*/


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
	// cnf := model.Init()
	// cnf.MakeFileFor()
	// fmt.Println(cnf)
}
