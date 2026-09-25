package sources

import (
	"context"
	"net/url"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

type inrecruitingFake struct {
	get  map[string]string
	post map[string]string
}

func (f inrecruitingFake) GetHTML(_ context.Context, u string) (*html.Node, error) {
	return html.Parse(strings.NewReader(f.get[u]))
}

func (f inrecruitingFake) PostFormWithHeaders(_ context.Context, _ string, _ map[string]string, values url.Values) (*html.Node, error) {
	return html.Parse(strings.NewReader(f.post[values.Get("page")]))
}

func TestInrecruitingFetchesPagedJSONLDPostings(t *testing.T) {
	listing := func(link string) string {
		return "<input id=\"url-for-announces\" value=\"https://tenant.example/app.php\"><a href=\"" + link + "\">role</a>"
	}
	detail := func(title string) string {
		return "<script type=\"application/ld+json\">{\"@type\":\"JobPosting\",\"title\":\"" + title + "\",\"description\":\"<p>Work</p>\",\"datePosted\":\"2026-09-25\",\"hiringOrganization\":{\"name\":\"Acme\"},\"jobLocation\":{\"address\":{\"addressLocality\":\"Milan\",\"addressCountry\":\"IT\"}}}</script>"
	}
	base := "https://tenant.example/en/career"
	fake := inrecruitingFake{
		get: map[string]string{
			base: listing("/jobs/first-role-123/en/"),
			"https://tenant.example/jobs/first-role-123/en/":  detail("First"),
			"https://tenant.example/jobs/second-role-456/en/": detail("Second"),
		},
		post: map[string]string{"2": listing("/jobs/second-role-456/en/"), "3": ""},
	}
	jobs, err := NewInrecruiting(fake).Fetch(context.Background(), CompanyEntry{Board: "tenant.example/en"})
	if err != nil {
		t.Fatal(err)
	}
	if len(jobs) != 2 {
		t.Fatalf("jobs = %d, want 2", len(jobs))
	}
	if jobs[0].ExternalID != "123" || jobs[1].ExternalID != "456" {
		t.Fatalf("unexpected ids: %#v", jobs)
	}
	if jobs[0].Company != "Acme" || jobs[0].Location != "Milan, IT" {
		t.Fatalf("unexpected mapped job: %#v", jobs[0])
	}
}

func TestInrecruitingRegistered(t *testing.T) {
	if _, ok := All(nil)["inrecruiting"]; !ok {
		t.Fatal("All() missing inrecruiting")
	}
}
