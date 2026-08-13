// https://github.com/Zeronetsec/Comet

package subdomain

import (
    "fmt"
    "strings"
    "time"
    "net"
    "sort"
    "net/http"
    "encoding/json"
)

func Fetch(
    domain string,
    timeout int,
    retries int,
) ([]string, error) {
    url := fmt.Sprintf(
        "https://crt.sh/?q=%%.%s&output=json",
        domain,
    )

    client := &http.Client{
        Timeout: time.Duration(timeout) * time.Second,
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

    for i := 0; i < retries; i++ {
        req, _ := http.NewRequest(
            "GET", url, nil,
        )

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

        time.Sleep(time.Duration(i+2) * time.Second)
    }

    if err != nil || resp == nil {
        return nil, fmt.Errorf(
            "failed to connect to api: %v",
            err,
        )
    }
    defer resp.Body.Close()

    if resp.StatusCode != 200 {
        return nil, fmt.Errorf(
            "server returned status code: %d",
            resp.StatusCode,
        )
    }

    var results []CrtResult
    err = json.NewDecoder(resp.Body).Decode(&results)
    if err != nil {
        return nil, err
    }

    uniqueSubs := make(map[string]bool)
    for _, res := range results {
        subs := strings.Split(res.NameValue, "\n")
        for _, sub := range subs {
            sub = strings.TrimSpace(sub)
            sub = strings.TrimPrefix(sub, "*.")
            if strings.HasSuffix(sub, domain) {
                uniqueSubs[sub] = true
            }
        }
    }

    var sortedSubs []string
    for sub := range uniqueSubs {
        sortedSubs = append(sortedSubs, sub)
    }
    sort.Strings(sortedSubs)

    return sortedSubs, nil
}

// Copyright (c) 2026 Zeronetsec