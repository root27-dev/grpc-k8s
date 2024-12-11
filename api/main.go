package main

import (
	"context"
	pb "github.com/root27-dev/grpc-k8s/pb"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"strconv"
)

func main() {

	conn, err := grpc.Dial("add-service:50051", grpc.WithInsecure())

	if err != nil {

		log.Fatalf("Failed to dial: %v", err)

	}

	addClient := pb.NewAddServiceClient(conn)

	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {

		queries := r.URL.Query()

		if len(queries) == 0 {

			w.Write([]byte("No queries found"))
			return

		}

		a, _ := strconv.ParseUint(queries.Get("a"), 10, 32)
		b, _ := strconv.ParseUint(queries.Get("b"), 10, 32)

		req := &pb.AddRequest{A: a, B: b}

		res, err := addClient.Add(context.Background(), req)

		if err != nil {

			w.Write([]byte("Failed to add: " + err.Error()))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`Sum is ` + strconv.FormatUint(res.Result, 10)))

	})

	http.ListenAndServe(":8080", nil)

}
