# go-generics/words

Go generic strings structures and algorithms.

`import "github.com/go-generics/words"`

## Pattern searching

### Z Algorithm

For input `s []T`, return `z []int`.

`z[i]` is the maximum length of subslice `s[i:]` that matches with `s`.

Time complexity: $\mathcal{O}(n)$

Memory complexity: $\mathcal{O}(1)$

https://cp-algorithms.com/string/z-function.html

https://www.hackerearth.com/practice/algorithms/string-algorithm/z-algorithm/tutorial/

```go
fmt.Print(words.Z([]rune("asdf$aasasdasdf")))  // [15 0 0 0 0 1 2 0 3 0 0 4 0 0 0]
fmt.Print(words.Z([]int{1, 2, 3, 1, 2, 1}))  // [6 0 0 2 0 1]
```
