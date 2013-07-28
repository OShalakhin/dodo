package main

import (
    "testing"
    "fmt"
)

// test existing domain to be not free
func TestExistingDomain(t *testing.T) {
    domain, isFree := check("google")
    fmt.Println(domain, isFree)
    if isFree == true {
        fmt.Println("this domain is not free")
        t.Fail()
    }
}

// test not existing domain to be free
func TestNotExistingDomain(t *testing.T) {
    domain, isFree := check("some.example")
    fmt.Println(domain, isFree)
    if isFree == false {
        fmt.Println("this domain is free")
        t.Fail()
    }
}
// TODO test files opening

// test show-free-only flag
func TestShowFreeOnly(t *testing.T) {
    showFreeOnly = false
    oneDomainCheck("google", showFreeOnly)
    oneDomainCheck("some.example", showFreeOnly)
}
