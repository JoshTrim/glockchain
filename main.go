package main

import (
	"github.com/JoshTrim/glockchain/internal/blockchain"
)

func main() {
	bc := blockchain.NewBlockchain()
	defer bc.DB.Close()

	cli := blockchain.CLI{bc}
	cli.Run()

}
