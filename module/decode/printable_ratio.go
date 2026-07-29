// https://github.com/Zeronetsec/Woofind

package decode

func printableRatio(data []byte) float64 {
    if len(data) == 0 {
        return 0
    }

    count := 0
    for _, b := range data {
        if ((b >= 32 && b <= 126) ||
            b == '\n' ||
            b == '\r' ||
            b == '\t') {
                count++
        }
    }

    return float64(count) / float64(len(data))
}

// Copyright (c) 2026 Zeronetsec