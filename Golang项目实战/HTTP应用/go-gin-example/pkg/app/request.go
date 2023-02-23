package app

import (
	"fmt"
	"github.com/astaxie/beego/validation"
)

func MarkErrors(errors []*validation.Error) {
	for _, err := range errors {
		fmt.Println(err.Key, err.Message)
	}

	return
}
