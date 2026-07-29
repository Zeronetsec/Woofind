// https://github.com/Zeronetsec/Woofind

package patternscan

import (
    "embed"
    "fmt"
    "strings"
    "sync"
    "io/fs"
    "path/filepath"
    "github.com/Zeronetsec/Woofind/utils/color"
    "github.com/Zeronetsec/Woofind/utils/invinput"
)

//go:embed patterns/patterns.txt
var defaultPatternsFS embed.FS
var idx int

func Scan(
    targetPath, patternFile string,
    threads int,
    caseInsensitive bool,
) {
    invinput.NotFound(targetPath)
    if patternFile != "" {
        invinput.NotFound(patternFile)
    }

    fmt.Printf(
        "%s[*] %sScanning: %s%s%s\n",
        color.B, color.N, color.GG, targetPath, color.N,
    )

    patterns, err := loadPatterns(patternFile)
    if err != nil {
        fmt.Printf(
            "%s[!] %sError to load patterns: %s%v%s\n",
            color.R, color.N, color.GG, err, color.N,
        )
        return
    }

    if len(patterns) == 0 {
        fmt.Printf(
            "%s[!] %sPatterns file: %s%s %sis empty line!\n",
            color.R, color.N, color.GG, patternFile, color.N,
        )
        return
    }

    if caseInsensitive {
        for i := range patterns {
            patterns[i] = strings.ToLower(patterns[i])
        }
    }

    printPattern := ""
    if patternFile == "" {
        printPattern = "runtime:module/patternscan/patterns/patterns.txt"
    } else {
        printPattern = patternFile
    }

    fmt.Printf(
        "%s[*] %sLoaded: %s%d %sfrom %s%s%s\n",
        color.B, color.N, color.GG, len(patterns),
        color.N, color.GG, printPattern, color.N,
    )

    fmt.Printf(
        "%s[*] %sThreads: %s%d%s\n",
        color.B, color.N, color.GG, threads, color.N,
    )

    fmt.Println()

    jobs := make(chan string, 100)
    var wg sync.WaitGroup
    var mu sync.Mutex

    for i := 0; i < threads; i++ {
        wg.Add(1)
        go worker(
            jobs, &wg,
            patterns, caseInsensitive, &mu,
        )
    }

    err = filepath.WalkDir(
        targetPath, func(
            path string, d fs.DirEntry, err error,
        ) error {
            if err != nil {
                return nil
            }
            if !d.IsDir() {
                jobs <- path
            }
            return nil
        },
    )

    if err != nil {
        fmt.Printf(
            "%s[!] %sFailed reading path: %s%v%s\n",
            color.R, color.N, color.GG, err, color.N,
        )
    }

    close(jobs)
    wg.Wait()

    if idx > 0 {
        fmt.Println()
    } else {
        fmt.Printf(
            "%s[!] %sNothing found!\n",
            color.R, color.N,
        )
        fmt.Println()
    }

    fmt.Printf(
        "%s[*] %sScanning done.\n",
        color.B, color.N,
    )
}

// Copyright (c) 2026 Zeronetsec