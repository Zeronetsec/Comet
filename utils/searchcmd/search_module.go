// Comet Framework

package searchcmd

import (
    "embed"
    "fmt"
    "strings"
    "io/fs"
    "encoding/json"
    "comet/utils/color"
)

//go:embed metadata/*
var MetadataFS embed.FS

func SearchModule(keyword string) {
    files, err := fs.Glob(MetadataFS, "metadata/*.json")
    if err != nil {
        fmt.Printf(
            "%s[!] %sError reading module metadata!\n",
            color.R, color.N,
        )
        return
    }

    keyword = strings.ToLower(keyword)
    var found bool

    for _, file := range files {
        data, err := MetadataFS.ReadFile(file)
        if err != nil {
            continue
        }

        var m Module
        if err := json.Unmarshal(data, &m); err != nil {
            continue
        }

        cmd := strings.ToLower(m.Command)
        desc := strings.ToLower(m.Description)

        if strings.Contains(cmd, keyword) || strings.Contains(desc, keyword) {
            args := ""
            if m.Args != "" {
                args = " " + m.Args
            }

            fmt.Printf(
                "    %s* %s%s%s%s%s\n",
                color.DG, color.GG, m.Command, color.CC, args, color.N,
            )

            fmt.Printf(
                "    %s└── %s%s%s\n",
                color.DG, color.WW, m.Description, color.N,
            )

            found = true
        }
    }

    if !found {
        fmt.Printf(
            "%s[!] %sNo module command found for keyword: %s%s%s\n",
            color.R, color.N, color.GG, keyword, color.N,
        )
    }
}

// Copyright (c) 2026 Zeronetsec