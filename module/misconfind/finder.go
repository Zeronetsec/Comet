// Comet Framework

package misconfind

import (
    "fmt"
    "embed"
    "comet/utils/color"
)

//go:embed patterns/*
var PatternsFS embed.FS

func Finder(args []string) {
    if len(args) == 0 {
        fmt.Printf(
            "%s[!] %sPath is required!\n",
            color.R, color.N,
        )
        return
    }

    rootPath := args[0]
    patterns, err := loadPatterns(
        PatternsFS, "patterns/malicious_patterns.txt",
    )

    if err != nil {
        fmt.Printf(
            "%s[!] %sFailed to load patterns: %s%v%s\n",
            color.R, color.N, color.GG, err, color.N,
        )
        return
    }

    if len(patterns) == 0 {
        fmt.Printf(
            "%s[!] %sNo patterns loaded!\n",
            color.R, color.N,
        )
        return
    }

    scanPath(rootPath, patterns)
}

// Copyright (c) 2026 Zeronetsec