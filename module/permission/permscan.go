// https://github.com/Zeronetsec/Woofind

package permission

import (
    "fmt"
    "strconv"
    "strings"
    "sync"
    "io/fs"
    "path/filepath"
    "github.com/Zeronetsec/Woofind/utils/color"
    "github.com/Zeronetsec/Woofind/utils/invinput"
)

var idx int

func PermScan(targetPath, permStr string, threads int) {
    invinput.NotFound(targetPath)
    fmt.Printf(
        "%s[*] %sScanning: %s%s%s\n",
        color.B, color.N, color.GG, targetPath, color.N,
    )

    if permStr == "" {
        permStr = "0777:7777:4755:4777:2755:2777:6755:6777"
    }

    targetPerms := make(map[uint32]bool)
    for _, p := range strings.Split(permStr, ":") {
        if val, err := strconv.ParseUint(p, 8, 32); err == nil {
            targetPerms[uint32(val)] = true
        }
    }

    if len(targetPerms) == 0 {
        fmt.Printf(
            "%s[!] %sInvalid permission input format!\n",
            color.R, color.N,
        )
        return
    }

    fmt.Printf(
        "%s[*] %sPermission: %s%s%s\n",
        color.B, color.N, color.GG, permStr, color.N,
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
            targetPerms, &mu,
        )
    }

    err := filepath.WalkDir(
        targetPath, func(
            path string, d fs.DirEntry, err error,
        ) error {
            if err != nil {
                return nil
            }
            jobs <- path 
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