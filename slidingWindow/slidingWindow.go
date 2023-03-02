package slidingWindow

import (
	"fmt"
	"io"
)

type slidingWindow struct {
	description string
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
