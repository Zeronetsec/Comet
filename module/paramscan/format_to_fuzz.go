// Comet Framework

package paramscan

import (
    "net/url"
)

func formatToFuzz(rawURL string) string {
    u, err := url.Parse(rawURL)
    if err != nil {
        return rawURL
    }

    qs := u.Query()
    for param := range qs {
        qs.Set(param, "FUZZ")
    }

    u.RawQuery = qs.Encode()
    final, _ := url.QueryUnescape(u.String())
    return final
}

// Copyright (c) 2026 Zeronetsec