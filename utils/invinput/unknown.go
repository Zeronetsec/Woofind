// https://github.com/Zeronetsec/Woofind

package invinput

import (
    "fmt"
    "woofind/utils/color"
)

func Unknown(input string) {
    fmt.Printf(
        "%s[!] %sInvalid command: %s%s%s\n",
        color.R, color.N, color.GG, input, color.N,
    )

    fmt.Printf(
        "%s[!] %sTry: %swoofind --help%s\n",
        color.R, color.N, color.GG, color.N,
    )
}

// Copyright (c) 2026 Zeronetsec