// https://github.com/Zeronetsec/Comet

package osint

func generateVariants(base string) []string {
    return []string{
        "https://" + base,
        "https://" + base + "/",
        "https://www." + base,
        "https://www." + base + "/",
        "http://" + base,
        "http://" + base + "/",
        "http://www." + base,
        "http://www." + base + "/",
    }
}

// Copyright (c) 2026 Zeronetsec