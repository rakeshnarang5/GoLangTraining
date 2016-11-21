package main

import (
	// "fmt"
	"errors"
	"net"
	"testing"
	"time"
)

type MyConn struct {
	buffer      []byte
	left, right int
}

func (c *MyConn) Read(b []byte) (int, error) {
	if c.left == c.right {
		return 0, errors.New("buffer empty. can't perform read operation.")
	}

	n := copy(b, c.buffer[c.left:c.right])

	if n == (c.right - c.left) {
		return n, nil
	} else {
		return n, errors.New("read op failed!")
	}

}

func (c *MyConn) Write(b []byte) (int, error) {

	n := copy(c.buffer[:], b)

	c.right = c.left + n

	if n == len(b) {
		return n, nil
	} else {
		return n, errors.New("Write op failed!")
	}
}

func (c *MyConn) Close() error {
	return nil
}

func (c *MyConn) LocalAddr() net.Addr {
	return nil
}

func (c *MyConn) RemoteAddr() net.Addr {
	return nil
}

func (c *MyConn) SetDeadline(t time.Time) error {
	return nil
}

func (c *MyConn) SetReadDeadline(t time.Time) error {
	return nil
}

func (c *MyConn) SetWriteDeadline(t time.Time) error {
	return nil
}

/*func TestReadMessage(t *testing.T) {

	//var object MyConn

	expected := "Hello, my name is Rakesh!\n"

	b := []byte("Hello, my name is Rakesh!\n")

	object.Write(b)

	actual := readMessage(&object)

	if actual != expected {
		t.Error("Test Failed!")
	}

}*/

func TestHelperSuccess(t *testing.T) {
	expected := "Chitty\n"

	actual := helper("Name\n")

	if actual != expected {
		t.Error("Test for \"success condition\" failed!")
	}
}

func TestHelperFailure(t *testing.T) {
	expected := "Could not identify this command.\n"

	actual := helper("blah\n")

	if actual != expected {
		t.Error("Test for \"failure condition\" failed!")
	}
}

func TestWriteMessage(t *testing.T) {
	var object MyConn

	object.buffer = make([]byte, 1024)

	object.left = 0
	object.right = 0

	n, err := writeMessage(&object, "Hello\n") //6
	// response will be hi, //3

	if err != nil {
		t.Error("Test Failed! error in write op occured")
	}

	if n != 3 {
		t.Error("Test Failed! n not equal to expected value")
	}

	/*expected := "Hi\n"

	b := make([]byte, 1024)

	_, err = object.Read(b)

	if err != nil {
		t.Error("Test Failed! error in read op occured")
	}

	actual := string(b[object.left:object.right])

	if actual != expected {
		t.Error("Test failed!")
	}*/

}

func TestRead(t *testing.T) {
	var object MyConn

	object.buffer = make([]byte, 1024)

	object.left = 0
	object.right = 0

	//expected := "Hello, my name is Rakesh Narang!\n" //expected is a string

	toWrite := []byte("Hello, my name is Rakesh Narang!\n")

	n1 := copy(object.buffer[0:], toWrite)

	object.right = object.left + n1

	/*_, err := object.Write(toWrite)

	if err != nil {
		t.Error("test failed! error in write op")
	}*/

	n, err := read(&object)

	if err != nil {
		t.Error("test failed! error occured in read op")
	}

	if n != n1 {
		t.Error("test failed! n not equal to expected value")
	}

}

// tested internal functions
// they work fine
// testing begins

/*func TestReadMessageInternal(t *testing.T) {

	var object MyConn

	object.buffer = make([]byte, 1024)

	expected := 6
	b := []byte("Hello\n") //6

	n, _ := object.Write(b)

	if n != expected {
		t.Error("no. bytes sent != no. of bytes read")
	}

}

func TestWriteMessageInternal(t *testing.T) {
	var object MyConn

	object.buffer = make([]byte, 1024)

	b := []byte("Hello\n") //6

	n, _ := object.Write(b)
	// n = no. of bytes written into object.buffer
	// now we will perform the read operation

	b2 := make([]byte, 10)

	n2, _ := object.Read(b2)
	// n2 = no. of bytes read from the buffer

	if n2 != n {
		t.Error("no. of bytes read from buffer != no of bytes written to it")
	}

}*/

//testing version 1 ends
