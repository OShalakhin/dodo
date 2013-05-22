// Dodo is a tiny console application to check quickly if domain name is
// already reserved for .com
package main

import (
    "fmt"
    "bufio"
    "log"
    "flag"
    "os"
    "os/exec"
    "regexp"

    "github.com/wsxiaoys/terminal/color"
)

var (
    file, name string
)

func init() {
    flag.StringVar(&file, "f", "", "path to the file with a list of domains to check")
}

func main() {
    flag.Parse()
    args := flag.Args()
    switch {
    case len(args) > 0:
        for i:=0; i < len(args); i++ {
            check(args[i])
        }
    case file != "":
        handle(file)
    }
}

func check(domain string) {
    domain = fmt.Sprint(domain, ".com")
    output, err := exec.Command("whois", domain).Output()
    if err != nil {
        log.Fatal(err)
    }
    re := regexp.MustCompile("No match")
    ending := fmt.Sprint("\t", domain, "\n")
    if re.FindString(string(output)) != "" {
        color.Print("@g[Open]", ending)
    } else {
        color.Print("@r[Taken]", ending)
    }
}

func handle(file string) {
    filename, err := os.Open(file)
    if err != nil {
        log.Fatal(err)
    }
    scanner := bufio.NewScanner(filename)
    for scanner.Scan() {
        check(scanner.Text())
    }
    defer filename.Close()
}
