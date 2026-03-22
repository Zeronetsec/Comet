// Comet Framework

package portscan

import (
    "fmt"
    "strings"
    "strconv"
    "comet/utils/color"
)

func ParseRange(input string) (int, int, error) {
    parts := strings.Split(input, ":")

    if len(parts) != 2 {
        fmt.Printf(
            "%s[!] %sInvalid format %s(%suse: %sstart%s:%send%s)%s\n",
            color.R, color.N, color.DG, color.N, color.GG, color.DG, color.GG, color.DG, color.N,
        )

        return 0, 0, fmt.Errorf("Invalid format")
    }

    start, err1 := strconv.Atoi(parts[0])
    end, err2 := strconv.Atoi(parts[1])

    if err1 != nil || err2 != nil {
        fmt.Printf(
            "%s[!] %sInvalid number!\n",
            color.R, color.N,
        )

        return 0, 0, fmt.Errorf("Invalid number")
    }

    if start < 1 || end > 65535 || start > end {
        fmt.Printf(
            "%s[!] %sInvalid range %s(%s1-65535%s)%s\n",
            color.R, color.N, color.DG, color.GG, color.DG, color.N,
        )

        return 0, 0, fmt.Errorf("Invalid range")
    }

    return start, end, nil
}

// Copyright (c) 2026 Zeronetsec