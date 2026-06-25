// https://github.com/Zeronetsec/Comet

package version

import (
    "fmt"
    "github.com/Zeronetsec/Comet/utils/color"
)

const (
    name = "Comet"
    version = "v0.1"
    creator = "Zeronetsec"
    homepage = "https://github.com/Zeronetsec/Comet"
)

func CometVersion() {
    fmt.Printf(
        "%sName: %s%s%s\n",
        color.N, color.GG, name, color.N,
    )

    fmt.Printf(
        "%sVersion: %s%s%s\n",
        color.N, color.GG, version, color.N,
    )

    fmt.Printf(
        "%sCreator: %s%s%s\n",
        color.N, color.GG, creator, color.N,
    )

    fmt.Printf(
        "%sHomepage: %s%s%s\n",
        color.N, color.GG, homepage, color.N,
    )
}

// Copyright (c) 2026 Zeronetsec