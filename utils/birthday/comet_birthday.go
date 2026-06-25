// https://github.com/Zeronetsec/Comet

package birthday

import (
    "fmt"
    "time"
    "github.com/Zeronetsec/Comet/utils/color"
)

func CometBirthDay() {
    birthDate := "03-22"
    now := time.Now().Format("01-02")
    if now == birthDate {
        fmt.Printf(
            "%s› %sHappy birthday for %scomet %s🎉\n",
            color.R, color.N, color.GG, color.N,
        )
        fmt.Println()
    }
}

// Copyright (c) 2026 Zeronetsec