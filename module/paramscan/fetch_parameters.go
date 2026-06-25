// https://github.com/Zeronetsec/Comet

package paramscan

import (
    "bufio"
    "fmt"
    "sort"
    "strings"
    "sync"
    "time"
    "net"
    "net/http"
    "net/url"
    "github.com/Zeronetsec/Comet/utils/color"
    "github.com/Zeronetsec/Comet/utils/logger"
)

func FetchParameters(
    target string,
    threads, timeout int,
    isFuzzing bool,
) {
    u, _ := url.Parse(target)
    domain := u.Host
    if domain == "" {
        domain = target
    }

    fmt.Printf(
        "%s[*] %sTarget: %s%s%s\n",
        color.B, color.N, color.GG, domain, color.N,
    )

    fmt.Printf(
        "%s[*] %sAPI: %sweb.archive.org%s\n",
        color.B, color.N, color.GG, color.N,
    )

    fmt.Printf(
        "%s[*] %sFuzz: %s%t%s\n",
        color.B, color.N, color.GG, isFuzzing, color.N,
    )
    fmt.Println()

    apiURL := fmt.Sprintf(
        "http://web.archive.org/cdx/search/cdx?url=%s/*&output=txt&fl=original&collapse=urlkey",
        domain,
    )

    client := &http.Client{
        Timeout: 60 * time.Second,
        Transport: &http.Transport{
            MaxIdleConns: 100,
            MaxIdleConnsPerHost: 100,
            DialContext: (&net.Dialer{
                Timeout: 15 * time.Second,
                KeepAlive: 30 * time.Second,
            }).DialContext,
            TLSHandshakeTimeout: 10 * time.Second,
            ResponseHeaderTimeout: 15 * time.Second,
        },
    }

    var resp *http.Response
    var err error
    maxApiRetries := 5

    for i := 0; i < maxApiRetries; i++ {
        req, _ := http.NewRequest("GET", apiURL, nil)
        req.Header.Set(
            "User-Agent",
            "https://github.com/Zeronetsec/Comet",
        )

        resp, err = client.Do(req)
        if err == nil && resp.StatusCode == 200 {
            break
        }

        if resp != nil {
            resp.Body.Close()
        }

        fmt.Printf(
            "%s[!] %sConnection unstable, retrying api %s(%s%d%s/%s%d%s)%s\n",
            color.R, color.N,
            color.DG, color.GG, i+1, color.DG,
            color.CC, maxApiRetries, color.DG, color.N,
        )

        time.Sleep(time.Duration(i+2) * time.Second)
    }

    if err != nil || resp == nil {
        fmt.Printf(
            "%s[!] %sFailed to connect to archives api: %s%v%s\n",
            color.R, color.N, color.GG, err, color.N,
        )
        return
    }
    defer resp.Body.Close()

    var patterns []string
    seen := make(map[string]bool)
    scanner := bufio.NewScanner(resp.Body)

    for scanner.Scan() {
        link := scanner.Text()
        if strings.Contains(link, "?") {
            var item string
            if !isFuzzing {
                item = link
            } else {
                item = formatToFuzz(link)
            }

            if !seen[item] {
                patterns = append(patterns, item)
                seen[item] = true
            }
        }
    }

    if len(patterns) == 0 {
        fmt.Printf(
            "%s[!] %sNo parameterized urls found!\n",
            color.R, color.N,
        )
        return
    }

    if !isFuzzing {
        fmt.Println()
        fmt.Printf(
            "%s[*] %sFound %s%d %surls:\n",
            color.B, color.N, color.GG, len(patterns), color.N,
        )

        sort.Strings(patterns)
        for _, p := range patterns {
            fmt.Printf(
                "%s[*] %sFound: %s%s%s\n",
                color.B, color.N, color.GG, p, color.N,
            )

            log := logger.NewLogger("paramscan")
            logMess := fmt.Sprintf(
                "Found: %s", p,
            )
            log.Log(":", logMess)
        }
        return
    }

    fmt.Println()
    fmt.Printf(
        "%s[*] %sstart parameter fuzzing %s(%s1-10s%s)%s\n",
        color.B, color.N, color.DG, color.CC, color.DG, color.N,
    )

    fmt.Printf(
        "%s[*] %sPatterns: %s%d%s\n",
        color.B, color.N, color.GG, len(patterns), color.N,
    )

    fmt.Printf(
        "%s[*] %sThreads: %s%d%s\n",
        color.B, color.N, color.GG, threads, color.N,
    )
    fmt.Println()

    var wg sync.WaitGroup
    sem := make(chan struct{}, threads)

    var mu sync.Mutex
    results := make(map[int][]string)

    fuzzClient := &http.Client{
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

    for _, p := range patterns {
        for val := 1; val <= 5; val++ {
            wg.Add(1)
            sem <- struct{}{}

            go func(pattern string, v int) {
                defer wg.Done()
                defer func() {
                    <-sem
                }()

                testURL := strings.ReplaceAll(
                    pattern, "FUZZ",
                    fmt.Sprintf("%d", v),
                )

                mu.Lock()
                fmt.Printf(
                    "%s[*] %sFuzz: %s%s%s\n",
                    color.B, color.N, color.GG, testURL, color.N,
                )
                mu.Unlock()

                for retry := 0; retry < 3; retry++ {
                    req, _ := http.NewRequest("GET", testURL, nil)
                    req.Header.Set(
                        "User-Agent",
                        "https://github.com/Zeronetsec/Comet",
                    )

                    res, err := fuzzClient.Do(req)
                    if err == nil {
                        if (res.StatusCode == 200 ||
                            res.StatusCode == 301 ||
                            res.StatusCode == 303) {
                            mu.Lock()
                            results[res.StatusCode] = append(
                                results[res.StatusCode],
                                testURL,
                            )

                            fmt.Printf(
                                "%s[+] %sFound: %s%s %s(%s%d%s)%s\n",
                                color.GG, color.N, color.GG, testURL,
                                color.DG, color.CC, res.StatusCode, color.DG, color.N,
                            )
                            mu.Unlock()
                        }
                        res.Body.Close()
                        break
                    }
                    time.Sleep(200 * time.Millisecond)
                }
            }(p, val)
        }
    }

    wg.Wait()
    summary(results)
}

// Copyright (c) 2026 Zeronetsec