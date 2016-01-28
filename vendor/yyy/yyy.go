package yyy

import (
	"fmt"

	"xxx"
)

type b struct{ value string }

var B = b{"This is B from yyy"}

var A = xxx.A

func init() {
	fmt.Printf("yyy.init(): xxx.A = %#v, xxx.B = %#v, A = %#v, B = %#v\n", xxx.A, xxx.A, A, B)
}
