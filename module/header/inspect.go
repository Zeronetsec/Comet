// https://github.com/Zeronetsec/Comet

package header

import (
    "fmt"
    "net"
    "strings"
    "time"
    "crypto/tls"
    "net/http"
    "github.com/Zeronetsec/Comet/utils/color"
    "github.com/Zeronetsec/Comet/utils/logger"
)

func Inspect(
    targetURL string,
    timeout time.Duration,
    followRedirect bool,
) {
    if !strings.HasPrefix(
        targetURL, "http://",
    ) && !strings.HasPrefix(
        targetURL, "https://",
    ) {
        targetURL = "https://" + targetURL
    }

    transport := &http.Transport{
        DialContext: (&net.Dialer{
            Timeout: timeout,
        }).DialContext,
        TLSClientConfig: &tls.Config{
            InsecureSkipVerify: true,
        },
    }

    client := &http.Client{
        Timeout: timeout,
        Transport: transport,
    }

    if !followRedirect {
        client.CheckRedirect = func(
            req *http.Request,
            via []*http.Request,
        ) error {
            return http.ErrUseLastResponse
        }
    }

    fmt.Printf(
        "%s[*] %sTarget: %s%s%s\n",
        color.B, color.N, color.GG, targetURL, color.N,
    )

    fmt.Printf(
        "%s[*] %sTimeout: %s%d%s\n",
        color.B, color.N, color.GG, timeout, color.N,
    )

    fmt.Printf(
        "%s[*] %sRedirect: %s%t%s\n",
        color.B, color.N, color.GG, followRedirect, color.N,
    )
    fmt.Println()

    startTime := time.Now()
    req, _ := http.NewRequest("HEAD", targetURL, nil)
    req.Header.Set(
        "User-Agent",
        "https://github.com/Zeronetsec/Comet",
    )

    resp, err := client.Do(req)
    if err != nil {
        req.Method = "GET"
        resp, err = client.Do(req)
        if err != nil {
            fmt.Printf(
                "%s[!] %sError connecting to target: %s%v%s\n",
                color.R, color.N, color.GG, err, color.N,
            )
            return
        }
        defer resp.Body.Close()
    } else {
        defer resp.Body.Close()
    }
    latency := time.Since(startTime)

    ipAddr := "Unknown"
    if addr, err := net.LookupHost(
        resp.Request.URL.Hostname(),
    ); err == nil && len(addr) > 0 {
        ipAddr = addr[0]
    }

    fmt.Printf(
        "%sCommon info:\n",
        color.N,
    )

    fmt.Printf(
        "%s* %sStatus: %s%d %s(%s%s%s)%s\n",
        color.DG, color.N, color.GG, resp.StatusCode,
        color.DG, color.CC, resp.Status, color.DG, color.N,
    )

    fmt.Printf(
        "%s* %sTarget IP: %s%s%s\n",
        color.DG, color.N, color.GG, ipAddr, color.N,
    )

    fmt.Printf(
        "%s* %sLatency: %s%s%s\n",
        color.DG, color.N, color.GG, latency.Round(time.Millisecond), color.N,
    )

    fmt.Println()
    fmt.Printf(
        "%sSSL/TLS Cert info:\n",
        color.N,
    )

    if resp.TLS != nil && len(resp.TLS.PeerCertificates) > 0 {
        cert := resp.TLS.PeerCertificates[0]
        daysLeft := int(time.Until(cert.NotAfter).Hours() / 24)
        fmt.Printf(
            "%s* %sSSL/TLS issuer: %s%s%s\n",
            color.DG, color.N, color.GG, cert.Issuer.CommonName, color.N,
        )

        fmt.Printf(
            "%s* %sCipher: %s%s%s\n",
            color.DG, color.N, color.GG, tls.CipherSuiteName(
                resp.TLS.CipherSuite,
            ), color.N,
        )

        if daysLeft < 0 {
            fmt.Printf(
                "%s* %sSSL expiry: %sEXPIRED%s\n",
                color.DG, color.N, color.R, color.N,
            )
        } else {
            fmt.Printf(
                "%s* %sSSL expiry: %s%d %sdays remaining %s(%s%s%s)%s\n",
                color.DG, color.N, color.GG, daysLeft, color.WW,
                color.DG, color.CC,
                cert.NotAfter.Format(
                    "2006-01-02",
                ), color.DG, color.N,
            )
        }
    } else {
        fmt.Printf(
            "%s* %sSSL/TLS: %sNone %s(%sHTTP Plain%s)%s\n",
            color.DG, color.N, color.R,
            color.DG, color.CC, color.DG, color.N,
        )
    }

    fmt.Println()
    fmt.Printf(
        "%sSecurity headers info:\n",
        color.N,
    )

    securityHeaders := []string{
        "Strict-Transport-Security",
        "X-Frame-Options",
        "X-Content-Type-Options",
        "Content-Security-Policy",
        "X-XSS-Protection",
        "Referrer-Policy",
        "Permissions-Policy",
    }

    for _, sh := range securityHeaders {
        val := resp.Header.Get(sh)
        if val != "" {
            fmt.Printf(
                "%s* %s%s: %s%s%s\n",
                color.DG, color.N, sh, color.GG, val, color.N,
            )
        } else {
            fmt.Printf(
                "%s* %s%s: %sUnknown%s\n",
                color.DG, color.N, sh, color.R, color.N,
            )
        }
    }

    cookies := resp.Cookies()
    fmt.Println()
    fmt.Printf(
        "%sDetected Cookies %s%d%s:\n",
        color.N, color.GG, len(cookies), color.N,
    )

    for _, cookie := range cookies {
        flags := []string{}
        if cookie.HttpOnly {
            flags = append(flags, "HttpOnly")
        }

        if cookie.Secure {
            flags = append(flags, "Secure")
        }

        if cookie.SameSite != http.SameSiteDefaultMode {
            flags = append(
                flags, fmt.Sprintf(
                    "SameSite=%d",
                    cookie.SameSite,
                ),
            )
        }

        flagStr := "None"
        if len(flags) > 0 {
            flagStr = strings.Join(flags, ", ")
        }
        fmt.Printf(
            "%s* %s%s: %s[%sFlags: %s%s%s]%s\n",
            color.DG, color.N, cookie.Name,
            color.DG, color.WW, color.CC, flagStr, color.DG, color.N,
        )
    }

    fmt.Println()
    fmt.Printf(
        "%sRaw Headers:\n",
        color.N,
    )

    log := logger.NewLogger("header")
    for key, values := range resp.Header {
        valueStr := strings.Join(values, ", ")
        fmt.Printf(
            "%s* %s%s: %s%s%s\n",
            color.DG, color.N, key, color.GG, valueStr, color.N,
        )

        logMess := fmt.Sprintf(
            "%s: %s",
            key, valueStr,
        )
        log.Log(":", logMess)
    }
}

// Copyright (c) 2026 Zeronetsec