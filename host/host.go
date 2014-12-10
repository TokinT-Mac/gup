package host

import "time"
import "code.google.com/p/go-uuid/uuid"
import "fmt"

type Host struct {
  id string
  name string
  checkins []*CheckIn
  failures []*Failure
  online bool
}

type CheckIn struct {
  time time.Time
}

type Failure struct {
  time time.Time
  reason string
}

type Manager interface {
  register(string) string
  //checkin(string) string
  //list() map[string]Host
}

type HostManager struct {
  hosts = map[string]*Host{}
}

func (h *HostManager) register(name string) (host Host) {
  hostId = uuid.New()
  host = Host{id: hostId, name: name}
  h.hosts[hostId] = &host
  return
}
