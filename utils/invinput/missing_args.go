// Comet Framework

package invinput

import (
    "fmt"
    "comet/utils/color"
)

func MissingArgs() {
    fmt.Printf(
        "%s[!] %sMissing arguments!\n",
        color.R, color.N,
    )

    fmt.Printf(
        "%s[!] %sTry: %scomet --help %sor %scomet --listcmd%s\n",
        color.R, color.N, color.GG, color.N, color.GG, color.N,
    )
}

// Copyright (c) 2026 Zeronetsec