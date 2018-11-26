package letter

// FreqMap maps letter with its occurance.
type FreqMap map[rune]int

// Frequency counts the frequency of letters in texts.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts the frequency of letters in text using parallel processes.
func ConcurrentFrequency(ss []string) FreqMap {
	ch := make(chan FreqMap, len(ss))
	for _, s := range ss {
		go func(s string) {
			ch <- Frequency(s)
		}(s)
	}
	freq := FreqMap{}
	for range ss {
		for k, v := range <-ch {
			freq[k] += v
		}
	}
	return freq
}
