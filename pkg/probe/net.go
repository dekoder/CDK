package probe

import (
	"context"
	"fmt"
	"github.com/Xyntax/CDK/conf"
	"golang.org/x/sync/semaphore"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

type PortScanner struct {
	ipRange string
	lock    *semaphore.Weighted
}

func ScanPort(ip string, port int, timeout time.Duration) bool {
	target := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.DialTimeout("tcp", target, timeout)

	if err != nil {
		if strings.Contains(err.Error(), "too many open files") {
			time.Sleep(timeout)
			ScanPort(ip, port, timeout)
		}
		return false
	}

	_ = conn.Close()
	return true
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
					if ScanPort(ip, port, conf.TCPScannerConf.Timeout) {
						fmt.Printf("open %s: %s:%d\n", p.Desc, ip, port)
					}
				}(port)
			}
		}
	}
}

func TCPScanExploitAPI(ipRange string) {
	ps := &PortScanner{
		ipRange: ipRange,
		lock:    semaphore.NewWeighted(conf.TCPScannerConf.MaxParallel),
	}
	ps.Start()
}

func TCPScanToolAPI(ipRange string) {
	ps := &PortScanner{
		ipRange: ipRange,
		lock:    semaphore.NewWeighted(conf.TCPScannerConf.MaxParallel),
	}
	ps.Start()
}
