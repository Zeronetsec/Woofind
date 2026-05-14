// https://github.com/Zeronetsec/Woofind

package version

import (
    "fmt"
    "woofind/utils/color"
)

const (
    tools = "Woofind"
    version = "v0.1"
    creator = "Zeronetsec"
    homepage = "https://github.com/Zeronetsec/Woofind"
)

func WoofindVersion() {
    fmt.Printf(
        "%sProject: %s%s%s\n",
        color.N, color.GG, tools, color.N,
    )

    fmt.Printf(
        "%sVersion: %s%s%s\n",
        color.N, color.GG, version, color.N,
    )

    fmt.Printf(
        "%sCreator: %s%s%s\n",
        color.N, color.GG, creator, color.N,
    )

    fmt.Printf(
        "%sHomepage: %s%s%s\n",
        color.N, color.GG, homepage, color.N,
    )
}