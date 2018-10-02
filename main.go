package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	// 표준 패키지 이외의 패키지를 가져올 경우 절대 경로로
	//structs and slices cleanly formatted package in our console
	"gochain/pkg/github.com/davecgh/go-spew/spew"
	//writing web handlers package
	"gochain/pkg/github.com/gorilla/mux"
	//lets us read from a .env file package
	"gochain/pkg/github.com/joho/godotenv"
)

// Block represents each 'item' in the gochain
type Block struct {
	Index     int
	Timestamp string
	BPM       int
	Hash      string
	PrevHash  string
}

// Blockchain is a series of validated Blocks
var Blockchain []Block

// Message takes incoming JSON payload for writing heart rate
type Message struct {
	BPM int
}

var mutex = &sync.Mutex{}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		t := time.Now()
		genesisBlock := Block{}
		genesisBlock = Block{0, t.String(), 0, calculateHash(genesisBlock), ""}
		spew.Dump(genesisBlock)

		mutex.Lock()
		Blockchain = append(Blockchain, genesisBlock)
		mutex.Unlock()
	}()
	log.Fatal(run())

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

	return muxRouter
}

// [ 블록 ]
//blocks
func getBlocks(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.MarshalIndent(Blockchain, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("get blocks success")
	io.WriteString(w, string(bytes))
}

//block/:hash

// [ 트렌젝션 ]
//transaction/:id

//sendTransaction

//transactionPool

//unspentTransactionOutputs

//myUnspentTransactionOutputs


// [ 마이닝 ]
//mineRawBlock

//mineBlock

//mineTransaction


// [ 지갑 (주소, 잔여금) ]
//balance

//address/:address

//address


// [ P2P ]
//peers

//addPeer

// SHA256 hasing
func calculateHash(block Block) string {
	record := strconv.Itoa(block.Index) + block.Timestamp + strconv.Itoa(block.BPM) + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}
