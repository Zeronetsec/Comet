// https://github.com/Zeronetsec/Comet

package sqlscan

import (
    "strings"
)

func checkVulnerable(body string) (bool, string) {
    for _, sig := range sqlSignatures {
        if strings.Contains(body, sig) {
            return true, sig
        }
    }
    return false, ""
}

// Copyright (c) 2026 Zeronetsec