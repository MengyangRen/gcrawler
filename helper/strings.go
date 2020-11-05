package helper

import (
	"fmt"
	"github.com/valyala/fastrand"
	"strconv"
	"strings"
)

func GenId() uint64 {

	str      := fmt.Sprintf("%d%d", Cputicks(), fastrand.Uint32n(9))
	val, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0
	}

	return val
}

//过滤关键词
func Filter(source string,rep string,new string)  string{

	return strings.Replace(source,rep,new,-1)
}

//过滤空格
func Space(str string)  string{

	return strings.Join(strings.Fields(str),"")

}
