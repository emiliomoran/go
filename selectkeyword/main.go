package selectkeyword

func emit(wordChannel chan string, done chan bool) {
	words := []string{"The", "quick", "brown", "fox"}

	i := 0

	for {
		select {
		case wordChannel <- words[i]:
			i += 1
			if i == len(words) {
				i = 0
			}

		case <-done:
			return
		}
	}
}
