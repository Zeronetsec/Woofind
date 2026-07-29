// https://github.com/Zeronetsec/Woofind

package list

import (
    "embed"
    "fmt"
    "path"
    "github.com/Zeronetsec/Woofind/utils/color"
)

//go:embed list/*
var listFS embed.FS

func Show(input string) {
    filePath := path.Join("list", input+".txt")
    data, err := listFS.ReadFile(filePath)
    if err != nil {
        fmt.Printf(
            "%s[!] %sInvalid list: %s%s%s\n",
            color.R, color.N, color.GG, input, color.N,
        )
        return
    }

    fmt.Printf(
        "%s[*] %sAvailable %s:\n",
        color.B, color.N, input,
    )
    fmt.Print(string(data))
}

// Copyright (c) 2026 Zeronetsec