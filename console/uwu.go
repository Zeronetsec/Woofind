// https://github.com/Zeronetsec/Woofind

package console

import (
    "time"
    "github.com/Zeronetsec/Woofind/utils/cursor"
    "github.com/Zeronetsec/Woofind/module/uwu"
)

type UWU struct{}
func (c UWU) Execute(args []string) {
    cursor.Hide()
    uwu.Nyanners(5 * time.Second)
    cursor.Visible()
}

// Copyright (c) 2026 Zeronetsec