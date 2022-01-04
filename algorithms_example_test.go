package words_test

import (
	"fmt"

	"github.com/go-generics/words"
)

func ExampleZ() {
	fmt.Println(words.Z([]rune("asdf$aasasdasdf")))
	fmt.Println(words.Z([]int{1, 2, 3, 1, 2, 1}))

	// Output:
	// [15 0 0 0 0 1 2 0 3 0 0 4 0 0 0]
	// [6 0 0 2 0 1]
}
