// Comet Framework

package dirfuzzer

import (
    "fmt"
    "strings"
    "sync"
    "net/http"
    "comet/utils/color"
)

func fuzz(
    client *http.Client,
    target, directory string,
    idx int, recursive bool,
    results *map[int][]string,
    dirs *[]string,
    mu *sync.Mutex,
) {
    directory = strings.TrimSpace(directory)
    if directory == "" {
        return
    }

    url := fmt.Sprintf("%s/%s", target, directory)
    req, _ := http.NewRequest("HEAD", url, nil)
    resp, err := client.Do(req)
    status := 0

    if err == nil {
        status = resp.StatusCode
        resp.Body.Close()
    }

    col := color.R
    if status == 200 {
        col = color.GG
    } else if status != 404 && status != 0 {
        col = color.YY
    }

    fmt.Printf(
        "%s[%s%d%s] %s%s%s/%s%s %s(%s%d%s)%s\n",
        color.DG, color.B, idx, color.DG,
        color.GG, target, color.DG, color.YY, directory,
        color.DG, col, status, color.DG, color.N,
	)

    if status == 200 || status == 301 || status == 302 {
        mu.Lock()
        (*results)[status] = append((*results)[status], url)
        mu.Unlock()

        if recursive {
            mu.Lock()
            *dirs = append(*dirs, directory)
            mu.Unlock()
        }
    }
}

// Copyright (c) 2026 Zeronetsec