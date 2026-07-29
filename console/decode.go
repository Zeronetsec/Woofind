// https://github.com/Zeronetsec/Woofind

package console

import (
    "os"
    "github.com/Zeronetsec/Woofind/module/decode"
    "github.com/Zeronetsec/Woofind/utils/invinput"
)

type Decode struct{}
func (c Decode) Execute(args []string) {
    if len(args) < 3 {
        invinput.MissingArgument()
        os.Exit(1)
    }

    input := args[2]
    limit := "unlimit"
    disable := ""

    for i := 3; i < len(args); i++ {
        if args[i] == "--limit" && i+1 < len(args) {
            limit = args[i+1]
            i++
        } else if args[i] == "--disable" && i+1 < len(args) {
            disable = args[i+1]
            i++
        }
    }

    decode.ExecBrute(input, limit, disable)
}

// Copyright (c) 2026 Zeronetsec