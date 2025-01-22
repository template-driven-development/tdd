package template

import "testing"

type testCase struct {
	input    string
	expected string
}

func test(t *testing.T, f func(string) string, fName string, testCases []testCase) {
	for _, testCase := range testCases {
		result := f(testCase.input)
		if result != testCase.expected {
			t.Errorf("expected %s(%q) to be %q, got %q", fName, testCase.input, testCase.expected, result)
		}
	}
}

func TestToCamelCase(t *testing.T) {
	test(t, toCamelCase, "toCamelCase", []testCase{
		{"hello world", "helloWorld"},
		{"Hello world", "helloWorld"},
		{"Hello World", "helloWorld"},
		{"hello World", "helloWorld"},
		{"HelloWorld", "helloWorld"},
		{"helloWorld", "helloWorld"},
		{"hello_world", "helloWorld"},
		{"Hello_world", "helloWorld"},
		{"Hello_World", "helloWorld"},
		{"hello_World", "helloWorld"},
		{"hello-world", "helloWorld"},
		{"Hello-world", "helloWorld"},
		{"Hello-World", "helloWorld"},
		{"hello-World", "helloWorld"},
	})
}

func TestToPascalCase(t *testing.T) {
	test(t, toPascalCase, "toPascalCase", []testCase{
		{"hello world", "HelloWorld"},
		{"Hello world", "HelloWorld"},
		{"Hello World", "HelloWorld"},
		{"hello World", "HelloWorld"},
		{"HelloWorld", "HelloWorld"},
		{"helloWorld", "HelloWorld"},
		{"hello_world", "HelloWorld"},
		{"Hello_world", "HelloWorld"},
		{"Hello_World", "HelloWorld"},
		{"hello_World", "HelloWorld"},
		{"hello-world", "HelloWorld"},
		{"Hello-world", "HelloWorld"},
		{"Hello-World", "HelloWorld"},
		{"hello-World", "HelloWorld"},
	})
}

func TestToSnakeCase(t *testing.T) {
	test(t, toSnakeCase, "toSnakeCase", []testCase{
		{"hello world", "hello_world"},
		{"Hello world", "hello_world"},
		{"Hello World", "hello_world"},
		{"hello World", "hello_world"},
		{"HelloWorld", "hello_world"},
		{"helloWorld", "hello_world"},
		{"hello_world", "hello_world"},
		{"Hello_world", "hello_world"},
		{"Hello_World", "hello_world"},
		{"hello_World", "hello_world"},
		{"hello-world", "hello_world"},
		{"Hello-world", "hello_world"},
		{"Hello-World", "hello_world"},
		{"hello-World", "hello_world"},
	})
}

func TestToKebabCase(t *testing.T) {
	test(t, toKebabCase, "toKebabCase", []testCase{
		{"hello world", "hello-world"},
		{"Hello world", "hello-world"},
		{"Hello World", "hello-world"},
		{"hello World", "hello-world"},
		{"HelloWorld", "hello-world"},
		{"helloWorld", "hello-world"},
		{"hello_world", "hello-world"},
		{"Hello_world", "hello-world"},
		{"Hello_World", "hello-world"},
		{"hello_World", "hello-world"},
		{"hello-world", "hello-world"},
		{"Hello-world", "hello-world"},
		{"Hello-World", "hello-world"},
		{"hello-World", "hello-world"},
	})
}
