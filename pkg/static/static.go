package static

func Scan(StaticChan chan string) {
	go func() { StaticChan <- "static" }()
}
