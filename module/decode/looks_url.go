// https://github.com/Zeronetsec/Woofind

package decode

import (
    "strings"
)

func looksURL(s string) bool {
    return strings.Contains(s, "%")
}

// Copyright (c) 2026 Zeronetsec