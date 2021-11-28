package utils

import (
	"strings"
)

func CompareJsonArray(result, data []map[string]interface{}, key1, key2 string) []map[string]interface{} {
	// list.result, data, 'pcode', 'goods_bcode'
	for _, v1 := range result {
		tt := v1[key1].(string)
		v1[key1] = strings.TrimSpace(tt)
	}

	for _, v2 := range data {
		tt := v2[key2].(string)
		v2[key2] = strings.TrimSpace(tt)		
	}
		

	vsf := make([]map[string]interface{}, 0);
	for _,k2 := range result {
		for _,v2 := range data {			
			if k2[key1] == v2[key2] {
				vsf = append(vsf, k2)
			}
		}
	}

	return vsf
}

