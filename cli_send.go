package blockchain6

import (
	"fmt"
	"log"
)

//send 转账
func (cli *CLI) send(from, to string, amount int) {
	if !ValidateAddress(from) {
		log.Panic("ERROR: 发送地址非法")
	}
	if !ValidateAddress(to) {
		log.Panic("ERROR: 接收地址非法")
	}

	bc := NewBlockchain() //打开数据库，读取区块链并构建区块链实例
	defer bc.Db.Close()   //转账完毕，关闭数据库

	UTXOSet := UTXOSet{bc}
	tx := NewUTXOTransaction(from, to, amount, &UTXOSet) //当前交易
	cbTx := NewCoinbaseTX(from, "")                      //挖矿奖励：挖矿由from发起，奖励给到from
	txs := []*Transaction{cbTx, tx}                      //将两个交易打包到一起

	newBlock := bc.MineBlock(txs) //挖出包含交易的区块，上链（写入区块链数据库）

	UTXOSet.Update(newBlock)

	fmt.Println("转账成功！")
}
