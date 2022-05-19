package utils

import (
	"fmt"
	"os"
)

func errHandle(err error, errMsg string) {
	if err != nil {
		fmt.Println(errMsg, err)
		os.Exit(-1)
	}

}
