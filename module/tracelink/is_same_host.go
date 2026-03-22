// Comet Framework

package tracelink

import (
    "net/url"
)

func isSameHost(base, link string) bool {
    b, err1 := url.Parse(base)
    l, err2 := url.Parse(link)

    if err1 != nil || err2 != nil {
        return false
    }

    return b.Host == l.Host
}

// Copyright (c) 2026 Zeronetsec