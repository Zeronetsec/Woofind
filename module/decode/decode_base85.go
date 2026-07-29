// https://github.com/Zeronetsec/Woofind

package decode

import (
    "bytes"
    "strings"
    "encoding/ascii85"
)

func decodeBase85(s string) ([]byte, error) {
    s = strings.TrimPrefix(s, "<~")
    s = strings.TrimSuffix(s, "~>")

    decoder := ascii85.NewDecoder(
        strings.NewReader(s),
    )

    var buf bytes.Buffer
    _, err := buf.ReadFrom(decoder)
    return buf.Bytes(), err
}

// Copyright (c) 2026 Zeronetsec