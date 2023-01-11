package latestblock

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/ethereum/go-ethereum/ethclient"
)


func GetLatestBlockNumber(rpcUrl string) (string, error) {

	log.Println("Getting latest block number")

	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	block_number := header.Number.String()

	log.Printf("Latest block number is: %s", block_number)

	return block_number, nil
}

type WrapperStruct struct {
    RpcUrl string
}

func (ws WrapperStruct) GetBlockHandler(w http.ResponseWriter, req *http.Request) {

	block, err := GetLatestBlockNumber(ws.RpcUrl)

	if err != nil {
		panic(err)
	}

	io.WriteString(w, "Latest block is: " + block)
}
