package maps

import "testing"

func assertStrings(t *testing.T, got, want string) {
	t.Helper()
	if want != got {
		t.Errorf("want %q but got %q", want, got)
	}

}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if want != got {
		t.Errorf("want %q but got %q", want, got)
	}

}

func TestDictionary(t *testing.T) {
	d := Dict{"red": "a colour"}

	t.Run("word exists", func(t *testing.T) {

		want := "a colour"
		got, _ := d.Definition("red")

		assertStrings(t, got, want)
	})

	t.Run("word does not exist", func(t *testing.T) {

		_, got := d.Definition("yellow")

		assertError(t, got, ErrNotFound)
	})

}
func TestAdd(t *testing.T) {
	dictionary := Dict{}
	dictionary.Add("test", "this is just a test")

	want := "this is just a test"
	got, err := dictionary.Definition("test")
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
