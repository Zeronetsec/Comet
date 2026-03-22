// Comet Framework

package decode

import (
    "fmt"
    "strings"
    "time"
    "os"
    "comet/utils/color"
)

var visited = make(map[string]bool)
var steps int

func ExecBrute(input string) {
    var data string

    if info, err := os.Stat(input); err == nil && !info.IsDir() {
        content, err := os.ReadFile(input)
        if err != nil {
            fmt.Printf(
                "%s[!] %sError reading file: %s%v%s\n",
                color.R, color.GG, err, color.N,
            )
            return
        }
        data = string(content)
    } else {
        data = input
        if strings.TrimSpace(data) == "" {
            fmt.Printf(
                "%s[!] %sInvalid input!\n",
                color.R, color.N,
            )
            return
        }
    }

    fmt.Printf(
        "%s[*] %sTrying decode\n",
        color.B, color.N,
    )

    start := time.Now()
    brute([]byte(data), 0)

    fmt.Printf(
        "%sFinished in %s%dms %s| Steps: %s%d%s\n",
        color.N, color.GG, time.Since(start).Milliseconds(), color.N, color.GG, steps, color.N,
    )
}

// Copyright (c) 2026 Zeronetsec