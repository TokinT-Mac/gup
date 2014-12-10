package host

import "testing"
import "fmt"

func  TestManager(t *testing.T) {
  manager := HostManager{
    hosts: map[string]*Host{},
        }
  resp := manager.register("test")
  fmt.Printl(resp.id)
}
