// https://github.com/Zeronetsec/Woofind

package birthday

import (
    "fmt"
    "time"
    "woofind/utils/color"
)

func WoofindBirthDay() {
    birthDate := "03-27"
    now := time.Now().Format("01-02")
    if now == birthDate {
        fmt.Printf(
            "%s› %sHappy birthday for %swoofind %s🎉\n",
            color.R, color.N, color.GG, color.N,
        )
        fmt.Println()
    }
}

// Copyright (c) 2026 Zeronetsec