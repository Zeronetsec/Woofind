// https://github.com/Zeronetsec/Woofind

package console

type Command interface {
    Execute(args []string)
}

// Copyright (c) 2026 Zeronetsec