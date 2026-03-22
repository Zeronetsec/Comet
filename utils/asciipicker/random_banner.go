// Comet Framework

package asciipicker

import (
    "embed"
    "fmt"
    "time"
    "io/fs"
    "math/rand"
    "comet/utils/color"
)

//go:embed ascii/*
var AsciiFS embed.FS

func RandomBanner() {
    files, err := fs.Glob(AsciiFS, "ascii/*.txt")
    if err != nil {
        fmt.Printf(
            "%s[!] %sFailed reading folder: %s%v%s\n",
            color.R, color.N, color.GG, err, color.N,
        )
        return
    }

    if len(files) == 0 {
        fmt.Printf(
            "%s[!] %sBanner not found!\n",
            color.R, color.N,
        )
        return
    }

    rand.Seed(time.Now().UnixNano())

    randomIndex := rand.Intn(len(files))
    randomFile := files[randomIndex]

    content, err := AsciiFS.ReadFile(randomFile)
    if err != nil {
        fmt.Printf(
            "%s[!] %sFailed read file: %s%s %s(%serror: %s%v%s)%s\n",
            color.R, color.N, color.GG, randomFile, color.DG, color.N,
            color.GG, err, color.DG, color.N,
        )
        return
    }

    fmt.Printf(
        "%s%s%s\n",
        color.B, string(content), color.N,
    )
}

// Copyright (c) 2026 Zeronetsec