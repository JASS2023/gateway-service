package main

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestMsgHandlerCinstruction(t *testing.T) {
	ConnectDB()
	input, err := os.ReadFile("jsonfile.json")
	if err != nil {
		t.Fatal(err)
	}
	str := fmt.Sprintf("%v", input)
	t.Log(str)
	var msg Message
	err = json.Unmarshal(input, &msg)
	if err != nil {
		t.Fatal(err)
	}
	id, err := msgHandler(msg)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(id)
	fmt.Println(id)
}

func TestMsgHandlerService(t *testing.T) {
	ConnectDB()
	input, err := os.ReadFile("jsonfile1.json")
	if err != nil {
		t.Fatal(err)
	}
	str := fmt.Sprintf("%v", input)
	t.Log(str)
	var msg Message
	err = json.Unmarshal(input, &msg)
	if err != nil {
		t.Fatal(err)
	}
	id, err := msgHandler(msg)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(id)
	fmt.Println(id)
}
