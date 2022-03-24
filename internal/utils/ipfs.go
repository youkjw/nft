package utils

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	client "github.com/ipfs/go-ipfs-api"
	nftstorage "github.com/nftstorage/go-client"
	"os"
	"path"
)

const IpfsServer = "http://175.178.183.199:5001"

func Store(c *gin.Context) (fileHash string, err error) {
	// Where your local node is running on localhost:5001
	//sh := client.NewShell(IpfsServer)

	dir, _, _ := GetCurrentPath()
	filepath := fmt.Sprintf("%s%s", path.Dir(dir), "/../uploads/1.png")
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "file not found: %v\n", err)
		return
	}

	var ctx = c.Request.Context()
	auth := context.WithValue(ctx, nftstorage.ContextAccessToken, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJkaWQ6ZXRocjoweEY1MjY1MDNBMDEzQjU3MTQzRDJjOEVkMmY4ZjRjMmRCODljMDgyNDgiLCJpc3MiOiJuZnQtc3RvcmFnZSIsImlhdCI6MTY0NzU4NzM3ODQxNCwibmFtZSI6Im15bmZ0In0.ieVfXrwD6XQKlaOeszwrN42Mn-A4kpqOv_cUOhd-aZM")
	configuration := nftstorage.NewConfiguration()
	api_client := nftstorage.NewAPIClient(configuration)
	resp, r, err := api_client.NFTStorageAPI.Store(auth).Body(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NFTStorageAPI.Store``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Store`: UploadResponse
	fmt.Fprintf(os.Stdout, "Response from `NFTStorageAPI.Store`: %v\n", resp)

	//fileHash, err = sh.Add(file, client.Pin(true), client.)
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "error: %s", err)
	//	os.Exit(1)
	//}

	return fileHash, nil
}

func GenerateKey(filename string) string {
	return ""
}
