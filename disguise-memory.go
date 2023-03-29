package main

import (
    "crypto/rand"
    "crypto/subtle"
    "fmt"
)

func main() {
    txtData := []byte("foobar")

    fmt.Printf("Original data: %s\n", string(txtData))

    DisguiseMemory(&txtData, 999)

    fmt.Printf("Overwritten data: %s\n", string(txtData))
}

func DisguiseMemory(pointer *[]byte, times int) {
    for i := 0; i < len(*pointer); i++ {
        (*pointer)[i] = 0
    }

    for i := 0; i < times; i++ {
        randomData := make([]byte, len(*pointer))
        rand.Read(randomData)
        subtle.ConstantTimeCopy(1, *pointer, randomData)
    }
}
