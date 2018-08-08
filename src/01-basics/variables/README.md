# Variables in Go

* short assignment statement `:=` can only be used inside a function

* *Zero Value* - Default value of declared but not initialized variable.
    * `bool` has `false` as zero value
    * `string` has `empty` as zero value
    * `int` and `float32/64` has `0` as zero value
    * `func` has `nil` as zero value

        ```go
        var addFunction func(int, int) // nil
        ```
    * `pointer` has `nil` as zero value
        ```go
            type linkedListNode struct {
                data int
                next *linkedListNode
            }

            var node = linkedListNode{data: 34}
            fmt.Println(node.next) // <nil>
        ```
    * `struct` zero value is composition of its member variables
        ```go
        type linkedListNode struct {
            data int
            next *linkedListNode
        }

        var node linkedListNode
        fmt.Println(node) // {0 <nil>}
        ```  

    