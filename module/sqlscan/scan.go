// https://github.com/Zeronetsec/Comet

package sqlscan

import (
    "fmt"
    "io"
    "net"
    "sync"
    "time"
    "io/fs"
    "net/http"
    "github.com/Zeronetsec/Comet/utils/color"
    "github.com/Zeronetsec/Comet/utils/logger"
)

func Scan(
    targetURL, wordlistPath string,
    embeddedFS fs.FS,
    threads, timeout int,
) {
    fmt.Printf(
        "%s[*] %sTarget: %s%s%s\n",
        color.B, color.N, color.GG, targetURL, color.N,
    )

    displayWordlist := wordlistPath
    if wordlistPath == "embedded" {
        displayWordlist = "embeded:wordlist/sqlscan/"
    }

    fmt.Printf(
        "%s[*] %sWordlist: %s%s%s\n",
        color.B, color.N, color.GG, displayWordlist, color.N,
    )

    payloads, err := LoadWordlist(wordlistPath, embeddedFS)
    if err != nil {
        fmt.Printf(
            "%s[!] %sError loading wordlist: %s%v%s\n",
            color.R, color.N, color.GG, err, color.N,
        )
        return
    }

    if len(payloads) == 0 {
        fmt.Printf(
            "%s[!] %sNo valid payloads found!%s\n",
            color.R, color.N, color.N,
        )
        return
    }

    fmt.Printf(
        "%s[*] %sLoaded: %s%d %spayloads\n",
        color.B, color.N, color.GG, len(payloads), color.N,
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
                Timeout: 10 * time.Second,
                KeepAlive: 30 * time.Second,
            }).DialContext,
            ResponseHeaderTimeout: 10 * time.Second,
        },
    }

    var wg sync.WaitGroup
    sem := make(
        chan struct{},
        threads,
    )

    var mu sync.Mutex
    vulnerableCount := 0

    for _, payload := range payloads {
        testURLs := buildTestURLs(targetURL, payload)
        for _, tURL := range testURLs {
            wg.Add(1)
            sem <- struct{}{}

            go func(testURL, p string) {
                defer wg.Done()
                defer func() {
                    <-sem
                }()

                req, _ := http.NewRequest(
                    "GET", testURL, nil,
                )

                req.Header.Set(
                    "User-Agent",
                    "https://github.com/Zeronetsec/Comet",
                )

                resp, err := client.Do(req)
                if err != nil {
                    return
                }
                defer resp.Body.Close()

                mu.Lock()
                fmt.Printf(
                    "%s[*] %sPayload: %s%s %s-> %s%s %s(%s%d%s)%s\n",
                    color.B, color.N, color.YY, p, color.DG,
                    color.GG, testURL, color.DG,
                    color.CC, resp.StatusCode, color.DG, color.N,
                )
                mu.Unlock()

                bodyBytes, _ := io.ReadAll(resp.Body)
                bodyStr := string(bodyBytes)

                isVuln, sig := checkVulnerable(bodyStr)
                if isVuln {
                    mu.Lock()
                    vulnerableCount++
                    fmt.Printf(
                        "%s[+] %sPossible sqli: %s%s %s(%s%s%s)%s\n",
                        color.GG, color.N, color.GG, testURL, color.DG,
                        color.CC, sig, color.DG, color.N,
                    )

                    log := logger.NewLogger("sqlscan")
                    log.Log(
                        ":", fmt.Sprintf(
                            "Possible sqli: %s (%s)",
                            testURL, sig,
                        ),
                    )
                    mu.Unlock()
                }
            }(tURL, payload)
        }
    }

    wg.Wait()
    fmt.Println()

    if vulnerableCount == 0 {
        fmt.Printf(
            "%s[!] %sNo potential sqli found.\n",
            color.R, color.N,
        )
    } else {
        fmt.Printf(
            "%s[*] %sFound: %s%d %spotential sqli points.\n",
            color.B, color.N, color.GG, vulnerableCount, color.N,
        )
    }
}

// Copyright (c) 2026 Zeronetsec