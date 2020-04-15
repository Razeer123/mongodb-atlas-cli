package cli

import "testing"

func TestGetHostNameAndPort(t *testing.T) {
	t.Run("valid parameter", func(t *testing.T) {
		host, port, err := getHostNameAndPort("test:2000")
		if err != nil {
			t.Fatalf("getHostNameAndPort unexpecteted err: %#v\n", err)
		}
		if host != "test" {
			t.Errorf("Expected '%s', got '%s'\n", "test", host)
		}
		if port != 2000 {
			t.Errorf("Expected '%d', got '%d'\n", 2000, port)
		}
	})
	t.Run("incomplete format", func(t *testing.T) {
		_, _, err := getHostNameAndPort("test")
		if err == nil {
			t.Fatal("getHostNameAndPort should return an error\n")
		}
	})
	t.Run("incomplete format", func(t *testing.T) {
		_, _, err := getHostNameAndPort(":test")
		if err == nil {
			t.Fatal("getHostNameAndPort should return an error\n")
		}
	})
}
