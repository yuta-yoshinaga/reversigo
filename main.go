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
	gob.Register(model.ReversiPlayInterfaceImpl{})
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
	rpi.SetmCallbacks(model.NewCallbacksJson())
	rpi.SetmDelegate(model.NewReversiPlayDelegate(model.NewReversiPlayInterfaceImpl()))

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
		rpi.GetmSetting().Copy(s)
		rpi.Reset()
		res.SetAuth("[SUCCESS]")
	} else if function == "reset" {
		rpi.Reset()
		res.SetAuth("[SUCCESS]")
	} else if function == "reversiPlay" {
		y, _ := strconv.Atoi(r.FormValue("y"))
		x, _ := strconv.Atoi(r.FormValue("x"))
		rpi.ReversiPlay(y, x)
		res.SetAuth("[SUCCESS]")
	}
	err = session.Save(r, w)
	if err != nil {
		panic(err)
	}
	// jsonヘッダーを出力
	w.Header().Set("Content-Type", "application/json")
	// jsonエンコード
	res.SetCallbacks(rpi.GetmCallbacks())
	outputJSONgo, err := json.Marshal(res)
	if err != nil {
		panic(err)
	}
	// jsonデータを出力
	w.Write(outputJSONgo)
}
