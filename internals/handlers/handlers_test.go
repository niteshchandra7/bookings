package handlers

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/niteshchandra7/bookings/internals/models"
)

type postData struct {
	key   string
	value string
}

// var reservation models.Reservation = models.Reservation{
// 	ID:        1,
// 	RoomID:    1,
// 	FirstName: "John",
// 	LastName:  "Smith",
// 	Room:      models.Room{},
// 	Email:     "me@here.com",
// 	Phone:     "10292999",
// 	StatDate:  time.Date(2022, time.April, 1, 0, 0, 0, 0, nil),
// 	EndDate:   time.Date(2022, time.April, 3, 0, 0, 0, 0, nil),
// }

var theTests = []struct {
	name   string
	url    string
	method string
	//params             []postData
	expectedStatusCode int
}{
	{"home", "/", "GET", http.StatusOK},
	{"about", "/about", "GET", http.StatusOK},
	{"gq", "/generals-quarters", "GET", http.StatusOK},
	{"ms", "/majors-suite", "GET", http.StatusOK},
	{"sa", "/search-availability", "GET", http.StatusOK},
	//{"mr", "/make-reservation", "GET", []postData{}, http.StatusOK},
	// {"post-search-avail", "/search-availability", "POST", []postData{
	// 	{key: "start", value: "2020-01-01"},
	// 	{key: "end", value: "2020-01-02"},
	// }, http.StatusOK},
	// {"post-search-avail-json", "/search-availability-json", "POST", []postData{
	// 	{key: "start", value: "2020-01-01"},
	// 	{key: "end", value: "2020-01-02"},
	// }, http.StatusOK},
	// {"make-reservation-post", "/make-reservation", "POST", []postData{
	// 	{key: "first_name", value: "Harry"},
	// 	{key: "last_name", value: "Kumar"},
	// 	{key: "email", value: "harry@gmail.com"},
	// 	{key: "phone", value: "900345612"},
	// }, http.StatusOK},
}

func TestHandlers(t *testing.T) {
	routes := getRoutes()
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	for _, e := range theTests {
		//if e.method == "GET" {
		resp, err := ts.Client().Get(ts.URL + e.url)
		if err != nil {
			t.Log(err)
			t.Fatal(err)
		}

		if resp.StatusCode != e.expectedStatusCode {
			t.Errorf("for %s, expected %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
		}
		// } else {
		// 	values := url.Values{}
		// 	// for _, x := range e.params {
		// 	// 	values.Add(x.key, x.value)
		// 	// }
		// 	//values.Add("csrf_token", models.TemplateData.CSRFToken)
		// 	resp, err := ts.Client().PostForm(ts.URL+e.url, values)
		// 	if err != nil {
		// 		t.Log(err)
		// 		t.Fatal(err)
		// 	}
		// 	if resp.StatusCode != e.expectedStatusCode {
		// 		t.Errorf("for %s, expected %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
		// 	}

		// }
	}
}

func TestRepository_Reservation(t *testing.T) {
	reservation := models.Reservation{
		RoomID: 1,
		Room: models.Room{
			ID:       1,
			RoomName: "Generals-Quarters",
		},
	}

	req, _ := http.NewRequest("GET", "/make-reservation", nil)
	ctx, _ := getContext(req)
	req = req.WithContext(ctx)

	rr := httptest.NewRecorder()
	session.Put(ctx, "reservation", reservation)
	handler := http.HandlerFunc(Repo.Reservation)

	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("Reservation handler wrong response code got %d wanted %d", rr.Code, http.StatusOK)
	}

	// test case where reservation is not in session

	req, _ = http.NewRequest("GET", "make-reservation", nil)
	ctx, _ = getContext(req)
	req = req.WithContext(ctx)
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusTemporaryRedirect {
		t.Errorf("Redirection failed")
	}

	// test with non existent room
	req, _ = http.NewRequest("GET", "make-reservation", nil)
	ctx, _ = getContext(req)
	req = req.WithContext(ctx)
	rr = httptest.NewRecorder()
	reservation.RoomID = 3
	session.Put(ctx, "reservation", reservation)

	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusTemporaryRedirect {
		t.Errorf("Redirection failed")
	}

}

func getContext(r *http.Request) (context.Context, error) {
	ctx, err := session.Load(r.Context(), r.Header.Get("X-Session"))
	if err != nil {
		return nil, err
	}
	return ctx, nil
}
