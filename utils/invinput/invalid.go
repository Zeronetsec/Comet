// Comet Framework

package invinput

import (
    "fmt"
    "comet/utils/color"
)

func Invalid() {
    fmt.Printf(
        "%s[!] %sInvalid input!\n",
        color.R, color.N,
    )

    fmt.Printf(
        "%s[!] %sTry: %scomet --help%s\n",
        color.R, color.N, color.GG, color.N,
    )
}

// Copyright (c) 2026 Zeronetsec