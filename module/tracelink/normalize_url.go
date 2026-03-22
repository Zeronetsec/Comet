// Comet Framework

package tracelink

import (
    "strings"
    "net/url"
)

func normalizeURL(base, href string) string {
    href = strings.TrimSpace(href)
    if href == "" || strings.HasPrefix(href, "javascript:") || strings.HasPrefix(href, "mailto:") {
        return ""
    }

    u, err := url.Parse(href)
    if err != nil {
        return ""
    }

    if u.IsAbs() {
        return u.String()
    }

    baseURL, err := url.Parse(base)
    if err != nil {
        return ""
    }

    return baseURL.ResolveReference(u).String()
}

// Copyright (c) 2026 Zeronetsec