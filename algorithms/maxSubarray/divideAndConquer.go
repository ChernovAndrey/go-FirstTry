package main

import (
	"fmt"
)

func findMaxCrossingSubarray(s []int, low int, mid int, high int) (int, int, int) { //sum,low,high
	leftSum := s[mid]
	sum := s[mid]
	maxLeft := mid
	for i := mid - 1; i >= 0; i-- {
		sum += s[i]
		if sum > leftSum {
			leftSum = sum
			maxLeft = i
		}
	}

	rightSum := s[mid+1]
	sum = s[mid+1]
	maxRight := mid + 1
	for i := mid + 2; i <= high; i++ {
		sum += s[i]
		if sum > rightSum {
			rightSum = sum
			maxRight = i
		}
	}
	return leftSum + rightSum, maxLeft, maxRight
}
func findMaxSubarray(s []int, low int, high int) (int, int, int) { //sum,low,high
	if high == low {
		return s[low], low, high
	}
	mid := (high + low) / 2
	leftSum, leftLow, leftHigh := findMaxSubarray(s, low, mid)
	rightSum, rightLow, rightHigh := findMaxSubarray(s, mid+1, high)
	crossSum, crossLow, crossHigh := findMaxCrossingSubarray(s, low, mid, high)
	if (crossSum >= leftSum) && (crossSum >= rightSum) {
		return crossSum, crossLow, crossHigh
	}

	if (leftSum >= rightSum) && (leftSum >= crossSum) {
		return leftSum, leftLow, leftHigh
	}
	return rightSum, rightLow, rightHigh
}

func main() {
	s := []int{13, -3, -25, 20, -3, -16, -23, 18, 20, -7, 12, -5, -22, 15, -4, 7}
	fmt.Println(findMaxSubarray(s, 0, len(s)-1))
}
