// https://github.com/Zeronetsec/Woofind

package main

import (
    "os"
    "strings"
    "github.com/Zeronetsec/Woofind/console"
)

func main() {
    args := os.Args[1:]
    input := strings.Join(args, " ")
    console.WoofindConsole(input)
}

// Copyright (c) 2026 Zeronetsec