package main

import("fmt")

func main22(){
	xs := []int{2,45,7,3,5,1,8,9}
	PairSum(xs, 9)
}

func PairSum(xs []int, sum int)  {
	m := make(map[int]int)

	for i :=0 ; i < len(xs) ; i++ {
		if _, ok := m[xs[i]] ; ok{
			fmt.Println(xs[i], ", ", m[xs[i]])
		}else{
			m[sum - xs[i]] = xs[i]
		}

	}

}