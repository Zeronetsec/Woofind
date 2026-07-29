// https://github.com/Zeronetsec/Woofind

package decode

import (
    "net/url"
)

func decodeURL(s string) ([]byte, error) {
    decoded, err := url.QueryUnescape(s)
    if err != nil {
        return nil, err
    }
    return []byte(decoded), nil
}

// Copyright (c) 2026 Zeronetsec