// Comet Framework

package decode

import (
    "fmt"
    "encoding/hex"
    "encoding/base32"
    "encoding/base64"
    "comet/utils/color"
)

func brute(data []byte, depth int) {
    input := string(data)

    decoders := []struct {
        name string
        check func(string) bool
        decode func(string) ([]byte, error)
    }{
        {
            "base16",
            looksBase16,
            func(s string) ([]byte, error) {
                return hex.DecodeString(s)
            },
        },
        {
            "base32",
            looksBase32,
            func(s string) ([]byte, error) {
                return base32.StdEncoding.DecodeString(s)
            },
        },
        {
            "base64",
            looksBase64,
            func(s string) ([]byte, error) {
                return base64.StdEncoding.DecodeString(s)
            },
        },
    }

    for _, d := range decoders {
        steps++

        if !d.check(input) {
            continue
        }

        decoded, err := d.decode(input)
        if err != nil {
            continue
        }

        str := string(decoded)
        if visited[str] {
            continue
        }

        visited[str] = true

        if !looksReasonable(decoded) {
            continue
        }

        fmt.Println()
        fmt.Printf(
            "%sLayer: %s%d%s\n",
            color.N, color.GG, depth+1, color.N,
        )

        fmt.Printf(
            "%sDecoder: %s%s%s\n",
            color.N, color.GG, d.name, color.N,
        )

        fmt.Printf(
            "%sLength: %s%d%s\n",
            color.N, color.GG, len(decoded), color.N,
        )

        fmt.Printf(
            "%sPrintable: %s%.1f%%%s\n",
            color.N, color.GG, printableRatio(decoded)*100, color.N,
        )

        fmt.Printf(
            "%sEntropy: %s%.2f%s\n",
            color.N, color.GG, entropy(decoded), color.N,
        )

        fmt.Printf(
            "%sResult: %s%s%s\n",
            color.N, color.GG, str, color.N,
        )

        brute(decoded, depth+1)
    }
}

// Copyright (c) 2026 Zeronetsec