package defectdojo

import (
	"fmt"
	"reflect"
	"runtime"
)

type p struct{}

func (d *DefectDojoAPI) errorWithStackTrace(err error) {
	if !d.verbose {
		return
	}
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	pkg := reflect.TypeOf(p{}).PkgPath()
	fmt.Printf("[%s] error: %s\nstack-trace:\n", pkg, err.Error())
	for {
		frame, ok := frames.Next()
		fmt.Printf("[%s][%s:%d] %s:\n", pkg, frame.File, frame.Line, frame.Function)
		if !ok {
			break
		}
	}
}
