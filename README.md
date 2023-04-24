# go-address-cleaner

Simple utility to format **french** usual addresses into regular ones, meaning without any special char, in Upper Case and with no additionnal spaces.

For exemple `1 Avenue des Champs-élysées 75000 Saint-Tropez` will become `1 AVENUE DES CHAMPS ELYSEES 75000 SAINT TROPEZ`

## Installation 
go get http://github.com/darthyoh/go-address-cleaner

## Usage

```
    import (
        cleaner "github.com/darthyoh/go-address-cleaner"
    )

    func main() {
        str := "1 Avenue des Champs-élysées 75000 Saint-Tropez"
        fmt.Println(cleaner.Clean(str)) // "1 AVENUE DES CHAMPS ELYSEES 75000 SAINT TROPEZ"
    }
    

    
```

