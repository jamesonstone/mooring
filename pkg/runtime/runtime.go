package runtime

func Scan(rChannel chan string) {
	go func() { rChannel <- "runtime" }()
}
