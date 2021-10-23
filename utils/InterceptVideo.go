/**
2 * @Author: lzq
3 * @Date: 2021/10/22 10:12
4 */

package utils

import "os/exec"

const AFTER_PATH = "D:\\video\\test\\after\\first.mp4"

func Intercept(path string, starTime, endTime string) error {
	cmdLine := "ffmpeg -i " + path + " -ss " + starTime + " -to " + endTime + " -acodec copy " + AFTER_PATH
	cmd := exec.Command("cmd", "/c", cmdLine)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil

}
