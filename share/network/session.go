package network

import (
	"io"
	"net"
	"strings"
	"sync"

	"github.com/rxjh-emu/server/share/event"
	"github.com/rxjh-emu/server/share/log"
	"github.com/rxjh-emu/server/share/models/character"
	"github.com/rxjh-emu/server/share/util"
)

// max buffer size
const MAX_RECV_BUFFER_SIZE = 4096

type Session struct {
	socket   net.Conn
	buffer   []byte
	Protocol Protocol

	UserIdx uint16
	AuthKey uint32
	DataEx  any
	Data    struct {
		AccountId     int32  // database account id
		Username      string // username
		Verified      bool   // version verification
		LoggedIn      bool   // auth verification
		CharVerified  bool   // character delete password verification
		CharacterList []character.Character
	}

	PeriodicJobs map[string]*PeriodicTask
	jobMutex     sync.Mutex
}

// Starts session goroutine
func (s *Session) Start() {
	// Create new receiving buffer
	s.buffer = make([]byte, MAX_RECV_BUFFER_SIZE)

	// Loop indefinitely for reading data
	for {
		// Read data into buffer
		n, err := s.socket.Read(s.buffer)
		if err != nil {
			if err != io.EOF {
				log.Error("Error reading: %v", err)
			}
			s.Close()
			break
		}

		// Loop through received data
		for i := 0; i < n; {
			log.Debugf("Raw data: %v", util.FormatHex(s.buffer[:n]))
			arg := s.Protocol.Read(s.buffer[:n])
			arg.Session = s

			// Trigger packet received event
			event.Trigger(event.PacketReceiveEvent, arg)

			i += s.Protocol.GetHeaderSize() + arg.Length
		}
	}
}

// Sends specified data to the client
func (s *Session) Send(writer *Writer) {
	data := writer.Finalize()
	data = *s.Protocol.Write(data)

	log.Debugf("Session send|UserIdx: %d", s.UserIdx)
	log.Debug(util.FormatHex(data))

	// send it...
	length, err := s.socket.Write(data)
	if err != nil {
		log.Error("Error sending packet: " + err.Error())
		return
	}

	// create new packet event argument
	arg := &PacketArgs{
		Session: s,
		Length:  length,
		Type:    writer.Type,
		Data:    data,
		Reader:  nil,
	}

	// trigger packet sent event
	event.Trigger(event.PacketSendEvent, arg)
}

// Returns session's remote endpoint
func (s *Session) GetEndPnt() string {
	return s.socket.RemoteAddr().String()
}

// Returns session's ip address
func (s *Session) GetIp() string {
	var ip = strings.Split(s.GetEndPnt(), ":")
	return ip[0]
}

// GetLocalEndPntIp returns local end point IP address.
// Local end point is server to which session is connected to.
func (s *Session) GetLocalEndPntIp() string {
	pnt := s.socket.LocalAddr().String()
	ip := strings.Split(pnt, ":")
	return ip[0]
}

// IsLocal checks if session's remote endpoint originated from private network.
func (s *Session) IsLocal() bool {
	return net.IP.IsPrivate(net.ParseIP(s.GetIp()))
}

// Closes session socket
func (s *Session) Close() {
	s.RemoveAllJobs()
	s.socket.Close()
	event.Trigger(event.ClientDisconnectEvent, s)
}

// AddJob adds a periodic task to the session's job list in a thread-safe manner.
// The job can be referenced and managed by its name.
func (s *Session) AddJob(name string, task *PeriodicTask) {
	s.jobMutex.Lock()
	defer s.jobMutex.Unlock()

	s.PeriodicJobs[name] = task
}

// RemoveJob stops and removes a periodic task from the session's job list
// based on its name in a thread-safe manner.
func (s *Session) RemoveJob(name string) {
	s.jobMutex.Lock()
	defer s.jobMutex.Unlock()

	if job, exists := s.PeriodicJobs[name]; exists {
		job.Stop()
		delete(s.PeriodicJobs, name)
	}
}

// RemoveAllJobs stops all periodic tasks and clears the job list
// for the session in a thread-safe manner.
func (s *Session) RemoveAllJobs() {
	s.jobMutex.Lock()
	defer s.jobMutex.Unlock()

	for name, job := range s.PeriodicJobs {
		job.Stop()
		delete(s.PeriodicJobs, name)
	}
}
