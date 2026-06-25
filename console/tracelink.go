// https://github.com/Zeronetsec/Comet

package console

import (
    "fmt"
    "os"
    "strconv"
    "github.com/Zeronetsec/Comet/utils/invinput"
    "github.com/Zeronetsec/Comet/utils/color"
    "github.com/Zeronetsec/Comet/module/tracelink"
)

type Tracelink struct{}
func (c Tracelink) Execute(args []string) {
    if len(args) < 3 {
        invinput.MissingArgument()
        os.Exit(1)
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
        fmt.Printf(
            "%s[!] %sInvalid threads value!\n",
            color.R, color.N,
        )
        os.Exit(1)
    }

    tracelink.Tracer(target, threads, recursive)
}

// Copyright (c) 2026 Zeronetsec