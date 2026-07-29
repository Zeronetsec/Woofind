// https://github.com/Zeronetsec/Woofind

package decode

import (
    "strings"
)

func looksBase89(s string) bool {
    if len(s) == 0 {
        return false
    }

    for _, c := range s {
        if !strings.ContainsRune(b89Alphabet, c) {
            return false
        }
    }

    return true
}

// Copyright (c) 2026 Zeronetsec