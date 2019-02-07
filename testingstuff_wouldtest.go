package main

// import "testing"

// func TestTableCalculate(t *testing.T) {
// 	var tests = []struct {
// 		input    int
// 		expected int
// 	}{
// 		{2, 4},
// 		{-1, 1},
// 		{0, 2},
// 		{99999, 100001},
// 	}

// 	for _, test := range tests {
// 		if output := Calculate(test.input); output != test.expected {
// 			t.Error("Test failed: {} inputted, {} expected, received: {}", test.input, test.expected, output)
// 		}
// 	}
// }

// type addResult struct {
// 	x        int
// 	y        int
// 	expected int
// }

// var addResults = []addResult{
// 	{1, 1, 2},
// }

// func TestAdd(t *testing.T) {
// 	for _, test := range addResults {
// 		result := Add(test.x, test.y)
// 		if result != test.expected {
// 			t.Fatal("Expected Result not given ")
// 		}
// 	}
// }

// func TestReadFile(t *testing.T) {
// 	data, err := ioutil.ReadFile("testdata/test.data")
// 	if err != nil {
// 		t.Fatal("Could not open file")
// 	}
// 	if string(data) != "hello world" {
// 		t.Fatal("String contents do not match expected")
// 	}
// }
