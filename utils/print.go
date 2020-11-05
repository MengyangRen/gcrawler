package utils

import "fmt"

/**
 *#30:黑
 *#31:红
 *#32:绿
 *#33:黄
 *#34:蓝色
 *#35:紫色
 *#36:深绿
 *#37:白色
 * utils.Error(errors.New("read pkg head error"))
 */
func Debug(str interface{}) {
	fmt.Println("\033[33m [Debug] \033[0m", str)
}

func Error(err interface{}) {
	fmt.Println("\033[31m [Error] \033[0m", err)
}

func Info(str interface{}) {
	fmt.Println("\033[32m [Info] \033[0m", str)
}
