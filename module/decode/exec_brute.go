// https://github.com/Zeronetsec/Woofind

package decode

import (
    "fmt"
    "os"
    "strconv"
    "strings"
    "time"
    "github.com/Zeronetsec/Woofind/utils/color"
)

var visited = make(map[string]bool)
var steps int
var idx int

func ExecBrute(
    input string,
    limitStr string,
    disableStr string,
) {
    disabledMap := make(map[string]bool)
    if disableStr != "" {
        parts := strings.Split(disableStr, ":")
        for _, p := range parts {
            disabledMap[strings.TrimSpace(p)] = true
        }
    }

    maxDepth := -1
    if limitStr != "unlimit" {
        val, err := strconv.Atoi(limitStr)
        if err != nil || val < 0 {
            fmt.Printf(
                "%s[!] %sInvalid limit value!\n",
                color.R, color.N,
            )
            maxDepth = -1
        } else {
            maxDepth = val
        }
    }

    tipe := ""
    var data string

    if info, err := os.Stat(
        input,
    ); err == nil && !info.IsDir() {
        content, err := os.ReadFile(input)
        if err != nil {
            fmt.Printf(
                "%s[!] %sError reading file: %s%v%s\n",
                color.R, color.GG, err, color.N,
            )
            return
        }
        tipe = "file"
        data = string(content)
    } else {
        tipe = "stdin"
        data = input
        if strings.TrimSpace(data) == "" {
            fmt.Printf(
                "%s[!] %sInvalid input!\n",
                color.R, color.N,
            )
            return
        }
    }

    fmt.Printf(
        "%s[*] %sDecode: %s%s%s:%s%s%s\n",
        color.B, color.N, color.BB, tipe, color.DG,
        color.GG, input, color.N,
    )

    fmt.Printf(
        "%s[*] %sLimit: %s%s%s\n",
        color.B, color.N, color.GG, limitStr, color.N,
    )

    disableStrPrint := "-"
    if disableStr != "" {
        disableStrPrint = disableStr
    }

    fmt.Printf(
        "%s[*] %sDisable: %s%s%s\n",
        color.B, color.N, color.GG, disableStrPrint, color.N,
    )

    fmt.Println()

    visited = make(map[string]bool)
    steps = 0
    idx = 0

    start := time.Now()
    visited[data] = true
    brute(
        []byte(data), 0,
        maxDepth, disabledMap,
    )

    if idx > 0 {
        fmt.Println()
    } else {
        fmt.Printf(
            "%s[!] %sNothing valid to decode!\n",
            color.R, color.N,
        )
        fmt.Println()
    }

    fmt.Printf(
        "%sFinished in %s%dms %s| %sSteps: %s%d%s\n",
        color.N, color.GG, time.Since(start).Milliseconds(), color.DG,
        color.N, color.GG, steps, color.N,
    )
}

// Copyright (c) 2026 Zeronetsec