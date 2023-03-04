package main

import (
	"github.com/reggieanim/leet/slidingWindow"
)

func main() {
	arr := []int{100, 200, 300, 400}
	slidingWindow.Get().MaxSumOfSubArrOfSizeK(arr, 4).Problem().Solution()
	slidingWindow.Get().SmallestSubArrayWithGivenSum(arr, 1000).Problem().Solution()
}
