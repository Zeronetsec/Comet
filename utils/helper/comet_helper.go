// Comet Framework

package helper

import (
    "embed"
    "fmt"
    "encoding/json"
    "io/fs"
    "comet/utils/color"
    "comet/utils/asciipicker"
    "comet/utils/birthday"
)

//go:embed metadata/*
var MetadataFS embed.FS

func CometHelper() {
    asciipicker.RandomBanner()
    birthday.CometBirthDay()

    fmt.Printf(
        "%sUsage: %scomet %s<command> [<args>]%s\n",
        color.N, color.GG, color.CC, color.N,
    )

    fmt.Println()
    fmt.Printf(
        "%sAvailable commands:\n",
        color.N,
    )

    files, err := fs.Glob(MetadataFS, "metadata/*.json")
    if err != nil {
        fmt.Printf(
            "%s[!] %sError reading config: %s%s%s\n",
            color.R, color.N, color.GG, err, color.N,
        )
        return
    }

    for _, file := range files {
        data, err := MetadataFS.ReadFile(file)
        if err != nil {
            continue 
        }

        var ch Helper
        err = json.Unmarshal(data, &ch)
        if err != nil {
            continue
        }

        args := ""
        if ch.Args != "" {
            args = " " + ch.Args
        }

        fmt.Printf(
            "    %s* %s%s%s%s%s\n",
            color.DG, color.GG, ch.Command, color.CC, args, color.N,
        )

        fmt.Printf(
            "    %s└── %s%s%s\n",
            color.DG, color.WW, ch.Description, color.N,
        )
    }
}

// Copyright (c) 2026 Zeronetsec