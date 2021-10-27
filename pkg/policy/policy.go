package policy

import (
	"io/ioutil"
	"log"
	"strings"
)

func Load(pathToDenyPolicyFile string) map[string]bool {
	log.Printf("Loading deny policy from %s", pathToDenyPolicyFile)
	denyPolicyMap := make(map[string]bool)
	denyList, err := ioutil.ReadFile(pathToDenyPolicyFile)
	if err != nil {
		log.Fatal(err)
	}

	data := strings.Split(strings.TrimSpace(string(denyList)), "\n")

	for _, l := range data {
		if strings.HasPrefix(l, "#") {
			// allow true and skip
			denyPolicyMap[l[1:]] = true
			continue
		}
		// allow false and deny
		denyPolicyMap[l] = false
	}
	return denyPolicyMap
}
