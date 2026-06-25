// https://github.com/Zeronetsec/Comet

package console

import (
    "fmt"
    "os"
    "strconv"
    "time"
    "github.com/Zeronetsec/Comet/utils/color"
    "github.com/Zeronetsec/Comet/utils/invinput"
    "github.com/Zeronetsec/Comet/module/header"
)

type Header struct{}
func (c Header) Execute(args []string) {
    if len(args) < 3 {
        invinput.MissingArgument()
        os.Exit(1)
    }

    targetURL := args[2]
    timeoutSec := 5 * time.Second
    followRedirect := false

    for i := 3; i < len(args); i++ {
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

        if args[i] == "--redirect" && i+1 < len(args) {
            val := args[i+1]
            if val == "false" {
                followRedirect = false
            } else if val == "true" {
                followRedirect = true
            } else {
                fmt.Printf(
                    "%s[!] %sInvalid redirect value: %s%s%s\n",
                    color.R, color.N, color.GG, val, color.N,
                )
                os.Exit(1)
            }
            i++
        }
    }

    header.Inspect(targetURL, timeoutSec, followRedirect)
}

// Copyright (c) 2026 Zeronetsec