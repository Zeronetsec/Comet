// Comet Framework

package tracelink

import (
    "fmt"
    "sort"
    "comet/utils/color"
)

func printSummary(results map[string][]string) {
    total := 0

    for _, urls := range results {
        total += len(urls)
    }

    if total == 0 {
        fmt.Printf(
            "%s[!] %sNo links found!\n",
            color.R, color.N,
        )
        return
    }

    fmt.Printf(
        "%s[*] %sTotal %s%d %sfound:\n",
        color.B, color.N, color.GG, total, color.N,
    )

    for _, urls := range results {
        sort.Strings(urls)
        fmt.Printf(
            "    %s* %s%d:%s\n",
            color.DG, color.WW, len(urls), color.N,
        )

        for i, u := range urls {
            prefix := "├──"
            if i == len(urls)-1 {
                prefix = "└──"
            }

            fmt.Printf(
                "        %s%s %s%s%s\n",
                color.DG, prefix, color.GG, u, color.N,
            )
        }
    }
}

// Copyright (c) 2026 Zeronetsec