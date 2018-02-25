// Package letter count frequency of letters in text
package letter

// ConcurrentFrequency counts frequency of letters in an array of texts concurrently
func ConcurrentFrequency(texts []string) FreqMap {
	// initialize a buffered channel of type FreqMap
	c := make(chan FreqMap, len(texts))

	// create a goroutine for each text in the input
	for _, text := range texts {
		// closure
		go func(in string) { c <- Frequency(in) }(text)
	}

	freq := FreqMap{}

	// merge all the FreqMaps in the channel into a single FreqMap
	for range texts {
		for k, v := range <-c {
			freq[k] += v
		}
	}
	return freq
}
