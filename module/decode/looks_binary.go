// https://github.com/Zeronetsec/Woofind

package decode

import (
    "strings"
)

func looksBinary(s string) bool {
    sClean := strings.ReplaceAll(s, " ", "")
    if len(sClean)%8 != 0 || len(sClean) == 0 {
        return false
    }

    for _, char := range sClean {
        if char != '0' && char != '1' {
            return false
        }
    }

    return true
}

// Copyright (c) 2026 Zeronetsec