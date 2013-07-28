// Dodo is a tiny console application to check quickly if domain name is
// already reserved for .com
package main

import (
    "fmt"
    "bufio"
    "sort"
    "log"
    "flag"
    "os"
    "os/exec"
    "regexp"

    "github.com/wsxiaoys/terminal/color"
)

var (
    file, name string
    showFreeOnly bool
)

func init() {
    flag.StringVar(&file, "f", "", "path to the file with a list of domains to check")
    flag.BoolVar(&showFreeOnly, "show-free-only", false, "show only available domains")
}

func main() {
    flag.Parse()
    args := flag.Args()
    switch {
    case len(args) > 0:
        for i:=0; i < len(args); i++ {
            oneDomainCheck(args[i], showFreeOnly)
        }
    case file != "":
        handle(file)
    }
}

func oneDomainCheck(domain string, flagShowFreeOnly bool) {
    checkedDomain, isFree := check(domain)
    if isFree == true {
        color.Print("@g[FREE]", checkedDomain)
    } else if flagShowFreeOnly == false {
        color.Print("@r[NOT FREE]", checkedDomain)
    }
}

func check(domain string) (checkedDomain string, isFree bool) {
    domain = fmt.Sprint(domain, ".com")
    output, err := exec.Command("whois", domain).Output()
    if err != nil {
        log.Fatal(err)
    }
    re := regexp.MustCompile("No match")
    checkedDomain = fmt.Sprint("\t", domain, "\n")
    if re.FindString(string(output)) != "" {
        isFree = true
    } else {
        isFree = false
    }
    return
}

func handle(file string) {
    filename, err := os.Open(file)
    defer filename.Close()
    if err != nil {
        log.Fatal(err)
    }
    scanner := bufio.NewScanner(filename)
    var domains []string
    for scanner.Scan() {
        domain, isFree := check(scanner.Text())
        if isFree == true {
            domains = append(domains, domain)
        } else if showFreeOnly == false {
            domains =  append(domains, domain)
        }
    }
    sort.Strings(domains)
    for i := 0; i < len(domains); i++ {
        if showFreeOnly == true {
            color.Print("@g[FREE]", domains[i])
        } else {
            fmt.Print(domains[i])
        }
    }
}
