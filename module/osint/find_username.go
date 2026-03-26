// Comet Framework

package osint

import (
    "embed"
    "time"
    "bufio"
    "strings"
    "sync"
    "fmt"
    "sort"
    "net"
    "net/http"
    "comet/utils/color"
)

//go:embed sites/*
var SitesFS embed.FS

var (
    osintList = "sites/osint_sites.txt"
    maxThreads = 100
    timeoutSec = 1 * time.Second
)

func FindUsername(username string) {
    file, err := SitesFS.Open(osintList)
    if err != nil {
        fmt.Printf(
            "%s[!] %sOsint list not found!\n",
            color.R, color.N,
        )
        return
    }
    defer file.Close()

    var domains []string
    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())
        if line != "" && !strings.HasPrefix(line, "#") {
            domains = append(domains, line)
        }
    }

    total := len(domains)
    fmt.Printf(
        "%s[*] %sScanning username: %s%s %son %s%d %sdomains\n",
        color.B, color.N, color.GG, username, color.N, color.GG, total, color.N,
    )
    fmt.Println()

    client := &http.Client{
        Timeout: timeoutSec,
        Transport: &http.Transport{
            DialContext: (&net.Dialer{
                Timeout: 5 * time.Second,
                KeepAlive: 5 * time.Second,
            }).DialContext,
            TLSHandshakeTimeout: 5 * time.Second,
            ResponseHeaderTimeout: 6 * time.Second,
            ExpectContinueTimeout: 1 * time.Second,
        },
    }

    var wg sync.WaitGroup
    sem := make(chan struct{}, maxThreads)

    var mu sync.Mutex
    results := make(map[int][]string)

    found := 0
    for _, domain := range domains {
        wg.Add(1)
        sem <- struct{}{}

        go func(d string) {
            defer wg.Done()
            url, status, ok := checkDomain(client, username, d)
            if !ok {
                <-sem
                return
            }

            mu.Lock()
            found++
            results[status] = append(results[status], url)
            current := found
            mu.Unlock()

            col := color.CC
            if status == 301 {
                col = color.YY
            } else if status == 302 {
                col = color.PP
            } else if status == 403 {
                col = color.R
            }

            fmt.Printf(
                "%s[+] %sFound %d: %s%s %s(%s%d%s)%s\n",
                color.B, color.N, current, color.GG, url,
                color.DG, col, status, color.DG, color.N,
            )
            <-sem
        }(domain)
    }

    wg.Wait()
    fmt.Println()

    if found == 0 {
        fmt.Printf(
            "%s[!] %sNo username found!\n",
            color.R, color.N,
        )
        return
    }

    fmt.Printf(
        "%s[*] %sTotal %s%d %sfound:\n",
        color.B, color.N, color.GG, found, color.N,
    )

    statusOrder := []int{200, 301, 302}
    for _, status := range statusOrder {
        urls, ok := results[status]
        if !ok || len(urls) == 0 {
            continue
        }

        sort.Strings(urls)

        fmt.Printf(
            "    %s* %s%d:%s\n",
            color.DG, color.WW, status, color.N,
        )

        for i, u := range urls {
            uprefix := "├──"
            if i == len(urls)-1 {
                uprefix = "└──"
            }

            fmt.Printf(
                "        %s%s %s%s%s\n",
                color.DG, uprefix, color.GG, u, color.N,
            )
        }
    }
}

// Copyright (c) 2026 Zeronetsec