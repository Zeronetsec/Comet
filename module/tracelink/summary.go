// https://github.com/Zeronetsec/Comet

package tracelink

import (
    "fmt"
    "sort"
    "github.com/Zeronetsec/Comet/utils/color"
    "github.com/Zeronetsec/Comet/utils/logger"
)

func summary(results map[string][]string) {
    total := 0
    for _, urls := range results {
        total += len(urls)
    }

    fmt.Println()
    if total == 0 {
        fmt.Printf(
            "%s[!] %sNo links found!\n",
            color.R, color.N,
        )
        return
    }

    log := logger.NewLogger("tracelink")
    for _, urls := range results {
        sort.Strings(urls)
        for _, u := range urls {
            fmt.Printf(
                "%s* %s%s%s\n",
                color.DG, color.GG, u, color.N,
            )

            logMess := fmt.Sprintf(
                "Found: %s",
                u,
            )
            log.Log(":", logMess)
        }
    }

    fmt.Println()
    fmt.Printf(
        "%s[*] %sTotal found: %s%d%s\n",
        color.B, color.N, color.GG, total, color.N,
    )
}

// Copyright (c) 2026 Zeronetsec