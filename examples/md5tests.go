package main

import "github.com/aibhstin/md5brute"
import "fmt"
import "os"

func main() {
    if len(os.Args) != 2 {
        os.Exit(1)
    }

    arg := os.Args[1]
    stringChan := make(chan string)

    go Check(stringChan, arg)


    for value := range stringChan {
        if value == "" {
            continue
        } else {
            fmt.Println(value)
            close(stringChan)
            os.Exit(0)
        }
    }
}

func Check(results chan string, target string) {
    var str = "a"
    for {
        go CheckStr(str, target, results)
        pStr, _ := md5brute.MutateCharArray([]rune(str), 0)
        //time.Sleep(1 * time.Second)
        //log.Printf("CURRENTLY: STR = %s\n", str)
        str = string(pStr)
    }
}

func CheckStr(str, target string, results chan string) {
    test := fmt.Sprintf("%x", md5brute.Hash([]rune(str)))
    //log.Printf("%x", target)
    if test == target {
        results <- str
    } else {
        results <- ""
    }
    /*
    if !md5brute.HashAndCompare(md5brute.Hash([]rune(str)), target) {
        results <- ""
    } else  {
        results <- str
    }*/
}
