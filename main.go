package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reversigo/model"
	"strconv"

	"github.com/gorilla/sessions"
)

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

func main() {
	//ディレクトリを指定する
	fs := http.FileServer(http.Dir("static"))
	//ルーティング設定。"/"というアクセスがきたらstaticディレクトリのコンテンツを表示させる
	http.Handle("/", fs)
	http.HandleFunc("/FrontController", frontController) //アクセスのルーティングを設定します。

	log.Println("Listening...")
	// 80ポートでサーバーを立ち上げる
	http.ListenAndServe(":80", nil)
}

func frontController(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")
	fmt.Println(session)
	rpi, ok := session.Values["reversiplay"].(*model.ReversiPlay)
	fmt.Println(ok)
	if rpi == nil {
		rpi = model.NewReversiPlay()
		session.Values["reversiplay"] = *rpi
		session.Save(r, w)
	}
	// rp := rpi.(*model.ReversiPlay)
	rp := rpi
	rp.SetmCallbacks(model.NewCallbacksJson())
	rp.SetmDelegate(model.NewReversiPlayDelegate(model.NewReversiPlayInterfaceImpl()))

	r.ParseForm()
	fmt.Println(r.Form)
	function := r.FormValue("func")
	res := model.NewResJson()
	if function == "setSetting" {
		s := model.NewReversiSetting()
		jsonText := r.FormValue("para")
		if err := json.Unmarshal([]byte(jsonText), &s); err != nil {
			log.Fatal(err)
		}
		// デコードしたデータを表示
		fmt.Println(s)
		rp.SetmSetting(s)
		rp.Reset()
		res.SetAuth("[SUCCESS]")
	} else if function == "reset" {
		rp.Reset()
		res.SetAuth("[SUCCESS]")
	} else if function == "reversiPlay" {
		y, _ := strconv.Atoi(r.FormValue("y"))
		x, _ := strconv.Atoi(r.FormValue("x"))
		rp.ReversiPlay(x, y)
		res.SetAuth("[SUCCESS]")
	}
	session.Save(r, w)
	// jsonヘッダーを出力
	w.Header().Set("Content-Type", "application/json")

	// jsonエンコード
	res.SetCallbacks(rp.GetmCallbacks())
	outputJson, err := json.Marshal(*res)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(outputJson))
	// jsonデータを出力
	w.Write(outputJson)
	// fmt.Fprint(w, string(outputJson))
}
