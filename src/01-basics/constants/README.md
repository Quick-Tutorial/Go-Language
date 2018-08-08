# const, iota in Go

* it is an auto-incrementing constants
* `iota` is very similer to C’s enum
* The values of iota is reset to 0 whenever the reserved word const appears in the source (i.e. each const block).
    ```go
    package main

    import "fmt"

    func main() {

        const (
            A = iota // 0
            B        // 1
            C = iota // 2
        )

        fmt.Println(A, B, C)

        // iota will reset to zero

        const (
            ROCK    = iota // 0
            PAPER          // 1
            SCISSOR        // 2
        )

        fmt.Println(ROCK, PAPER, SCISSOR)
    }
    ```


## https://blog.golang.org/constants
