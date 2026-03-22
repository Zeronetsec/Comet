// Comet Framework

package dumpstring

import (
    "unicode"
)

func isPrintable(s string) string {
    for _, r := range s {
        if !unicode.IsPrint(r) {
            return "No"
        }
    }

    return "Yes"
}

// Copyright (c) 2026 Zeronetsec