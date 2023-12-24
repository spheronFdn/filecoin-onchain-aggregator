package constants

import (
	"time"
)

var TopMiners map[string][]string = map[string][]string{
	"Mainnet": {"t03112", "t038057", "t035120", "t023097", "t034056", "t013799", "t028772", "t037295", "t038298", "t01158"},
	"Testnet": {"t01024", "t06024", "t01024", "t023097", "t034056", "t013799", "t028772", "t037295", "t038298", "t01158"},
	"Devnet":  {"f01000"},
}

var Polling_Rate time.Duration = time.Second * 10 // seconds

var Min_Deal_Duration = 518400
