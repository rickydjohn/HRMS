package sessions

import (
	"encoding/json"
	"errors"
	"strings"
)

func Begin() SessInterface {
	return &inSession{
		session: make(map[string]session),
	}
}

func (s *inSession) Create(uuid, user, ip string, uid int) bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	ipaddr := strings.Split(ip, ":")
	s.session[uuid] = session{Ip: ipaddr[0], Uname: user, Uid: uid}
	return true
}

func (s *inSession) Delete(uuid string) error {
	s.lock.Lock()
	defer s.lock.Unlock()
	_, k := s.session[uuid]
	if !k {
		return errors.New("The SSID provided does not exist: " + uuid)
	}
	delete(s.session, uuid)
	return nil
}

func (s *inSession) IsValid(uuid, ip string) error {
	s.lock.Lock()
	defer s.lock.Unlock()
	v, k := s.session[uuid]
	if !k {
		return errors.New("invalid session ID")
	}
	ipaddr := strings.Split(ip, ":")
	if v.Ip != ipaddr[0] {
		return errors.New("request is originating from a different source")
	}
	return nil
}

func (s *inSession) GetUID(uuid string) (int, error) {
	v, k := s.session[uuid]
	if !k {
		return 0, errors.New("Not a valid session ID")
	}
	return v.Uid, nil
}

func (s *inSession) AllSessions() ([]byte, error) {
	tmp := make(map[string]session)
	for i, v := range s.session {
		tmp[i] = v
	}
	bt, err := json.Marshal(tmp)
	if err != nil {
		return nil, err
	}
	return bt, err
}
