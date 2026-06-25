// https://github.com/Zeronetsec/Comet

package paramscan

import (
    "fmt"
    "sort"
    "github.com/Zeronetsec/Comet/utils/color"
    "github.com/Zeronetsec/Comet/utils/logger"
)

func summary(results map[int][]string) {
    totalFound := 0
    for _, urls := range results {
        totalFound += len(urls)
    }

    fmt.Println()
    if totalFound == 0 {
        fmt.Printf(
            "%s[!] %sNo active parameters found!\n",
            color.R, color.N,
        )
        return
    }

    log := logger.NewLogger("paramscan")
    statuses := []int{200, 301, 302}
    for _, status := range statuses {
        urls, ok := results[status]
        if !ok || len(urls) == 0 {
            continue
        }

        sort.Strings(urls)
        for _, u := range urls {
            fmt.Printf(
                "%s* %s%s%s:%s%d%s\n",
                color.DG, color.GG, u, color.DG,
                color.CC, status, color.N,
            )

            logMess := fmt.Sprintf(
                "Found: %s:%d", u, status,
            )
            log.Log(":", logMess)
        }
    }

    fmt.Println()
    fmt.Printf(
        "%s[*] %sTotal active parameters found: %s%d%s:\n",
        color.B, color.N, color.GG, totalFound, color.N,
    )
}

// Copyright (c) 2026 Zeronetsec