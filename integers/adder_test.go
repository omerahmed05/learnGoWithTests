package integers

import "testing"
import "fmt"

func TestAdd(t *testing.T) {
	got := Add(1, 1)
	expect := 2

	if got != expect {
		t.Errorf("got %d but expect %d", got, expect)
	}
}

func ExampleAdd() {
	result := Add(1, 5)
	fmt.Println(result)
	// Output: 7
}