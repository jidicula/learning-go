package main

import (
	"errors"
	"fmt"
	"net/http"
)

// LogOutput is a logger.
func LogOutput(message string) {
	fmt.Println(message)
}

type SimpleDataStore struct {
	userData map[string]string
}

// UserNameForID fetches username from an ID.
func (sds SimpleDataStore) UserNameForID(userID string) (string, bool) {
	name, ok := sds.userData[userID]
	return name, ok
}

// NewSimpleDataStore creates an instance of a `SimpleDataStore`.
func NewSimpleDataStore() SimpleDataStore {
	return SimpleDataStore{
		userData: map[string]string{
			"1": "Fred", "2": "Mary", "3": "Pat",
		},
	}
}

// Don't want to explicitly depend on SimpleDataStore
type DataStore interface {
	UserNameForID(userId string) (string, bool)
}

// Don't want to explicitly depend on LogOutput, so we define an interface...
type Logger interface {
	Log(message string)
}

// And an adaptor that implements the interface
type LoggerAdaptor func(message string)

// Log is an adaptor method that meets the Logger interface.
func (lg LoggerAdaptor) Log(message string) {
	lg(message)
}

// biz logic. Note that nothing in this struct actually mentions types, only
// interfaces, so there's no dependency on them. We can easily swap in new
// implementations from a different provider when we want to.
// The fields in here are also unexported.
type SimpleLogic struct {
	l  Logger
	ds DataStore
}

// SayHello says hello.
func (sl SimpleLogic) SayHello(userID string) (string, error) {
	sl.l.Log("in SayHello for " + userID)
	name, ok := sl.ds.UserNameForID(userID)
	if !ok {
		return "", errors.New("unknown user")
	}
	return "Hello, " + name, nil
}

// SayGoodbye says goodbye.
func (sl SimpleLogic) SayGoodbye(userID string) (string, error) {
	sl.l.Log("in SayGoodbye for " + userID)
	name, ok := sl.ds.UserNameForID(userID)
	if !ok {
		return "", errors.New("unknown user")
	}
	return "Goodbye, " + name, nil
}

// NewSimpleLogic is a SimpleLogic factory.
func NewSimpleLogic(l Logger, ds DataStore) SimpleLogic {
	return SimpleLogic{
		l:  l,
		ds: ds,
	}
}

// API layer.

// Interface is owned by client code, so method set is customized to needs of
// client code.
type Logic interface {
	SayHello(userID string) (string, error)
}

type Controller struct {
	l     Logger
	logic Logic
}

// SayHello handles a greeting query.
func (c Controller) SayHello(w http.ResponseWriter, r *http.Request) {
	c.l.Log("In SayHello")
	userID := r.URL.Query().Get("user_id")
	message, err := c.logic.SayHello(userID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	_, _ = w.Write([]byte(message))
}

// NewController is a Controller factory.
func NewController(l Logger, logic Logic) Controller {
	return Controller{
		l:     nil,
		logic: logic,
	}
}

func main() {
	l := LoggerAdaptor(LogOutput)
	ds := NewSimpleDataStore()
	logic := NewSimpleLogic(l, ds)
	c := NewController(l, logic)
	http.HandleFunc("/hello", c.SayHello)
	_ = http.ListenAndServe(":8080", nil)
}
