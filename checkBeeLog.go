package main

import (
	"fmt"
	"log"
	"strings"

	"golang.org/x/exp/slices"
)

func checkBeeLog(data string, resultMap *map[string]bool, logName string, keyFrases []string) {
	fmt.Println("Start CHECK " + logName + ".log")
	resultName := "check_" + logName + "_log"

	for len(data) > 0 {
		//fmt.Println(len(data))
		before, after, _ := strings.Cut(data, "\n")
		for i := 0; i < len(keyFrases); i++ {
			findFrase := strings.Index(before, keyFrases[i])
			if findFrase > -1 {
				fmt.Println("Found " + logName + ".log: " + keyFrases[i])
				log.Println("Found " + logName + ".log: " + keyFrases[i])
				//fmt.Println("i:", i)
				keyFrases = slices.Delete(keyFrases, i, 1)
				i = len(keyFrases)
				//fmt.Println(keyFrases)
			}
		}
		if len(keyFrases) > 0 {
			data = after
		} else {
			(*resultMap)[resultName] = true
			fmt.Println("CHECK " + logName + ".log finished with SUCCESS")
			data = ""
		}

	}

	if len(keyFrases) > 0 {
		(*resultMap)[resultName] = false
		fmt.Println("Check " + logName + ".log not found:")
		fmt.Println(keyFrases)
		log.Println("Check " + logName + ".log not found:")
		log.Println(keyFrases)
	}

}

/*
"msg"="bee version" "version"="1.17.4-038dbfb9"
"msg"="swarm public key"
"msg"="pss public key"
"msg"="using ethereum address"
"msg"="using overlay address"
"msg"="starting with a disabled chain backend"
"level"="info" "logger"="node" "msg"="using chain with network network" "chain_id"=100 "network_id"=1
"level"="info" "logger"="node" "msg"="starting debug server" "address"="127.0.0.1:1635"

"msg"="using datadir" "path"="/Users/zolmac/Library/Application Support/Swarm Desktop/data-dir"
"logger"="node/storer" "msg"="localstore sharky .DIRTY file exists: starting recovery due to previous dirty exit" ???
"logger"="node/storer" "msg"="localstore sharky recovery finished"
"logger"="node" "msg"="starting in ultra-light mode"
"logger"="node/multiresolver" "msg"="connected" "tld"="" "endpoint"="https://cloudflare-eth.com"
"level"="info" "logger"="node" "msg"="starting api server" "address"="127.0.0.1:1633"





*/
