// https://github.com/Zeronetsec/Woofind

package cursor

import (
    "fmt"
)

func Visible() {
    fmt.Print("\x1b[?25h")
}

// Copyright (c) 2026 Zeronetsec