package main

import (
	"log"
	"net/http"
	"os"
	"time"

	//writing web handlers package
	"github.com/gorilla/mux"

	"./cmd"
)

func main() {

	/*[고민 필요]*/
	//서버는 서버대로 돌리고 CLI는 CLI대로 돌리는
	//go루틴 체계를 생각을 해보는것도 좋을듯
	//Cli의 명령을 http 라우터와 연동시킬 필요가 있음
	/***************************************/
	//bc := core.NewBlockchain("")
	//log.Println("get blocks success", bc)
	//defer bc.Db.Close()

	//httpPort := os.Getenv("PORT")
	//log.Println("HTTP Server Listening on port :", httpPort)
	//
	//nodeID := os.Getenv("NODE_ID")
	//log.Println("HTTP Server Listening on nodeID :", nodeID)

	cli := cmd.CLI{}
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
	//muxRouter.HandleFunc("/blocks", getBlocks).Methods("GET")
	//muxRouter.HandleFunc("/block", getBlock).Methods("GET")
	//muxRouter.HandleFunc("/tx", getTx).Methods("GET")
	//muxRouter.HandleFunc("/tx/to", sendTx).Methods("POST")
	//muxRouter.HandleFunc("/tx/pool", getTxPool).Methods("GET")
	//muxRouter.HandleFunc("/utxo", getUtxo).Methods("GET")
	//muxRouter.HandleFunc("/utxo/me", getMyUtxo).Methods("GET")
	//muxRouter.HandleFunc("/utxo/address", getUtxoByAddress).Methods("GET")
	//muxRouter.HandleFunc("/balance", getBalance).Methods("GET")
	//muxRouter.HandleFunc("/address", getAddress).Methods("GET")
	//muxRouter.HandleFunc("/p2p", getPeers).Methods("GET")
	//muxRouter.HandleFunc("/p2p/address", addPeer).Methods("POST")

	return muxRouter
}
