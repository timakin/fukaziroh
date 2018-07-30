package main

import (
	"bytes"
	"os"
	"time"

	vegeta "github.com/tsenart/vegeta/lib"
)

func main() {
	rate := uint64(100) // per second
	duration := 50 * time.Second
	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "GET",
		URL:    "http://localhost:8000/v1",
	})
	attacker := vegeta.NewAttacker(vegeta.Workers(50))

	rs := make(vegeta.Results, 5000)
	for res := range attacker.Attack(targeter, rate, duration, "GET user's campaigns") {
		rs = append(rs, *res)
	}

	var buf bytes.Buffer
	rep := vegeta.NewPlotReporter("Plot Report", &rs)
	rep.Report(&buf)
	// pp.Println(buf.String())
	file, err := os.Create(`./rep.html`)
	if err != nil {
		// Openエラー処理
	}
	defer file.Close()
	file.Write(buf.Bytes())
	// pp.Println(metrics.)
	// fmt.Printf("99th percentile: %s\n", metrics.Latencies.P99)
}
