// https://github.com/Zeronetsec/Woofind

package decode

func decodeRot13(s string) ([]byte, error) {
    out := make([]byte, len(s))

    for i := 0; i < len(s); i++ {
        c := s[i]
        if c >= 'a' && c <= 'z' {
            out[i] = 'a' + (c-'a'+13)%26
        } else if c >= 'A' && c <= 'Z' {
            out[i] = 'A' + (c-'A'+13)%26
        } else {
            out[i] = c
        }
    }

    return out, nil
}

// Copyright (c) 2026 Zeronetsec