package static

import (
	"io/ioutil"
	"log"
	"strings"
)

func Scan(StaticChan chan string, pathToDockerfile string, policyMap map[string]bool) {
	dockerFileMap := make(map[string]int)
	dockerfile, err := ioutil.ReadFile(pathToDockerfile)
	if err != nil {
		log.Fatal(err)
	}

	d := strings.Split(strings.TrimSpace(string(dockerfile)), "\n")

	for i, v := range d {
		if len(v) < 1 {
			continue
		}
		dockerFileMap[v] = i
		log.Println(i, v)
	}

	log.Print(dockerFileMap)

	go func() { StaticChan <- "static" }()
}
