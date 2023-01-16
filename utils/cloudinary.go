package utils

import (
	"context"
	"fmt"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"mime/multipart"
	"strings"
)

type FileInfo struct {
	fileName, fileFolder string
}

func UploadImage(cld *cloudinary.Cloudinary, ctx context.Context, fileHeader *multipart.FileHeader, fileInfo FileInfo) {
	file, _ := fileHeader.Open()
	if "" == fileInfo.fileName {
		fileInfo.fileName = fileHeader.Filename
	}
	if "" == fileInfo.fileFolder {
		fileInfo.fileFolder = fileHeader.Filename[:strings.Index(fileHeader.Filename, ".")]
	}
	fmt.Println(fileInfo.fileFolder)
	fmt.Println(fileInfo.fileName)
	// Upload the image.
	// Set the asset's public ID and allow overwriting the asset with new versions
	resp, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{
		PublicID:       fileInfo.fileName,
		Transformation: "f_auto,q_auto",
		UniqueFilename: api.Bool(true),
		Folder:         "matrimonial/" + fileInfo.fileFolder,
	})
	if err != nil {
		fmt.Println("error")
	}

	// Log the delivery URL
	fmt.Printf("****2. Upload an image****Delivery URL: %+v\n", resp)
}
