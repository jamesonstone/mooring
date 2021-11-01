package static

import "strings"

// runCmdAnalysis analyzes Dockerfile RUN commands within the final container
// looking for high risk packages
//
func runCmdAnalysis(finalContainerContents []string) []string {
	var tokenizedRunStatements []string
	for _, v := range finalContainerContents {
		if strings.HasPrefix(v, "RUN") {
			for _, y := range strings.Split(v, " ") {
				if len(y) > 0 {
					tokenizedRunStatements = append(tokenizedRunStatements, y)
				}
			}
		}
	}

	return tokenizedRunStatements
}
