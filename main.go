package main

import (
	"fmt"
	"github.com/yerkesh/onelab_hw3/atoi"
	"github.com/yerkesh/onelab_hw3/itoa"
	"github.com/yerkesh/onelab_hw3/reverse"
)



func main() {
	fmt.Println(itoa.MyItoa(5432))
	fmt.Println(reverse.Reverse("Алматы"))
	fmt.Println(atoi.MyAtoi("569"))
}
