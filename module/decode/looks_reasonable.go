// https://github.com/Zeronetsec/Woofind

package decode

func looksReasonable(data []byte) bool {
    return printableRatio(data) >= 0.8 && len(data) >= 3
}

// Copyright (c) 2026 Zeronetsec