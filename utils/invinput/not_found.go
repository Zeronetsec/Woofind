// https://github.com/Zeronetsec/Woofind

package invinput

import (
    "fmt"
    "os"
    "strings"
    "github.com/Zeronetsec/Woofind/utils/color"
)

func NotFound(input string) bool {
    _, err := os.Stat(input)
    if err == nil {
        return true
    }

    if os.IsNotExist(err) {
        if strings.HasSuffix(input, "/") {
            fmt.Printf(
                "%s[!] %sDirectory: %s%s %snot found!\n",
                color.R, color.N, color.GG, input, color.N,
            )
            os.Exit(1)
        } else {
            fmt.Printf(
                "%s[!] %sFile: %s%s %snot found!\n",
                color.R, color.N, color.GG, input, color.N,
            )
            os.Exit(1)
        }
    } else {
        fmt.Printf(
            "%s[!] %sError accessing path: %s%s %s(%s%v%s)%s\n",
            color.R, color.N, color.GG, input, color.DG,
            color.CC, err, color.DG,
            color.N,
        )
        os.Exit(1)
    }

    return false
}

// Copyright (c) 2026 Zeronetsec