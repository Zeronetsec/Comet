// https://github.com/Zeronetsec/Comet

package uwu

import (
    "fmt"
    "time"
)

func Nyan(duration time.Duration) {
    faces := []string{
        "(｡◕‿◕｡)",
        "(≧◡≦)",
        "ʕ•ᴥ•ʔ",
        "(・ω・)",
        "(๑˃ᴗ˂)ﻭ",
        "(ง'̀-'́)ง",
        "(=^･ω･^=)",
    }

    delay := 200 * time.Millisecond
    end := time.After(duration)
    nyaa := 0

    for {
        select {
            case <-end:
                fmt.Print("\x1b[K")
                return
            default:
                fmt.Printf(
                    "\r%s\x1b[K",
                    faces[nyaa%len(faces)],
                )
            time.Sleep(delay)
            nyaa++
        }
    }
}

// Copyright (c) 2026 Zeronetsec