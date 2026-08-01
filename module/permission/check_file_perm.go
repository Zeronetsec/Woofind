// https://github.com/Zeronetsec/Woofind

package permission

import (
    "fmt"
    "sync"
    "os"
    "github.com/Zeronetsec/Woofind/utils/color"
    "github.com/Zeronetsec/Woofind/utils/logger"
)

func checkFilePerm(
    filePath string,
    targetPerms map[uint32]bool,
    mu *sync.Mutex,
) {
    info, err := os.Stat(filePath)
    if err != nil {
        return
    }

    log := logger.NewLogger("permission")
    mode := info.Mode()
    perm := uint32(mode.Perm())

    if mode&os.ModeSetuid != 0 {
        perm |= 04000
    }

    if mode&os.ModeSetgid != 0 {
        perm |= 02000
    }

    if mode&os.ModeSticky != 0 {
        perm |= 01000
    }

    if targetPerms[perm] {
        strPerm := fmt.Sprintf("%04o", perm)
        tipe := "file"
        if info.IsDir() {
            tipe = "dir"
        }

        mu.Lock()
        fmt.Printf(
            "%s[+] %sFound: %s%s%s:%s%s %s(%s%s%s)%s\n",
            color.GG, color.N, color.BB, tipe, color.DG,
            color.GG, filePath, color.DG,
            color.CC, strPerm, color.DG,
            color.N,
        )

        log.Log(":", fmt.Sprintf(
            "Found: %s:%s (%s)",
            tipe, filePath, strPerm,
        ))

        mu.Unlock()
        idx = 1
    }
}

// Copyright (c) 2026 Zeronetsec