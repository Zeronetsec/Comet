// https://github.com/Zeronetsec/Comet

package console

import (
    "fmt"
    "os"
    "strconv"
    "github.com/Zeronetsec/Comet/module/paramscan"
    "github.com/Zeronetsec/Comet/utils/color"
    "github.com/Zeronetsec/Comet/utils/invinput"
)

type Paramscan struct{}
func (c Paramscan) Execute(args []string) {
    if len(args) < 3 {
        invinput.MissingArgument()
        os.Exit(1)
    }

    target := args[2]
    timeout := 15
    threads := 100
    retry := 5
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
            case "--retry":
                if i+1 < len(args) {
                    r, err := strconv.Atoi(args[i+1])
                    if err == nil {
                        retry = r
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

    if retry <= 0 {
        fmt.Printf(
            "%s[!] %sInvalid retry value!\n",
            color.R, color.N,
        )
        os.Exit(1)
    }

    paramscan.FetchParameters(
        target,
        threads,
        timeout,
        retry,
        fuzz,
    )
}

// Copyright (c) 2026 Zeronetsec