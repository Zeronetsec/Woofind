// https://github.com/Zeronetsec/Woofind

package console

import (
    "os"
    "woofind/utils/invinput"
)

func WoofindConsole(input string) {
    args := os.Args
    if len(args) < 2 {
        invinput.Invalid()
        return
    }

    commands := map[string]Command{
        "--sysinfo": Sysinfo{},
        "--procinfo": Procinfo{},
        "--checkroot": Checkroot{},
        "--dumpstring": Dumpstring{},
        "--decode": Decode{},
        "--misconfind": Misconfind{},
        "--uwu": UWU{},
        "--version": Version{},
        "--help": Helper{},
    }

    if cmd, ok := commands[args[1]]; ok {
        cmd.Execute(args)
    } else {
        invinput.Unknown(args[1])
    }
}

// Copyright (c) 2026 Zeronetsec