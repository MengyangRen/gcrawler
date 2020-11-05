package validator

import (
	//"fmt"
	"errors"
	"math"
	"reflect"
	"strconv"
	"strings"

	"github.com/modern-go/reflect2"
	"github.com/valyala/fasthttp"
)

/*
type Tags struct {
	Rule string `json:"rule"`
	Min  int    `json:"min"`
	Max  int    `json:"max"`
	Msg  string `json:"msg"`
}
*/

func Bind(ctx *fasthttp.RequestCtx, objs interface{}) error {

	rt := reflect2.TypeOf(objs)
	rtElem := rt

	if rt.Kind() != reflect.Ptr {
		return errors.New("argument 2 should be map or ptr")
	}

	rt = rt.(reflect2.PtrType).Elem()
	rtElem = rt

	if rtElem.Kind() != reflect.Struct {
		return errors.New("non-structure type not supported yet")
	}

	s := rtElem.(reflect2.StructType)

	for i := 0; i < s.NumField(); i++ {

		f := s.Field(i)

		min := int64(0)
		max := int64(math.MaxInt64)
		name := f.Tag().Get("name")
		rule := f.Tag().Get("rule")
		msg := f.Tag().Get("msg")
		def := f.Tag().Get("default")
		nums := len(def)

		if len(name) == 0 {
			name = strings.ToLower(f.Name())
		}

		if v, err := strconv.ParseInt(f.Tag().Get("min"), 10, 64); err == nil {
			min = v
		}
		if v, err := strconv.ParseInt(f.Tag().Get("max"), 10, 64); err == nil {
			max = v
		}

		defaultVal := ""
		if string(ctx.Method()) == "GET" {
			defaultVal = string(ctx.QueryArgs().Peek(name))
		} else if string(ctx.Method()) == "POST" {
			defaultVal = string(ctx.PostArgs().Peek(name))
		}

		//fmt.Println("name = ", string(ctx.PostBody()))

		if defaultVal == "" {
			if nums > 0 {
				defaultVal = def
			} else if rule == "none" {
				continue
			} else {
				return errors.New(name + " not found")
			}
		}

		if rule == "digit" {
			if !CheckStringDigit(defaultVal) || !checkIntScope(defaultVal, min, max) {
				return errors.New(msg)
			}
		} else if rule == "sDigit" {
			if !CheckStringCommaDigit(defaultVal) || !CheckStringLength(defaultVal, int(min), int(max)) {
				return errors.New(msg)
			}
		} else if rule == "sAlpha" {
			if !CheckStringCommaAlpha(defaultVal) || !CheckStringLength(defaultVal, int(min), int(max)) {
				return errors.New(msg)
			}
		} else if rule == "url" {
			if !CheckUrl(defaultVal) {
				return errors.New(msg)
			}
		} else if rule == "alnum" {
			if !CheckStringAlnum(defaultVal) || !CheckStringLength(defaultVal, int(min), int(max)) {
				return errors.New(msg)
			}
		} else if rule == "priv" {
			if !isPriv(defaultVal) {
				return errors.New(msg)
			}
		} else if rule == "dateTime" {
			if !CheckDateTime(defaultVal) {
				return errors.New(msg)
			}
		} else if rule == "date" {
			if !CheckDate(defaultVal) {
				return errors.New(msg)
			}
		} else if rule == "time" {
			if !checkTime(defaultVal) {
				return errors.New(msg)
			}
		} else if rule == "chn" {
			if !CheckStringCHN(defaultVal) {
				return errors.New(msg)
			}
		} else if rule == "module" {
			if !CheckStringModule(defaultVal) || !CheckStringLength(defaultVal, int(min), int(max)) {
				return errors.New(msg)
			}
		}

		switch f.Type().Kind() {
		case reflect.Bool:
			if val, err := strconv.ParseBool(defaultVal); err == nil {
				f.UnsafeSet(reflect2.PtrOf(objs), reflect2.PtrOf(val))
			}
		case reflect.Int:
			if val, err := strconv.Atoi(defaultVal); err == nil {
				f.UnsafeSet(reflect2.PtrOf(objs), reflect2.PtrOf(val))
			}
		case reflect.Int8:
			if val, err := strconv.ParseInt(defaultVal, 10, 8); err == nil {
				f.UnsafeSet(reflect2.PtrOf(objs), reflect2.PtrOf(val))
			}
		case reflect.Int16:
			if val, err := strconv.ParseInt(defaultVal, 10, 16); err == nil {
				f.UnsafeSet(reflect2.PtrOf(objs), reflect2.PtrOf(val))
			}
		case reflect.Int32:
			if val, err := strconv.ParseInt(defaultVal, 10, 32); err == nil {
				f.UnsafeSet(reflect2.PtrOf(objs), reflect2.PtrOf(val))
			}
		case reflect.Int64:
			if val, err := strconv.ParseInt(defaultVal, 10, 64); err == nil {
				f.UnsafeSet(reflect2.PtrOf(objs), reflect2.PtrOf(val))
			}
		case reflect.Uint:
			if val, err := strconv.ParseUint(defaultVal, 10, 64); err == nil {
				f.UnsafeSet(reflect2.PtrOf(objs), reflect2.PtrOf(val))
			}
		case reflect.Uint8:
			if val, err := strconv.ParseUint(defaultVal, 10, 8); err == nil {
				f.UnsafeSet(reflect2.PtrOf(objs), reflect2.PtrOf(val))
			}
		case reflect.Uint16:
			if val, err := strconv.ParseUint(defaultVal, 10, 16); err == nil {
				f.UnsafeSet(reflect2.PtrOf(objs), reflect2.PtrOf(val))
			}
		case reflect.Uint32:
			if val, err := strconv.ParseUint(defaultVal, 10, 32); err == nil {
				f.UnsafeSet(reflect2.PtrOf(objs), reflect2.PtrOf(val))
			}
		case reflect.Uint64:
			if val, err := strconv.ParseUint(defaultVal, 10, 64); err == nil {
				f.UnsafeSet(reflect2.PtrOf(objs), reflect2.PtrOf(val))
			}
		case reflect.Uintptr:
			if val, err := strconv.ParseUint(defaultVal, 10, 64); err == nil {
				f.UnsafeSet(reflect2.PtrOf(objs), reflect2.PtrOf(val))
			}
		case reflect.Float32:
			if val, err := strconv.ParseFloat(defaultVal, 32); err == nil {
				f.UnsafeSet(reflect2.PtrOf(objs), reflect2.PtrOf(val))
			}
		case reflect.Float64:
			if val, err := strconv.ParseFloat(defaultVal, 64); err == nil {
				f.UnsafeSet(reflect2.PtrOf(objs), reflect2.PtrOf(val))
			}
		case reflect.String:
			f.UnsafeSet(reflect2.PtrOf(objs), reflect2.PtrOf(defaultVal))
		}
		//fmt.Println("MatchType = ", f.MatchType().Kind())
	}

	return nil
}
