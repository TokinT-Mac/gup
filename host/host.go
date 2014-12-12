package host

import "time"
import "code.google.com/p/go-uuid/uuid"
//import "fmt"

type Host struct {
  id string
  name string
  checkins map[string]CheckIn
  failures map[string]Failure
  online bool
}

type CheckIn struct {
  time time.Time
  id string
}

type Failure struct {
  time time.Time
  id string
  reason string
}

type Manager interface {
  register(string) string
  checkin(string) string
  list() map[string]Host
}

func NewHost(name string) Host {
  host := new(Host)
  host.id = uuid.New()
  host.checkins = make(map[string]CheckIn)
  host.failures = make(map[string]Failure)
  return *host
}

func NewCheckIn() (checkin CheckIn) {
  checkin = CheckIn{time: time.Now().UTC(), id: uuid.New()}
  return checkin
}

func NewFailure(reason string) (failure Failure) {
  failure = Failure{time: time.Now().UTC(), id: uuid.New(), reason: reason}
  return failure
}

type HostManager struct {
  hosts map[string]*Host
}

func (h *HostManager) register(name string) (host Host) {
  host = NewHost(name)
  h.hosts[host.id] = &host
  return
}

func (h HostManager) getHost(id string) *Host {
  host := h.hosts[id]
  return host
}

func (h *HostManager) checkin(id string) time.Time {
  check := NewCheckIn()
  host := h.hosts[id]
  host.checkins[check.id] = check

  return check.time
}
