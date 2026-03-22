// Comet Framework

package tracelink

import (
    "fmt"
    "sync"
)

func Tracer(target string, threads int, recursive bool) {
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
    fmt.Println()
    printSummary(results)
}

// Copyright (c) 2026 Zeronetsec