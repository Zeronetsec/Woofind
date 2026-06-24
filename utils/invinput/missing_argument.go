// https://github.com/Zeronetsec/Woofind

package invinput

import (
    "fmt"
    "github.com/Zeronetsec/Woofind/utils/color"
)

func MissingArgument() {
    fmt.Printf(
        "%s[!] %sMissing argument!\n",
        color.R, color.N,
    )

    fmt.Printf(
        "%s[!] %sTry: %swoofind --help%s\n",
        color.R, color.N, color.GG, color.N,
    )
}

// Copyright (c) 2026 Zeronetsec