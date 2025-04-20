package integration_test

import (
	"net"
	"os"
	"os/exec"
	"testing"
	"time"
)

func TestTCPServer(t *testing.T) {
	// Start the server binary
	cmd := exec.Command("go", "run", "../../cmd/server/main.go")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		t.Fatalf("Failed to start server: %v", err)
	}
	defer cmd.Process.Kill()

	// Wait for server to start
	time.Sleep(500 * time.Millisecond)

	// --- Test Connection ---
	t.Run("Connect", func(t *testing.T) {
		conn, err := net.Dial("tcp", ":6000")
		if err != nil {
			t.Fatalf("Connection failed: %v", err)
		}
		conn.Close()
	})

	// // --- Test Message Handling ---
	// t.Run("MessageExchange", func(t *testing.T) {
	// 	conn, err := net.Dial("tcp", ":6000")
	// 	if err != nil {
	// 		t.Fatal(err)
	// 	}
	// 	defer conn.Close()
	//
	// 	if _, err := conn.Write([]byte("ping\n")); err != nil {
	// 		t.Fatal(err)
	// 	}
	//
	// 	buf := make([]byte, 1024)
	// 	n, err := conn.Read(buf)
	// 	if err != nil {
	// 		t.Fatal(err)
	// 	}
	// 	t.Logf("Received: %q", string(buf[:n]))
	// })
}
