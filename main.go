package main

import (
	"encoding/gob"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reversigo/model"
	"strconv"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	// key   = []byte("super-secret-key")
	// store = sessions.NewCookieStore(key)
	store       *sessions.FilesystemStore = sessions.NewFilesystemStore("", securecookie.GenerateRandomKey(64))
	sessionName                           = "cookie-name"
)

func main() {
	// ディレクトリを指定する
	fs := http.FileServer(http.Dir("static"))
	// ルーティング設定。"/"というアクセスがきたらstaticディレクトリのコンテンツを表示させる
	http.Handle("/", fs)
	// アクセスのルーティングを設定します。
	http.HandleFunc("/FrontController", frontController)
	log.Println("Listening...")
	// 80ポートでサーバーを立ち上げる
	http.ListenAndServe(":80", nil)
}

func frontController(w http.ResponseWriter, r *http.Request) {
	store.MaxLength(4 * 1024 * 1024)
	gob.Register(model.ReversiPlay{})
	gob.Register(model.ReversiSetting{})
	session, err := store.Get(r, sessionName)
	if err != nil {
		// 不正なセッションだった場合は作り直す
		session, err = store.New(r, sessionName)
	}
	rpi, ok := session.Values["reversiplay"].(model.ReversiPlay)
	fmt.Println(ok)
	if ok == false {
		rpi = *model.NewReversiPlay()
		session.Values["reversiplay"] = rpi
		err = session.Save(r, w)
		if err != nil {
			panic(err)
		}
	}
	rsi, ok2 := session.Values["reversisetting"].(model.ReversiSetting)
	fmt.Println(ok2)
	if ok2 == false {
		rsi = *model.NewReversiSetting()
		session.Values["reversisetting"] = rsi
		err = session.Save(r, w)
		if err != nil {
			panic(err)
		}
	}
	rp := rpi
	rp.GetmSetting().Copy(&rsi)
	rp.SetmCallbacks(model.NewCallbacksJson())
	rp.SetmDelegate(model.NewReversiPlayDelegate(model.NewReversiPlayInterfaceImpl()))

	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Printf("%p ", &rsi)
	fmt.Println(rsi)
	function := r.FormValue("func")
	res := model.NewResJson()
	if function == "setSetting" {
		s := model.NewReversiSetting()
		jsonText := r.FormValue("para")
		if err := json.Unmarshal([]byte(jsonText), &s); err != nil {
			log.Fatal(err)
		}
		rsi.Copy(s)
		session.Values["reversisetting"] = rsi
		rp.GetmSetting().Copy(s)
		rp.Reset()
		res.SetAuth("[SUCCESS]")
	} else if function == "reset" {
		rp.Reset()
		res.SetAuth("[SUCCESS]")
	} else if function == "reversiPlay" {
		y, _ := strconv.Atoi(r.FormValue("y"))
		x, _ := strconv.Atoi(r.FormValue("x"))
		rp.ReversiPlay(y, x)
		res.SetAuth("[SUCCESS]")
	}
	fmt.Printf("%p ", &rsi)
	fmt.Println(rsi)
	session.Values["reversiplay"] = rpi
	err = session.Save(r, w)
	if err != nil {
		panic(err)
	}

	rsi2, ok3 := session.Values["reversisetting"].(model.ReversiSetting)
	fmt.Println(ok3)
	if ok3 == false {
	} else {
		fmt.Printf("%p ", &rsi2)
		fmt.Println(rsi2)
	}

	// jsonヘッダーを出力
	w.Header().Set("Content-Type", "application/json")
	// jsonエンコード
	res.SetCallbacks(rp.GetmCallbacks())
	outputJson, err := json.Marshal(res)
	if err != nil {
		panic(err)
	}
	// jsonデータを出力
	w.Write(outputJson)
}
