package sources

import (
	"context"
	"fmt"
	"net/url"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

// inrecruiting adapts Intervieweb/Inrecruiting public career sites. A board is the
// public host and locale path, such as inrecruiting.intervieweb.it/ansaldoenergia/it
// or bonatti.intervieweb.it/en. The initial HTML renders the first vacancy page and
// contains a CSRF-authorized endpoint for subsequent pages. Details carry JobPosting JSON-LD.
type inrecruiting struct{ http inrecruitingHTTP }

type inrecruitingHTTP interface {
	HTMLGetter
	HeaderFormPoster
}

func NewInrecruiting(c inrecruitingHTTP) Source { return inrecruiting{http: c} }

func (inrecruiting) Provider() string { return "inrecruiting" }

const inrecruitingMaxPages = 200

func (s inrecruiting) Fetch(ctx context.Context, e CompanyEntry) ([]Job, error) {
	base, err := inrecruitingCareerURL(e.Board)
	if err != nil {
		return nil, err
	}
	root, err := s.http.GetHTML(ctx, base)
	if err != nil {
		return nil, fmt.Errorf("inrecruiting: listing %s: %w", e.Board, err)
	}
	endpoint := inrecruitingPagingEndpoint(root)
	links := inrecruitingJobLinks(base, root)
	if endpoint != "" {
		for page := 2; page <= inrecruitingMaxPages; page++ {
			next, err := s.http.PostFormWithHeaders(ctx, endpoint, map[string]string{"referer": base}, url.Values{
				"act1": {"vacancyListCareer"}, "page": {fmt.Sprint(page)},
			})
			if err != nil {
				break
			}
			pageLinks := inrecruitingJobLinks(base, next)
			if len(pageLinks) == 0 {
				break
			}
			links = append(links, pageLinks...)
		}
	}
	return fetchDetails(links, defaultDetailWorkers, func(loc string) (Job, bool) {
		return s.detail(ctx, e, loc)
	}), nil
}

func inrecruitingCareerURL(board string) (string, error) {
	board = strings.Trim(strings.TrimSpace(board), "/")
	if board == "" || !strings.Contains(board, ".") {
		return "", fmt.Errorf("inrecruiting: invalid board %q: want host/tenant/locale", board)
	}
	return "https://" + board + "/career", nil
}

func inrecruitingPagingEndpoint(root *html.Node) string {
	var endpoint string
	walk(root, func(n *html.Node) bool {
		if n.Type == html.ElementNode && n.Data == "input" && Attr(n, "id") == "url-for-announces" {
			endpoint = Attr(n, "value")
			return false
		}
		return true
	})
	return endpoint
}

var inrecruitingJobIDPattern = regexp.MustCompile("/jobs/[^/?#]+-(\\d+)/(?:[a-z]{2})/?(?:[?#].*)?$")

func inrecruitingJobID(loc string) string { return firstSubmatch(inrecruitingJobIDPattern, loc) }

func inrecruitingJobLinks(base string, root *html.Node) []string {
	u, err := url.Parse(base)
	if err != nil {
		return nil
	}
	return jobLinks(u, root, func(href string) bool { return inrecruitingJobID(href) != "" })
}

func (s inrecruiting) detail(ctx context.Context, e CompanyEntry, loc string) (Job, bool) {
	root, err := s.http.GetHTML(ctx, loc)
	if err != nil {
		return Job{}, false
	}
	var p inrecruitingPosting
	if !ldJobPosting(root, &p) || inrecruitingJobID(loc) == "" {
		return Job{}, false
	}
	location := p.JobLocation.Address.Location()
	return Job{
		ExternalID: inrecruitingJobID(loc), URL: loc, Title: p.Title,
		Company:  firstNonEmpty(p.HiringOrganization.Name, e.Company),
		Location: location, Description: sanitizeHTML(html.UnescapeString(p.Description)),
		Remote: isRemote(location), PostedAt: parseRFC3339(p.DatePosted),
	}, true
}

type inrecruitingPosting struct {
	Title              string
	Description        string
	DatePosted         string
	HiringOrganization struct{ Name string }
	JobLocation        schemaPlace
}
