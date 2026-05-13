// https://github.com/Zeronetsec/Woofind

package console

import (
    "time"
    "woofind/utils/invinput"
    "woofind/utils/uwu"
    "woofind/utils/version"
    "woofind/utils/helper"
    "woofind/utils/cursor"
    "woofind/module/sysinfo"
    "woofind/module/procinfo"
    "woofind/module/checkroot"
    "woofind/module/dumpstring"
    "woofind/module/decode"
    "woofind/module/misconfind"
)

type Command interface {
    Execute(args []string)
}

type Sysinfo struct{}
func (c Sysinfo) Execute(args []string) {
    cursor.Hide()
    sysinfo.MachineInfo()
    cursor.Visible()
}

type Procinfo struct{}
func (c Procinfo) Execute(args []string) {
    cursor.Hide()
    procinfo.ShowProcInfo()
    cursor.Visible()
}

type Checkroot struct{}
func (c Checkroot) Execute(args []string) {
    cursor.Hide()
    checkroot.GetCheck()
    cursor.Visible()
}

type Dumpstring struct{}
func (c Dumpstring) Execute(args []string) {
    if len(args) < 3 {
        invinput.MissingArgs()
        return
    }

    cursor.Hide()
    dumpstring.Analyzer(args[2])
    cursor.Visible()
}

type Decode struct{}
func (c Decode) Execute(args []string) {
    if len(args) < 3 {
        invinput.MissingArgs()
        return
    }

    cursor.Hide()
    decode.ExecBrute(args[2])
    cursor.Visible()
}

type Misconfind struct{}
func (c Misconfind) Execute(args []string) {
    if len(args) < 3 {
        invinput.MissingArgs()
        return
    }

    cursor.Hide()
    misconfind.Finder(args[2:])
    cursor.Visible()
}

type UWU struct{}
func (c UWU) Execute(args []string) {
    cursor.Hide()
    uwu.Nyanners(5 * time.Second)
    cursor.Visible()
}

type Version struct{}
func (c Version) Execute(args []string) {
    cursor.Hide()
    version.WoofindVersion()
    cursor.Visible()
}

type Helper struct{}
func (c Helper) Execute(args []string) {
    cursor.Hide()
    helper.WoofindHelper()
    cursor.Visible()
}

// Copyright (c) 2026 Zeronetsec