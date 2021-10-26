package main

import (
	"log"

	"github.com/jstone28/mooring/pkg/policy"
	"github.com/jstone28/mooring/pkg/runtime"
	"github.com/jstone28/mooring/pkg/static"
)

func main() {
	var (
		sChannel           = make(chan string)
		rChannel           = make(chan string)
		denyPolicyDocument = "./DENY.txt"
		pathToDockerfile   = "./Dockerfile"
		policyMap          = policy.Load(denyPolicyDocument)
	)

	static.Scan(sChannel, pathToDockerfile, policyMap)
	runtime.Scan(rChannel)

	s := <-sChannel
	r := <-rChannel

	log.Println(r, s)
}
