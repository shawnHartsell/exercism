package luhn

import (
	"fmt"
	"strings"
	"time"
)

// Valid determines if the input is a Luhn sequence
func Valid(input string) bool {
	isLuhn, err := isLuhnSequence(generateLuhnNums(input))

	if err != nil {
		fmt.Printf("%v\n", err)
	}

	return isLuhn
}

// reads a stream of numbers and determines if is a luhn sequence
// it assumes the int stream is ordered from right to left
func isLuhnSequence(nums <-chan int, err <-chan error) (bool, error) {
	sum := 0
	double := false
	for {
		select {
		case n, ok := <-nums:
			//if the channel has closed do the final
			//calculation
			if !ok {
				fmt.Printf("%v\n", sum)
				return sum%10 == 0, nil
			}

			if double {
				n = n * 2
				if n > 9 {
					n = n - 9
				}
			}

			double = !double
			sum += n
		case e := <-err:
			return false, e
		case <-time.After(1 * time.Second):
			return false, fmt.Errorf("timeout exceeded")
		}
	}
}

/*parse the input into a stream of ints to be processed by the
luhn calculator. If there is a non-numeric value an error is emmitted
and processing stops
*/
func generateLuhnNums(input string) (<-chan int, <-chan error) {
	numsChan := make(chan int)
	errChan := make(chan error)
	go func() {
		if (len(strings.Trim(input, " "))) <= 1 {
			errChan <- fmt.Errorf("(%v) input must have at least 2 chars", input)
		}

		for i := len(input) - 1; i >= 0; i-- {
			r := input[i]
			if r == ' ' {
				continue
			}

			if r <= 47 || r > 57 {
				errChan <- fmt.Errorf("found invalid rune (%v)", string(r))
				return
			}

			numsChan <- int(r - '0')
		}

		close(numsChan)
	}()

	return numsChan, errChan
}
