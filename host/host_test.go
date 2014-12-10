package host

import "testing"
import "fmt"

func  TestManager(t *testing.T) {
  manager := new(HostManager)
  resp := manager.register("test")
  fmt.Printl(resp.id)
}
