// https://github.com/Zeronetsec/Comet

package subdomain

import (
    "fmt"
    "github.com/Zeronetsec/Comet/utils/color"
    "github.com/Zeronetsec/Comet/utils/logger"
)

func Find(domain string, timeout int, retries int) {
    fmt.Printf(
        "%s[*] %sTarget: %s%s%s\n",
        color.B, color.N, color.GG, domain, color.N,
    )

    fmt.Printf(
        "%s[*] %sAPI: %scrt.sh%s\n",
        color.B, color.N, color.GG, color.N,
    )

    fmt.Println()

    subs, err := Fetch(domain, timeout, retries)
    if err != nil {
        fmt.Printf(
            "%s[!] %sError: %s%v%s\n",
            color.R, color.N, color.GG, err, color.N,
        )
        return
    }

    if len(subs) == 0 {
        fmt.Printf(
            "%s[!] %sNo subdomains found for: %s%s%s\n",
            color.R, color.N, color.GG, domain, color.N,
        )
        return
    }

    log := logger.NewLogger("subdomain")
    for _, sub := range subs {
        fmt.Printf(
            "%s* %s%s%s\n",
            color.DG, color.GG, sub, color.N,
        )

        log.Log(
            ":", fmt.Sprintf(
                "Found: %s", sub,
            ),
        )
    }

    fmt.Println()
    fmt.Printf(
        "%s[*] %sTotal subdomains found: %s%d%s\n",
        color.B, color.N, color.GG, len(subs), color.N,
    )
}

// Copyright (c) 2026 Zeronetsec