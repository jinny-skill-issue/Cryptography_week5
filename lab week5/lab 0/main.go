package main

import (
    "crypto/md5"
    "crypto/sha1"
    "crypto/sha256"
    "crypto/sha3"
    "crypto/sha512"
    "encoding/hex"
    "fmt"
)

func hashMD5(data []byte) string {
    sum := md5.Sum(data)
    return hex.EncodeToString(sum[:])
}

func hashSHA1(data []byte) string {
    sum := sha1.Sum(data)
    return hex.EncodeToString(sum[:])
}

func hashSHA256(data []byte) string {
    sum := sha256.Sum256(data)
    return hex.EncodeToString(sum[:])
}

func hashSHA512(data []byte) string {
    sum := sha512.Sum512(data)
    return hex.EncodeToString(sum[:])
}

func hashSHA3_256(data []byte) string {
    sum := sha3.Sum256(data)
    return hex.EncodeToString(sum[:])
}

func hashSHA3_512(data []byte) string {
    sum := sha3.Sum512(data)
    return hex.EncodeToString(sum[:])
}

func proofMe(txt1, txt2 string) {
    fmt.Println("======== Name + Hashing Program ========")
    fmt.Printf("Please input value 1: %s\n", txt1)
    fmt.Printf("Please input value 2: %s\n", txt2)

    algorithms := []struct {
        name string
        fn   func([]byte) string
    }{
        {"MD5", hashMD5},
        {"SHA1", hashSHA1},
        {"SHA256", hashSHA256},
        {"SHA512", hashSHA512},
        {"SHA3-256", hashSHA3_256},
        {"SHA3-512", hashSHA3_512},
    }

    for _, alg := range algorithms {
        h1 := alg.fn([]byte(txt1))
        h2 := alg.fn([]byte(txt2))

        result := "No Match!"
        if h1 == h2 {
            result = "Match!"
        }

        fmt.Printf("Hash (%s): %s / %s => %s\n", alg.name, h1, h2, result)
    }
}

func main() {
    var txt1, txt2 string

    fmt.Print("Please input value 1: ")
    fmt.Scan(&txt1)
    fmt.Print("Please input value 2: ")
    fmt.Scan(&txt2)

    proofMe(txt1, txt2)
}