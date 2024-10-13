package error_utils

import (
	"fmt"
	"runtime"
)

func HandleError(err error) (newErr error) {
	if err != nil {
		// notice that we're using 1, so it will actually log the where
		// the error happened, 0 = this function, we don't want that.
		pc, filename, line, _ := runtime.Caller(1)

		return fmt.Errorf("[error] in %s[%s:%d] %v", runtime.FuncForPC(pc).Name(), filename, line, err)
	}
	return
}
