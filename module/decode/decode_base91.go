// https://github.com/Zeronetsec/Woofind

package decode

import (
    "bytes"
)

var b91Alphabet = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!#$%&()*+,./:;<=>?@[]^_`{|}~\"")
var b91DecTable [256]int

func decodeBase91(s string) ([]byte, error) {
    var out bytes.Buffer
    b, n, v := 0, 0, -1

    for i := 0; i < len(s); i++ {
        char := s[i]
        if b91DecTable[char] == -1 {
            continue
        }

        if v < 0 {
            v = b91DecTable[char]
        } else {
            v += b91DecTable[char] * 91
            b |= v << n
            if (v & 8191) > 88 {
                n += 13
            } else {
                n += 14
            }

            for {
                out.WriteByte(byte(b & 255))
                b >>= 8
                n -= 8
                if n <= 7 {
                    break
                }
            }
            v = -1
        }
    }

    if v > -1 {
        out.WriteByte(byte((b | v<<n) & 255))
    }

    return out.Bytes(), nil
}

// Copyright (c) 2026 Zeronetsec