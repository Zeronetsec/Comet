// https://github.com/Zeronetsec/Comet

package tracelink

import (
    "sync"
    "fmt"
    "github.com/Zeronetsec/Comet/utils/color"
)

func Tracer(
    target string,
    threads int,
    recursive bool,
) {
    fmt.Printf(
        "%s[*] %sTarget: %s%s%s\n",
        color.B, color.N, color.GG, target, color.N,
    )

    fmt.Printf(
        "%s[*] %sThreads: %s%d%s\n",
        color.B, color.N, color.GG, threads, color.N,
    )

    fmt.Printf(
        "%s[*] %sRecursive: %s%t%s\n",
        color.B, color.N, color.GG, recursive, color.N,
    )

    results := make(map[string][]string)
    seen := make(map[string]struct{})

    var mu sync.Mutex
    var wg sync.WaitGroup
    var crawl func(string)

    sem := make(chan struct{}, threads)
    crawl = func(u string) {
        defer wg.Done()
        links := fetchLinks(u)

        for _, link := range links {
            mu.Lock()
            if _, ok := seen[link]; ok {
                mu.Unlock()
                continue
            }

            seen[link] = struct{}{}
            results[u] = append(results[u], link)
            mu.Unlock()

            if recursive && isSameHost(target, link) {
                wg.Add(1)
                sem <- struct{}{}

                go func(next string) {
                    defer func() {
                        <-sem
                    }()
                    crawl(next)
                }(link)
            }
        }
    }

    wg.Add(1)
    sem <- struct{}{}
    go func() {
        defer func() {
            <-sem
        }()
        crawl(target)
    }()

    wg.Wait()
    summary(results)
}

// Copyright (c) 2026 Zeronetsec