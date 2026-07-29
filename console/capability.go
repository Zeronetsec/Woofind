// https://github.com/Zeronetsec/Woofind

package console

import (
    "os"
    "strconv"
    "strings"
    "github.com/Zeronetsec/Woofind/module/capability"
    "github.com/Zeronetsec/Woofind/utils/invinput"
)

type Capability struct{}
func (c Capability) Execute(args []string) {
    if len(args) < 3 {
        invinput.MissingArgument()
        os.Exit(1)
    }

    targetPath := args[2]
    capList := ""
    threads := 100

    for i := 3; i < len(args); i++ {
        switch args[i] {
            case "--threads":
                if i+1 < len(args) {
                    t, err := strconv.Atoi(args[i+1])
                    if err == nil && t > 0 {
                        threads = t
                    }
                    i++
                }
            default:
                if (capList == "" &&
                    !strings.HasPrefix(args[i], "--")) {
                        capList = args[i]
                }
        }
    }

    capability.Getcap(targetPath, capList, threads)
}

// Copyright (c) 2026 Zeronetsec