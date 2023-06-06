package bench

import (
	"encoding/json"
	"log"
	"marshal/dtos"
	"testing"

	"google.golang.org/protobuf/proto"
)

func BenchmarkSmall(b *testing.B) {
	t := dtos.SmallType{
		Id:     123456,
		Field1: "ABC123",
		Field2: "ABC123",
	}

	protoBytes, err := proto.Marshal(&t)
	if err != nil {
		log.Fatalf("proto marshal error: %v", err)
	}

	jsonBytes, err := json.Marshal(&t)
	if err != nil {
		log.Fatalf("json marshal error: %v", err)
	}

	b.ResetTimer()

	b.Run("proto-marshal", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := proto.Marshal(&t)

			if err != nil {
				log.Fatalf("proto marshal error: %v", err)
			}
		}
	})

	b.Run("json-marshal", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := json.Marshal(&t)

			if err != nil {
				log.Fatalf("json marshal error: %v", err)
			}
		}
	})

	b.Run("proto-unmarshal", func(b *testing.B) {
		lt := dtos.LargeType{}

		for i := 0; i < b.N; i++ {
			if err := proto.Unmarshal(protoBytes, &lt); err != nil {
				log.Fatalf("proto unmarshal error: %v", err)
			}
		}
	})

	b.Run("json-unmarshal", func(b *testing.B) {
		lt := dtos.LargeType{}

		for i := 0; i < b.N; i++ {
			if err := json.Unmarshal(jsonBytes, &lt); err != nil {
				log.Fatalf("json unmarshal error: %v", err)
			}
		}
	})
}

func BenchmarkMedium(b *testing.B) {
	t := dtos.MediumType{
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

	protoBytes, err := proto.Marshal(&t)
	if err != nil {
		log.Fatalf("proto marshal error: %v", err)
	}

	jsonBytes, err := json.Marshal(&t)
	if err != nil {
		log.Fatalf("json marshal error: %v", err)
	}

	b.ResetTimer()

	b.Run("proto-marshal", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := proto.Marshal(&t)

			if err != nil {
				log.Fatalf("proto marshal error: %v", err)
			}
		}
	})

	b.Run("json-marshal", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := json.Marshal(&t)

			if err != nil {
				log.Fatalf("json marshal error: %v", err)
			}
		}
	})

	b.Run("proto-unmarshal", func(b *testing.B) {
		lt := dtos.LargeType{}

		for i := 0; i < b.N; i++ {
			if err := proto.Unmarshal(protoBytes, &lt); err != nil {
				log.Fatalf("proto unmarshal error: %v", err)
			}
		}
	})

	b.Run("json-unmarshal", func(b *testing.B) {
		lt := dtos.LargeType{}

		for i := 0; i < b.N; i++ {
			if err := json.Unmarshal(jsonBytes, &lt); err != nil {
				log.Fatalf("json unmarshal error: %v", err)
			}
		}
	})
}

func BenchmarkLarge(b *testing.B) {
	t := dtos.LargeType{
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

	t.Nested = append(t.Nested, &dtos.MediumType{
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

	protoBytes, err := proto.Marshal(&t)
	if err != nil {
		log.Fatalf("proto marshal error: %v", err)
	}

	jsonBytes, err := json.Marshal(&t)
	if err != nil {
		log.Fatalf("json marshal error: %v", err)
	}

	b.ResetTimer()

	b.Run("proto-marshal", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := proto.Marshal(&t)

			if err != nil {
				log.Fatalf("proto marshal error: %v", err)
			}
		}
	})

	b.Run("json-marshal", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := json.Marshal(&t)

			if err != nil {
				log.Fatalf("json marshal error: %v", err)
			}
		}
	})

	b.Run("proto-unmarshal", func(b *testing.B) {
		lt := dtos.LargeType{}

		for i := 0; i < b.N; i++ {
			if err := proto.Unmarshal(protoBytes, &lt); err != nil {
				log.Fatalf("proto unmarshal error: %v", err)
			}
		}
	})

	b.Run("json-unmarshal", func(b *testing.B) {
		lt := dtos.LargeType{}

		for i := 0; i < b.N; i++ {
			if err := json.Unmarshal(jsonBytes, &lt); err != nil {
				log.Fatalf("json unmarshal error: %v", err)
			}
		}
	})
}
