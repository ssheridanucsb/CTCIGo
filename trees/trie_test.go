package trees

import "testing"

func TestTrie2Prefix(t *testing.T) {
	T := NewTrie()
	T.insert("hello")
	T.insert("help")
	T.insert("hallo")
	T.insert("hip")
	T.insert("hummus")
	T.insert("hum")
	output := make(chan string)
	autocomplete("hu", *T, output)

	for i := 0; i < 2; i++ {
		w := <-output
		if w != "hummus" && w != "hum" {
			t.Error("incorrect word")
		}
	}
}

func TestTrie(t *testing.T) {
	T := NewTrie()
	T.insert("hello")
	T.insert("help")
	T.insert("hallo")
	T.insert("hip")
	T.insert("hummus")
	T.insert("hum")
	output := make(chan string)
	autocomplete("h", *T, output)

	wordMap := map[string]bool{"hello": true, "help": true, "hallo": true, "hip": true, "hummus": true, "hum": true}
	for i := 0; i < 6; i++ {
		w := <-output
		_, ok := wordMap[w]
		if !ok {
			t.Error("incorrect word")
		}
	}
}
