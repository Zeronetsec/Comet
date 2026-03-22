// Comet Framework

package version

import (
    "fmt"
    "comet/utils/color"
)

const (
    tools = "Comet"
    version = "v1.0"
    homepage = "https://github.com/Zeronetsec/Comet"
)

func CometVersion() {
    fmt.Printf(
        "%sProject: %s%s%s\n",
        color.N, color.GG, tools, color.N,
    )

    fmt.Printf(
        "%sVersion: %s%s%s\n",
        color.N, color.GG, version, color.N,
    )

    fmt.Printf(
        "%sHomepage: %s%s%s\n",
        color.N, color.GG, homepage, color.N,
    )
}

// Copyright (c) 2026 Zeronetsec