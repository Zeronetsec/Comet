// Comet Framework

package decode

import (
    "strings"
)

func looksBase32(s string) bool {
    for _, c := range s {
        if !strings.Contains("ABCDEFGHIJKLMNOPQRSTUVWXYZ234567=", string(c)) {
            return false
        }
    }
    return true
}

// Copyright (c) 2026 Zeronetsec