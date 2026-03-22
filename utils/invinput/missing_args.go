// Comet Framework

package invinput

import (
    "fmt"
    "comet/utils/color"
)

func MissingArgs() {
    fmt.Printf(
        "%s[!] %sInvalid input!\n",
        color.R, color.N,
    )

    fmt.Printf(
        "%s[!] %sMissing arguments!\n",
        color.R, color.N,
    )
}

// Copyright (c) 2026 Zeronetsec