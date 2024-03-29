// Package main imlements a client for movieinfo service
package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/NaseerLodge/CloudNativeCourse/labs/lab5/movieapi"
	"google.golang.org/grpc"
)

const (
	address      = "localhost:50051"
	defaultTitle = "Pulp fiction"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := movieapi.NewMovieInfoClient(conn)

	// Contact the server and print out its response.Narayana PU College
	title := defaultTitle
	if len(os.Args) > 1 {
		title = os.Args[1]
	}
	// Timeout if server doesn't respond
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetMovieInfo(ctx, &movieapi.MovieRequest{Title: title})
	if err != nil {
		log.Fatalf("could not get movie info: %v", err)
	}
	log.Printf("Movie Info for %s %d %s %v", title, r.GetYear(), r.GetDirector(), r.GetCast())

	//Giving the input MoviData to SetMovie Info with the title, year, director and Cast
	status, err := c.SetMovieInfo(ctx, &movieapi.MovieData{Title: "Top Gun Maverick", Year: 2022, Director: "Joseph Kosinski", Cast: []string{"Tom Cruise, Miles Teller, Jennifer Connelly"}})

	//Error Check
	if err != nil {
		log.Fatalf("could not set movie info: %v", err)
		log.Fatalf("SetMovieInfo Status: %v", status)
	}

	//Aftter setting the MovieData, now using GetMoveInfo to output the details
	r1, err := c.GetMovieInfo(ctx, &movieapi.MovieRequest{Title: "Top Gun Maverick"})
	if err != nil {
		log.Fatalf("could not get movie info: %v", err)
	}
	log.Printf("Movie Info for Top Gun Maverick: %d %s %v", r1.GetYear(), r1.GetDirector(), r1.GetCast())
}
