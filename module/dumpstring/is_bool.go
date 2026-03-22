// Comet Framework

package dumpstring

import (
    "regexp"
)

func isBool(s string) bool {
    match, _ := regexp.MatchString(`^(true|false|0|1)$`, s)
    return match
}

// Copyright (c) 2026 Zeronetsec