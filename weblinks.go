package weblinks

import (
	"fmt"
	"math"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

var (
	linkRe = regexp.MustCompile(`<([^>]+)>; rel="([^"]+)"`)
)

// Create creates weblinks
func Create(baseURL string, page, pageSize, totalCount int) (*WebLinks, error) {
	if totalCount == 0 {
		totalCount = pageSize
	}
	parsedURL, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	lastPage := int(math.Ceil(float64(totalCount / pageSize)))
	res := WebLinks{
		Self:  paginateURL(parsedURL, page, pageSize),
		First: paginateURL(parsedURL, 1, pageSize),
		Last:  paginateURL(parsedURL, lastPage, pageSize),
	}

	if page > 1 {
		res.Prev = paginateURL(parsedURL, page-1, pageSize)
	}

	if totalCount > page*pageSize {
		nextPage := page + 1
		if nextPage == 1 {
			nextPage = 1
		}
		res.Next = paginateURL(parsedURL, nextPage, pageSize)
	}

	return &res, nil
}

// Parse parses a Link header
func Parse(linkHeader string) (*WebLinks, error) {
	res := &WebLinks{}
	lines := strings.Split(linkHeader, ",\n")

	for _, line := range lines {
		if !linkRe.MatchString(line) {
			return nil, errors.New("invalid line in header: " + line)
		}
		matches := linkRe.FindStringSubmatch(line)
		rel := matches[2]
		link, err := url.Parse(matches[1])
		if err != nil {
			return nil, errors.Wrapf(err, "could not parse url from rel:%s of linkHeader", rel)
		}

		switch rel {
		case "self":
			res.Self = link
			continue
		case "prev":
			res.Prev = link
			continue
		case "next":
			res.Next = link
			continue
		case "first":
			res.First = link
			continue
		case "last":
			res.Last = link
			continue
		}
	}
	return res, nil
}

type WebLinks struct {
	Self  *url.URL
	Next  *url.URL
	Prev  *url.URL
	First *url.URL
	Last  *url.URL
}

func (w *WebLinks) LinkHeader() string {
	items := []string{}

	if w.Self != nil {
		items = append(items, fmt.Sprintf(`<%s>; rel="%s"`, w.Self.String(), "self"))
	}
	if w.Next != nil {
		items = append(items, fmt.Sprintf(`<%s>; rel="%s"`, w.Next.String(), "next"))
	}
	if w.Prev != nil {
		items = append(items, fmt.Sprintf(`<%s>; rel="%s"`, w.Prev.String(), "prev"))
	}
	if w.First != nil {
		items = append(items, fmt.Sprintf(`<%s>; rel="%s"`, w.First.String(), "first"))
	}
	if w.Last != nil {
		items = append(items, fmt.Sprintf(`<%s>; rel="%s"`, w.Last.String(), "last"))
	}
	return strings.Join(items, ",\n")
}

func paginateURL(baseURL *url.URL, page, pageSize int) *url.URL {
	if page < 1 {
		return nil
	}

	// have to use a new URL because we can't manipulate the original
	newURL := url.URL(*baseURL)
	q := newURL.Query()
	q.Set("page", strconv.Itoa(page))
	q.Set("page_size", strconv.Itoa(pageSize))
	newURL.RawQuery = q.Encode()
	return &newURL
}
