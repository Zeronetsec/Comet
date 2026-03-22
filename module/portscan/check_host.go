// Comet Framework

package portscan

import (
    "time"
    "net"
    "fmt"
    "os/exec"
)

func checkHost(ip string) bool {
    cmd := exec.Command("ping", "-c", "1", "-W", "1", ip)
    if err := cmd.Run(); err == nil {
        return true
    }

    ports := []int{80, 443, 22, 21}
    for _, port := range ports {
        address := fmt.Sprintf("%s:%d", ip, port)
        conn, err := net.DialTimeout(
            "tcp", address, 300*time.Millisecond,
        )

        if err == nil {
            conn.Close()
            return true
        }
    }

    return false
}

// Copyright (c) 2026 Zeronetsec