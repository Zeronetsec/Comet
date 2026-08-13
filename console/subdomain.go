// https://github.com/Zeronetsec/Comet

package console

import (
    "os"
    "strconv"
    "github.com/Zeronetsec/Comet/module/subdomain"
    "github.com/Zeronetsec/Comet/utils/invinput"
)

type Subdomain struct{}
func (c Subdomain) Execute(args []string) {
    if len(args) < 3 {
        invinput.MissingArgument()
        os.Exit(1)
    }

    timeout := 60
    retries := 5
    targetDomain := ""

    for i := 2; i < len(args); i++ {
        switch args[i] {
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
                        retries = r
                    }
                    i++
                }
            default:
                if (targetDomain == "" &&
                    args[i][0] != '-') {
                    targetDomain = args[i]
                }
        }
    }

    if targetDomain == "" {
        invinput.MissingArgument()
        os.Exit(1)
    }

    subdomain.Find(
        targetDomain,
        timeout,
        retries,
    )
}

// Copyright (c) 2026 Zeronetsec