// https://github.com/Zeronetsec/Comet

package osint

import (
    "fmt"
    "strings"
    "net/http"
    "github.com/Zeronetsec/Comet/utils/color"
)

func checkDomain(
    client *http.Client, username, domain string,
) (string, int, bool) {
    domain = strings.TrimSpace(domain)
    if domain == "" || strings.HasPrefix(domain, "#") {
        return "", 0, false
    }

    base := strings.ReplaceAll(domain, "{}", username)
    variants := generateVariants(base)

    for _, url := range variants {
        fmt.Printf(
            "%s[*] %sChecking: %s%s%s\n",
            color.B, color.N, color.GG, url, color.N,
        )

        req, _ := http.NewRequest("GET", url, nil)
        req.Header.Set(
            "User-Agent",
            "https://github.com/Zeronetsec/Comet",
        )

        resp, err := client.Do(req)
        if err != nil {
            continue
        }

        status := resp.StatusCode
        resp.Body.Close()

        if (status == 200 ||
            status == 301 ||
            status == 302 ||
            status == 403) {
            return url, status, true
        }
    }

    return "", 0, false
}

// Copyright (c) 2026 Zeronetsec