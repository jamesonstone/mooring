package static

import (
	"io/ioutil"
	"log"
	"strings"
)

// Scan is the top level function for the static analysis
// StaticChan is the channel to send the static analysis results
// pathToDockerfile is the path to the dockerfile; default to the root directory
// policyMap is the map of policies to check against
func Scan(staticChan chan string, pathToDockerfile string, policyMap map[string]bool) {
	dockerfile, err := ioutil.ReadFile(pathToDockerfile)
	if err != nil {
		log.Fatal(err)
	}

	deniedPackagesList := policyAnalysis(fetchFinalContainerImage(strings.Split(string(dockerfile), "\n")), policyMap)

	go func() { if len(deniedPackagesList) > 0 {
		staticChan <- "❗️  The following package should be removed form the final container image: " + strings.Join(deniedPackagesList, " ")
	 } else {
		staticChan <- "✨  No vulnerable packages found in the final container image"
	}
	}()

}

func policyAnalysis(finalContainerContents []string, denyPolicyMap map[string]bool) []string {
	var (
		tokenizedRunStatements []string
		denyList               []string
	)

	for _, v := range finalContainerContents {
		if strings.HasPrefix(v, "RUN") {
			for _, y := range strings.Split(v, " ") {
				if len(y) > 0 {
					tokenizedRunStatements = append(tokenizedRunStatements, y)
				}
			}
		}
	}

	for _, v := range tokenizedRunStatements {
		if _, ok := denyPolicyMap[v]; ok {
			if !denyPolicyMap[v] {
				denyList = append(denyList, v)
			}
		} else {
			continue
		}
	}

	log.Print(denyList)
	return denyList
}

// fetchFinalContainerImage is a helper function to fetch the final container image by searching FROM commands
// for the final one
// dockerfile is the dockerfile as a string
// returns the final container image as a string slice
func fetchFinalContainerImage(dockerfile []string) []string {
	var (
		fromIndex           int
		finalContainerImage []string
	)

	for i, v := range dockerfile {
		if strings.HasPrefix(v, "FROM") {
			fromIndex = i
		}
	}

	for _, v := range dockerfile[fromIndex:] {
		if len(v) > 0 {
			finalContainerImage = append(finalContainerImage, v)
		}
	}
	return finalContainerImage
}
