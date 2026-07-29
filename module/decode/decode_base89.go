// https://github.com/Zeronetsec/Woofind

package decode

import (
    "errors"
    "strings"
    "math/big"
)

const (
    b89Alphabet = "!#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[]^_`abcdefghijklmnopqrstuvwxyz{|}~"
)

func decodeBase89(s string) ([]byte, error) {
    num := big.NewInt(0)
    base := big.NewInt(89)

    for _, char := range s {
        idx := strings.IndexRune(b89Alphabet, char)
        if idx == -1 {
            return nil, errors.New(
                "invalid base89 character!",
            )
        }

        val := big.NewInt(int64(idx))
        num.Mul(num, base)
        num.Add(num, val)
    }

    return num.Bytes(), nil
}

// Copyright (c) 2026 Zeronetsec