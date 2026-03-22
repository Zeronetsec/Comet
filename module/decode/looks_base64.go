// Comet Framework

package decode

import (
    "strings"
)

func looksBase64(s string) bool {
    for _, c := range s {
        if !strings.Contains("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/=", string(c)) {
            return false
        }
    }
    return true
}

// Copyright (c) 2026 Zeronetsec