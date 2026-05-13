// https://github.com/Zeronetsec/Woofind

package invinput

import (
    "fmt"
    "woofind/utils/color"
)

func MissingArgs() {
    fmt.Printf(
        "%s[!] %sMissing arguments!\n",
        color.R, color.N,
    )

    fmt.Printf(
        "%s[!] %sTry: %swoofind --help%s\n",
        color.R, color.N, color.GG, color.N,
    )
}

// Copyright (c) 2026 Zeronetsec