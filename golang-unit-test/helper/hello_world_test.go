package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("ade")
	}
}

func BenchmarkTable(b *testing.B) {
	users := []struct {
		name    string
		request string
	}{
		{
			name:    "ade",
			request: "ade",
		},
		{
			name:    "Bayu",
			request: "Bayu",
		},
		{
			name:    "rizki",
			request: "rizki",
		},
	}

	for _, user := range users {
		b.Run(user.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(user.request)
			}
		})
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("ade", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("ade")
		}
	})
	b.Run("Bayu", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Bayu")
		}
	})

}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "ade",
			request:  "ade",
			expected: "Hello ade",
		},
		{
			name:     "Bayu",
			request:  "Kheneddy",
			expected: "Hello Bayu",
		},
		{
			name:     "Pake",
			request:  "Kobo",
			expected: "Hello Pake",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("ade", func(t *testing.T) {
		result := HelloWorld("ade")
		assert.Equal(t, "Hello ade", result, "Result is not Hello ade")

		fmt.Println("Success for TestSubTest ade")
	})

	t.Run("Bayu", func(t *testing.T) {
		result := HelloWorld("Kheneddy")
		assert.Equal(t, "Hello Bayu", result, "Result is not Hello Bayu")

		fmt.Println("Success for TestSubTest Bayu")
	})
}

func TestMain(m *testing.M) {
	fmt.Println("Sebelum Unit Test")

	m.Run()

	fmt.Println("Setelah Unit Test")
}

// nama function harus berawalan Test
func TestHelloWorldBayy(t *testing.T) {
	result := HelloWorld("Muhammad")
	if result != "Hello Bayu" {
		// panic("Result is not Hello Muhammad")
		t.Fatal("Result is not Hello Muhammad")
	}
	// tidak akan mengeksekusi baris selanjutnya
	fmt.Println("Success for TestHelloWorldBayy")
}

func TestHelloWorldBayu(t *testing.T) {
	result := HelloWorld("Muhammad")
	if result != "Hello Bayu Dwi S" {
		// panic("Result is not Hello Muhammad")
		t.Error("Result is not Hello Muhammad")
	}

	fmt.Println("Success for TestHelloWorldBayu")
}

func TestHelloWorldPakeAssert(t *testing.T) {
	result := HelloWorld("Kobo")
	assert.Equal(t, "Hello Pake", result, "Result is not Hello Pake")

	fmt.Println("Success for TestHelloWorldPake")
}

func TestHelloWorldPakeRequire(t *testing.T) {
	result := HelloWorld("Kobo")
	require.Equal(t, "Hello Pake", result, "Result is not Hello Pake")

	fmt.Println("Success for TestHelloWorldPake")
}

func TestSkip(t *testing.T) {
	fmt.Println(runtime.GOOS)
	if runtime.GOOS == "darwin" {
		t.Skip("Can't run on MacOS")
	}
	result := HelloWorld("Kobo")
	require.Equal(t, "Hello Kobo", result, "Result is not Hello Kobo")
}
