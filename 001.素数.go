package main

import (
	"math"
	"fmt"
)

/*
    埃氏法, 线性查找素数

    1. 要得到自然数n以内的全部素数，只需把(2, √n)的所有素数的倍数剔除，剩下的就是素数
    2. 如何拿到(2, √n)之前所有的素数呢? 只需要找到(2, n的四次方根)之间的素数, 继续递归, 最终坍缩为2, 而2确实是素数
    
    因此只需要从2开始迭代即可
*/

func main(){
	ret := countPrime(100)
	fmt.Println(ret)
}


func countPrime(n int) []int{
    result := make([]int, 0)
	limit := int(math.Sqrt(float64(n)))

	vector := make([]bool, n+1)

    // 设置所有数为素数
	for i := range vector{
		vector[i] = true
	}

	vector[0] = false
	vector[1] = false

    // 将素数的倍数标为false
	for i:=2;i<=limit;i++{
		if vector[i] {
			for j:=2;i*j<=n;j++{
				vector[i*j] = false
			}
		}
	}

	for i := range vector{

		if vector[i]{
			result = append(result, i)
		}
	}
	return result

}
