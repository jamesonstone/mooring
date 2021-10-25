package policy

import (
	"io/ioutil"
	"log"
	"strings"
)

func Load(pathToPolicyFile string) {
	policyData, err := ioutil.ReadFile(pathToPolicyFile)
	if err != nil {
		log.Fatal(err)
	}

	data := strings.Split(strings.TrimSpace(string(policyData)), "\n")

	for _, l := range data {
		if strings.HasPrefix(l, "#") {
			continue
		}

		if strings.HasPrefix(l, "allow") {
			log.Println(l)
		} else if strings.HasPrefix(l, "deny") {
			deny(l)
		} else {
			log.Fatal("Unknown policy statement: " + l)
		}
	}

}
