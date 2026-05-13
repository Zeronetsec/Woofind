// https://github.com/Zeronetsec/Woofind

package helper

import (
    "embed"
    "fmt"
    "encoding/json"
    "io/fs"
    "woofind/utils/color"
    "woofind/utils/birthday"
)

//go:embed metadata/*
var MetadataFS embed.FS

func WoofindHelper() {
    birthday.WoofindBirthDay()

    fmt.Printf(
        "%sUsage: %swoofind %s<command> [<args>]%s\n",
        color.N, color.GG, color.CC, color.N,
    )

    fmt.Println()
    fmt.Printf(
        "%sAvailable commands:\n",
        color.N,
    )

    files, err := fs.Glob(MetadataFS, "metadata/*.json")
    if err != nil {
        fmt.Printf(
            "%s[!] %sError reading config: %s%s%s\n",
            color.R, color.N, color.GG, err, color.N,
        )
        return
    }

    for _, f := range files {
        data, err := MetadataFS.ReadFile(f)
        if err != nil {
            continue
        }

        var hp Helper
        err = json.Unmarshal(data, &hp)
        if err != nil {
            continue
        }

        args := ""
        if hp.Args != "" {
            args = " " + hp.Args
        }

        fmt.Printf(
            "    %s› %s%s%s%s%s\n",
            color.R, color.GG, hp.Command, color.CC, args, color.N,
        )

        fmt.Printf(
            "    %s└──%s› %s%s%s\n",
            color.DG, color.B, color.WW, hp.Description, color.N,
        )
    }
}

// Copyright (c) 2026 Zeronetsec