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

}
