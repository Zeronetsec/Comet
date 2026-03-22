// Comet Framework

package osint

import (
    "strings"
    "net/http"
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
        req, _ := http.NewRequest("GET", url, nil)
        req.Header.Set("User-Agent", "https://github.com/Zeronetsec/Comet")

        resp, err := client.Do(req)
        if err != nil {
            continue
        }

        status := resp.StatusCode
        resp.Body.Close()

        if status == 200 || status == 301 || status == 302 || status == 403 {
            return url, status, true
        }
    }

    return "", 0, false
}

// Copyright (c) 2026 Zeronetsec