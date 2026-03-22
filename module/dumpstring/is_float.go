// Comet Framework

package dumpstring

import (
    "regexp"
)

func isFloat(s string) bool {
    match, _ := regexp.MatchString(`^-?\d+\.\d+$`, s)
    return match
}

// Copyright (c) 2026 Zeronetsec