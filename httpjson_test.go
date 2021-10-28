package httpjson

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"testing"
)

func testFn(fn responseFn, expectedStatus int, t *testing.T) {
	rr := httptest.NewRecorder()
	expected, e := json.Marshal(map[string]string{"foo": "bar"})

	if e != nil {
		t.Fatal(e)
	}
	fn(rr, map[string]string{"foo": "bar"})

	result := rr.Result()

	buf := &bytes.Buffer{}
	buf.ReadFrom(result.Body)
	actual := buf.String()

	if actual != string(expected) {
		fmt.Printf("Expected %s, but got %s\n", expected, actual)
		t.Fail()
	}

	if result.StatusCode != expectedStatus {
		fmt.Printf(
			"Expected status code to be %d, but got %d\n",
			expectedStatus,
			result.StatusCode,
		)
		t.Fail()
	}
}

func TestOK(t *testing.T) {
	testFn(OK, 200, t)
}

func TestAccepted(t *testing.T) {
	testFn(Accepted, 202, t)
}

func TestBadrequest(t *testing.T) {
	testFn(BadRequest, 400, t)
}

func TestUnauthorized(t *testing.T) {
	testFn(Unauthorized, 401, t)
}

func TestForbidden(t *testing.T) {
	testFn(Forbidden, 403, t)
}

func TestNotFound(t *testing.T) {
	testFn(NotFound, 404, t)
}

func TestMethodNotAllowed(t *testing.T) {
	testFn(MethodNotAllowed, 405, t)
}

func TestInternalServerError(t *testing.T) {
	testFn(InternalServerError, 500, t)
}
