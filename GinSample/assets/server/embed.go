package server

import (
	"GoSamples/GinSample/assets"
	"embed"
)

//go:embed static/*
var content embed.FS

func init() {
	assets.Register(content)
}
