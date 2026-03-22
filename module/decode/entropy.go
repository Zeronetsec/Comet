// Comet Framework

package decode

import (
    "math"
)

func entropy(data []byte) float64 {
    if len(data) == 0 {
        return 0
    }

    freq := make(map[byte]int)
    for _, b := range data {
        freq[b]++
    }

    var e float64
    l := float64(len(data))

    for _, c := range freq {
        p := float64(c) / l
        e -= p * math.Log2(p)
    }

    return e
}

// Copyright (c) 2026 Zeronetsec