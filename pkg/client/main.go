package main

import (
	"context"
	grpc2 "file_storage/internal/handler/grpc"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "localhost:8000", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := grpc2.NewFileStorageClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	//r, err := c.SayHello(ctx, &kitty.HelloRequest{Name: *name})
	//if err != nil {
	//	log.Fatalf("could not greet: %v", err)
	//}

	//log.Printf("Greeting: %s", r.GetMessage())

	// upload photo
	path := ".\\examplePhotoForUpload\\4_8.jpg"
	postPhoto, err := os.ReadFile(path)
	if err != nil {
		log.Printf("Error of read file: %s", err)
		return
	}
	log.Printf("len: %v", len(postPhoto))

	file, err := os.Open(path)
	if err != nil {
		log.Printf("Error of read file: %s", err)
		return
	}
	fileName := filepath.Base(file.Name())
	fmt.Println("fileName")
	fmt.Println(fileName)

	for i := 0; i < 306; i++ {
		go func() {
			_, err = c.PostPhoto(ctx, &grpc2.PostPhotoRequest{
				Name:  fileName,
				Photo: postPhoto,
			})
			if err != nil {
				fmt.Printf("Error of post file: %s\n", err)
				return
			}
			fmt.Printf("ok\n")
		}()
	}
	time.Sleep(10 * time.Second)

	//files, err := c.GetAllPhotosInfo(ctx, &kitty.Empty{})
	//if err != nil {
	//	log.Printf("could not get all photos info: %v", err)
	//}
	//
	//fmt.Println("Имя файла | Дата создания | Дата обновления")
	//for _, fileInfo := range files.FileInfo {
	//	fmt.Printf("%s | %v | %v \n", fileInfo.Name, fileInfo.Created.AsTime().Format(time.RFC822), fileInfo.Edited.AsTime().Format(time.RFC822))
	//
	//}

	//files.FileInfo[0].Created.AsTime()
	//fmt.Println("Photo:")
	//fmt.Println(files)
	//fmt.Println(files.FileInfo[0].Created.AsTime())

	//photo, err := c.GetPhoto(ctx, &kitty.GetPhotoRequest{Id: uint64(7)})
	//if err != nil {
	//	log.Printf("could not get all photos: %v", err)
	//	return
	//}
	//fmt.Println(photo)

	//get list uploaded photo

}
