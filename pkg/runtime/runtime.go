package runtime

import "log"

func Scan(containerRegistryURIWithTag string, rChannel chan string) {
	if len(containerRegistryURIWithTag) > 0 {
		log.Println("Prebuilt image detected. Skipping to docker scan")
	} else {

	}

	scanWithDockerScan(containerRegistryURIWithTag)

	go func() { rChannel <- string("hey") }()
}

func buildTemp() {
	// check for image locally; build if it doesn't exist

	// fail if we can't build
}

func scanWithDockerScan(containerImageURI string) {
	// func Scan(rChannel chan string, containerRegistryURI, versionTag string) {
	// out, error := exec.Command("docker", "scan", containerRegistryURI + ":" + versionTag).Output()
	// if error != nil {
	// 	log.Fatal("Container re-tagging failed: ", error.Error())
	// }
}
