package forms

import (
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)
	if !form.Valid() {
		t.Error("invalid form")
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)
	form.Required("a", "b", "c")

	if form.Valid() {
		t.Error("form shows valid when required fields are missing")
	}

	postedData := url.Values{}

	postedData.Add("a", "a")
	postedData.Add("b", "a")
	postedData.Add("c", "a")

	r = httptest.NewRequest("POST", "/whatever", nil)
	r.PostForm = postedData
	form = New(r.PostForm)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("form shows invalid when required fields are present")
	}

}

func TestMin_Length(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	postedData := url.Values{}
	postedData.Add("a", "abc")
	r.PostForm = postedData
	r.Form = postedData
	form := New(r.PostForm)

	ok := form.MinLength("a", 5, r)
	if ok {
		t.Error("min Length validation failed")
	}

	ok = form.MinLength("a", 3, r)
	if !ok {
		t.Error("min Length validation failed")
	}

}

func TestIsEmail(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	postedData := url.Values{}
	postedData.Add("email", "abc")
	r.PostForm = postedData
	r.Form = postedData
	form := New(r.PostForm)

	ok := form.IsEmail("email")
	if ok {
		t.Error("form field is not an email")
	}

	postedData.Add("email1", "xyz@gmail.com")
	r.PostForm = postedData
	r.Form = postedData
	ok = form.IsEmail("email1")
	if !ok {
		t.Error("form field shows invalid email when field has valid email")
	}

}

func TestForm_Has(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	postedData := url.Values{}
	postedData.Add("email", "abc")
	r.PostForm = postedData
	r.Form = postedData
	form := New(r.PostForm)

	ok := form.Has("email1", r)
	if ok {
		t.Error("shows form field does not has email1 but form does not have email1 field")
	}

	postedData.Add("email1", "xyz@gmail.com")
	r.PostForm = postedData
	r.Form = postedData
	ok = form.Has("email1", r)
	if !ok {
		t.Error("form field shows invalid email1 when field has valid email1")
	}

}
