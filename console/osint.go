// https://github.com/Zeronetsec/Comet

package console

import (
    "fmt"
    "os"
    "strconv"
    "time"
    "github.com/Zeronetsec/Comet/utils/color"
    "github.com/Zeronetsec/Comet/utils/invinput"
    "github.com/Zeronetsec/Comet/module/osint"
)

type Osint struct{}
func (c Osint) Execute(args []string) {
    if len(args) < 3 {
        invinput.MissingArgument()
        os.Exit(1)
    }

    username := args[2]
    maxThreads := 100
    timeoutSec := 1 * time.Second

    for i := 3; i < len(args); i++ {
        if args[i] == "--threads" && i+1 < len(args) {
            t, err := strconv.Atoi(args[i+1])
            if err == nil && t > 0 {
                maxThreads = t
            } else {
                fmt.Printf(
                    "%s[!] %sInvalid thread value: %s%s%s\n",
                    color.R, color.N, color.GG, args[i+1], color.N,
                )
                os.Exit(1)
            }
            i++
        }

        if args[i] == "--timeout" && i+1 < len(args) {
            t, err := strconv.Atoi(args[i+1])
            if err == nil && t > 0 {
                timeoutSec = time.Duration(t) * time.Second
            } else {
                fmt.Printf(
                    "%s[!] %sInvalid timeout value: %s%s%s\n",
                    color.R, color.N, color.GG, args[i+1], color.N,
                )
                os.Exit(1)
            }
            i++
        }
    }

    osint.FindUsername(username, maxThreads, timeoutSec)
}

// Copyright (c) 2026 Zeronetsec