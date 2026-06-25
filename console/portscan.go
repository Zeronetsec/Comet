// https://github.com/Zeronetsec/Comet

package console

import (
    "fmt"
    "os"
    "github.com/Zeronetsec/Comet/utils/invinput"
    "github.com/Zeronetsec/Comet/utils/color"
    "github.com/Zeronetsec/Comet/module/portscan"
)

type Portscan struct{}
func (c Portscan) Execute(args []string) {
    if len(args) < 3 {
        invinput.MissingArgument()
        os.Exit(1)
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
            os.Exit(1)
        }

        if st > 0 {
            start = st
        }

        if en > 0 {
            end = en
        }
    }

    portscan.ScanPort(ip, start, end)
}

// Copyright (c) 2026 Zeronetsec