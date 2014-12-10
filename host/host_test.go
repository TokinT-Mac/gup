package host

import "testing"
import "fmt"
func  TestingManager(t *testing.T) {
  manager = new(HostManager)
  resp = manager.register('test')
  fmt.Printl(resp)
}
