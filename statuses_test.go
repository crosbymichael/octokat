package octokat

import (
	"net/http"
	"testing"

	"github.com/bmizerany/assert"
)

func TestStatuses(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/repos/jingweno/gh/statuses/740211b9c6cd8e526a7124fe2b33115602fbc637", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		respondWith(w, loadFixture("statuses.json"))
	})

	repo := Repo{UserName: "jingweno", Name: "gh"}
	sha := "740211b9c6cd8e526a7124fe2b33115602fbc637"
	statuses, _ := client.Statuses(repo, sha, nil)
	assert.Equal(t, 2, len(statuses))

	firstStatus := statuses[0]
	assert.Equal(t, "pending", firstStatus.State)
	assert.Equal(t, "The Travis CI build is in progress", firstStatus.Description)
	assert.Equal(t, "https://travis-ci.org/jingweno/gh/builds/11911500", firstStatus.TargetURL)
}

func TestCombinedStatus(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/repos/jingweno/gh/commits/740211b9c6cd8e526a7124fe2b33115602fbc637/status", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		respondWith(w, loadFixture("combined_status.json"))
	})

	repo := Repo{UserName: "jingweno", Name: "gh"}
	sha := "740211b9c6cd8e526a7124fe2b33115602fbc637"
	combinedStatus, _ := client.CombinedStatus(repo, sha, nil)
	assert.Equal(t, "pending", combinedStatus.State)
	assert.Equal(t, 2, combinedStatus.TotalCount)
	assert.Equal(t, 2, len(combinedStatus.Statuses))

	firstStatus := combinedStatus.Statuses[0]
	assert.Equal(t, "pending", firstStatus.State)
	assert.Equal(t, "The Travis CI build is in progress", firstStatus.Description)
	assert.Equal(t, "https://travis-ci.org/jingweno/gh/builds/11911500", firstStatus.TargetURL)
}

func TestSetStatus(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/repos/erikh/test/statuses/5456e827a6b4930a3b1d727b7bd407b0b499c08f", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		respondWith(w, loadFixture("status_response.json"))
	})

	repo := Repo{UserName: "erikh", Name: "test"}
	sha := "5456e827a6b4930a3b1d727b7bd407b0b499c08f"

	opts := &StatusOptions{
		State:       "success",
		Description: "this is only a test",
		URL:         "http://google.com",
		Context:     "docci",
	}

	status, err := client.SetStatus(repo, sha, opts)

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "this is only a test", status.Description)
	assert.Equal(t, "success", status.State)
}
