//
// Copyright (C) 2022 lazywhite <lazywhite@qq.com>
//
// Distributed under terms of the MIT license.
//

package main

import (
	"fmt"
)

func main(){
	
	a := 10
	b := 20 
    /*
    a = 10 ^ 20
    b = 10 ^ 20 ^ 20 = 10
    a = 10 ^ 20 ^ 10 = 20
    */

	a = a ^ b  
	b = a ^ b
	a = a ^ b
	fmt.Println(a, b)
}
