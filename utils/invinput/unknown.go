// Comet Framework

package invinput

import (
    "fmt"
    "comet/utils/color"
)

func Unknown(input string) {
    fmt.Printf(
        "%s[!] %sUnknown command: %s%s%s\n",
        color.R, color.N, color.GG, input, color.N,
    )

    fmt.Printf(
        "%s[!] %sTry: %scomet --listcmd %sor %scomet --searchcmd %s<keyword>%s\n",
        color.R, color.N, color.GG, color.N, color.GG, color.CC, color.N,
    )
}

// Copyright (c) 2026 Zeronetsec