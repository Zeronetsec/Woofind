// https://github.com/Zeronetsec/Woofind

package decode

import (
    "errors"
    "strings"
)

func decodeMorse(s string) ([]byte, error) {
    var out strings.Builder
    s = strings.ReplaceAll(s, "/", "  ")
    words := strings.Split(s, "  ")

    for i, word := range words {
        letters := strings.Split(
            strings.TrimSpace(word), " ",
        )
        for _, letter := range letters {
            if letter == "" {
                continue
            }
            if char, ok := morseMap[letter]; ok {
                out.WriteString(char)
            } else {
                return nil, errors.New(
                    "invalid morse character: " + letter,
                )
            }
        }
        if i < len(words)-1 {
            out.WriteString(" ")
        }
    }

    return []byte(out.String()), nil
}

// Copyright (c) 2026 Zeronetsec