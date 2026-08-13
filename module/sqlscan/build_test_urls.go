// https://github.com/Zeronetsec/Comet

package sqlscan

import (
    "net/url"
)

func buildTestURLs(
    baseURL string,
    payload string,
) []string {
    u, err := url.Parse(baseURL)
    if err != nil {
        return []string{baseURL + payload}
    }

    q := u.Query()
    if len(q) == 0 {
        return []string{baseURL + payload}
    }

    var urls []string
    for key, values := range q {
        for i, val := range values {
            qCopy := url.Values{}
            for k, v := range q {
                qCopy[k] = append([]string(nil), v...)
            }
            qCopy[key][i] = val + payload

            uCopy, _ := url.Parse(baseURL)
            uCopy.RawQuery = qCopy.Encode()
            urls = append(urls, uCopy.String())
        }
    }

    return urls
}

// Copyright (c) 2026 Zeronetsec