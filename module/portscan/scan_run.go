// Comet Framework

package portscan

import (
    "fmt"
    "net"
    "time"
    "comet/utils/color"
    "comet/utils/logger"
)

func scanRun(ip string, port int) bool {
    log := logger.NewLogger("portscan")
    address := fmt.Sprintf("%s:%d", ip, port)
    conn, err := net.DialTimeout(
        "tcp", address, 200*time.Millisecond,
    )

    if err != nil {
        return false
    }

    defer conn.Close()

    fmt.Printf(
        "%s[+] %sPort: %s%d %sopen\n",
        color.GG, color.N, color.GG, port, color.N,
    )

    logMess := fmt.Sprintf(
        "Port: %s:%d open", ip, port,
    )

    log.Log(":", logMess)
    return true
}

// Copyright (c) 2026 Zeronetsec