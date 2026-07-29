// https://github.com/Zeronetsec/Woofind

package decode

import (
    "strings"
)

func looksBase32(s string) bool {
    s = strings.TrimSpace(s)
    if len(s) == 0 || len(s)%8 != 0 {
        return false
    }

    s = strings.ToUpper(s)
    for i, c := range s {
        if !strings.Contains(
            "ABCDEFGHIJKLMNOPQRSTUVWXYZ234567=",
            string(c),
        ) {
            return false
        }

        if c == '=' && i < len(s)-6 {
            return false
        }
    }

    return true
}

// Copyright (c) 2026 Zeronetsec