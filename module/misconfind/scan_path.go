// Comet Framework

package misconfind

import (
    "os"
    "fmt"
    "path/filepath"
    "comet/utils/color"
)

func scanPath(root string, patterns []string) {
    filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            fmt.Printf(
                "%s[!] %s%s: %s%v%s\n",
                color.R, color.N, path, color.GG, err, color.N,
            )
            return nil
        }

        checkPermission(path, info)

        if !info.IsDir() {
            checkFileContent(path, patterns)
        }

        return nil
    })
}

// Copyright (c) 2026 Zeronetsec