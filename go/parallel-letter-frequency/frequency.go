package letter

import "sync"

//FreqMap represents a frequency map for
//counting the number of characters in a string
type FreqMap map[rune]int

//Frequency counts the number of letters that
//are in the specified string
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

//ConcurrentFrequency creates a FreqMap for all
//of the provided strings concurrently by generating a
//goroutine for each
func ConcurrentFrequency(input []string) FreqMap {
	m := FreqMap{}
	//create an unbuffered channel to avoid using locks
	//while merging worker results to a final map.
	//alternatively, a buffered channel and a sync.Map could be used
	c := make(chan FreqMap)

	//create N workers to create FreqMap for each string
	var wg sync.WaitGroup
	for _, i := range input {
		wg.Add(1)
		go func(s string) {
			m := Frequency(s)
			c <- m
			wg.Done()
		}(i)
	}

	//start a goroutine that closes the channel once
	//all of the workers have completed
	go func() {
		wg.Wait()
		close(c)
	}()

	//Merge the results of each worker into a single
	//FreqMap
	for subMap := range c {
		for k, v := range subMap {
			m[k] += v
		}
	}

	return m
}
