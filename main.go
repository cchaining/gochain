package main

import (
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	//writing web handlers package
	"github.com/gorilla/mux"

	//"github.com/cchaining/gochain/gc_cmd"
	//"gc_cmd"
	//"core"
	"./core"
	"./cmd"
)


var mutex = &sync.Mutex{}

func main() {

	/*[고민 필요]*/
	//서버는 서버대로 돌리고 CLI는 CLI대로 돌리는
	//go루틴 체계를 생각을 해보는것도 좋을듯
	//Cli의 명령을 http 라우터와 연동시킬 필요가 있음
	/***************************************/
	bc := core.NewBlockchain()
	log.Println("get blocks success", bc)
	defer bc.Db.Close()

	cli := cmd.CLI{bc}
	cli.Run()
	/***************************************/

	//err := godotenv.Load()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//go func() {
	//	bc := gc_core.NewBlockchain()
	//	spew.Dump(bc)
	//
	//	mutex.Lock()
	//	mutex.Unlock()
	//}()
	//log.Fatal(run())

}
// web server
func run() error {
	mux := makeMuxRouter()
	httpPort := os.Getenv("PORT")
	log.Println("HTTP Server Listening on port :", httpPort)
	s := &http.Server{
		Addr:           ":" + httpPort,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

// create handlers
func makeMuxRouter() http.Handler {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/blocks", getBlocks).Methods("GET")
	muxRouter.HandleFunc("/block", getBlock).Methods("GET")
	muxRouter.HandleFunc("/tx", getTx).Methods("GET")
	muxRouter.HandleFunc("/tx/to", sendTx).Methods("POST")
	muxRouter.HandleFunc("/tx/pool", getTxPool).Methods("GET")
	muxRouter.HandleFunc("/utxo", getUtxo).Methods("GET")
	muxRouter.HandleFunc("/utxo/me", getMyUtxo).Methods("GET")
	muxRouter.HandleFunc("/utxo/address", getUtxoByAddress).Methods("GET")
	muxRouter.HandleFunc("/balance", getBalance).Methods("GET")
	muxRouter.HandleFunc("/address", getAddress).Methods("GET")
	muxRouter.HandleFunc("/p2p", getPeers).Methods("GET")
	muxRouter.HandleFunc("/p2p/address", addPeer).Methods("POST")

	return muxRouter
}

// [ 블록 ]
func getBlocks(w http.ResponseWriter, r *http.Request) {
	//bytes, err := json.MarshalIndent(Blockchain, "", "  ")
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}
	//log.Println("get blocks success")
	//io.WriteString(w, string(bytes))
}

func getBlock(w http.ResponseWriter, r *http.Request) {

}

// [ 트렌젝션 ]
func getTx(w http.ResponseWriter, r *http.Request) {

}

func sendTx(w http.ResponseWriter, r *http.Request) {

}

func getTxPool(w http.ResponseWriter, r *http.Request) {

}

func getUtxo(w http.ResponseWriter, r *http.Request) {

}

func getMyUtxo(w http.ResponseWriter, r *http.Request) {

}

func getUtxoByAddress(w http.ResponseWriter, r *http.Request) {

}

// [ 마이닝 ]
func mineRawBlock() {

}

func mineBlock() {

}

func mineTx() {

}


// [ 지갑 (주소, 잔여금) ]
func getBalance(w http.ResponseWriter, r *http.Request) {

}

func getAddress(w http.ResponseWriter, r *http.Request) {

}


// [ P2P ]
func getPeers(w http.ResponseWriter, r *http.Request) {

}

func addPeer(w http.ResponseWriter, r *http.Request) {

}

func stop() {

}
