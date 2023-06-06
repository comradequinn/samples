package main

import (
	"log"
	"marshal/dtos"

	"google.golang.org/protobuf/proto"
)

func main() {
	lt := dtos.LargeType{
		Id:     123456,
		Field1: "ABC123",
		Field2: "ABC123",
		Field3: "ABC123",
		Field4: "ABC123",
		Field5: "ABC123",
		Field6: "ABC123",
		Field7: "ABC123",
		Field8: "ABC123",
	}

	lt.Nested = append(lt.Nested, &dtos.MediumType{
		Id:     123456,
		Field1: "ABC123",
		Field2: "ABC123",
		Field3: "ABC123",
		Field4: "ABC123",
		Field5: "ABC123",
		Field6: "ABC123",
		Field7: "ABC123",
		Field8: "ABC123",
	},
		&dtos.MediumType{
			Id:     123456,
			Field1: "ABC123",
			Field2: "ABC123",
			Field3: "ABC123",
			Field4: "ABC123",
			Field5: "ABC123",
			Field6: "ABC123",
			Field7: "ABC123",
			Field8: "ABC123",
		})

	out, err := proto.Marshal(&lt)

	if err != nil {
		log.Fatalf("marshal error: %v", err)
	}

	log.Printf("marshalled large type: %v", out)

	lt = dtos.LargeType{}

	if err := proto.Unmarshal(out, &lt); err != nil {
		log.Fatalf("unmarshal error: %v", err)
	}

	log.Printf("unmarshalled large type: %+v", &lt)
}
