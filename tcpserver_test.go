package main

import (
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

func TestReadMessage(t *testing.T) {
	var object MyConn
	object.buffer = []byte("1234")
	object.left = 0
	object.right = 0
	err := readMessage(&object)
	if err == nil {
		t.Error("error occured in read operation")
	}
}

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
	if err != nil {
		t.Error("Test Failed! error in write op occured")
	}
	if n != 3 {
		t.Error("Test Failed! n not equal to expected value")
	}
}

func TestRead(t *testing.T) {
	var object MyConn
	object.buffer = make([]byte, 1024)
	object.left = 0
	object.right = 0
	toWrite := []byte("Hello, my name is Rakesh Narang!\n")
	n1 := copy(object.buffer[0:], toWrite)
	object.right = object.left + n1
	n, err := read(&object)
	if err != nil {
		t.Error("test failed! error occured in read op")
	}
	if n != n1 {
		t.Error("test failed! n not equal to expected value")
	}
}
