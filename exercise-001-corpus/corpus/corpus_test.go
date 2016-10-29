package corpus
import (
    // "fmt"
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestWordCount(t *testing.T) {
    assert := assert.New(t)

    // Sad Path
    _, err := WordCount("nonexistent.txt")
    assert.NotNil(err)

    // Happy Path
    result, err := WordCount("test.txt")
    assert.Nil(err)
    assert.Equal(len(result), 5)
    assert.Equal(result[0].Word, "lots")
    assert.Equal(result[0].Count, 7)
    assert.Equal(result[1].Word, "this")
    assert.Equal(result[1].Count, 2)
    assert.Equal(result[2].Word, "is")
    assert.Equal(result[2].Count, 2)
    assert.Equal(result[3].Word, "of")
    assert.Equal(result[3].Count, 1)
    assert.Equal(result[4].Word, "punctuation")
    assert.Equal(result[4].Count, 1)
}

// What else can we benchmark in this function?
func BenchmarkWordCount(b *testing.B) {
    for n := 0; n < b.N; n++ {
        WordCount("7oldsamr.txt")
    }
}