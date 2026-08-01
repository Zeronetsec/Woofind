// https://github.com/Zeronetsec/Woofind

package decode

func looksBase122(s string) bool {
    if len(s) == 0 {
        return false
    }

    hasNonASCII := false
    for _, c := range s {
        if c == 0x00 {
            return false
        }

        if c > 127 {
            hasNonASCII = true
        }
    }

    return hasNonASCII
}

// Copyright (c) 2026 Zeronetsec