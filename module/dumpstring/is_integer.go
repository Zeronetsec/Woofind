// https://github.com/Zeronetsec/Woofind

package dumpstring

import (
    "regexp"
)

func isInteger(s string) bool {
    match, _ := regexp.MatchString(`^-?\d+$`, s)
    return match
}

// Copyright (c) 2026 Zeronetsec