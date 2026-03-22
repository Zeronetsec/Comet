// Comet Framework

package dumpstring

import (
    "math"
)

func entropy(s string) float64 {
    if len(s) == 0 {
        return 0
    }

    freq := make(map[rune]int)
    for _, r := range s {
        freq[r]++
    }

    var e float64
    length := float64(len(s))

    for _, count := range freq {
        p := float64(count) / length
        e -= p * (math.Log2(p))
    }

    return e
}

// Copyright (c) 2026 Zeronetsec