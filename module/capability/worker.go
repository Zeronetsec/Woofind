// https://github.com/Zeronetsec/Woofind

package capability

import (
    "sync"
)

func worker(
    jobs <-chan string,
    wg *sync.WaitGroup,
    targetCaps map[int]string,
    mu *sync.Mutex,
) {
    defer wg.Done()
    for filePath := range jobs {
        checkCapability(
            filePath, targetCaps,
            mu,
        )
    }
}

// Copyright (c) 2026 Zeronetsec