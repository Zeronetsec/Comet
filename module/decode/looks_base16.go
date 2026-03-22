// Comet Framework

package decode

import (
    "strings"
)

func looksBase16(s string) bool {
    for _, c := range s {
        if !strings.Contains("0123456789abcdefABCDEF", string(c)) {
            return false
        }
    }
    return true
}

// Copyright (c) 2026 Zeronetsec