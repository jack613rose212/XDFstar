package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	//x := make(map[int]int)
	//for i := 0; i < 30; i++ {
	//	x[i+100] = i
	//}
	//map并发不安全  range出来的数据是无序的
	//for k:= range x {
	//fmt.Println(k,v)
	//}
	//	for k,v:= range x {
	//		fmt.Println(k,v)
	//	}

	//for v:= range x {  //得到的是key  通过key得到values
	//	fmt.Println(v)
	//}
	//fmt.Println(x[119]) //119是key  通过了key得到的values

	//var Url =`https://chaoshi.tmall.com/?targetPage=index&ali_trackid=2%3Amm_123270724_271200191_109813450244%3A1575946902_241_713103899`
	////创建请求
	//req,err:=http.NewRequest("get",Url,strings.NewReader(body))
	//
	////设置请求头
	//req.Header
	//
	////发送请求
	//resp,err:=http.DefaultClient.Do(req)
	//r.Resp.Body
	//

	//var Url =`https://chaoshi.tmall.com/?targetPage=index&ali_trackid=2%3Amm_123270724_271200191_109813450244%3A1575946902_241_713103899`
	//怎么带入 怎么处理初始化
	url1 := `https://chaoshi.tmall.com`
	//queryValues = `targetPage=index&ali_trackid=2%3Amm_123270724_271200191_109813450244%3A1575946902_241_713103899`
	queryValues := url.Values{ //相当于？后面的参数
		"targetPage":  {"index"}, //kv结构  {"参数"}  切片的初始化{切片包裹对象}
		"ali_trackid": {"2%3Amm_123270724_271200191_109813450244%3A1575946902_241_713103899"},
	}
	//
	for i := 1; i <= 10000; i++ {
		_, err := DownloadString(url1, queryValues)
		if err != nil {
			fmt.Println("错误:", err.Error())
			return
		}

		//var sss  string
		//json.Unmarshal(body,sss)
		//fmt.Println(sss,i)
		//fmt.Println(i,string(body))
		fmt.Println("success!!", i)
	}

	//_,err:=DownloadString(url1, queryValues)
	//if err != nil {
	//	fmt.Println("错误:",err.Error())
	//	return
	//}
	//
	fmt.Println("success OK!")
	//var sss  string
	//json.Marshal(body)
	//fmt.Println(sss)
}

//GET http://127.0.0.1:8080/ping?param1=value1&param2=123 HTTP/1.1
func DownloadString(remoteUrl string, queryValues url.Values) (body []byte, err error) {
	client := &http.Client{};
	body = nil;
	uri, err := url.Parse(remoteUrl);
	if (err != nil) {
		return;
	}
	if (queryValues != nil) {
		values := uri.Query();
		if (values != nil) {
			for k, v := range values {
				queryValues[k] = v;
			}
		}
		uri.RawQuery = queryValues.Encode(); //?号后
	}
	reqest, err := http.NewRequest("GET", uri.String(), nil);
	reqest.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8");
	reqest.Header.Add("Accept-Encoding", "gzip, deflate");
	reqest.Header.Add("Accept-Language", "zh-cn,zh;q=0.8,en-us;q=0.5,en;q=0.3");
	reqest.Header.Add("Connection", "keep-alive");
	reqest.Header.Add("Host", uri.Host);
	reqest.Header.Add("Referer", uri.String());
	reqest.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:12.0) Gecko/20100101 Firefox/12.0");
	response, err := client.Do(reqest)
	defer response.Body.Close();
	if (err != nil) {
		return;
	}

	if response.StatusCode == 200 {
		fmt.Println("sucess!!200")
		switch response.Header.Get("Content-Encoding") {
		case "gzip":
			reader, _ := gzip.NewReader(response.Body)
			for {
				buf := make([]byte, 1024)
				n, err := reader.Read(buf)

				if err != nil && err != io.EOF {
					panic(err)
				}

				if n == 0 {
					break
				}
				body = append(body, buf...);
			}
		default:
			body, _ = ioutil.ReadAll(response.Body)

		}
	}
	return;
}
