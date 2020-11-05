package utils

/**
 * To change this template, choose Tools | Templates
 * and open the template in the editor.
 *
 * 定义项目基础工具方法
 *
 * @author  m.y <j01428@kok.work>
 * @date
 * @package utils
 * @description
 *
 */
import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const (
	BYTE = 1 << (10 * iota)
	KILOBYTE
	MEGABYTE
	GIGABYTE
	TERABYTE
	PETABYTE
	EXABYTE
)

//THREAD-THROW-PANIC-37405
func CatchPanic(s string) {
	if err := recover(); err != nil {
		Error(fmt.Sprintf("%s Error:%v", s, err))
	}
}

func RandSleep(sec int) {
	time.Sleep(time.Duration(rand.Intn(sec)) * time.Second)
}

//随机毫秒睡眠
func Usleep() {
	time.Sleep(100 * time.Millisecond * time.Duration(rand.Int31n(100)))
}

//json转map
func Json2Map(s string) (bm map[int]map[string]string, r error) {
	err := json.Unmarshal([]byte(s), &bm)
	if err != nil {
		Error(fmt.Sprintf("json->map err=%v", err))
		return nil, r
	}
	return bm, nil
}

//map转json
func Map2Json(m interface{}) string {
	data, err := json.Marshal(m)
	if err != nil {
		Error(fmt.Sprintf("map->json err=%v", err))
	}
	return string(data)
}

//json转结构体
func Json2Struct(data []byte, v interface{}) error {
	err := json.Unmarshal(data, v)
	if err != nil {
		Error(err)
	}
	return err
}

//结构体转json
func Struct2Json(s interface{}) string {
	data, err := json.Marshal(s)
	if err != nil {
		Error(err)
	}
	return string(data)
}

//多个字典合并
func MapMerge(maps ...map[int]map[string]string) map[int]map[string]string {
	newMap := make(map[int]map[string]string)
	index := 0
	for _, item := range maps {
		for _, v := range item {
			newMap[index] = v
			index++
		}
	}
	return newMap
}

//查询数组是否一个元素
func InStrArray(search string, arr []string) bool {
	for _, item := range arr {
		if item == search {
			return true
		}
	}
	return false
}

//结构体转Json
func StructToJson(s interface{}) string {
	data, err := json.Marshal(s)
	if err != nil {
		Error(err)
	}
	return string(data)
}

func StructToByteJson(s interface{}) []byte {
	data, err := json.Marshal(s)
	if err != nil {
		Error(err)
	}
	return data
}

func JsonToStruct(data []byte, v interface{}) error {
	err := json.Unmarshal(data, v)
	if err != nil {
		Error(err)
	}
	return err
}

// HeapSys：程序向应用程序申请的内存  m.HeapSys
// HeapAlloc：堆上目前分配的内存   ms.Alloc
// HeapIdle：堆上目前没有使用的内存 ms.HeapIdle
// HeapReleased：回收到操作系统的内存 ms.HeapRelease
func ReadMemoryStats() {
	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)

	log.Printf("\n程序向应用程序申请的内存(HeapSys):%v\n堆上内存分配(Alloc):%v\n堆未使用的内存(HeapIdle):%v\n堆释放的内存(HeapReleased):%v",
		ByteSize(ms.HeapSys), ByteSize(ms.Alloc), ByteSize(ms.HeapIdle), ByteSize(ms.HeapReleased))
}

func ByteSize(bytes uint64) string {
	unit := ""
	value := float64(bytes)

	switch {
	case bytes >= EXABYTE:
		unit = "E"
		value = value / EXABYTE
	case bytes >= PETABYTE:
		unit = "P"
		value = value / PETABYTE
	case bytes >= TERABYTE:
		unit = "T"
		value = value / TERABYTE
	case bytes >= GIGABYTE:
		unit = "G"
		value = value / GIGABYTE
	case bytes >= MEGABYTE:
		unit = "M"
		value = value / MEGABYTE
	case bytes >= KILOBYTE:
		unit = "K"
		value = value / KILOBYTE
	case bytes >= BYTE:
		unit = "B"
	case bytes == 0:
		return "0"
	}

	result := strconv.FormatFloat(value, 'f', 1, 64)
	result = strings.TrimSuffix(result, ".0")
	return result + unit
}
