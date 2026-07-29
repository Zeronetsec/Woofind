// https://github.com/Zeronetsec/Woofind

package decode

func looksBase91(s string) bool {
    if len(s) == 0 {
        return false
    }

    for _, c := range s {
        if c > 255 || b91DecTable[c] == -1 {
            return false
        }
    }

    return true
}

// Copyright (c) 2026 Zeronetsec