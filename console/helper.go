// https://github.com/Zeronetsec/Woofind

package console

import (
    "github.com/Zeronetsec/Woofind/module/helper"
)

type Helper struct{}
func (c Helper) Execute(args []string) {
    helper.WoofindHelper()
}

// Copyright (c) 2026 Zeronetsec