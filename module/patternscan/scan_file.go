// https://github.com/Zeronetsec/Woofind

package patternscan

import (
    "fmt"
    "sync"
    "os"
    "bufio"
    "strings"
    "github.com/Zeronetsec/Woofind/utils/logger"
    "github.com/Zeronetsec/Woofind/utils/color"
)

func scanFile(
    filePath string,
    patterns []string,
    caseInsensitive bool,
    mu *sync.Mutex,
) {
    file, err := os.Open(filePath)
    if err != nil {
        return
    }
    defer file.Close()

    log := logger.NewLogger("patternscan")
    scanner := bufio.NewScanner(file)
    lineNum := 1

    idx = 0
    for scanner.Scan() {
        originalLine := scanner.Text()
        checkLine := originalLine

        if caseInsensitive {
            checkLine = strings.ToLower(originalLine)
        }

        for _, p := range patterns {
            if strings.Contains(checkLine, p) {
                mu.Lock()
                fmt.Printf(
                    "%s[+] %sFound: %s%s%s:%s%d %s-> %s%s%s\n",
                    color.GG, color.N, color.GG, filePath, color.DG,
                    color.CC, lineNum, color.DG,
                    color.YY, p, color.N,
                )

                log.Log(":", fmt.Sprintf(
                    "Found: %s:%d -> %s",
                    filePath, lineNum, p,
                ))

                mu.Unlock()
                idx++
            }
        }
        lineNum++
    }
}

// Copyright (c) 2026 Zeronetsec