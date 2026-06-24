// https://github.com/Zeronetsec/Woofind

package console

import (
    "github.com/Zeronetsec/Woofind/module/version"
)

type Version struct{}
func (c Version) Execute(args []string) {
    version.WoofindVersion()
}

// Copyright (c) 2026 Zeronetsec