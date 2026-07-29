// https://github.com/Zeronetsec/Woofind

package capability

import (
    "fmt"
    "strings"
    "sync"
    "io/fs"
    "path/filepath"
    "github.com/Zeronetsec/Woofind/utils/color"
    "github.com/Zeronetsec/Woofind/utils/invinput"
)

var idx int

func Getcap(targetPath, capStr string, threads int) {
    invinput.NotFound(targetPath)
    fmt.Printf(
        "%s[*] %sScanning: %s%s%s\n",
        color.B, color.N, color.GG, targetPath, color.N,
    )

    if capStr == "" {
        capStr = "cap_setuid:cap_setgid:cap_net_raw:cap_sys_admin:cap_sys_ptrace:cap_dac_override"
    }

    targetCaps := make(map[int]string)
    for _, c := range strings.Split(capStr, ":") {
        c = strings.ToLower(strings.TrimSpace(c))
        if bit, ok := capMap[c]; ok {
            targetCaps[bit] = c
        } else {
            fmt.Printf(
                "%s[!] %sInvalid capability: %s%s%s\n",
                color.YY, color.N, color.GG, c, color.N,
            )
        }
    }

    if len(targetCaps) == 0 {
        fmt.Printf(
            "%s[!] %sNothing valid capability to scan!\n",
            color.R, color.N,
        )
        return
    }

    fmt.Printf(
        "%s[*] %sCapability: %s%s%s\n",
        color.B, color.N, color.GG, capStr, color.N,
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
            targetCaps, &mu,
        )
    }

    err := filepath.WalkDir(
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
            "%s[!] %sFailed to read path: %s%v%s\n",
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