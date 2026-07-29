// https://github.com/Zeronetsec/Woofind

package console

import (
    "github.com/Zeronetsec/Woofind/module/dumpenvaddr"
)

type DumpEnvAddr struct{}
func (c DumpEnvAddr) Execute(args []string) {
    shiftAddr := ""
    for i := 2; i < len(args); i++ {
        switch args[i] {
            case "--shiftaddr":
                if i+1 < len(args) {
                    shiftAddr = args[i+1]
                    i++
                }
        }
    }

    dumpenvaddr.MemoryDump(shiftAddr)
}

// Copyright (c) 2026 Zeronetsec