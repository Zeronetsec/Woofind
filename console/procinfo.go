// https://github.com/Zeronetsec/Woofind

package console

import (
    "github.com/Zeronetsec/Woofind/module/procinfo"
)

type Procinfo struct{}
func (c Procinfo) Execute(args []string) {
    procinfo.ShowProcInfo()
}

// Copyright (c) 2026 Zeronetsec