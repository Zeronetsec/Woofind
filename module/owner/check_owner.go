// https://github.com/Zeronetsec/Woofind

package owner

import (
    "fmt"
    "os"
    "syscall"
    "sync"
    "github.com/Zeronetsec/Woofind/utils/color"
    "github.com/Zeronetsec/Woofind/utils/logger"
)

func checkOwner(
    filePath string,
    targetUIDs map[uint32]string,
    mu *sync.Mutex,
) {
    info, err := os.Stat(filePath)
    if err != nil {
        return
    }

    stat, ok := info.Sys().(*syscall.Stat_t)
    if !ok {
        return
    }

    log := logger.NewLogger("owner")
    fileUID := uint32(stat.Uid)

    if ownerName, found := targetUIDs[fileUID]; found {
        tipe := "file"
        if info.IsDir() {
            tipe = "dir"
        }

        mu.Lock()
        fmt.Printf(
            "%s[+] %sFound: %s%s%s:%s%s %s(%s%s%s)%s\n",
            color.GG, color.N, color.BB, tipe, color.DG,
            color.GG, filePath, color.DG,
            color.CC, ownerName, color.DG,
            color.N,
        )

        log.Log(":", fmt.Sprintf(
            "Found: %s:%s (%s)",
            tipe, filePath, ownerName,
        ))

        mu.Unlock()
        idx = 1
    }
}

// Copyright (c) 2026 Zeronetsec