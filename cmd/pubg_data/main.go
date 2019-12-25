package main

import (
	"fmt"
	"os"
	"time"

	utils "github.com/pubg_go/pubg_last_id/utils"
)

func main() {
	start := time.Now()
	playerName := os.Args[1]
	_, lastid := utils.GetLastID(playerName)
	//fmt.Printf("Account id: %v\nLast match id: %v\n", accid, lastid)
	telURL := utils.GetTelemetryURL(lastid)
	all := utils.GetKillersVictims(telURL)
	for i := range all {
		if all[i].KillerName == playerName || all[i].VictimName == playerName {
			fmt.Println(all[i])
		}
	}
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Printf("Execution took %v\n", elapsed)
}