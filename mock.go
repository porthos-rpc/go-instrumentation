package instrumentation

import (
	"time"
)

func Mock() Application {
	return &mockApplication{}
}

type mockApplication struct {
}

func (m *mockApplication) StartTransaction(name string) Transaction {
	return &mockTransaction{}
}

func (m *mockApplication) Shutdown(timeout time.Duration) {}

type mockTransaction struct {
}

func (m *mockTransaction) StartSegment(name string) Segment {
	return &mockSegment{}
}

func (m *mockTransaction) End() error {
	return nil
}

func (m *mockTransaction) Ignore() error {
	return nil
}

func (m *mockTransaction) NoticeError(err error) error {
	return nil
}

func (m *mockTransaction) AddAttribute(key string, value interface{}) error {
	return nil
}

type mockSegment struct {
}

func (m *mockSegment) End() error {
	return nil
}
