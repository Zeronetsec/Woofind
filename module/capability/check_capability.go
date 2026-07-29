// https://github.com/Zeronetsec/Woofind

package capability

import (
    "fmt"
    "strings"
    "sync"
    "syscall"
    "encoding/binary"
    "github.com/Zeronetsec/Woofind/utils/color"
    "github.com/Zeronetsec/Woofind/utils/logger"
)

func checkCapability(
    filePath string,
    targetCaps map[int]string,
    mu *sync.Mutex,
) {
    log := logger.NewLogger("capability")
    dest := make([]byte, 128)
    sz, err := syscall.Getxattr(
        filePath, "security.capability", dest,
    )

    if err != nil || sz < 20 {
        return
    }

    permitted0 := binary.LittleEndian.Uint32(dest[4:8])
    var permitted1 uint32
    if sz >= 20 {
        permitted1 = binary.LittleEndian.Uint32(dest[12:16])
    }

    var foundCaps []string
    for bit, name := range targetCaps {
        idx := bit / 32
        shift := bit % 32
        if idx == 0 {
            if (permitted0 & (1 << shift)) != 0 {
                foundCaps = append(foundCaps, name)
            }
        } else if idx == 1 {
            if (permitted1 & (1 << shift)) != 0 {
                foundCaps = append(foundCaps, name)
            }
        }
    }

    if len(foundCaps) > 0 {
        mu.Lock()
        fmt.Printf(
            "%s[+] %sFound: %s%s %s(%s%s%s)%s\n",
            color.GG, color.N, color.GG, filePath, color.DG,
            color.CC, strings.Join(foundCaps, ","), color.DG,
            color.N,
        )

        log.Log(":", fmt.Sprintf(
            "Found: %s (%s)",
            filePath, strings.Join(foundCaps, ","),
        ))

        mu.Unlock()
        idx = 1
    }
}

// Copyright (c) 2026 Zeronetsec