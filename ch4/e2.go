package main
import (
    "fmt"
    "os"
    "crypto/sha256"
    "crypto/sha512"
)

func main() {
    sMap := map[string]string{
        "sha256":"sha256",
        "sha512":"sha512",
        "sha384":"sha384",
    }

    str := os.Args[1]
    ty  := "all"
    if len(os.Args) > 2 && os.Args[2] != "" {
        ty = os.Args[2]
    }
    if _, ok := sMap[ty]; !ok {
        fmt.Printf("%s\t%x\n", str, sha256.Sum256([]byte(str)))
        fmt.Printf("%s\t%x\n", str, sha512.Sum512([]byte(str)))
        fmt.Printf("%s\t%x\n", str, sha512.Sum384([]byte(str)))
    } else {
        if ty == "sha256" {
            fmt.Printf("%s\t%x\n", str, sha256.Sum256([]byte(str)))
        } else if ty == "sha512" {
            fmt.Printf("%s\t%x\n", str, sha512.Sum512([]byte(str)))
        } else if ty == "sha384" {
            fmt.Printf("%s\t%x\n", str, sha512.Sum384([]byte(str)))
        }
    }
}
