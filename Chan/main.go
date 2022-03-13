package main

import (
	"fmt"
	"net"
	"time"
)

func main(){
	msg := []byte{0,0}
	var server net.UDPAddr
	server.IP = []byte{192,168,3,1}
	server.Port = 5351
	conn, err := net.DialUDP("udp", nil, &server)
	if err != nil {
		return
	}
	defer conn.Close()

	// 16 bytes is the maximum result size.
	result := make([]byte, 16)

	var finalTimeout time.Time

		finalTimeout = time.Now().Add(time.Second*30)


	needNewDeadline := true

	var tries uint
	for tries = 0; (tries < 9 && finalTimeout.IsZero()) || time.Now().Before(finalTimeout); {
		if needNewDeadline {
			nextDeadline := time.Now().Add((250 << tries) * time.Millisecond)
			err = conn.SetDeadline(minTime(nextDeadline, finalTimeout))
			if err != nil {
				return
			}
			needNewDeadline = false
		}
		_, err = conn.Write(msg)
		if err != nil {
			return
		}
		var bytesRead int
		var remoteAddr *net.UDPAddr
		bytesRead, remoteAddr, err = conn.ReadFromUDP(result)
		if err != nil {
			if err.(net.Error).Timeout() {
				tries++
				needNewDeadline = true
				continue
			}
			return
		}
		if !remoteAddr.IP.Equal(server.IP) {
			// Ignore this packet.
			// Continue without increasing retransmission timeout or deadline.
			continue
		}
		// Trim result to actual number of bytes received
		if bytesRead < len(result) {
			result = result[:bytesRead]
		}
		return
	}
	err = fmt.Errorf("Timed out trying to contact gateway")
	return
}

func minTime(a, b time.Time) time.Time {
	if a.IsZero() {
		return b
	}
	if b.IsZero() {
		return a
	}
	if a.Before(b) {
		return a
	}
	return b
}