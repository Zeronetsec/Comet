// https://github.com/Zeronetsec/Comet

package console

import (
    "os"
    "strconv"
    "github.com/Zeronetsec/Comet/module/corsscan"
    "github.com/Zeronetsec/Comet/utils/invinput"
)

type CorsScan struct{}
func (c CorsScan) Execute(args []string) {
    if len(args) < 3 {
        invinput.MissingArgument()
        os.Exit(1)
    }

    target := ""
    origin := "https://evil.com"
    customHeader := ""
    threads := 100
    timeout := 10

    for i := 2; i < len(args); i++ {
        switch args[i] {
            case "--origin":
                if i+1 < len(args) {
                    origin = args[i+1]
                    i++
                }
            case "--header":
                if i+1 < len(args) {
                    customHeader = args[i+1]
                    i++
                }
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
            default:
                if (target == "" &&
                    args[i][0] != '-') {
                        target = args[i]
                }
        }
    }

    if target == "" {
        invinput.MissingArgument()
        os.Exit(1)
    }

    corsscan.Scan(
        target,
        origin,
        customHeader,
        threads,
        timeout,
    )
}

// Copyright (c) 2026 Zeronetsec