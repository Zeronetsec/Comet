// https://github.com/Zeronetsec/Comet

package console

import (
    "fmt"
    "os"
    "strconv"
    "github.com/Zeronetsec/Comet/utils/color"
    "github.com/Zeronetsec/Comet/utils/invinput"
    "github.com/Zeronetsec/Comet/module/paramscan"
)

type Paramscan struct{}
func (c Paramscan) Execute(args []string) {
    if len(args) < 3 {
        invinput.MissingArgument()
        os.Exit(1)
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
        fmt.Printf(
            "%s[!] %sInvalid threads value!\n",
            color.R, color.N,
        )
        os.Exit(1)
    }

    if timeout <= 0 {
        fmt.Printf(
            "%s[!] %sInvalid timeout value!\n",
            color.R, color.N,
        )
        os.Exit(1)
    }

    paramscan.FetchParameters(
        target,
        threads,
        timeout,
        fuzz,
    )
}

// Copyright (c) 2026 Zeronetsec