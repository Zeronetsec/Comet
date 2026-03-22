// Comet Framework

package console

import (
    "embed"
    "fmt"
    "time"
    "strconv"
    "comet/utils/invinput"
    "comet/utils/color"
    "comet/utils/listcmd"
    "comet/utils/searchcmd"
    "comet/utils/uwu"
    "comet/utils/version"
    "comet/utils/helper"
    "comet/utils/cursor"
    "comet/module/sysinfo"
    "comet/module/portscan"
    "comet/module/procinfo"
    "comet/module/checkroot"
    "comet/module/dumpstring"
    "comet/module/decode"
    "comet/module/misconfind"
    "comet/module/dirfuzzer"
    "comet/module/osint"
    "comet/module/paramscan"
    "comet/module/tracelink"
)

//go:embed wordlist/*
var WordlistFS embed.FS

type Command interface {
    Execute(args []string)
}

type Sysinfo struct{}
func (c Sysinfo) Execute(args []string) {
    cursor.Hide()
    sysinfo.MachineInfo()
    cursor.Visible()
}

type Portscan struct{}
func (c Portscan) Execute(args []string) {
    if len(args) < 3 {
        invinput.MissingArgs()
        return
    }

    ip := args[2]
    start := 1
    end := 65535

    if len(args) >= 4 {
        st, en, err := portscan.ParseRange(args[3])
        if err != nil {
            fmt.Printf(
                "%s[!] %sError parsing range: %s%v%s\n",
                color.R, color.N, color.GG, err, color.N,
            )
            return
        }

        if st > 0 {
            start = st
        }

        if en > 0 {
            end = en
        }
    }

    cursor.Hide()
    portscan.ScanPort(ip, start, end)
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

type Dirfuzzer struct{}
func (c Dirfuzzer) Execute(args []string) {
    if len(args) < 3 {
        invinput.MissingArgs()
        return
    }

    target := args[2]
    wordlist := "wordlist/common.txt"
    timeout := 10
    threads := 100
    recursive := false

    for i := 3; i < len(args); i++ {
        switch args[i] {
            case "--recursive":
                recursive = true
            case "--threads":
                if i+1 < len(args) {
                    t, err := strconv.Atoi(args[i+1])
                    if err == nil {
                        threads = t
                    }
                    i++
                }
            case "--timeout":
                if i+1 < len(args) {
                    t, err := strconv.Atoi(args[i+1])
                    if err == nil {
                        timeout = t
                    }
                    i++
                }
            case "--wordlist":
                if i+1 < len(args) {
                    wordlist = args[i+1]
                    i++
                }
        }
    }

    if threads <= 0 {
        threads = 100
    }

    if timeout <= 0 {
        timeout = 10
    }

    cursor.Hide()
    dirfuzzer.ExecFuzzing(target, WordlistFS, wordlist, timeout, recursive, threads)
    cursor.Visible()
}

type Osint struct{}
func (c Osint) Execute(args []string) {
    if len(args) < 3 {
        invinput.MissingArgs()
        return
    }

    cursor.Hide()
    osint.FindUsername(args[2])
    cursor.Visible()
}

type Paramscan struct{}
func (c Paramscan) Execute(args []string) {
    if len(args) < 3 {
        invinput.MissingArgs()
        return
    }

    target := args[2]
    timeout := 10
    threads := 100
    fuzz := false

    for i := 3; i < len(args); i++ {
        switch args[i] {
            case "--fuzz":
                fuzz = true
            case "--threads":
                if i+1 < len(args) {
                    t, err := strconv.Atoi(args[i+1])
                    if err == nil {
                        threads = t
                    }
                    i++
                }
            case "--timeout":
                if i+1 < len(args) {
                    t, err := strconv.Atoi(args[i+1])
                    if err == nil {
                        timeout = t
                    }
                    i++
                }
        }
    }

    if threads <= 0 {
        threads = 100
    }

    if timeout <= 0 {
        timeout = 10
    }

    cursor.Hide()
    paramscan.FetchParameters(target, threads, timeout, fuzz)
    cursor.Visible()
}

type Tracelink struct{}
func (c Tracelink) Execute(args []string) {
    if len(args) < 3 {
        invinput.MissingArgs()
        return
    }

    target := args[2]
    threads := 100
    recursive := false

    for i := 3; i < len(args); i++ {
        switch args[i] {
            case "--threads":
                if i+1 < len(args) {
                    t, err := strconv.Atoi(args[i+1])
                    if err == nil {
                        threads = t
                    }
                    i++
                }
            case "--recursive":
                recursive = true
        }
    }

    if threads <= 0 {
        threads = 100
    }

    cursor.Hide()
    tracelink.Tracer(target, threads, recursive)
    cursor.Visible()
}

type Listcmd struct{}
func (c Listcmd) Execute(args []string) {
    cursor.Hide()
    listcmd.Lister()
    cursor.Visible()
}

type Searchcmd struct{}
func (c Searchcmd) Execute(args []string) {
    if len(args) < 3 {
        invinput.MissingArgs()
        return
    }

    cursor.Hide()
    searchcmd.SearchModule(args[2])
    cursor.Visible()
}

type Uwu struct{}
func (c Uwu) Execute(args []string) {
    cursor.Hide()
    uwu.Nyan(5 * time.Second)
    cursor.Visible()
}

type Version struct{}
func (c Version) Execute(args []string) {
    cursor.Hide()
    version.CometVersion()
    cursor.Visible()
}

type Helper struct{}
func (c Helper) Execute(args []string) {
    cursor.Hide()
    helper.CometHelper()
    cursor.Visible()
}

// Copyright (c) 2026 Zeronetsec