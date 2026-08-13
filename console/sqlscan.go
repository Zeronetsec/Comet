// https://github.com/Zeronetsec/Comet

package console

import (
    "os"
    "strconv"
    "github.com/Zeronetsec/Comet/module/sqlscan"
    "github.com/Zeronetsec/Comet/utils/invinput"
)

type Sqlscan struct{}
func (c Sqlscan) Execute(args []string) {
    if len(args) < 3 {
        invinput.MissingArgument()
        os.Exit(1)
    }

    targetURL := ""
    wordlist := "embedded"
    threads := 100
    timeout := 15

    for i := 2; i < len(args); i++ {
        switch args[i] {
            case "--payload":
                if i+1 < len(args) {
                    wordlist = args[i+1]
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
                if (targetURL == "" &&
                    args[i][0] != '-') {
                        targetURL = args[i]
                }
        }
    }

    if targetURL == "" {
        invinput.MissingArgument()
        os.Exit(1)
    }

    sqlscan.Scan(
        targetURL,
        wordlist,
        WordlistFS,
        threads,
        timeout,
    )
}

// Copyright (c) 2026 Zeronetsec