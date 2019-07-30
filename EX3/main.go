package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"github.com/Ann-Geo/Protobuf_course/EX3/addressbookpb"

	"github.com/golang/protobuf/proto"
)

func main() {
	p1 := createMessage()

	readAndWriteDemo(p1)
}

func createMessage() *addressbookpb.Person {
	p1 := addressbookpb.Person{
		Name:  "Steve",
		Id:    123,
		Email: "steve@yahoo.com",
		Phones: []*addressbookpb.Person_PhoneNumber{
			&addressbookpb.Person_PhoneNumber{
				Number: "12345678",
				Type:   addressbookpb.Person_MOBILE,
			},
			&addressbookpb.Person_PhoneNumber{
				Number: "9845670",
				Type:   addressbookpb.Person_HOME,
			},
		},
	}

	return &p1

}

func readAndWriteDemo(p1 proto.Message) {
	writeToFile("simple.bin", p1)
	p2 := &addressbookpb.Person{}
	readFromFile("simple.bin", p2)
	fmt.Println("Read the content:", p2)
}

func writeToFile(fname string, p1 proto.Message) error {
	out, err := proto.Marshal(p1)
	if err != nil {
		log.Fatalln("Can't serialise to bytes", err)
		return err
	}

	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Can't write to file", err)
		return err
	}

	fmt.Println("Data has been written!")
	return nil
}

func readFromFile(fname string, p2 proto.Message) error {
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Something went wrong when reading the file", err)
		return err
	}

	err2 := proto.Unmarshal(in, p2)
	if err2 != nil {
		log.Fatalln("Couldn't put the bytes into the protocol buffers struct", err2)
		return err2
	}

	return nil
}
