// https://github.com/Zeronetsec/Comet

package corsscan

import (
    "fmt"
    "net"
    "strings"
    "sync"
    "time"
    "net/http"
    "github.com/Zeronetsec/Comet/utils/color"
    "github.com/Zeronetsec/Comet/utils/logger"
)

func Scan(
    targetInput, origin, customHeader string,
    threads, timeout int,
) {
    targets, err := readTarget(targetInput)
    if err != nil {
        fmt.Printf(
            "%s[!] %sFailed to load targets: %s%v%s\n",
            color.R, color.N, color.GG, err, color.N,
        )
        return
    }

    if len(targets) == 0 {
        fmt.Printf(
            "%s[!] %sNo targets to scan!\n",
            color.R, color.N,
        )
        return
    }

    fmt.Printf(
        "%s[*] %sLoaded: %s%d %starget\n",
        color.B, color.N, color.GG, len(targets), color.N,
    )

    fmt.Printf(
        "%s[*] %sOrigin: %s%s%s\n",
        color.B, color.N, color.GG, origin, color.N,
    )

    if customHeader != "" {
        fmt.Printf(
            "%s[*] %sHeader: %s%s%s\n",
            color.B, color.N, color.GG, customHeader, color.N,
        )
    }

    fmt.Printf(
        "%s[*] %sThreads: %s%d%s\n",
        color.B, color.N, color.GG, threads, color.N,
    )

    fmt.Println()

    client := &http.Client{
        Timeout: time.Duration(timeout) * time.Second,
        Transport: &http.Transport{
            MaxIdleConns: threads,
            MaxIdleConnsPerHost: threads,
            DialContext: (&net.Dialer{
                Timeout: 5 * time.Second,
                KeepAlive: 30 * time.Second,
            }).DialContext,
            TLSHandshakeTimeout: 5 * time.Second,
            ResponseHeaderTimeout: 5 * time.Second,
        },
    }

    var wg sync.WaitGroup
    sem := make(
        chan struct{},
        threads,
    )

    var mu sync.Mutex
    vulnerableCount := 0

    for _, tURL := range targets {
        wg.Add(1)
        sem <- struct{}{}

        go func(targetURL string) {
            defer wg.Done()
            defer func() {
                <-sem
            }()

            req, err := http.NewRequest(
                "GET", targetURL, nil,
            )

            if err != nil {
                return
            }

            req.Header.Set(
                "User-Agent",
                "https://github.com/Zeronetsec/Comet",
            )

            req.Header.Set(
                "Origin", origin,
            )

            if customHeader != "" {
                parts := strings.SplitN(
                    customHeader, ":", 2,
                )

                if len(parts) == 2 {
                    req.Header.Set(
                        strings.TrimSpace(parts[0]),
                        strings.TrimSpace(parts[1]),
                    )
                }
            }

            resp, err := client.Do(req)
            if err != nil {
                return
            }
            resp.Body.Close()

            mu.Lock()
            fmt.Printf(
                "%s[*] %sScanning: %s%s %s-> %s%d%s\n",
                color.B, color.N, color.GG, targetURL, color.DG,
                color.CC, resp.StatusCode, color.N,
            )
            mu.Unlock()

            acao := resp.Header.Get(
                "Access-Control-Allow-Origin",
            )

            acac := resp.Header.Get(
                "Access-Control-Allow-Credentials",
            )

            if acao == origin || acao == "*" {
                mu.Lock()
                vulnerableCount++
                vulnSeverity := "low/Mid"
                if acac == "true" && acao != "*" {
                    vulnSeverity = "high/critical"
                }

                fmt.Printf(
                    "%s[+] %sCORS vuln %s(%s%s%s)%s: %s%s %s[%sACAO: %s%s %s| %sACAC: %s%s%s]%s\n",
                    color.GG, color.N, color.DG,
                    color.YY, vulnSeverity, color.DG,
                    color.N, color.GG, targetURL, color.DG,
                    color.WW, color.CC, acao, color.DG,
                    color.WW, color.CC, acac, color.DG,
                    color.N,
                )

                log := logger.NewLogger("corsscan")
                log.Log(
                    ":", fmt.Sprintf(
                        "CORS vuln (%s): %s [ACAO: %s | ACAC: %s]",
                        vulnSeverity, targetURL, acao, acac,
                    ),
                )
                mu.Unlock()
            }
        }(tURL)
    }

    wg.Wait()
    fmt.Println()

    if vulnerableCount == 0 {
        fmt.Printf(
            "%s[!] %sNo cors misconfigurations found!\n",
            color.R, color.N,
        )
    } else {
        fmt.Printf(
            "%s[*] %sFound: %s%d %svulnerable endpoints!\n",
            color.B, color.N, color.GG, vulnerableCount, color.N,
        )
    }
}

// Copyright (c) 2026 Zeronetsec