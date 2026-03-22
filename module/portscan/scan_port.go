// Comet Framework

package portscan

import (
    "fmt"
    "sync"
    "comet/utils/color"
)

func ScanPort(ip string, start, end int) {
    if !checkHost(ip) {
        fmt.Printf(
            "%s[!] %sHost: %s%s %sis down\n",
            color.R, color.N, color.GG, ip, color.N,
        )
        return
    }

    fmt.Printf(
        "%s[*] %sScanning port on %s%s %s(%s%d-%d%s)%s\n",
        color.B, color.N, color.GG, ip, color.DG, color.CC, start, end, color.DG, color.N,
    )

    var wg sync.WaitGroup
    sem := make(chan struct{}, 200)

    var openPorts int
    var mu sync.Mutex

    fmt.Println()
    for port := start; port <= end; port++ {
        wg.Add(1)
        sem <- struct{}{}

        go func(p int) {
            defer wg.Done()

            if scanRun(ip, p) {
                mu.Lock()
                openPorts++
                mu.Unlock()
            }

            <-sem
        }(port)
    }

    wg.Wait()

    if openPorts == 0 {
        fmt.Printf(
            "%s[*] %sNo port running on %s%s%s\n",
            color.B, color.N, color.GG, ip, color.N,
        )
    } else {
        fmt.Printf(
            "\n%s[*] %sScanning complete %s(%s%d %sport open%s)%s\n",
            color.B, color.N, color.DG, color.GG, openPorts, color.WW, color.DG, color.N,
        )
    }
}

// Copyright (c) 2026 Zeronetsec