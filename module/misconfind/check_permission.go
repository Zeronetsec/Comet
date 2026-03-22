// Comet Framework

package misconfind

import (
    "fmt"
    "os"
    "comet/utils/color"
)

func checkPermission(path string, info os.FileInfo) {
    mode := info.Mode().Perm()

    if mode&0002 != 0 {
        fmt.Printf(
            "%s[Warning] %sWorld-writable: %s%s %s(%s%o%s)%s\n",
            color.YY, color.N, color.GG, path, color.DG, color.CC, mode, color.DG, color.N,
        )
    }

    if mode&0004 != 0 {
        fmt.Printf(
            "%s[Info] %sWorld-readable: %s%s %s(%s%o%s)%s\n",
            color.B, color.N, color.GG, path, color.DG, color.CC, mode, color.DG, color.N,
        )
    }

    if mode&0111 != 0 && !info.IsDir() {
        fmt.Printf(
            "%s[Info] %sExecutable file: %s%s %s(%s%o%s)%s\n",
            color.B, color.N, color.GG, path, color.DG, color.CC, mode, color.DG, color.N,
        )
    }
}

// Copyright (c) 2026 Zeronetsec