// https://github.com/Zeronetsec/Woofind

package varaddr

import (
    "fmt"
    "os"
    "strconv"
    "unsafe"
    "github.com/Zeronetsec/Woofind/utils/color"
)

func MemoryAddress(varName string, shiftAddrStr string) {
    val := os.Getenv(varName)
    if val == "" {
        fmt.Printf(
            "%s[!] %sEnvironment variable: %s%s %sis empty value!\n",
            color.R, color.N, color.GG, varName, color.N,
        )
        return
    }

    var shift int64 = 0
    var err error

    if shiftAddrStr != "" {
        shift, err = strconv.ParseInt(
            shiftAddrStr, 10, 64,
        )
        if err != nil {
            shift = 0
        }
    }

    origPtr := uintptr(unsafe.Pointer(&val))
    shiftedPtr := origPtr + uintptr(shift)

    fmt.Printf(
        "%s%s%s: %s%s %s-> %s0x%x%s:%s0x%x%s\n",
        color.BB, varName, color.N,
        color.GG, val, color.DG,
        color.CC, origPtr, color.DG,
        color.YY, shiftedPtr, color.N,
    )
}

// Copyright (c) 2026 Zeronetsec