// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package main

import (
	"context"
	"fmt"

	"github.com/lrstanley/go-ytdlp"
)

func main() {
	dl := ytdlp.New().
		Verbose().
		NoProgress().
		NoCallHome().
		NoCheckCertificates().
		RejectTitle(".*preview/.*").
		FormatSort("res,ext:mp4:m4a").
		RecodeVideo("mp4").
		NoPlaylist().
		NoOverwrites().
		Continue().
		Output("%(extractor)s - %(title)s.%(ext)s")

	res, err := dl.Run(context.TODO(), "https://www.youtube.com/watch?v=dQw4w9WgXcQ")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", res)
}