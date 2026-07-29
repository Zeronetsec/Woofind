// https://github.com/Zeronetsec/Woofind

package console

import (
    "os"
    "github.com/Zeronetsec/Woofind/utils/invinput"
)

func WoofindConsole(input string) {
    args := os.Args
    if len(args) < 2 {
        invinput.MissingArgument()
        os.Exit(1)
    }

    commands := map[string]Command{
        "--sysinfo": Sysinfo{},
        "--procinfo": Procinfo{},
        "--checkroot": Checkroot{},
        "--dumpstring": Dumpstring{},
        "--decode": Decode{},
        "--uwu": UWU{},
        "--version": Version{},
        "--help": Helper{},
        "--pattern-scan": PatternScan{},
        "--permission": Permission{},
        "--owner": Owner{},
        "--capability": Capability{},
        "--varaddr": VarAddr{},
        "--dumpenvaddr": DumpEnvAddr{},
        "--useraudit": UserAudit{},
        "--list": List{},
    }

    if cmd, ok := commands[args[1]]; ok {
        cmd.Execute(args)
    } else {
        invinput.InvalidOption(args[1])
        os.Exit(1)
    }
}

// Copyright (c) 2026 Zeronetsec