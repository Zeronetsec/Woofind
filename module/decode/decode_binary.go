// https://github.com/Zeronetsec/Woofind

package decode

import (
    "errors"
    "strconv"
    "strings"
)

func decodeBinary(s string) ([]byte, error) {
    s = strings.ReplaceAll(s, " ", "")
    if len(s)%8 != 0 {
        return nil, errors.New(
            "the binary length is not a multiple of 8!",
        )
    }

    var decoded []byte
    for i := 0; i < len(s); i += 8 {
        chunk := s[i : i+8]
        val, err := strconv.ParseUint(chunk, 2, 8)
        if err != nil {
            return nil, err
        }
        decoded = append(decoded, byte(val))
    }

    return decoded, nil
}

// Copyright (c) 2026 Zeronetsec