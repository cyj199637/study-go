package chapter02

import (
	"study-go/common"
)

/**
var a map[int]string -> 이런 식으로 초기화를 해버리면 a 에는 nil이 할당된다.
var a = make(map[int]string) -> 이렇게 초기화를 해야 empty map이 생성된다.
*/
type SampleDict map[string]string

func (s SampleDict) Search(key string) (string, error) {
	value, exists := s[key]
	if !exists {
		return "", common.ErrNotFound
	}
	return value, nil
}

/**
hashmap 은 기본적으로 * 가 포함되어 있기 때문에
변경 메소드를 정의할 때 receiver에 *를 붙이지 않아도 된다.
*/
func (s SampleDict) Add(key string, value string) error {
	_, err := s.Search(key)
	if err == nil {
		return common.ErrAlreadyExists
	}
	s[key] = value
	return nil
}

func (s SampleDict) Update(key string, value string) error {
	_, err := s.Search(key)
	if err != nil {
		return err
	}
	s[key] = value
	return nil
}

func (s SampleDict) Delete(key string) error {
	_, err := s.Search(key)
	if err != nil {
		return err
	}
	delete(s, key)
	return nil
}
