// https://github.com/Zeronetsec/Woofind

package decode

import (
    "errors"
    "strings"
    "math/big"
)

const (
    b58Alphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
)

func decodeBase58(s string) ([]byte, error) {
    num := big.NewInt(0)
    base := big.NewInt(58)

    for _, char := range s {
        idx := strings.IndexRune(b58Alphabet, char)
        if idx == -1 {
            return nil, errors.New(
                "invalid base58 character!",
            )
        }

        val := big.NewInt(int64(idx))
        num.Mul(num, base)
        num.Add(num, val)
    }

    var leadingZeros int
    for _, char := range s {
        if char == '1' {
            leadingZeros++
        } else {
            break
        }
    }

    res := num.Bytes()
    if leadingZeros > 0 {
        zeros := make([]byte, leadingZeros)
        res = append(zeros, res...)
    }

    return res, nil
}

// Copyright (c) 2026 Zeronetsec