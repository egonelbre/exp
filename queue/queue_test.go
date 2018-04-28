// +build ignore

package queue

import "testing"

func testBlocking(t *testing.T, q Blocking)       {}
func testNonblocking(t *testing.T, q Nonblocking) {}

func benchBlocking(b *testing.B, q Blocking)       {}
func benchNonblocking(b *testing.B, q Nonblocking) {}
