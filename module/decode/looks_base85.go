// https://github.com/Zeronetsec/Woofind

package decode

import (
    "strings"
)

func looksBase85(s string) bool {
    s = strings.TrimPrefix(s, "<~")
    s = strings.TrimSuffix(s, "~>")
    if len(s) == 0 {
        return false
    }

    for _, c := range s {
        if c < '!' || c > 'u' {
            if (c != 'z' &&
                c != 'y' &&
                c != ' ' &&
                c != '\n' &&
                c != '\r') {
                    return false
            }
        }
    }

    return true
}

// Copyright (c) 2026 Zeronetsec