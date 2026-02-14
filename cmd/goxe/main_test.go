package main

import (
	"net"
	"os"
	"testing"
	"time"
)

func BenchmarkServerUDP(b *testing.B) {
	go main()

	oldArgs := os.Args
	os.Args = []string{oldArgs[0]}
	defer func() { os.Args = oldArgs }()
	time.Sleep(100 * time.Millisecond)

	conn, err := net.Dial("udp", "127.0.0.1:1729")
	if err != nil {
		b.Fatal(err)
	}
	defer conn.Close()

	logs := [][]byte{
		[]byte("null_12 Aug 22, 2022 23:59:59 debug: Gateway connection trace %d 128.24.49.22"),
		[]byte("null_12 Aug 22, 2022 23:59:59 info: User Logged %d 128.24.49.22"),
		[]byte("null_12 Aug 22, 2022 23:59:59 notice: System update available %d 128.24.49.22"),
		[]byte("null_12 Aug 22, 2022 23:59:59 warning: High memory usage %d 128.24.49.22"),
		[]byte("null_12 Aug 22, 2022 23:59:59 warn: Disk space low %d 128.24.49.22"),
		[]byte("null_12 Aug 22, 2022 23:59:59 error: Gateway terminated %d 128.24.49.22"),
		[]byte("null_12 Aug 22, 2022 23:59:59 critical: Hardware failure detected %d 128.24.49.22"),
		[]byte("null_12 Aug 22, 2022 23:59:59 alert: Intrusion attempt blocked %d 128.24.49.22"),
		[]byte("null_12 Aug 22, 2022 23:59:59 emergency: Kernel panic, system halting %d 128.24.49.22"),
	}

	b.ResetTimer()
	b.ReportAllocs()
	for b.Loop() {
		log := logs[b.N%len(logs)]
		_, err := conn.Write(log)
		if err != nil {
			b.Errorf("Error enviando log: %v", err)
		}
	}
}
