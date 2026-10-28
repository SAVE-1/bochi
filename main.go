package main

import (
	"context"
	"fmt"

	"github.com/qdrant/go-client/qdrant"
)

// localhost:6333/dashboard

func main() {

	client, err := qdrant.NewClient(&qdrant.Config{
		Host: "localhost",
		Port: 6334, // gRPC port, not 6333
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	// ctx, _ := context.WithTimeout(context.Background(), time.Duration(time.Second*1))

	// 768 must match your embedding model's output size
	client.CreateCollection(context.Background(), &qdrant.CreateCollection{
		CollectionName: "test_collection",
		VectorsConfig: qdrant.NewVectorsConfig(&qdrant.VectorParams{
			Size:     4,
			Distance: qdrant.Distance_Dot,
		}),
	})

	operationInfo, err := client.Upsert(context.Background(), &qdrant.UpsertPoints{
		CollectionName: "test_collection",
		Points: []*qdrant.PointStruct{
			{
				Id:      qdrant.NewIDNum(1),
				Vectors: qdrant.NewVectors(0.05, 0.61, 0.76, 0.74),
				Payload: qdrant.NewValueMap(map[string]any{"city": "Berlin"}),
			},
			{
				Id:      qdrant.NewIDNum(2),
				Vectors: qdrant.NewVectors(0.19, 0.81, 0.75, 0.11),
				Payload: qdrant.NewValueMap(map[string]any{"city": "London"}),
			},
			{
				Id:      qdrant.NewIDNum(3),
				Vectors: qdrant.NewVectors(0.36, 0.55, 0.47, 0.94),
				Payload: qdrant.NewValueMap(map[string]any{"city": "Moscow"}),
			},
			// Truncated
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(operationInfo)

}
