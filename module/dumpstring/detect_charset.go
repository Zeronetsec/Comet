// Comet Framework

package dumpstring

import (
    "unicode"
    "sort"
    "strings"
)

func detectCharset(s string) string {
    set := map[string]bool{}

    for _, r := range s {
        switch {
            case unicode.IsLower(r):
                set["a-z"] = true
            case unicode.IsUpper(r):
                set["A-Z"] = true
            case unicode.IsDigit(r):
                set["0-9"] = true
            default:
                set["symbol"] = true
        }
    }

    var keys []string
    for k := range set {
        keys = append(keys, k)
    }

    sort.Strings(keys)
    return strings.Join(keys, ", ")
}

// Copyright (c) 2026 Zeronetsec