package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	cclient "github.com/IHaveNothingg/cclientwtf"
	tls "github.com/refraction-networking/utls"
)

//Creating JSON
type Location struct {
	Postalcode   string `json:"postalCode"`
	City         string `json:"city"`
	State        string `json:"state"`
	Isziplocated bool   `json:"isZipLocated"`
}
type Payload struct {
	Offerid               string   `json:"offerId"`
	Quantity              int      `json:"quantity"`
	Location              Location `json:"location"`
	Shipmethoddefaultrule string   `json:"shipMethodDefaultRule"`
	Storeids              []int    `json:"storeIds"`
}

func main() {
	//Creating a client using my custom ClientHello
	client, err := cclient.NewClient(tls.HelloChrome_Auto)
	if err != nil {
		log.Fatal(err)
	}
	//Making the JSON
	data := Payload{
		"00572D4606B5433A9C7F26C494B1E8A7",
		1,
		Location{"", "Ridgecrest", "CA", true},
		"SHIP_RULE_1",
		[]int{1600},
	}
	//Turning it into JSON
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	bodySend := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "https://www.walmart.com/api/v3/cart/:CRT/items", bodySend)
	if err != nil {
		log.Fatal(err)
	}

	//I added header-order so this looks confusing but the Header-Order doesn't get sent it only is used to order them
	type headersmap map[string][]string
	newHeaders := http.Header{}
	Headers := headersmap{
		"Upgrade-Insecure-Requests": {"1"},
		"Host":                      {"www.walmart.com"},
		"sec-ch-ua-mobile":          {"?0"},
		"Origin":                    {"https://www.walmart.com"},
		"Sec-Fetch-Site":            {"same-origin"},
		"Sec-Fetch-Mode":            {"cors"},
		"Sec-Fetch-Dest":            {"empty"},
		"accept":                    {"application/json"},
		"Referer":                   {"https://www.walmart.com/ip/Charlotte-Pipe-Schedule-40-PVC-Solid-Pipe-3-in-Dia-2-ft-Plain-End-260-psi/149527172"},
		"Connection":                {"keep-alive"},
		"User-Agent":                {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.114 Safari/537.36"},
		"content-type":              {"application/json"},
		"Content-Length":            {fmt.Sprint(bodySend.Size())},
		"Accept-Encoding":           {"deflate, br"},
		"Accept-Language":           {"en-US,en;q=0.9"},
		"Cookie":                    {`DL=93555%2C%2C%2Cip%2C93555%2C%2C; TB_Latency_Tracker_100=1; TB_Navigation_Preload_01=1; vtc=dGRTVr3MD40EEB_Sx8vrNY; TBV=7; _pxvid=68db174b-9783-11eb-9f5a-0242ac120006; _gcl_au=1.1.358264655.1617787622; viq=Walmart; _fbp=fb.1.1617787622955.1847770411; tb_sw_supported=true; TB_SFOU-100=1; cbp=149527172-1617788255706; athrvi=RVI~h8e99a84; TS013ed49a=01538efd7c3928fbbafcf3a394d8fe8f3ed3bcb170db0add995709f73128c29eebdaf52553a8abd264d663718e092ac5ef39a63a89; ACID=f339c440-263d-44dc-ab4e-cae23c39e2a2; hasACID=1; CRT=69155c1f-bc75-45fb-92f6-9aa5f23c3195; hasCRT=1; auth=MTAyOTYyMDE4PbS9aNBy5CEcPtr9Vu2bkly4khdgqIgN7OK%2F80ffaa5t3ve271roxGI9QrGLk68Qv%2Bkj5G%2BZafOVl11ZmazaMHvE5WPVxaRfB%2FkmJ%2FY2gVrW73PeijviHYLr4LMsvN9N767wuZloTfhm7Wk2KcjyglM949MaUzwsNnQKx2EXSLmSE7JPshSufjQAS9KqQ3UO50th%2B%2FdoqROpHahhs62qP662R4I83zDF8cM4PhieW4EUMk70P8glgOEpLOprhDfMywI05adPtwc9%2Fm5r1ONHR0VGeoSa7Vi0msmbjLkwibZosDwIjxTHcvMkDyHQRn2iXiHnA8R0xkXPF2A1qVbCZ5rEXH4RA%2BJj%2Bj3r2K%2FGywYkvFWGSe8iGChOfXukMXEolg71858uOlU1K7Mm9bkIfg%3D%3D; type=GUEST; cart-item-count=1; next-day=1617924600|true|false|1617969600|1617862937; location-data=93555%3ARidgecrest%3ACA%3A%3A1%3A1|18g%3B%3B1.53||7|1|; TB_DC_Flap_Test=0; tb-c30=scus-t1; bstc=XFgK1_9aen-Yu4eqNv6LTM; mobileweb=0; xpa=OPRTJ|hWKTc|s-848|vZriO; exp-ck=OPRTJ1hWKTc1s-8481vZriO2; TS01b0be75=01538efd7c775e28060ee605a8687aa58b88012da6bfd2bf4faf28cbca47be977a3f3770d9f689b4ae4264879aa407a4f8f0680009; akavpau_p1=1617863537~id=ad11014571d6182fa0bdffbdc77b90b2; xpm=1%2B1617862937%2BdGRTVr3MD40EEB_Sx8vrNY~%2B0; com.wm.reflector="reflectorid:0000000000000000000000@lastupd:1617862950084@firstcreate:1617787620449"; akavpau_p8=1617863550~id=1ebd5c95398d7115f2d0396eb1c9a68e; s_sess_2=c32_v%3DS2H%2Cnull%3B%20prop32%3DS2H-V%2CS2H; _px3=0e7bc5e89367add3adca2bd001e07702654c12402be398cc7a1cdbb6df1af615:xaa7b2SvbaeFVEkHPrSo40M4gjoSXu7HxuHVT/27Q9ZEYkfMZd6Jj/bbgA6LUR1IFQugs1Efz7VpwGhg7g0TKQ==:1000:qs51jdlNrrRVGu6mRDLihiXOQJRz3iQZOCmTKIaxUZjSaOR5cPA/hR3IFHOgg84hdtNGPq83MCtZVHyAtDhvwx/EtTnig5I1OA09XxzuLkACl19jzNgTSBZTAtDGo96MHX2h0xNyuKd4thnXTT5N/66jvKi0vLvjP0J20Ps7VnM=; _uetsid=69234de0978311eb814d2d0558565e41; _uetvid=692394a0978311ebbe8f118ab8730f2d; _pxde=ca95ac842d1c641fdfec509ff7af240dfacfa8ff9552f3f632ef2581ed1de5dc:eyJ0aW1lc3RhbXAiOjE2MTc4NjMwMTg2NjYsImZfa2IiOjAsImlwY19pZCI6W119`},

		"Header-Order:": {"Host", "Connection", "Content-Length", "accept", "User-Agent", "content-type", "Origin", "sec-ch-ua", "Sec-Fetch-Site", "Sec-Fetch-Mode", "Sec-Fetch-Dest", "Referer", "Accept-Encoding", "Accept-Language", "Cookie"},
	}
	//Adding all the headers to a *http.Header map
	for i, x := range Headers {
		newHeaders[i] = x
	}
	//Setting the headers equal to that map
	req.Header = newHeaders
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)

	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)

	}
	fmt.Println(string(body))
}
