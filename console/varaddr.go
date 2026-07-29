// https://github.com/Zeronetsec/Woofind

package console

import (
    "os"
    "github.com/Zeronetsec/Woofind/module/varaddr"
    "github.com/Zeronetsec/Woofind/utils/invinput"
)

type VarAddr struct{}
func (c VarAddr) Execute(args []string) {
    if len(args) < 3 {
        invinput.MissingArgument()
        os.Exit(1)
    }

    varName := args[2]
    shiftAddr := ""

    for i := 3; i < len(args); i++ {
        switch args[i] {
            case "--shiftaddr":
                if i+1 < len(args) {
                    shiftAddr = args[i+1]
                    i++
                }
        }
    }

    varaddr.MemoryAddress(varName, shiftAddr)
}

// Copyright (c) 2026 Zeronetsec