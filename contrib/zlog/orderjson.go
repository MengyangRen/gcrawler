package zlog

import (
	"bytes"
	"encoding/json"
)

type KeyVal struct {
	Key string
	Val interface{}
}

// Define an ordered map
type OrderedMap struct {
	vals   []KeyVal
	keymap map[string]int
}

func NewOrderMap() *OrderedMap {
	return &OrderedMap{
		vals:   make([]KeyVal, 0),
		keymap: make(map[string]int),
	}
}

// Implement the json.Marshaler interface
func (o OrderedMap) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer

	buf.WriteString("{")
	for i, kv := range o.vals {
		if i != 0 {
			buf.WriteString(",")
		}
		// marshal key
		key, err := json.Marshal(kv.Key)
		if err != nil {
			return nil, err
		}
		buf.Write(key)
		buf.WriteString(":")
		// marshal value
		if tmp, ok := kv.Val.(error); ok {
			kv.Val = tmp.Error() //对于error类型特殊处理
		}
		val, err := json.Marshal(kv.Val)
		if err != nil {
			return nil, err
		}
		buf.Write(val)
	}
	buf.WriteString("}")
	return buf.Bytes(), nil
}
func (o *OrderedMap) Set(key string, val interface{}) {
	newval := KeyVal{Key: key, Val: val}
	if v, ok := o.keymap[key]; ok {
		o.vals[v] = newval
	} else {
		o.vals = append(o.vals, newval)
		o.keymap[key] = len(o.vals) - 1
	}

}
func (o OrderedMap) Get(key string) (val interface{}, ok bool) {
	if v, ok := o.keymap[key]; ok {
		return o.vals[v].Val, true
	}
	return nil, false
}
func (o *OrderedMap) AddVals(vals *OrderedMap) {
	if vals == nil {
		return
	}
	for _, val := range vals.vals {
		o.Set(val.Key, val.Val)
	}
}
