// https://github.com/Zeronetsec/Woofind

package console

import (
    "os"
    "strconv"
    "strings"
    "github.com/Zeronetsec/Woofind/module/owner"
    "github.com/Zeronetsec/Woofind/utils/invinput"
)

type Owner struct{}
func (c Owner) Execute(args []string) {
    if len(args) < 3 {
        invinput.MissingArgument()
        os.Exit(1)
    }

    targetPath := args[2]
    ownerList := ""
    threads := 100
    force := false

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
            case "--force":
                force = true
            default:
                if (ownerList == "" &&
                    !strings.HasPrefix(args[i], "--")) {
                        ownerList = args[i]
                }
        }
    }

    owner.OwnScan(targetPath, ownerList, threads, force)
}

// Copyright (c) 2026 Zeronetsec