// https://github.com/Zeronetsec/Woofind

package dumpenvaddr

import (
    "bytes"
    "fmt"
    "strings"
    "os/exec"
    "github.com/Zeronetsec/Woofind/module/varaddr"
    "github.com/Zeronetsec/Woofind/utils/color"
)

func MemoryDump(shiftAddr string) {
    fmt.Printf(
        "%s[*] %sDumping environment variables...\n",
        color.B, color.N,
    )

    cmd := exec.Command("env")
    var out bytes.Buffer
    cmd.Stdout = &out

    err := cmd.Run()
    if err != nil {
        fmt.Printf(
            "%s[!] %sFailed to execute: %senv %s(%s%v%s)%s\n",
            color.R, color.N, color.GG, color.DG,
            color.CC, err, color.DG, color.N,
        )
        return
    }

    lines := strings.Split(out.String(), "\n")
    count := 0

    for _, line := range lines {
        if line == "" {
            continue
        }

        parts := strings.SplitN(line, "=", 2)
        if len(parts) == 2 {
            varName := parts[0]
            varaddr.MemoryAddress(varName, shiftAddr)
            count++
        }
    }
}

// Copyright (c) 2026 Zeronetsec