// https://github.com/Zeronetsec/Woofind

package console

import (
    "github.com/Zeronetsec/Woofind/module/checkroot"
)

type Checkroot struct{}
func (c Checkroot) Execute(args []string) {
    checkroot.GetCheck()
}

// Copyright (c) 2026 Zeronetsec