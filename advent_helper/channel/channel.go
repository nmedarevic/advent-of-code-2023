package channel

func IterateThroughInputChannel(lines <-chan []string, cb func(string)) {
	for task := range lines {
		for _, line := range task {
			cb(line)
		}
	}
}
