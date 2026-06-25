// https://github.com/Zeronetsec/Comet

package hostsearch

import (
    "bufio"
    "fmt"
    "strings"
    "time"
    "net/http"
    "github.com/Zeronetsec/Comet/utils/color"
    "github.com/Zeronetsec/Comet/utils/logger"
)

type HostResult struct {
    Subdomain string
    IPAddress string
}

func Scan(domain string) {
    fmt.Printf(
        "%s[*] %sTarget: %s%s%s\n",
        color.B, color.N, color.GG, domain, color.N,
    )

    fmt.Printf(
        "%s[*] %sAPI: %sapi.hackertarget.com%s\n",
        color.B, color.N, color.GG, color.N,
    )
    fmt.Println()

    url := fmt.Sprintf(
        "https://api.hackertarget.com/hostsearch/?q=%s",
        domain,
    )

    client := &http.Client{
        Timeout: 20 * time.Second,
    }

    resp, err := client.Get(url)
    if err != nil {
        fmt.Printf(
            "%s[!] %sConnection error: %s%v%s\n",
            color.R, color.N, color.GG, err, color.N,
        )
        return
    }
    defer resp.Body.Close()

    if resp.StatusCode != 200 {
        fmt.Printf(
            "%s[!] %sServer returned error code: %s%d%s\n",
            color.R, color.N, color.GG, resp.StatusCode, color.N,
        )
        return
    }

    var list []HostResult
    maxSubLen := 15

    scanner := bufio.NewScanner(resp.Body)
    for scanner.Scan() {
        line := scanner.Text()
        if strings.Contains(line, "API count exceeded") {
            fmt.Printf(
                "%s[!] %sRate limit reached!\n",
                color.R, color.N,
            )
            return
        }

        parts := strings.Split(line, ",")
        if len(parts) == 2 {
            sub := strings.TrimSpace(parts[0])
            ip := strings.TrimSpace(parts[1])

            list = append(
                list,
                HostResult{
                    Subdomain: sub,
                    IPAddress: ip,
                },
            )

            if len(sub) > maxSubLen {
                maxSubLen = len(sub)
            }
        }
    }

    if len(list) == 0 {
        fmt.Printf(
            "%s[!] %sNo data found or target invalid!\n",
            color.R, color.N,
        )
        return
    }

    for _, item := range list {
        fmt.Printf(
            "%s* %s%s%s:%s%s%s\n",
            color.DG, color.GG, item.Subdomain, color.DG,
            color.CC, item.IPAddress, color.N,
        )

        log := logger.NewLogger("hostsearch")
        logMess := fmt.Sprintf(
            "Found: %s:%s",
            item.Subdomain, item.IPAddress,
        )
        log.Log(":", logMess)
    }

    fmt.Println()
    fmt.Printf(
        "%s[*] %sTotal found: %s%d %shost\n",
        color.B, color.N, color.GG, len(list), color.N,
    )
}

// Copyright (c) 2026 Zeronetsec