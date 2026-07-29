// https://github.com/Zeronetsec/Woofind

package patternscan

import (
    "sync"
)

func worker(
    jobs <-chan string,
    wg *sync.WaitGroup,
    patterns []string,
    caseInsensitive bool,
    mu *sync.Mutex,
) {
    defer wg.Done()
    for filePath := range jobs {
        scanFile(
            filePath, patterns,
            caseInsensitive, mu,
        )
    }
}

// Copyright (c) 2026 Zeronetsec