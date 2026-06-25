// https://github.com/Zeronetsec/Comet

package subdomain

import (
    "fmt"
    "sort"
    "strings"
    "time"
    "encoding/json"
    "net/http"
    "github.com/Zeronetsec/Comet/utils/color"
    "github.com/Zeronetsec/Comet/utils/logger"
)

type CrtResult struct {
    NameValue string `json:"name_value"`
}

func Find(domain string) {
    fmt.Printf(
        "%s[*] %sTarget: %s%s%s\n",
        color.B, color.N, color.GG, domain, color.N,
    )

    fmt.Printf(
        "%s[*] %sAPI: %scrt.sh%s\n",
        color.B, color.N, color.GG, color.N,
    )
    fmt.Println()

    url := fmt.Sprintf(
        "https://crt.sh/?q=%%.%s&output=json",
        domain,
    )

    client := &http.Client{
        Timeout: 30 * time.Second,
    }

    resp, err := client.Get(url)
    if err != nil {
        fmt.Printf(
            "%s[!] %sError connecting to server: %s%v%s\n",
            color.R, color.N, color.GG, err, color.N,
        )
        return
    }
    defer resp.Body.Close()

    if resp.StatusCode != 200 {
        fmt.Printf(
            "%s[!] %sServer returned status code: %s%d%s\n",
            color.R, color.N, color.GG, resp.StatusCode, color.N,
        )
        return
    }

    var results []CrtResult
    err = json.NewDecoder(resp.Body).Decode(&results)
    if err != nil {
        fmt.Printf(
            "%s[!] %sFailed to parse JSON response: %s%v%s\n",
            color.R, color.N, color.GG, err, color.N,
        )
        return
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

    if len(sortedSubs) == 0 {
        fmt.Printf(
            "%s[!] %sNo subdomains found for: %s%s%s\n",
            color.R, color.N, color.GG, domain, color.N,
        )
        return
    }

    log := logger.NewLogger("subdomain")
    for _, sub := range sortedSubs {
        fmt.Printf(
            "%s* %s%s%s\n",
            color.DG, color.GG, sub, color.N,
        )

        logMess := fmt.Sprintf("Found: %s", sub)
        log.Log(":", logMess)
    }

    fmt.Println()
    fmt.Printf(
        "%s[*] %sTotal Subdomains Found: %s%d%s\n",
        color.B, color.N, color.GG, len(sortedSubs), color.N,
    )
}

// Copyright (c) 2026 Zeronetsec