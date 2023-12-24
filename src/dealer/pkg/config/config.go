package config

import (
	"os"
)

var ENV string = os.Getenv("ENV")                 // Use later, create a mapping for DEV,TEST and PROD
var WORKING_DIR string = os.Getenv("WORKING_DIR") // /tmp/import
var FIL_NODE string = os.Getenv("FIL_NODE")       // localhost:1234
var PRIV_KEY string = os.Getenv("PRIV_KEY")
var BEARER_TOKEN_FILECOIN_NODE = os.Getenv("BEARER_TOKEN_FILECOIN_NODE")
var DEAL_CLIENT_CONTRACT = os.Getenv("DEAL_CLIENT_CONTRACT")
var CHAIN_RPC = os.Getenv("CHAIN_RPC")
var MAX_BATCH_SIZE = os.Getenv("MAX_BATCH_SIZE")

// var WORKING_DIR string = "/tmp/import" // /tmp/import
// var FIL_NODE string = "localhost:1234"  // localhost:1234
// var CHAIN_RPC = "http://127.0.0.1:8545"
