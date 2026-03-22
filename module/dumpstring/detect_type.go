// Comet Framework

package dumpstring

import (
    "strings"
)

func detectType(s string) string {
    if s == "" {
        return "Undefined"
    }

    if isInteger(s) {
        return "Integer"
    }

    if isFloat(s) {
        return "Float"
    }

    if strings.EqualFold(s, "true") || strings.EqualFold(s, "false") {
        return "Boolean-like"
    }

    return "String"
}

// Copyright (c) 2026 Zeronetsec