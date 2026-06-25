// https://github.com/Zeronetsec/Comet

package banner

import (
    "embed"
    "fmt"
    "github.com/Zeronetsec/Comet/utils/color"
)

//go:embed ascii/*.txt
var asciiFS embed.FS

func Show() {
    data, err := asciiFS.ReadFile("ascii/comet_devil.txt")
    if err != nil {
        fmt.Printf(
            "%s[!] %sError loading banner!\n",
            color.R, color.N,
        )
        return
    }

    fmt.Printf(
        "%s%s%s\n",
        color.B, string(data), color.N,
    )
}

// Copyright (c) 2026 Zeronetsec