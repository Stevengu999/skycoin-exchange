package main

import (
	"flag"
	"os"
	"path/filepath"

	"github.com/golang/glog"
	"github.com/skycoin/skycoin-exchange/src/server"
	"github.com/skycoin/skycoin/src/cipher"
)

var sk = "38d010a84c7b9374352468b41b076fa585d7dfac67ac34adabe2bbba4f4f6257"

func registerFlags(cfg *server.Config) {
	flag.IntVar(&cfg.Port, "port", 8080, "server listen port")
	flag.IntVar(&cfg.Fee, "fee", 10000, "transaction fee in satoish")
	flag.StringVar(&cfg.DataDir, "dataDir", ".skycoin-exchange", "data directory")
	flag.StringVar(&cfg.WalletName, "wltName", "server.wlt", "server's wallet file name")
	flag.StringVar(&cfg.Seed, "s", "test", "wallet's seed")
	flag.StringVar(&cfg.AcntName, "acntName", "account.data", "accounts file name")
	flag.IntVar(&cfg.UtxoPoolSize, "poolsize", 1000, "utxo pool size")

	// set the log dir
	// check if the log dir is exist, create it if not exist.
	logDir := filepath.Join(cfg.DataDir, "log")
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		// make directory
		os.Mkdir(logDir, 0700)
	}

	flag.Set("log_dir", logDir)
	flag.Set("alsologtostderr", "true")
}

func main() {
	cfg := server.Config{}
	registerFlags(&cfg)
	flag.Parse()
	key, err := cipher.SecKeyFromHex(sk)
	if err != nil {
		glog.Fatal(err)
	}
	cfg.Seckey = key
	s := server.New(cfg)
	s.Run()
}
