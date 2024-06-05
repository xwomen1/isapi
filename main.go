package main

import (
	"fmt"
	"isapi/hikivision"
	"os"
	"os/exec"
)

func main() {
	// Thông tin kết nối
	URL := "192.168.100.200"
	userName := "admin"
	passWord := "123456aA"

	dataHikvision, err := hikivision.ConnectToHikvisionDevice(URL, userName, passWord)
	if err != nil {
		fmt.Println("Error logging into Hikvision:", err)
		return
	}

	startTime := "2024-06-05T00:00:00Z"
	endTime := "2024-06-05T23:59:59Z"
	playbackURI := fmt.Sprintf("rtsp://192.168.100.200/Streaming/tracks/701/?starttime=%s&endtime=%s&name=00010000832000100&size=6194940", startTime, endTime)

	dataReq := hikivision.DownloadRequest{
		PlaybackURI: playbackURI,
	}

	_, _, err = dataHikvision.PostDownloadVideo(&dataReq)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Video downloaded successfully.")

	inputFile := "videos/response.mp4"
	segmentTime := 60.0 // mỗi đoạn video dài 60 giây
	outputPattern := "output%03d.ts"
	err = segmentVideo(inputFile, segmentTime, outputPattern)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Video segmented successfully")
		currentDir, err := os.Getwd()
		if err != nil {
			fmt.Println("Error getting current directory:", err)
			return
		}
		fmt.Println("Current working directory:", currentDir)
	}

	fmt.Println("Video đã được cắt thành các đoạn có độ dài 60 giây.")
}

// Ham cat video thanh cac file.ts
func segmentVideo(inputFile string, segmentTime float64, outputPattern string) error {
	segmentTimeStr := fmt.Sprintf("%.2f", segmentTime)

	cmd := exec.Command("ffmpeg", "-i", inputFile, "-c", "copy", "-f", "segment", "-segment_time", segmentTimeStr, outputPattern)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error segmenting video: %w", err)
	}

	return nil
}
