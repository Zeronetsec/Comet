// Comet Framework

package procinfo

import (
    "strings"
    "strconv"
    "os/exec"
)

func getTTYMap() map[int32]string {
    ttyMap := make(map[int32]string)

    cmd := exec.Command("ps", "-A", "-o", "pid,tty")
    output, err := cmd.Output()
    if err != nil {
        return ttyMap
    }

    lines := strings.Split(string(output), "\n")

    for _, line := range lines[1:] {
        fields := strings.Fields(line)
        if len(fields) < 2 {
            continue
        }

        pid, err := strconv.Atoi(fields[0])
        if err != nil {
            continue
        }

        ttyMap[int32(pid)] = fields[1]
    }

    return ttyMap
}

// Copyright (c) 2026 Zeronetsec