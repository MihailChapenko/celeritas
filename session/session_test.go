package session

import (
	"github.com/alexedwards/scs/v2"
	"reflect"
	"testing"
)

func TestSession_InitSession(t *testing.T) {
	s := &Session{
		CookieName:     "celeritas",
		CookieLifeTime: "100",
		CookiePersist:  "true",
		CookieDomain:   "localhost",
		SessionType:    "cookie",
	}

	var sm *scs.SessionManager
	var sessKind reflect.Kind
	var sessType reflect.Type

	sess := s.InitSession()
	rv := reflect.ValueOf(sess)

	for rv.Kind() == reflect.Ptr || rv.Kind() == reflect.Interface {
		sessKind = rv.Kind()
		sessType = rv.Type()

		rv = rv.Elem()
	}

	if !rv.IsValid() {
		t.Error("Invalid type or kind; kind:", rv.Kind(), "type:", rv.Type())
	}

	if sessKind != reflect.ValueOf(sm).Kind() {
		t.Error("wrong type returned. Expected:", reflect.ValueOf(sm).Kind(), "and got:", sessKind)
	}

	if sessType != reflect.ValueOf(sm).Type() {
		t.Error("wrong type returned. Expected:", reflect.ValueOf(sm).Type(), "and got:", sessType)
	}
}
