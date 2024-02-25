package helper

import (
	"encoding/json"
	"runtime"
)

func GetFuncName() string {
	pc, _, _, _ := runtime.Caller(1)
	funcDetails := runtime.FuncForPC(pc)
	return funcDetails.Name()
}

// Dump to json using json marshal
func Dump(i interface{}) string {
	return string(ToByte(i))
}

// ToByte parses any to byte
func ToByte(i interface{}) []byte {
	bt, _ := json.Marshal(i)
	return bt
}
