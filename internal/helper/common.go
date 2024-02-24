package helper

import (
	"runtime"
)

func GetFuncName() string {
	pc, _, _, _ := runtime.Caller(1)
	funcDetails := runtime.FuncForPC(pc)
	return funcDetails.Name()
}
