// Comet Framework

package paramscan

import (
    "fmt"
    "sort"
    "comet/utils/color"
)

func displaySummary(results map[int][]string) {
    totalFound := 0
    for _, urls := range results {
        totalFound += len(urls)
    }

    if totalFound == 0 {
        fmt.Printf(
            "%s[!] %sNo active parameters found!\n",
            color.R, color.N,
        )
        return
    }

    fmt.Printf(
        "%s[*] %sTotal %s%d %sactive parameters found:\n",
        color.B, color.N, color.GG, totalFound, color.N,
    )

    statuses := []int{200, 301, 302}
    for _, status := range statuses {
        urls, ok := results[status]
        if !ok || len(urls) == 0 {
            continue
        }

        sort.Strings(urls)
        fmt.Printf(
            "    %s* %s%d:%s\n",
            color.DG, color.WW, status, color.N,
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