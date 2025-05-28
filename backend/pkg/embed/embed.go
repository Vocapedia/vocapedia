package embed

import "embed"

////go:embed all:dist/*
var distfs embed.FS

//go:embed all:statics/*
var staticsFS embed.FS


func DistFS() embed.FS {
	return distfs
}

func StaticsFS() embed.FS {
	return staticsFS
}
