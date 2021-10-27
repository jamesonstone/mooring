package runtime

func Scan(rChannel chan string) {
	out := "hello, runtime"
	// func Scan(rChannel chan string, containerRegistryURI, versionTag string) {
	// out, error := exec.Command("docker", "scan", containerRegistryURI + ":" + versionTag).Output()
	// if error != nil {
	// 	log.Fatal("Container re-tagging failed: ", error.Error())
	// }

	go func() { rChannel <- string(out) }()
}
