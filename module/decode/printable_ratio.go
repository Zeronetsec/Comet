// Comet Framework

package decode

import (
    "unicode"
)

func printableRatio(data []byte) float64 {
    count := 0
    for _, b := range data {
        if unicode.IsPrint(rune(b)) {
            count++
        }
    }

    if len(data) == 0 {
        return 0
    }

    return float64(count) / float64(len(data))
}

// Copyright (c) 2026 Zeronetsec