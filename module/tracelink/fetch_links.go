// Comet Framework

package tracelink

import (
    "net"
    "time"
    "net/http"
    "golang.org/x/net/html"
)

func fetchLinks(target string) []string {
    client := &http.Client{
        Timeout: 10 * time.Second,
        Transport: &http.Transport{
            DialContext: (&net.Dialer{
                Timeout: 5 * time.Second,
                KeepAlive: 5 * time.Second,
            }).DialContext,
            TLSHandshakeTimeout: 5 * time.Second,
            ResponseHeaderTimeout: 6 * time.Second,
            ExpectContinueTimeout: 1 * time.Second,
        },
    }

    resp, err := client.Get(target)
    if err != nil {
        return nil
    }
    defer resp.Body.Close()

    z := html.NewTokenizer(resp.Body)
    var links []string

    for {
        tt := z.Next()
        switch tt {
            case html.ErrorToken:
                return links
            case html.StartTagToken, html.SelfClosingTagToken:
                t := z.Token()

                if t.Data == "a" {
                    for _, attr := range t.Attr {
                        if attr.Key == "href" {
                            link := normalizeURL(target, attr.Val)
                            if link != "" {
                                links = append(links, link)
                            }
                        }
                    }
                }
        }
    }
}

// Copyright (c) 2026 Zeronetsec