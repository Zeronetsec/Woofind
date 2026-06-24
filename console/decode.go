// https://github.com/Zeronetsec/Woofind

package console

import (
    "os"
    "github.com/Zeronetsec/Woofind/utils/invinput"
    "github.com/Zeronetsec/Woofind/module/decode"
)

type Decode struct{}
func (c Decode) Execute(args []string) {
    if len(args) < 3 {
        invinput.MissingArgument()
        os.Exit(1)
    }

    decode.ExecBrute(args[2])
}

// Copyright (c) 2026 Zeronetsec