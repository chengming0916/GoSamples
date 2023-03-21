package main

import (
	"GoSamples/ServiceSample/pkg/crypto"
	"math/rand"
	"time"

	// _ "GoSamples/ServiceSample/pkg/metrics"
	_ "GoSamples/ServiceSample/assets"
)

func main() {
	crypto.DefaultSalt = "SampleServer"

	rand.Seed(time.Now().UnixNano())

	Execute()
}
