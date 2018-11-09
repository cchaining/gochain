/*
 * cmd 패키지 :
 *
 * cli 파일을 담고 있다.
 * 커맨드라인에서 명령어를 통해 블록체인 코어를 실행하고 조절하는 go 파일을 담고있다.
 * ethereum으로 비교하자면 console 패키지와 같다.
 */
package cmd

import (
	"flag"
	"fmt"
	"log"
	"os"
)

//__[필기]__
/*
 * field 가 앞글자 대문자 이면 public
 * 소문자이면 private 외부에서 호출할 수 없다
 * CLI responsible for processing command line arguments
 */
type CLI struct{}

func (cli *CLI) printUsage() {
	fmt.Println("사용법 : ")
	fmt.Println("  createblockchain -address ADDRESS            - 블록체인을 생성하고 제네시스 블록의 보상을 해당 주소로 보낸다.")
	fmt.Println("  createwallet                                 - 새로운 키를 생성하고 지갑 파일로 저장한다.")
	fmt.Println("  getbalance -address ADDRESS                  - 해당 주소의 밸런스를 확인한다.")
	fmt.Println("  listaddresses                                - 모든 지갑 파일의 주소목록을 보여준다.")
	fmt.Println("  printchain                                   - 블록체인의 모든 블록을 보여준다.")
	fmt.Println("  reindexutxo                                  - UTXO Set을 재설정 한다.")
	fmt.Println("  send -from FROM -to TO -amount AMOUNT -mine  - 일정량의 코인을 FROM 주소에서 TO 주소로 보낸다. Mine on the same node, when -mine is set.")
	fmt.Println("  startnode -miner ADDRESS                     - 시스템환경에 셋팅한 NODE_ID의 특정 노드는 마이닝을 시작한다.")
}

func (cli *CLI) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}

// Run parses command line arguments and processes commands
func (cli *CLI) Run() {
	cli.validateArgs()

	nodeID := os.Getenv("NODE_ID")
	log.Println("HTTP Server nodeID : ", nodeID, " 서버가 열렸습니다. ")
	if nodeID == "" {
		fmt.Printf("NODE_ID env. 가 설정되지 않았습니다.")
		os.Exit(1)
	}

	getBalanceCmd := flag.NewFlagSet("getbalance", flag.ExitOnError)
	createBlockchainCmd := flag.NewFlagSet("createblockchain", flag.ExitOnError)
	createWalletCmd := flag.NewFlagSet("createwallet", flag.ExitOnError)
	listAddressesCmd := flag.NewFlagSet("listaddresses", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)
	reindexUTXOCmd := flag.NewFlagSet("reindexutxo", flag.ExitOnError)
	sendCmd := flag.NewFlagSet("send", flag.ExitOnError)
	startNodeCmd := flag.NewFlagSet("startnode", flag.ExitOnError)

	getBalanceAddress := getBalanceCmd.String("address", "", "The address to get balance for")
	createBlockchainAddress := createBlockchainCmd.String("address", "", "The address to send genesis block reward to")
	sendFrom := sendCmd.String("from", "", "Source wallet address")
	sendTo := sendCmd.String("to", "", "Destination wallet address")
	sendAmount := sendCmd.Int("amount", 0, "Amount to send")
	sendMine := sendCmd.Bool("mine", false, "Mine immediately on the same node")
	startNodeMiner := startNodeCmd.String("miner", "", "Enable mining mode and send reward to ADDRESS")

	switch os.Args[1] {
	case "getbalance":
		err := getBalanceCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "createblockchain":
		err := createBlockchainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "createwallet":
		err := createWalletCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "listaddresses":
		err := listAddressesCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "printchain":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "reindexutxo":
		err := reindexUTXOCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "send":
		err := sendCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "startnode":
		err := startNodeCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		cli.printUsage()
		os.Exit(1)
	}

	if getBalanceCmd.Parsed() {
		if *getBalanceAddress == "" {
			getBalanceCmd.Usage()
			os.Exit(1)
		}
		cli.getBalance(*getBalanceAddress, nodeID)
	}

	if createBlockchainCmd.Parsed() {
		if *createBlockchainAddress == "" {
			createBlockchainCmd.Usage()
			os.Exit(1)
		}
		cli.createBlockchain(*createBlockchainAddress, nodeID)
	}

	if createWalletCmd.Parsed() {
		cli.createWallet(nodeID)
	}

	if listAddressesCmd.Parsed() {
		cli.listAddresses(nodeID)
	}

	if printChainCmd.Parsed() {
		cli.printChain(nodeID)
	}

	if reindexUTXOCmd.Parsed() {
		cli.reindexUTXO(nodeID)
	}

	if sendCmd.Parsed() {
		if *sendFrom == "" || *sendTo == "" || *sendAmount <= 0 {
			sendCmd.Usage()
			os.Exit(1)
		}

		cli.send(*sendFrom, *sendTo, *sendAmount, nodeID, *sendMine)
	}

	if startNodeCmd.Parsed() {
		nodeID := os.Getenv("NODE_ID")
		if nodeID == "" {
			startNodeCmd.Usage()
			os.Exit(1)
		}
		cli.startNode(nodeID, *startNodeMiner)
	}
}
