// Comet Framework

package sysinfo

import (
    "fmt"
    "strings"
    "os/exec"
    "comet/utils/color"
)

func execShell(cmd string, args ...string) string {
    command := exec.Command(cmd, args...)
    output, err := command.Output()
    if err != nil {
        return fmt.Sprintf(
            "%s[!] %sError: %s%s%s\n",
            color.R, color.N, color.GG, err, color.N,
        )
    }

    return strings.TrimSpace(string(output))
}

// Copyright (c) 2026 Zeronetsec