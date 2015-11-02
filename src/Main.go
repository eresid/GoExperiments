package main

import (
	"fmt"
)


// https://codility.com/
func main() {
	elements := []int{3, 1, 2, 4, 3}

	fmt.Println("len = ", len(elements))
	//result := Solution(elements)
	//fmt.Println("result = ", result)
}

//func Solution(A []int) int {
//	sumOfAll := ((len(A) + 1) * (len(A) + 2)) / 2;
//	for i := 0; i < len(A); i++ {
//		sumOfAll -= A[i];
//	}
//	return sumOfAll;
//}

func Solution(X int, Y int, D int) int {
	//return ((Y - X) + (D - 1)) / D
	distance := Y - X;
	ostatok := distance % D;
	if ostatok == 0 {
		return distance / D;
	} else {
		return distance / D + 1;
	}
}

//func Solution(A []int) int {
//	sumPre := A[0];
//	sumPost := 0;
//	for i := 1; i < len(A); i++ {
//		sumPost += A[i];
//	}
//	difMin := abs(sumPost - sumPre);
//	tempSub := 0;
//	for i := 1; i < len(A) - 1; i++ {
//		sumPre += A[i];
//		sumPost -= A[i];
//		tempSub = abs(sumPost - sumPre);
//		if (tempSub < difMin) {
//			difMin = tempSub;
//			// idx = i+1;
//		}
//	}
//	return difMin;
//}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func sumNat() {

}