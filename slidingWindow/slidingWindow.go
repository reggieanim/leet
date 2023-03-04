package slidingWindow

import (
	"fmt"
	"io"
	"math"
)

type slidingWindow struct {
	description string
	problem     string
	solution    interface{}
}

func Get() *slidingWindow {
	des := `The basic idea of the sliding window pattern is to maintain a window of elements in the input that can either grow or shrink as needed to solve the problem.
The sliding window pattern typically involves two pointers, which can be thought of as the edges of the window. The window starts out small (e.g. with a single element) and moves across the input data, growing or shrinking as needed based on the problem constraints. At each step, the solution is updated based on the current window.

The sliding window pattern is useful for solving problems that involve finding a subarray or substring that satisfies a certain condition. Examples of such problems include finding the maximum or minimum sum of a subarray, finding a substring with a certain property (e.g. all characters are unique), or finding the longest substring with no more than K distinct characters.

The sliding window pattern can lead to efficient solutions because it avoids redundant computations by reusing information from the previous window. The time complexity of the sliding window pattern is typically O(n), where n is the size of the input data.`
	sw := &slidingWindow{
		description: des,
	}
	return sw
}

func (sw *slidingWindow) Description(dst io.Writer) {
	fmt.Fprint(dst, sw.description)
}

func (sw *slidingWindow) Problem() *slidingWindow {
	fmt.Println(sw.problem)
	return sw
}

func (sw *slidingWindow) Solution() *slidingWindow {
	fmt.Println(sw.solution)
	return sw
}

// Given an array of integers Arr of size N and a number K. Return the maximum sum of a subarray of size K.

func (sw *slidingWindow) MaxSumOfSubArrOfSizeK(arr []int, k int) *slidingWindow {
	sw.problem = "Given an array of integers Arr of size N and a number K. Return the maximum sum of a subarray of size K."
	maxSum := 0
	sum := 0
	windowStart := 0
	// Create a loop which simulates and expanding window
	for windowEnd := 0; windowEnd < len(arr); windowEnd++ {
		// Get character in the window end
		rightChar := arr[windowEnd]

		// Add to sum tracker
		sum += rightChar

		// When we get to the size of window determined by k
		if windowEnd >= k-1 {
			// Get the maximum sum
			maxSum = int(math.Max(float64(sum), float64(maxSum)))
			// Get character at the start of window
			leftChar := arr[windowStart]

			// Subtract from sum tracker
			sum -= leftChar

			// Slide
			windowStart++
		}
	}
	sw.solution = maxSum
	return sw
}
