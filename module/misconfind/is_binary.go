// Comet Framework

package misconfind

import (
    "os"
    "io"
)

func isBinary(file *os.File) bool {
    buf := make([]byte, 8000)
    n, err := file.Read(buf)
    if err != nil && err != io.EOF {
        return true
    }

    for i := 0; i < n; i++ {
        if buf[i] == 0 {
            return true
        }
    }

    return false
}

// Copyright (c) 2026 Zeronetsec