// Comet Framework

package listcmd

import (
    "embed"
    "fmt"
    "encoding/json"
    "io/fs"
    "comet/utils/color"
)

//go:embed metadata/*
var MetadataFS embed.FS

func Lister() {
    fmt.Printf(
        "%sModule commands:\n",
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

        var mod Module
        err = json.Unmarshal(data, &mod)
        if err != nil {
            continue
        }

        args := ""
        if mod.Args != "" {
            args = " " + mod.Args
        }

        fmt.Printf(
            "    %s* %s%s%s%s%s\n",
            color.DG, color.GG, mod.Command, color.CC, args, color.N,
        )

        fmt.Printf(
            "    %s└── %s%s%s\n",
            color.DG, color.WW, mod.Description, color.N,
        )
    }
}

// Copyright (c) 2026 Zeronetsec