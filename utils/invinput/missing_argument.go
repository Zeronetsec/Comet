// https://github.com/Zeronetsec/Comet

package invinput

import (
    "fmt"
    "github.com/Zeronetsec/Comet/utils/color"
)

func MissingArgument() {
    fmt.Printf(
        "%s[!] %sMissing argument!\n",
        color.R, color.N,
    )

    fmt.Printf(
        "%s[!] %sTry: %scomet --help%s\n",
        color.R, color.N, color.GG, color.N,
    )
}

// Copyright (c) 2026 Zeronetsec