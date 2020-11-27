package nmap

import (
	"context"
	"fmt"
	"github.com/Xyntax/CDK/conf"
	"log"
	"net"
	"strings"
	"sync"
	"time"

	"golang.org/x/sync/semaphore"
)

type PortScanner struct {
	ipRange string
	lock    *semaphore.Weighted
}

//func Ulimit() int64 {
//	out, err := exec.Command("ulimit", "-n").Output()
//	if err != nil {
//		panic(err)
//	}
//
//	s := strings.TrimSpace(string(out))
//
//	i, err := strconv.ParseInt(s, 10, 64)
//	if err != nil {
//		panic(err)
//	}
//
//	return i
//}

func ScanPort(ip string, port int, timeout time.Duration) {
	target := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.DialTimeout("tcp", target, timeout)

	if err != nil {
		if strings.Contains(err.Error(), "too many open files") {
			time.Sleep(timeout)
			ScanPort(ip, port, timeout)
		} else {
			fmt.Println(port, "closed")
		}
		return
	}

	conn.Close()
	fmt.Println(port, "open")
}

func (ps *PortScanner) Start() {
	wg := sync.WaitGroup{}
	defer wg.Wait()

	portFromTo, _ := GetTaskPortList()
	base, start, end, err := GetTaskIPList(ps.ipRange)
	if err != nil {
		log.Println("error found when gene ip list to scan task")
		log.Fatal(err)
	}

	// iterate ip in task list
	for p := start; p <= end; p++ {
		ip := base + "." + fmt.Sprintf("%d", p)
		// iterate port in task list
		for _, p := range portFromTo {
			// iterate port from A-B
			for port := p.From; port <= p.To; port++ {
				ps.lock.Acquire(context.TODO(), 1)
				wg.Add(1)
				go func(port int) {
					defer ps.lock.Release(1)
					defer wg.Done()
					fmt.Println(ip,port,p.Desc,conf.TCPScannerConf.Timeout)
					ScanPort(ip, port, conf.TCPScannerConf.Timeout)
				}(port)
			}
		}
	}
}

func TCPScanToolAPI(ipRange string) {
	ps := &PortScanner{
		ipRange:   ipRange,
		lock: semaphore.NewWeighted(conf.TCPScannerConf.MaxParallel),
	}
	ps.Start()
}
