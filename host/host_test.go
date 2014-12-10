package host

import "testing"
import "fmt"

func  TestManager(t *testing.T) {
  manager := HostManager{
    hosts: map[string]Host{},
        }
  resp := manager.register("test")
  fmt.Println(resp)
  resp2 := manager.getHost(resp.id)
  fmt.Println(resp2)
  if resp.id != resp2.id {
    t.Fail()
  }
}
