package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Dates struct {
	Code int	`json:"company"`
	Cid int		`josn:"cid"`
	Userip string`josn:"userip"`
	Data Inner`josn:"data"`
}
type Inner struct {
	Expiration int`josn:"expiration"`
	Items []map[string]string`josn:"items"`
}

func download(d Dates){
	vk:=d.Data.Items[0]["vkey"]
	pa:=d.Data.Items[0]["songmid"]
	fmt.Println(vk,pa)
	url:="http://dl.stream.qqmusic.qq.com/C400params.m4a?vkey=vvv&guid=9082027038&uin=0&fromtag=66"
	url=strings.Replace(url,"vvv",vk,-1)
	url=strings.Replace(url,"params",pa,-1)
	resp,err:=http.Get(url)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(url)
	name:=pa+".mp3"
	defer resp.Body.Close()
	r,err:=ioutil.ReadAll(resp.Body)
	if err!=nil{
		fmt.Println(err)
	}
	ioutil.WriteFile(name,r,0740)
}

func getMusic(mid string){
	var url string;
	url="https://c.y.qq.com/base/fcgi-bin/fcg_music_express_mobile3.fcg?&jsonpCallback=MusicJsonCallback&cid=205361747&songmid=params&filename=C400params.m4a&guid=9082027038"
	url=strings.Replace(url,"params",mid,-1)
	client := &http.Client{
	}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("user-agent", `Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.26 Safari/537.36 Core/1.63.6726.400 QQBrowser/10.2.2265.400`)
	req.Header.Add("referer",`https://y.qq.com/portal/player.html`)
	req.Header.Add("content-type",`application/x-www-form-urlencoded; charset=UTF-8`)

	resp, _ := client.Do(req)
	defer resp.Body.Close()
	//body, _ := ioutil.ReadAll(resp.Body)
	//d:=Dates{}
	//err:=json.Unmarshal(body,&d)
	//if err!=nil{
	//	fmt.Println(err)
	//}
	//fmt.Println(string(body))
	d:=Dates{}
	err:=json.NewDecoder(resp.Body).Decode(&d)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Printf("%+v\n",d)
	download(d)
}

func ParseUrl(url string){
	var mid []string
	mid=strings.Split(url, "/")
	res:=mid[len(mid)-1]
	res=strings.SplitN(res,".",2)[0]
	fmt.Println(res)
	getMusic(res)
}

func main(){
	var url string;
	url="https://y.qq.com/n/yqq/song/0048OU664d4J3G.html"
	ParseUrl(url)
}
