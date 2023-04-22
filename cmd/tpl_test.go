package cmd

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/nvima/httpcli/util"
)

var testServer *httptest.Server

func TestMain(m *testing.M) {
	testServer = startTestHTTPServer()
	exitCode := m.Run()
	testServer.Close()
	os.Exit(exitCode)
}

func startTestHTTPServer() *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/v1/chat/completions":
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				http.Error(w, "Bad Request", http.StatusBadRequest)
			}

			var data ChatGptPayload
			err = json.Unmarshal(body, &data)
			if err != nil {
				http.Error(w, "Bad Request", http.StatusBadRequest)
			}

			w.Write([]byte(`{"choices": [{"message": {"content": "` + data.Messages[0].Content + `"}}]}`))
		case "/v2/translate":
			w.Write([]byte(`{"translations": [{"text": "Pomme"}]}`))
		case "/emptydata":
			w.Write([]byte(`{"response":"Hello"}`))
		default:
			http.Error(w, "Not found", http.StatusNotFound)
		}
	}))
	os.Setenv("TEST_SERVER_URL", server.URL)
	return server
}

func TestGitDiff(t *testing.T) {
	b := bytes.NewBufferString("")
	rootCmd.SetOut(b)
	rootCmd.SetArgs([]string{"gitdiff", "--config", "./testdata/config.yaml"})
	rootCmd.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	expected := "Hello"
	if string(out) != expected {
		t.Fatalf("expected \"%s\" got \"%s\"", expected, string(out))
	}
}

func TestTranslate(t *testing.T) {
	b := bytes.NewBufferString("")
	rootCmd.SetOut(b)
	rootCmd.SetArgs([]string{"translate", "--config", "./testdata/config.yaml"})
	rootCmd.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	expected := "Pomme"
	if string(out) != expected {
		t.Fatalf("expected \"%s\" got \"%s\"", expected, string(out))
	}
}

func TestEmptyData(t *testing.T) {
	b := bytes.NewBufferString("")
	rootCmd.SetOut(b)
	rootCmd.SetArgs([]string{"emptydata", "--config", "./testdata/config.yaml"})
	rootCmd.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	expected := "Hello"
	if string(out) != expected {
		t.Fatalf("expected \"%s\" got \"%s\"", expected, string(out))
	}
}

func TestEmptyOutput(t *testing.T) {
	b := bytes.NewBufferString("")
	rootCmd.SetOut(b)
	rootCmd.SetArgs([]string{"emptyoutput", "--config", "./testdata/config.yaml"})
	rootCmd.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	expected := []byte(`{"response":"Hello"}`)
	if string(out) != string(expected) {
		t.Fatalf("expected \"%s\" got \"%s\"", expected, string(out))
	}
}

func TestEmptyUrl(t *testing.T) {
	b := bytes.NewBufferString("")
	rootCmd.SetOut(b)
	rootCmd.SetArgs([]string{"emptyurl", "--config", "./testdata/config.yaml"})
	err := rootCmd.Execute()
	if err.Error() != util.NO_URL_ERR.Msg() {
		t.Fatalf("expected \"%s\" got \"%s\"", util.NO_URL_ERR.Msg(), err.Error())
	}
}
