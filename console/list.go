// https://github.com/Zeronetsec/Woofind

package console

import (
    "os"
    "github.com/Zeronetsec/Woofind/module/list"
    "github.com/Zeronetsec/Woofind/utils/invinput"
)

type List struct{}
func (c List) Execute(args []string) {
    if len(args) < 3 {
        invinput.MissingArgument()
        os.Exit(1)
    }

    list.Show(os.Args[2])
}

// Copyright (c) 2026 Zeronetsec