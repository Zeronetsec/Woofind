// https://github.com/Zeronetsec/Woofind

package misconfind

import (
    "fmt"
    "os"
    "github.com/Zeronetsec/Woofind/utils/color"
    "github.com/Zeronetsec/Woofind/utils/logger"
)

func checkPermission(path string, info os.FileInfo) {
    mode := info.Mode().Perm()
    log := logger.NewLogger("misconfind")
    logMess := ""

    if mode&0002 != 0 {
        fmt.Printf(
            "%s[Warning] %sWorld-writable: %s%s %s(%s%o%s)%s\n",
            color.YY, color.N, color.GG, path, color.DG, color.CC, mode, color.DG, color.N,
        )

        logMess = fmt.Sprintf(
            "Warning world-writable: %s (%o)",
            path, mode,
        )

        log.Log(":", logMess)
    }

    if mode&0004 != 0 {
        fmt.Printf(
            "%s[Info] %sWorld-readable: %s%s %s(%s%o%s)%s\n",
            color.B, color.N, color.GG, path, color.DG, color.CC, mode, color.DG, color.N,
        )

        logMess = fmt.Sprintf(
            "info world-readable: %s (%o)",
            path, mode,
        )

        log.Log(":info", logMess)
    }

    if mode&0111 != 0 && !info.IsDir() {
        fmt.Printf(
            "%s[Info] %sExecutable file: %s%s %s(%s%o%s)%s\n",
            color.B, color.N, color.GG, path, color.DG, color.CC, mode, color.DG, color.N,
        )

        logMess = fmt.Sprintf(
            "info executable files: %s (%o)",
            path, mode,
        )

        log.Log(":info", logMess)
    }
}

// Copyright (c) 2026 Zeronetsec