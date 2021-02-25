package main

import (
	"fmt"
	pkgRename "go_demos/pkg"
)

func init() {
	fmt.Println("main init")
}

func packageFun() {
	fmt.Println(pkgRename.Mode)
	fmt.Println(pkgRename.Add(3,4))
}
