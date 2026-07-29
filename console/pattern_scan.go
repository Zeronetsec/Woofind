// https://github.com/Zeronetsec/Woofind

package console

import (
    "os"
    "strconv"
    "strings"
    "github.com/Zeronetsec/Woofind/module/patternscan"
    "github.com/Zeronetsec/Woofind/utils/invinput"
)

type PatternScan struct{}
func (c PatternScan) Execute(args []string) {
    if len(args) < 3 {
        invinput.MissingArgument()
        os.Exit(1)
    }

    targetPath := args[2]
    patternFile := ""
    threads := 100
    caseInsensitive := false

    for i := 3; i < len(args); i++ {
        switch args[i] {
            case "--threads":
                if i+1 < len(args) {
                    t, err := strconv.Atoi(args[i+1])
                    if err == nil && t > 0 {
                        threads = t
                    }
                    i++
                }
            case "--case-insensitive":
                if i+1 < len(args) {
                    if args[i+1] == "true" {
                        caseInsensitive = true
                    }
                    i++
                }
            default:
                if (patternFile == "" &&
                    !strings.HasPrefix(args[i], "--")) {
                        patternFile = args[i]
                }
        }
    }

    patternscan.Scan(
        targetPath,
        patternFile,
        threads,
        caseInsensitive,
    )
}

// Copyright (c) 2026 Zeronetsec