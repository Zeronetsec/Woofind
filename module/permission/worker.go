// https://github.com/Zeronetsec/Woofind

package permission

import (
    "sync"
)

func worker(
    jobs <-chan string,
    wg *sync.WaitGroup,
    targetPerms map[uint32]bool,
    mu *sync.Mutex,
) {
    defer wg.Done()
    for filePath := range jobs {
        checkFilePerm(
            filePath, targetPerms,
            mu,
        )
    }
}

// Copyright (c) 2026 Zeronetsec