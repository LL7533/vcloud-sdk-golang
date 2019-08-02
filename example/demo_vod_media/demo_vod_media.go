package main

import (
	"encoding/json"
	"fmt"

	"github.com/TTvcloud/vcloud-sdk-golang/base"
	"github.com/TTvcloud/vcloud-sdk-golang/service/vod"
)

func searchVideo() {

}

func deleteVideo() {

}

func modifyVideoInfo() {

}

func setVideoPlayStatus() {
	// call below method if you dont set ak and sk in ～/.vcloud/config
	vod.DefaultInstance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	// or set ak and ak as follow
	//vod.DefaultInstance.SetAccessKey("")
	//vod.DefaultInstance.SetSecretKey("")

	resp, code, _ := vod.DefaultInstance.SetVideoPublishStatus("space", "vidxxxxx", "Published")
	fmt.Println(code)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
}

func DescribeVideoInfos() {

}

func main() {
	setVideoPlayStatus()
}
