// https://github.com/Zeronetsec/Woofind

package owner

import (
    "fmt"
    "strconv"
    "strings"
    "sync"
    "io/fs"
    "os/user"
    "path/filepath"
    "github.com/Zeronetsec/Woofind/utils/color"
    "github.com/Zeronetsec/Woofind/utils/invinput"
)

var idx int

func OwnScan(
    targetPath, ownerStr string,
    threads int,
    force bool,
) {
    invinput.NotFound(targetPath)
    fmt.Printf(
        "%s[*] %sScanning: %s%s%s\n",
        color.B, color.N, color.GG, targetPath, color.N,
    )

    if ownerStr == "" {
        ownerStr = "root:www-data:admin"
    }

    targetUIDs := make(map[uint32]string)
    for _, u := range strings.Split(ownerStr, ":") {
        u = strings.TrimSpace(u)
        if u == "" {
            continue
        }

        usr, err := user.Lookup(u)
        if err == nil {
            uid, _ := strconv.ParseUint(usr.Uid, 10, 32)
            targetUIDs[uint32(uid)] = usr.Username
        } else {
            if uid, err2 := strconv.ParseUint(u, 10, 32); err2 == nil {
                if usrById, err3 := user.LookupId(u); err3 == nil {
                    targetUIDs[uint32(uid)] = usrById.Username
                } else {
                    targetUIDs[uint32(uid)] = u
                }
            } else {
                if !force {
                    fmt.Printf(
                        "%s[!] %sUser: %s%s %snot found, try using UID!\n",
                        color.YY, color.N, color.GG, u, color.N,
                    )
                }
            }
        }
    }

    if len(targetUIDs) == 0 {
        fmt.Printf(
            "%s[!] %sNothing valid UID to scan!\n",
            color.R, color.N,
        )
        return
    }

    fmt.Printf(
        "%s[*] %sForce: %s%t%s\n",
        color.B, color.N, color.GG, force, color.N,
    )

    fmt.Printf(
        "%s[*] %sOwner: %s%s%s\n",
        color.B, color.N, color.GG, ownerStr, color.N,
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
            targetUIDs, &mu,
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