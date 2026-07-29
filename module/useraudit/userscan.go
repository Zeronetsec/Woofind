// https://github.com/Zeronetsec/Woofind

package useraudit

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "github.com/Zeronetsec/Woofind/utils/color"
    "github.com/Zeronetsec/Woofind/utils/invinput"
    "github.com/Zeronetsec/Woofind/utils/logger"
)

var idx int

func UserScan(passwdFile, shadowFile, shellStr string) {
    invinput.NotFound(passwdFile)
    invinput.NotFound(shadowFile)

    idx = 0
    fmt.Printf(
        "%s[*] %sScanning: %spasswd%s:%sshadow%s\n",
        color.B, color.N, color.GG, color.DG,
        color.GG, color.N,
    )

    if shellStr == "" {
        shellStr = "bash:zsh:fish:sh:dash:mksh:nushell"
    }
    targetShells := strings.Split(shellStr, ":")

    fmt.Printf(
        "%s[*] %sPasswd: %s%s%s\n",
        color.B, color.N, color.GG, passwdFile, color.N,
    )

    fmt.Printf(
        "%s[*] %sShadow: %s%s%s\n",
        color.B, color.N, color.GG, shadowFile, color.N,
    )

    fmt.Printf(
        "%s[*] %sShell: %s%s%s\n",
        color.B, color.N, color.GG, shellStr, color.N,
    )

    fmt.Println()
    log := logger.NewLogger("useraudit")

    fileP, err := os.Open(passwdFile)
    if err != nil {
        fmt.Printf(
            "%s[!] %sFailed to read passwd: %s%v%s\n",
            color.R, color.N, color.GG, err, color.N,
        )
    } else {
        defer fileP.Close()
        scanner := bufio.NewScanner(fileP)
        lineNum := 1
        for scanner.Scan() {
            line := strings.TrimSpace(scanner.Text())
            if (line == "" ||
                strings.HasPrefix(line, "#")) {
                    lineNum++
                    continue
            }

            parts := strings.Split(line, ":")
            if len(parts) >= 7 {
                username := parts[0]
                uid := parts[2]
                shell := parts[6]

                if uid == "0" {
                    fmt.Printf(
                        "%s[+] %sFound: %s%s %s-> %suser%s:%s%s %s(%s0%s)%s\n",
                        color.GG, color.N, color.GG, passwdFile, color.DG,
                        color.BB, color.DG,
                        color.GG, username, color.DG,
                        color.CC, color.DG,
                        color.N,
                    )

                    log.Log(":", fmt.Sprintf(
                        "Found: %s -> user:%s (0)",
                        passwdFile, username,
                    ))
                    idx++
                }

                for _, ts := range targetShells {
                    ts = strings.TrimSpace(ts)
                    if ts != "" && (strings.HasSuffix(
                        shell, "/"+ts,
                    ) || shell == ts) {
                        fmt.Printf(
                            "%s[+] %sFound: %s%s %s-> %sshell%s:%s%s %s(%s%s%s)%s\n",
                            color.GG, color.N, color.GG, passwdFile, color.DG,
                            color.BB, color.DG,
                            color.GG, username, color.DG,
                            color.CC, shell, color.DG,
                            color.N,
                        )

                        log.Log(":", fmt.Sprintf(
                            "Found: %s -> shell:%s (%s)",
                            passwdFile, username, shell,
                        ))
                        idx++
                        break
                    }
                }
            }
            lineNum++
        }
    }

    fileS, err := os.Open(shadowFile)
    if err != nil {
        if os.IsPermission(err) {
            fmt.Printf(
                "%s[!] %sFailed to read shadow: %spermission denied%s\n",
                color.R, color.N, color.GG, color.N,
            )
        } else {
            fmt.Printf(
                "%s[!] %sFailed to read shadow: %s%v%s\n",
                color.R, color.N, color.GG, err, color.N,
            )
        }
    } else {
        defer fileS.Close()
        scanner := bufio.NewScanner(fileS)
        lineNum := 1
        for scanner.Scan() {
            line := strings.TrimSpace(scanner.Text())
            if line == "" || strings.HasPrefix(
                line, "#",
            ) {
                lineNum++
                continue
            }

            parts := strings.Split(line, ":")
            if len(parts) >= 2 {
                username := parts[0]
                pwdHash := parts[1]
                if pwdHash == "" {
                    fmt.Printf(
                        "%s[+] %sFound: %s%s %s-> %suser%s:%s%s %s(%snopassword%s)%s\n",
                        color.GG, color.N, color.GG, shadowFile, color.DG,
                        color.BB, color.DG,
                        color.GG, username, color.DG,
                        color.CC, color.DG,
                        color.N,
                    )

                    log.Log(":", fmt.Sprintf(
                        "Found: %s -> user:%s (nopassword)",
                        shadowFile, username,
                    ))
                    idx++
                }
            }
            lineNum++
        }
    }

    if idx > 0 {
        fmt.Println()
    } else {
        fmt.Printf(
            "%s[!] %sNothing found!\n",
            color.R, color.N,
        )
        fmt.Println()
    }

    fmt.Printf(
        "%s[*] %sScanning done.\n",
        color.B, color.N,
    )
}

// Copyright (c) 2026 Zeronetsec