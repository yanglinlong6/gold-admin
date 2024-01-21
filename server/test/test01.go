package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func MinioConnection() (*minio.Client, error) {
	ctx := context.Background()
	// endpoint := "minio.linux008.com"
	endpoint := "149.129.78.23:9000"
	accessKeyID := "gold"
	secretAccessKey := "r4285613u8t#2!"
	useSSL := false
	// Initialize minio client object.
	minioClient, errInit := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if errInit != nil {
		log.Fatalln(errInit)
	}

	// Make a new bucket called dev-minio.
	bucketName := "gold"
	location := "gold"

	err := minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("Successfully created %s\n", bucketName)
	}
	return minioClient, errInit
}

func main() {
	// 创建 MinIO 客户端对象
	// minioClient, err := minio.New("149.129.78.23:9000", &minio.Options{
	// Create minio connection.
	minioClient, err := MinioConnection()
	if err != nil {
		fmt.Println(err)
		return
	}

	// 上传文件到 MinIO
	objectName := "2023-11-03 21.41.28.jpg"
	filePath := "./2023-11-03 21.41.28.jpg"
	contentType := "application/octet-stream"
	data, err := ioutil.ReadFile(filePath)
	println(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = minioClient.FPutObject(context.Background(), "gold", objectName,
		filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("File uploaded successfully")
}
