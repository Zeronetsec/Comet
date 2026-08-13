// https://github.com/Zeronetsec/Comet

package dnslookup

import (
    "bufio"
    "fmt"
    "net"
    "strings"
    "time"
    "net/http"
    "github.com/Zeronetsec/Comet/utils/color"
    "github.com/Zeronetsec/Comet/utils/logger"
)

func Scan(domain string, timeout int, retries int) {
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
        "https://api.hackertarget.com/dnslookup/?q=%s",
        domain,
    )

    client := &http.Client{
        Timeout: time.Duration(timeout) * time.Second,
        Transport: &http.Transport{
            DialContext: (&net.Dialer{
                Timeout: 15 * time.Second,
                KeepAlive: 30 * time.Second,
            }).DialContext,
            TLSHandshakeTimeout: 10 * time.Second,
            ResponseHeaderTimeout: 15 * time.Second,
        },
    }

    var resp *http.Response
    var err error

    for i := 0; i < retries; i++ {
        req, _ := http.NewRequest("GET", url, nil)
        req.Header.Set(
            "User-Agent",
            "https://github.com/Zeronetsec/Comet",
        )

        resp, err = client.Do(req)
        if err == nil && resp.StatusCode == 200 {
            break
        }

        if resp != nil {
            resp.Body.Close()
        }

        fmt.Printf(
            "%s[!] %sConnection unstable, retrying api %s(%s%d%s/%s%d%s)%s\n",
            color.R, color.N,
            color.DG, color.GG, i+1, color.DG,
            color.CC, retries, color.DG, color.N,
        )

        time.Sleep(time.Duration(i+2) * time.Second)
    }

    if err != nil || resp == nil {
        fmt.Printf(
            "%s[!] %sFailed to connect to api: %s%v%s\n",
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

    scanner := bufio.NewScanner(resp.Body)
    hasOutput := false

    for scanner.Scan() {
        line := scanner.Text()
        if strings.Contains(line, "API count exceeded") {
            fmt.Printf(
                "%s[!] %sRate limit reached!\n",
                color.R, color.N,
            )
            return
        }

        if strings.TrimSpace(line) == "" {
            continue
        }

        log := logger.NewLogger("dnslookup")
        parts := strings.SplitN(line, ":", 2)
        if len(parts) == 2 {
            key := strings.TrimSpace(parts[0])
            value := strings.TrimSpace(parts[1])

            fmt.Printf(
                "%s%s: %s%s%s\n",
                color.N, key, color.GG, value, color.N,
            )

            logMess := fmt.Sprintf(
                "%s: %s",
                key, value,
            )
            log.Log(":", logMess)

            hasOutput = true
        } else {
            fmt.Printf(
                "%s%s%s\n",
                color.YY, line, color.N,
            )

            logMess := fmt.Sprintf(
                "%s",
                line,
            )
            log.Log(":", logMess)
        }
    }

    if !hasOutput {
        fmt.Printf(
            "%s[!] %sNo dns records found or target invalid!\n",
            color.R, color.N,
        )
    }
}

// Copyright (c) 2026 Zeronetsec