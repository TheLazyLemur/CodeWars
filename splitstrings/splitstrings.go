package splitstrings

import (
	"strings"
)

func Solution(str string) []string {
	arr := strings.Split(str, "")
	res := make([]string, 0)

	for i := 0; i < len(arr); i+=2 {
		if i+1 < len(arr){
			v := arr[i] + arr[i+1]
			res = append(res, v)
		}else{
			res = append(res, arr[i]+"_")
		}
	}

	return res
}
