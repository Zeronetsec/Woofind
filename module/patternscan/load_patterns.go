// https://github.com/Zeronetsec/Woofind

package patternscan

import (
    "os"
    "bufio"
    "strings"
    "io/fs"
)

func loadPatterns(patternPath string) ([]string, error) {
    var scanner *bufio.Scanner
    var file fs.File
    var err error

    if patternPath == "" {
        file, err = defaultPatternsFS.Open(
            "patterns/patterns.txt",
        )
    } else {
        file, err = os.Open(patternPath)
    }

    if err != nil {
        return nil, err
    }
    defer file.Close()

    scanner = bufio.NewScanner(file)
    var patterns []string

    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())
        if line == "" || strings.HasPrefix(line, "#") {
            continue
        }
        patterns = append(patterns, line)
    }

    return patterns, scanner.Err()
}

// Copyright (c) 2026 Zeronetsec