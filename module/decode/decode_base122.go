// https://github.com/Zeronetsec/Woofind

package decode

import (
    "errors"
)

func decodeBase122(s string) ([]byte, error) {
    data := []byte(s)
    var out []byte
    var curData int
    var curBits uint

    for i := 0; i < len(data); i++ {
        b := int(data[i])
        var val int

        if b == 0xC2 || b == 0xC3 {
            if i+1 >= len(data) {
                return nil, errors.New(
                    "invalid base122 sequence!",
                )
            }
            i++

            b2 := int(data[i])
            if b == 0xC2 {
                val = b2 & 0x7F
            } else {
                val = (b2 & 0x7F) | 0x40
            }
        } else {
            val = b
        }

        curData = (curData << 7) | val
        curBits += 7
        if curBits >= 8 {
            curBits -= 8
            out = append(out, byte(curData>>curBits))
            curData &= (1 << curBits) - 1
        }
    }

    return out, nil
}

// Copyright (c) 2026 Zeronetsec