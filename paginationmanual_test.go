// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcombeeperdesktopapigo_test

import (
	"context"
	"os"
	"testing"

	"github.com/beeper/desktop-api-go"
	"github.com/beeper/desktop-api-go/internal/testutil"
	"github.com/beeper/desktop-api-go/option"
)

func TestManualPagination(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := githubcombeeperdesktopapigo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAccessToken("My Access Token"),
	)
	page, err := client.Messages.SearchMessages(context.TODO(), githubcombeeperdesktopapigo.MessageSearchMessagesParams{
		Limit: githubcombeeperdesktopapigo.Int(20),
		Query: githubcombeeperdesktopapigo.String("meeting"),
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	for _, message := range page.Data {
		t.Logf("%+v\n", message.ID)
	}
	// Prism mock isn't going to give us real pagination
	page, err = page.GetNextPage()
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if page != nil {
		for _, message := range page.Data {
			t.Logf("%+v\n", message.ID)
		}
	}
}
