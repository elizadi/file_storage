package main

import (
	"context"
	grpc2 "file_storage/internal/handler/grpc"
	"flag"
	"fmt"
	logger "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:8000", "the address to connect to")
)

var rootCmd = &cobra.Command{
	Use:   "file_storage",
	Short: "A brief description of your CLI application",
	Long: `A longer description that explains your CLI application in detail, 
    including available commands and their usage.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to file_storage! Use --help for usage.")
	},
}

var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "upload file",
	Long:  `Upload file (photo) to the server`,
	Run: func(cmd *cobra.Command, args []string) {
		// Set up a connection to the server.
		conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			logger.WithError(err).Errorln("did not connect")
			return
		}
		defer conn.Close()
		c := grpc2.NewFileStorageClient(conn)

		// Contact the server and print out its response.
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// Save photo
		path := args[0]
		postPhoto, err := os.ReadFile(path)
		if err != nil {
			logger.WithError(err).Errorln("Error of read file")
			return
		}

		file, err := os.Open(path)
		if err != nil {
			logger.WithError(err).Errorln("Error of open file")
			return
		}
		fileName := filepath.Base(file.Name())

		_, err = c.PostPhoto(ctx, &grpc2.PostPhotoRequest{
			Name:  fileName,
			Photo: postPhoto,
		})
		if err != nil {
			logger.WithError(err).Errorln("Error of post file")
			return
		}
	},
}

var getAllInfoCmd = &cobra.Command{
	Use:   "getAllInfo",
	Short: "Get info about all files",
	Long:  `Get info (such as id, name, date of created and date of edited) about all saved files`,
	Run: func(cmd *cobra.Command, args []string) {
		// Set up a connection to the server.
		conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			logger.WithError(err).Errorln("did not connect")
			return
		}
		defer conn.Close()
		c := grpc2.NewFileStorageClient(conn)

		// Contact the server and print out its response.
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// Get info about all saved photos
		files, err := c.GetAllPhotosInfo(ctx, &grpc2.Empty{})
		if err != nil {
			logger.WithError(err).Errorln("could not get all photos info")
		}

		fmt.Println("ID | Имя файла | Дата создания | Дата обновления")
		for _, fileInfo := range files.FileInfo {
			fmt.Printf("%d | %s | %v | %v \n", fileInfo.Id, fileInfo.Name, fileInfo.Created.AsTime().Format(time.RFC822), fileInfo.Edited.AsTime().Format(time.RFC822))
		}
	},
}

var getPhotoCmd = &cobra.Command{
	Use:   "getPhoto",
	Short: "Get photo by id",
	Long: `Get photo saved earlier by it's id. First arguments id, second argument path where to save the file.
Example: getPhoto 1 "same/path/to/save/file"`,
	Run: func(cmd *cobra.Command, args []string) {
		// Set up a connection to the server.
		conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			logger.WithError(err).Errorln("did not connect")
			return
		}
		defer conn.Close()
		c := grpc2.NewFileStorageClient(conn)

		// Contact the server and print out its response.
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// Get photo by ID
		id, err := strconv.ParseUint(args[0], 0, 64)
		photo, err := c.GetPhoto(ctx, &grpc2.GetPhotoRequest{Id: id})
		if err != nil {
			logger.WithError(err).Errorln("could not get photo")
			return
		}

		path := filepath.Join(args[1], photo.Name)
		err = os.WriteFile(path, photo.Photo, 0644)
		if err != nil {
			logger.WithError(err).Errorln("Error save file")
			return
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func Init() {
	rootCmd.AddCommand(uploadCmd)
	rootCmd.AddCommand(getAllInfoCmd)
	rootCmd.AddCommand(getPhotoCmd)
}

func main() {
	Init()
	Execute()
}
