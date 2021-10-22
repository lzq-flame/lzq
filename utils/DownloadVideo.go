/**
2 * @Author: lzq
3 * @Date: 2021/10/22 9:59
4 */

package utils

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const PATH = "D:\\video\\test\\first.mp4"

func DownloadFromUrl(url string) (string, error) {

	fmt.Println("Downloading", url, "to", PATH)

	// TODO: check file existence first with io.IsExist
	output, err := os.Create(PATH)
	if err != nil {
		log.Println("Error while creating", PATH, "-", err)
		return "", err
	}
	defer output.Close()

	response, err := http.Get(url)
	if err != nil {
		log.Println("Error while downloading", url, "-", err)
		return "", err
	}
	defer response.Body.Close()

	n, err := io.Copy(output, response.Body)
	if err != nil {
		log.Println("Error while downloading", url, "-", err)
		return "", err
	}

	log.Println(n, "bytes downloaded.")
	return PATH, err
}
