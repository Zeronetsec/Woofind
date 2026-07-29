// https://github.com/Zeronetsec/Woofind

package console

import (
    "os"
    "github.com/Zeronetsec/Woofind/module/useraudit"
)

type UserAudit struct{}
func (c UserAudit) Execute(args []string) {
    prefix := os.Getenv("PREFIX")
    if prefix == "" {
        prefix = "/usr"
    }

    passwdFile := prefix + "/etc/passwd"
    shadowFile := prefix + "/etc/shadow"
    shellStr := ""

    for i := 2; i < len(args); i++ {
        switch args[i] {
            case "--passwd":
                if i+1 < len(args) {
                    passwdFile = args[i+1]
                    i++
                }
            case "--shadow":
                if i+1 < len(args) {
                    shadowFile = args[i+1]
                    i++
                }
            case "--shell":
                if i+1 < len(args) {
                    shellStr = args[i+1]
                    i++
                }
        }
    }

    useraudit.UserScan(
        passwdFile, shadowFile, shellStr,
    )
}

// Copyright (c) 2026 Zeronetsec