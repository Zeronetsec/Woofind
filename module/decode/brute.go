// https://github.com/Zeronetsec/Woofind

package decode

import (
    "fmt"
    "strings"
    "encoding/base32"
    "encoding/base64"
    "encoding/hex"
    "github.com/Zeronetsec/Woofind/utils/color"
)

func brute(
    data []byte,
    depth int,
    maxDepth int,
    disabledMap map[string]bool,
) {
    if maxDepth != -1 && depth >= maxDepth {
        return
    }

    input := strings.TrimSpace(string(data))
    if input == "" {
        return
    }

    decoders := []struct {
        name string
        check func(string) bool
        decode func(string) ([]byte, error)
    }{
        {
            "base16",
            looksBase16,
            func(s string) ([]byte, error) {
                return hex.DecodeString(s)
            },
        },

        {
            "base32",
            looksBase32,
            func(s string) ([]byte, error) {
                return base32.StdEncoding.DecodeString(s)
            },
        },

        {
            "base64",
            looksBase64,
            func(s string) ([]byte, error) {
                return base64.StdEncoding.DecodeString(s)
            },
        },

        {
            "base85",
            looksBase85,
            decodeBase85,
        },

        {
            "base89",
            looksBase89,
            decodeBase89,
        },

        {
            "base91",
            looksBase91,
            decodeBase91,
        },

        {
            "base122",
            looksBase122,
            decodeBase122,
        },

        {
            "binary",
            looksBinary,
            decodeBinary,
        },

        {
            "base58",
            looksBase58,
            decodeBase58,
        },

        {
            "rot13",
            looksRot13,
            decodeRot13,
        },

        {
            "url",
            looksURL,
            decodeURL,
        },

        {
            "morse",
            looksMorse,
            decodeMorse,
        },
    }

    for _, d := range decoders {
        if disabledMap[d.name] {
            continue
        }

        steps++

        if !d.check(input) {
            continue
        }

        decoded, err := d.decode(input)
        if err != nil {
            continue
        }

        str := strings.TrimSpace(string(decoded))
        if visited[str] {
            continue
        }

        visited[str] = true
        decoded = []byte(str)

        if !looksReasonable(decoded) {
            continue
        }

        fmt.Printf(
            "%s[*] %sBrute: %slay%s=%s%d %sdec%s=%s%s %slen%s=%s%d %sprtl%s=%s%.1f%% %sent%s=%s%.2f %sres%s=%s%s%s\n",
            color.B, color.N,
            color.GG, color.DG, color.CC, depth+1,
            color.GG, color.DG, color.CC, d.name,
            color.GG, color.DG, color.CC, len(decoded),
            color.GG, color.DG, color.CC, printableRatio(decoded)*100,
            color.GG, color.DG, color.CC, entropy(decoded),
            color.GG, color.DG, color.YY, str, color.N,
        )

        idx = 1
        brute(decoded, depth+1, maxDepth, disabledMap)
    }
}

// Copyright (c) 2026 Zeronetsec