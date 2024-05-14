package main

import (
	"bufio"
	"log"
	"os"
	"simple_lsp/rpc"
)

func main() {
    // change this path to adjust log file localtion
    logger := getLogger("/Users/zhisu/Documents/resources/side_projects/simple_lsp/log.txt")
    logger.Println("LSP server started")

    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(rpc.Split)

    for scanner.Scan() {
        msg := scanner.Text()
        handleMessage(logger, msg)
    }
}

func handleMessage(logger *log.Logger, msg any) {
    logger.Println(msg)
}

func getLogger(filename string) *log.Logger {
    logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
    if err != nil {
	panic("Invalid path name")
    }

    return log.New(logfile, "[simple_lsp]", log.Ldate|log.Ltime|log.Lshortfile)
}
