// Comet Framework

package dumpstring

import (
    "fmt"
    "crypto/md5"
    "crypto/sha1"
    "crypto/sha256"
    "crypto/sha512"
    "encoding/base32"
    "encoding/base64"
    "encoding/hex"
    "comet/utils/color"
)

func Analyzer(val string) {
    fmt.Printf(
        "%s[*] %sAnalyze strings\n",
        color.B, color.N,
    )

    fmt.Println()

    fmt.Printf(
        "%sValue: %s%s%s\n",
        color.N, color.GG, val, color.N,
    )

    fmt.Printf(
        "%sLength (char): %s%d%s\n",
        color.N, color.GG, len([]rune(val)), color.N,
    )

    fmt.Printf(
        "%sLength (byte): %s%d%s\n",
        color.N, color.GG, len(val), color.N,
    )

    fmt.Printf(
        "%sType: %s%s%s\n",
        color.N, color.GG, detectType(val), color.N,
    )

    fmt.Printf(
        "%sPrintable: %s%s%s\n",
        color.N, color.GG, isPrintable(val), color.N,
    )

    fmt.Printf(
        "%sCharset: %s%s%s\n",
        color.N, color.GG, detectCharset(val), color.N,
    )

    fmt.Printf(
        "%sIs integer: %s%s%s\n",
        color.N, color.GG, yesNo(isInteger(val)), color.N,
    )

    fmt.Printf(
        "%sIs float: %s%s%s\n",
        color.N, color.GG, yesNo(isFloat(val)), color.N,
    )

    fmt.Printf(
        "%sIs boolean: %s%s%s\n",
        color.N, color.GG, yesNo(isBool(val)), color.N,
    )

    fmt.Printf(
        "%sBase16: %s%s%s\n",
        color.N, color.GG, hex.EncodeToString([]byte(val)), color.N,
    )

    fmt.Printf(
        "%sBase32: %s%s%s\n",
        color.N, color.GG, base32.StdEncoding.EncodeToString([]byte(val)), color.N,
    )

    fmt.Printf(
        "%sBase64: %s%s%s\n",
        color.N, color.GG, base64.StdEncoding.EncodeToString([]byte(val)), color.N,
    )

    fmt.Printf(
        "%sMD5: %s%x%s\n",
        color.N, color.GG, md5.Sum([]byte(val)), color.N,
    )

    fmt.Printf(
        "%sSHA1: %s%x%s\n",
        color.N, color.GG, sha1.Sum([]byte(val)), color.N,
    )

    fmt.Printf(
        "%sSHA256: %s%x%s\n",
        color.N, color.GG, sha256.Sum256([]byte(val)), color.N,
    )

    fmt.Printf(
        "%sSHA512: %s%x%s\n",
        color.N, color.GG, sha512.Sum512([]byte(val)), color.N,
    )

    fmt.Printf(
        "%sBinary: %s%s%s\n",
        color.N, color.GG, toBinary(val), color.N,
    )

    fmt.Printf(
        "\n%s[*] %sShanon entropy: %s%.4f bits/byte%s\n",
        color.B, color.N, color.GG, entropy(val), color.N,
    )
}

// Copyright (c) 2026 Zeronetsec