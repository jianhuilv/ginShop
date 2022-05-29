package utils

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"mime/multipart"
	"os"
)

func handleErrorFake(err error) {
	fmt.Println("Error:", err)
	os.Exit(-1)
}

func GetImageFake(pid string) string {
	// Endpoint以杭州为例，其它Region请按实际情况填写。
	endpoint := "http://oss-cn-guangzhou.aliyuncs.com"
	// 阿里云账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM用户进行API访问或日常运维，请登录RAM控制台创建RAM用户。
	accessKeyId := "******"
	accessKeySecret := "******"
	bucketName := "******"
	//objectName := "test2"
	// <yourLocalFileName>由本地文件路径加文件名包括后缀组成，例如/users/local/myfile.txt。
	//localFileName := "C:\\Users\\69060\\Desktop\\image\\发个帖子而已\\34b3497ff21fbc6580d3e977882059e36f2742d8.png"
	// 创建OSSClient实例。
	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		handleError(err)
	}
	// 获取存储空间。
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		handleError(err)
	}

	// 上传文件。
	// 指定图片名称。如果图片不在Bucket根目录，需携带文件完整路径，例如exampledir/example.jpg。
	ossImageName := pid
	// 生成带签名的URL，并指定过期时间为600s。
	signedURL, err := bucket.SignURL(ossImageName, oss.HTTPGet, 600, oss.Process("image/format,png"))
	if err != nil {
		handleError(err)
	} else {
		fmt.Println(signedURL)
	}
	return signedURL
}

func UploadImageFake(pid string, fileHeader *multipart.FileHeader) error {
	// Endpoint以杭州为例，其它Region请按实际情况填写。
	endpoint := "http://oss-cn-guangzhou.aliyuncs.com"
	// 阿里云账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM用户进行API访问或日常运维，请登录RAM控制台创建RAM用户。
	accessKeyId := "*********"
	accessKeySecret := "*********"
	bucketName := "lvjianhui"
	//objectName := "test2"
	// <yourLocalFileName>由本地文件路径加文件名包括后缀组成，例如/users/local/myfile.txt。
	//localFileName := "C:\\Users\\69060\\Desktop\\image\\发个帖子而已\\34b3497ff21fbc6580d3e977882059e36f2742d8.png"
	// 创建OSSClient实例。
	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		handleError(err)
	}
	// 获取存储空间。
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		handleError(err)
	}
	fd, err := fileHeader.Open()
	// 上传文件。
	// 指定图片名称。如果图片不在Bucket根目录，需携带文件完整路径，例如exampledir/example.jpg。
	err = bucket.PutObject(pid, fd)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	return err
}
