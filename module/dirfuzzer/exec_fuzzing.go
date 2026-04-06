// Comet Framework

package dirfuzzer

import (
    "fmt"
    "net"
    "sort"
    "strings"
    "sync"
    "time"
    "io/fs"
    "net/http"
    "comet/utils/color"
    "comet/utils/logger"
)

func ExecFuzzing(target string, fsys fs.FS, wordlist string, timeout int, recursive bool, threads int) {
    target = strings.TrimRight(target, "/")

    scanner, closeFn, total, err := openWordlist(fsys, wordlist)
    if err != nil {
        fmt.Printf(
            "%s[!] %sWordlist: %s%s %snot found!\n",
            color.R, color.N, color.GG, wordlist, color.N,
        )
        return
    }
    defer closeFn()

    client := &http.Client{
        Timeout: time.Duration(timeout) * time.Second,
        Transport: &http.Transport{
            DialContext: (&net.Dialer{
                Timeout: 2 * time.Second,
                KeepAlive: 2 * time.Second,
            }).DialContext,
            TLSHandshakeTimeout: 2 * time.Second,
            ResponseHeaderTimeout: 3 * time.Second,
            ExpectContinueTimeout: 1 * time.Second,
        },
    }

    _, err = client.Head(target)
    if err != nil {
        fmt.Printf(
            "%s[!] %sTarget: %s%s %snot responding!\n",
            color.R, color.N, color.GG, target, color.N,
        )
        return
    }

    fmt.Printf(
        "%s[*] %sStart fuzzing\n",
        color.B, color.N,
    )

    fmt.Println()
    fmt.Printf(
        "%sTarget: %s%s%s\n",
        color.N, color.GG, target, color.N,
    )

    fmt.Printf(
        "%sWordlist: %s%s %s(%s%d%s)%s\n",
        color.N, color.GG, wordlist, color.DG, color.CC, total, color.DG, color.N,
    )

    fmt.Printf(
        "%sTimeout: %s%d%s\n",
        color.N, color.GG, timeout, color.N,
    )

    fmt.Printf(
        "%sThreads: %s%d%s\n",
        color.N, color.GG, threads, color.N,
    )

    fmt.Printf(
        "%sRecursive: %s%v%s\n",
        color.N, color.GG, recursive, color.N,
    )
    fmt.Println()

    var wg sync.WaitGroup
    sem := make(chan struct{}, threads)

    var mu sync.Mutex
    results := make(map[int][]string)

    var dirs []string

    idx := 1
    for scanner.Scan() {
        dir := scanner.Text()

        wg.Add(1)
        sem <- struct{}{}

        go func(i int, d string) {
            defer wg.Done()
            fuzz(client, target, d, i, recursive, &results, &dirs, &mu)
            <-sem
        }(idx, dir)
        idx++
    }

    wg.Wait()

    if recursive {
        for _, dir := range dirs {
            fmt.Printf(
                "%s[*] %sRecursively fuzzing: %s%s/%s%s\n",
                color.B, color.N, color.GG, target, dir, color.N,
            )

            subScanner, subClose, _, err := openWordlist(fsys, wordlist)
            if err != nil {
                continue
            }

            idx := 1
            for subScanner.Scan() {
                sub := strings.TrimSpace(subScanner.Text())

                wg.Add(1)
                sem <- struct{}{}

                go func(i int, path string) {
                    defer wg.Done()
                    fuzz(client, target, path, i, false, &results, &dirs, &mu)
                    <-sem
                }(idx, dir+"/"+sub)
                idx++
            }

            wg.Wait()
            subClose()
        }
    }

    totalFound := 0
    for _, urls := range results {
        totalFound += len(urls)
    }

    fmt.Println()
    if totalFound == 0 {
        fmt.Printf(
            "%s[!] %sNo directory found!\n",
            color.R, color.N,
        )
        return
    }

    fmt.Printf(
        "%s[*] %sTotal %s%d %sfound:\n",
        color.B, color.N, color.GG, totalFound, color.N,
    )

    log := logger.NewLogger("dirfuzzer")
    statuses := []int{200, 301, 302}
    for _, status := range statuses {
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

            logMess := fmt.Sprintf(
                "Found: %s:%d", u, status,
            )

            log.Log(":", logMess)
        }
    }
}

// Copyright (c) 2026 Zeronetsec