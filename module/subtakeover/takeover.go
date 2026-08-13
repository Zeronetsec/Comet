// https://github.com/Zeronetsec/Comet

package subtakeover

import (
    "fmt"
    "io"
    "net"
    "strings"
    "sync"
    "time"
    "net/http"
    "github.com/Zeronetsec/Comet/module/subdomain"
    "github.com/Zeronetsec/Comet/utils/color"
    "github.com/Zeronetsec/Comet/utils/logger"
)

func Takeover(
    domain string,
    threads, timeout, retries int,
) {
    fmt.Printf(
        "%s[*] %sTarget: %s%s%s\n",
        color.B, color.N, color.GG, domain, color.N,
    )

    fmt.Printf(
        "%s[*] %sFetching subdomains...\n",
        color.B, color.N,
    )

    subs, err := subdomain.Fetch(domain, 60, retries)
    if err != nil {
        fmt.Printf(
            "%s[!] %sError fetching subdomains: %s%v%s\n",
            color.R, color.N, color.GG, err, color.N,
        )
        return
    }

    if len(subs) == 0 {
        fmt.Printf(
            "%s[!] %sNo subdomains found!\n",
            color.R, color.N,
        )
        return
    }

    fmt.Printf(
        "%s[*] %sFound: %s%d %ssubdomains\n",
        color.B, color.N, color.GG, len(subs), color.N,
    )

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

    for _, sub := range subs {
        wg.Add(1)
        sem <- struct{}{}

        go func(target string) {
            defer wg.Done()
            defer func() {
                <-sem
            }()

            var resp *http.Response
            var reqErr error

            for i := 0; i < retries; i++ {
                testURL := "http://" + target
                req, _ := http.NewRequest(
                    "GET", testURL, nil,
                )

                req.Header.Set(
                    "User-Agent",
                    "https://github.com/Zeronetsec/Comet",
                )

                resp, reqErr = client.Do(req)
                if reqErr != nil {
                    testURL = "https://" + target
                    req, _ = http.NewRequest(
                        "GET", testURL, nil,
                    )

                    req.Header.Set(
                        "User-Agent",
                        "https://github.com/Zeronetsec/Comet",
                    )

                    resp, reqErr = client.Do(req)
                }

                if reqErr == nil && resp != nil {
                    break
                }

                time.Sleep(time.Duration(i+1) * time.Second)
            }

            if reqErr != nil || resp == nil {
                return
            }
            defer resp.Body.Close()

            mu.Lock()
            fmt.Printf(
                "%s[*] %sScanning: %s%s %s-> %s%d%s\n",
                color.B, color.N, color.DG, target, color.DG,
                color.CC, resp.StatusCode, color.N,
            )
            mu.Unlock()

            bodyBytes, _ := io.ReadAll(resp.Body)
            bodyStr := string(bodyBytes)

            isVuln := false
            var provider string

            for prov, sig := range Fingerprints {
                if strings.Contains(bodyStr, sig) {
                    isVuln = true
                    provider = prov
                    break
                }
            }

            if isVuln {
                mu.Lock()
                vulnerableCount++
                fmt.Printf(
                    "%s[+] %sTakeover: %s%s %s(%s%s%s)%s\n",
                    color.GG, color.N, color.GG, target, color.DG,
                    color.CC, provider, color.DG, color.N,
                )

                log := logger.NewLogger("subtakeover")
                log.Log(
                    ":", fmt.Sprintf(
                        "Takeover: %s (%s)",
                        target, provider,
                    ),
                )
                mu.Unlock()
            }
        }(sub)
    }

    wg.Wait()
    fmt.Println()

    if vulnerableCount == 0 {
        fmt.Printf(
            "%s[!] %sNo takeover vulnerability found!\n",
            color.R, color.N,
        )
    } else {
        fmt.Printf(
            "%s[*] %sFound: %s%d %spotential takeovers!\n",
            color.B, color.N, color.GG, vulnerableCount, color.N,
        )
    }
}

// Copyright (c) 2026 Zeronetsec