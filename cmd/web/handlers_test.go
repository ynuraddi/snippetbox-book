package main

import (

	// New import
	"net/http"
	"testing"
)

func TestPing(t *testing.T) {
	app := newTestApplication(t)
	ts := newTestServer(t, app.routes())
	defer ts.Close()
	code, _, body := ts.get(t, "/ping")
	if code != http.StatusOK {
		t.Errorf("want %d; got %d", http.StatusOK, code)
	}
	if string(body) != "OK" {
		t.Errorf("want body to equal %q", "OK")
	}
}

// func TestPing(t *testing.T) {
// 	// Create a new instance of our application struct. For now, this just
// 	// contains a couple of mock loggers (which discard anything written to
// 	// them).
// 	app := &application{
// 		errorLog: log.New(ioutil.Discard, "", 0),
// 		infoLog:  log.New(ioutil.Discard, "", 0),
// 	}
// 	// We then use the httptest.NewTLSServer() function to create a new test
// 	// server, passing in the value returned by our app.routes() method as the
// 	// handler for the server. This starts up a HTTPS server which listens on a
// 	// randomly-chosen port of your local machine for the duration of the test.
// 	// Notice that we defer a call to ts.Close() to shutdown the server when
// 	// the test finishes.
// 	ts := httptest.NewTLSServer(app.routes())
// 	defer ts.Close()
// 	// The network address that the test server is listening on is contained
// 	// in the ts.URL field. We can use this along with the ts.Client().Get()
// 	// method to make a GET /ping request against the test server. This
// 	// returns a http.Response struct containing the response.
// 	rs, err := ts.Client().Get(ts.URL + "/ping")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	// We can then check the value of the response status code and body using
// 	// the same code as before.
// 	if rs.StatusCode != http.StatusOK {
// 		t.Errorf("want %d; got %d", http.StatusOK, rs.StatusCode)
// 	}
// 	defer rs.Body.Close()
// 	body, err := ioutil.ReadAll(rs.Body)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	if string(body) != "OK" {
// 		t.Errorf("want body to equal %q", "OK")
// 	}
// }
