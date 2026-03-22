// Comet Framework

package dumpstring

import (
    "strings"
    "fmt"
)

func toBinary(s string) string {
    var result []string
    for _, b := range []byte(s) {
        result = append(result, fmt.Sprintf("%08b", b))
    }
    return strings.Join(result, " ")
}

// Copyright (c) 2026 Zeronetsec