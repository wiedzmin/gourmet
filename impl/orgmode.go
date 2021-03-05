package impl

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/wiedzmin/gourmet/util/orgmode"
)

type Headline struct {
	URL       string
	Title     string
	Level     int // Named "Position" in util/orgmode
	Timestamp time.Time
	Tags      []string
	Priority  string
}

var timestampFormat = "2006-01-02 15:04"

func getFlattenedHeadings(tree *orgmode.Tree) []orgmode.Node {
	var result []orgmode.Node
	for _, subtree := range tree.Subtrees {
		result = append(result, *subtree.Nodes[0])
		if len(subtree.Subtrees) > 0 {
			result = append(result, getFlattenedHeadings(subtree)...)
		}
	}
	return result
}

func parseHeadline(data string, loc *time.Location) (*Headline, error) {
	var result Headline
	dataRemainder := data
	tagsRegexp := regexp.MustCompile(`.*?\s+(:[A-Za-z0-9_@#%:]+:\s*$)`)
	timestampRegexp := regexp.MustCompile(` \[(\d{4}-\d{2}-\d{2} .+? \d{2}:\d{2})\]$`)
	headingRegexp := regexp.MustCompile(`\[\[(.*)\]\[(.*)\]\]$`)
	prioRegexp := regexp.MustCompile(`^\[\#([A-Z]+)\]`)
	if m, mi := tagsRegexp.FindStringSubmatch(dataRemainder), tagsRegexp.FindStringSubmatchIndex(dataRemainder); m != nil {
		result.Tags = strings.FieldsFunc(m[1], func(r rune) bool { return r == ':' })
		dataRemainder = strings.TrimSpace(dataRemainder[:mi[2]])
	}
	if m, mi := timestampRegexp.FindStringSubmatch(dataRemainder), timestampRegexp.FindStringSubmatchIndex(dataRemainder); m != nil {
		tsChunks := strings.Split(m[1], " ")
		newTS := strings.Join([]string{tsChunks[0], tsChunks[2]}, " ")
		ts, err := time.ParseInLocation(timestampFormat, newTS, loc)
		if err != nil {
			return nil, err
		}
		result.Timestamp = ts
		dataRemainder = strings.TrimSpace(dataRemainder[:mi[2]-2])
	}
	if m := prioRegexp.FindStringSubmatch(dataRemainder); m != nil {
		result.Priority = m[1]
		dataRemainder = strings.TrimSpace(dataRemainder[4:])
	}
	if m := headingRegexp.FindStringSubmatch(dataRemainder); m != nil {
		result.URL, result.Title = m[1], m[2]
	} else {
		result.URL = dataRemainder
		result.Title = ""
	}
	return &result, nil
}

func importOrg(orgFile string) error {
	orgTree := orgmode.TreeFromFile(orgFile)
	nodes := getFlattenedHeadings(orgTree)
	loc, _ := time.LoadLocation("Europe/Moscow") // FIXME: parameterize
	for _, n := range nodes {
		parsedHeadline, err := parseHeadline(n.Headline, loc)
		if err != nil {
			return err
		}
		// FIXME: move to tests (exploratory/debug)
		fmt.Printf("prio: |%v| :: headline tags: %v :: timestamp: %s :: url: %s :: title: %s\n", parsedHeadline.Priority, parsedHeadline.Tags, parsedHeadline.Timestamp, parsedHeadline.URL, parsedHeadline.Title)
	}
	return nil
}
