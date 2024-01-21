package upload

import (
	"context"
	"errors"
	"log"
	"mime/multipart"
	"path"
	"strings"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/minio/minio-go/pkg/credentials"
	"github.com/minio/minio-go/v7"
	"go.uber.org/zap"
)

type MinIO struct{}

//@author: [piexlmax](https://github.com/piexlmax)
//@author: [ccfish86](https://github.com/ccfish86)
//@author: [SliverHorn](https://github.com/SliverHorn)
//@object: *Local
//@function: UploadFile
//@description: 上传文件
//@param: file *multipart.FileHeader
//@return: string, string, error

func (*MinIO) UploadFile(file *multipart.FileHeader) (string, string, error) {
	// Create minio connection.
	minioClient, err := MinioConnection()
	if err != nil {
		global.GVA_LOG.Error("function MinioConnection() failed", zap.Any("err", err.Error()))
		return "", "", errors.New("function MinioConnection() failed, err:" + err.Error())
	}

	// 读取文件后缀
	ext := path.Ext(file.Filename)
	// 读取文件名并加密
	name := strings.TrimSuffix(file.Filename, ext)
	name = utils.MD5V([]byte(name))

	// Get Buffer from file
	buffer, err := file.Open()

	if err != nil {
		global.GVA_LOG.Error("function MinioConnection() failed", zap.Any("err", err.Error()))
		return "", "", errors.New("function MinioConnection() failed, err:" + err.Error())
	}
	defer buffer.Close()

	// objectName := file.Filename
	fileBuffer := buffer
	fileSize := file.Size
	// 拼接新文件名
	filename := name + "_" + time.Now().Format("20060102150405") + ext
	info, err := minioClient.PutObject(context.Background(), "gold", filename, fileBuffer, fileSize, minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		global.GVA_LOG.Error("function MinioConnection() failed", zap.Any("err", err.Error()))
		return "", "", errors.New("function MinioConnection() failed, err:" + err.Error())
	}
	return info.Location, info.Key, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@author: [ccfish86](https://github.com/ccfish86)
//@author: [SliverHorn](https://github.com/SliverHorn)
//@object: *Local
//@function: DeleteFile
//@description: 删除文件
//@param: key string
//@return: error

func (*MinIO) DeleteFile(key string) error {
	// Create minio connection.
	minioClient, err := MinioConnection()
	if err != nil {
		global.GVA_LOG.Error("function MinioConnection() failed", zap.Any("err", err.Error()))
		return errors.New("function MinioConnection() failed, err:" + err.Error())
	}
	bucketName := "gold"
	objectName := key
	err = minioClient.RemoveObject(context.Background(), bucketName, objectName, minio.RemoveObjectOptions{})
	if err != nil {
		return errors.New("function MinioConnection() failed, err:" + err.Error())
	}
	return nil
}

func MinioConnection() (*minio.Client, error) {
	ctx := context.Background()
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
