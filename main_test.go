package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func Test_isPrime(t *testing.T) {
	primeTests := []struct {
		name     string
		num      int
		expected bool
		msg      string
	}{
		{"prime", 7, true, "7 is prime number"},
		{"Not prime", 8, false, "8 is not a prime number because it is divided by 2"},
		{"1", 1, false, "number less than 2 can't be a prime number"},
		{"0", 0, false, "number less than 2 can't be a prime number"},
		{"Number = 2", 2, true, "2 is a prime number"},
	}

	for _, v := range primeTests {
		result, msg := isPrime(v.num)
		if v.expected && !result {
			t.Errorf("%s :expected that 7 is a prime number and should be %v but got %v meaning it is not a prime number", v.name, v.expected, result)
		}

		if !v.expected && result {
			t.Errorf("%s :expected that 8 is not prime number and should be %v but got %v meaning which is wrong", v.name, v.expected, result)
		}

		if msg != v.msg {
			t.Error("wrong error message: ", msg)
		}
	}
}

func Test_prompt(t *testing.T) {
	//save original stdout
	oldStdout := os.Stdout
	// create read aned write pipe for capturing our stdout
	r, w, _ := os.Pipe()
	// reinitialize our stdout with the writer(w) from the pipe
	os.Stdout = w

	prompt()
	//  close writer
	_ = w.Close()
	// Restore the original os.Stdout
	os.Stdout = oldStdout

	out, _ := io.ReadAll(r)

	got := string(out)

	if got != "-> " {
		t.Errorf("expected -> but got %v", got)
	}

}

func Test_intro(t *testing.T) {
	oldStdout := os.Stdout

	r, w, _ := os.Pipe()

	os.Stdout = w

	intro()

	_ = w.Close()

	os.Stdout = oldStdout

	out, _ := io.ReadAll(r)

	got := string(out)

	if !strings.Contains(got, "Is it Prime?") {
		t.Errorf("intro text does not exist")
	}
}

// when simulating user input
func Test_checkNumbers(t *testing.T) {

	checkTests := []struct {
		name         string
		input        string
		expectedMsg  string
		expectedBool bool
	}{
		{"quiting case", "q", "", true},
		{"wrong input", "gfhjg", "enter whole number", false},
		{"empty", "", "enter whole number", false},
		{"0", "0", "number less than 2 can't be a prime number", false},
		{"1", "1", "number less than 2 can't be a prime number", false},
		{"2", "2", "2 is a prime number", true},
		{"7", "7", "7 is prime number", true},
	}

	for _, v := range checkTests {
		input := strings.NewReader(v.input)
		reader := bufio.NewScanner(input)

		gotMsg, gotBool := checkNumbers(reader)
		if gotMsg != v.expectedMsg && gotBool != v.expectedBool {
			t.Errorf("something went wrong expectedMsg = %v but got = %v and expected bool = %v but got %v", v.expectedMsg, gotMsg, v.expectedBool, gotBool)
		}
	}
}

func Test_readUserInput(t *testing.T) {
	// create a chan and an io.reader type(basically something that satisfies the io reader interface)
	quitChan := make(chan bool)

	var stdin bytes.Buffer

	stdin.Write([]byte("1\nq\n")) // simulating the user typing 1 the pressing enter then typing q then pressing again

	go readUserInput(&stdin, quitChan)
	<-quitChan

	close(quitChan)
}
