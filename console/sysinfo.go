// https://github.com/Zeronetsec/Woofind

package console

import (
    "github.com/Zeronetsec/Woofind/module/sysinfo"
)

type Sysinfo struct{}
func (c Sysinfo) Execute(args []string) {
    sysinfo.MachineInfo()
}

// Copyright (c) 2026 Zeronetsec