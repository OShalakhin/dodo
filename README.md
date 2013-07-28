# DODO check domains available in .com

## Usage

Simple usage:

```bash
dodo bamboo

[NO] bamboo.com
```

or check text in files:

```bash
dodo -f domains_to_test.txt

[NOT FREE] bamboo.com
[NOT FREE] google.com
[FREE]     example.com

dodo -show-free-only=true -f domains_to_test.txt
[FREE]     example.com
```

## Installation

Download dodo.

```bash
go get dodo
go install dodo
```
