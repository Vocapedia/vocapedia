package embed

import "embed"

//go:embed all:dist/*
var distfs embed.FS

func StaticsFS() embed.FS {
	return distfs
}
