// https://github.com/Zeronetsec/Woofind

package misconfind

import (
    "os"
    "bufio"
    "strings"
    "fmt"
    "github.com/Zeronetsec/Woofind/utils/color"
    "github.com/Zeronetsec/Woofind/utils/logger"
)

func checkFileContent(path string, patterns []string) {
    file, err := os.Open(path)
    if err != nil {
        return
    }
    defer file.Close()

    if isBinary(file) {
        return
    }

    file.Seek(0, 0)

    log := logger.NewLogger("misconfind")
    scanner := bufio.NewScanner(file)
    lineNum := 0

    for scanner.Scan() {
        lineNum++
        line := strings.ToLower(scanner.Text())

        for _, p := range patterns {
            if strings.Contains(line, p) {
                fmt.Printf(
                    "%s[Alert] %sPattern: %s%s %sin %s%s%s:%s%d%s\n",
                    color.R, color.N, color.GG, p, color.N, color.GG, path, color.DG, color.GG, lineNum, color.N,
                )

                logMess := fmt.Sprintf(
                    "Alert pattern: %s in %s (%d)",
                    p, path, lineNum,
                )

                log.Log(":", logMess)
            }
        }
    }
}

// Copyright (c) 2026 Zeronetsec