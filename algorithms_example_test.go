package words_test

import (
	"fmt"

	"github.com/go-generics/words"
)

func ExampleZ() {
	fmt.Print(words.Z("asdf$aasasdasdf"))

	// Output:
	// [15 0 0 0 0 1 2 0 3 0 0 4 0 0 0]
}
