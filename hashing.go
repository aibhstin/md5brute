package md5brute

import "crypto/md5"
import "fmt"

func Hash(data []rune) []byte {
    sum := md5.Sum([]byte(string(data)))
    return sum[:]
}

func HashAndCompare(sum, target []byte) bool {
    return fmt.Sprintf("%x", sum) == fmt.Sprintf("%x", target)
}


