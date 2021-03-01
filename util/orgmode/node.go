package orgmode

import (
	"bufio"
	"encoding/json"
	"os"
	"regexp"
)

// /*
//  * A Node models an org-mode headline with a following section
//  * a section can be comprised of multiple lines
//  * position is the headline's asterisk count
// */

type Node struct {
	Headline string   `json:"headline"`
	Position int      `json:"position"`
	Section  []string `json:"sections"`
	parent   *Node
}

func (node *Node) findParent(nodes []*Node) *Node {
	if len(nodes) == 0 {
		return nil
	} else if nodes[len(nodes)-1].Position < node.Position {
		return nodes[len(nodes)-1]
	} else {
		nodes = nodes[0 : len(nodes)-1]
		return node.findParent(nodes)
	}
}

func (node Node) toJson() (string, error) {
	json, err := json.Marshal(node)
	if err != nil {
		return "", err
	}
	return string(json), nil
}

func nodesFromFile(path string) []*Node {
	file, _ := os.Open(path)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var node *Node
	var nodes []*Node
	var headline string
	var position int
	var section string
	var isBlock bool
	var isCodeBlock bool
	var isTable bool

	for scanner.Scan() {
		line := scanner.Text()

		r := regexp.MustCompile(`\A(\*+)\ (.*)`) // should use \S
		submatch := r.FindStringSubmatch(line)

		if len(submatch) > 1 {
			isBlock = false
			isCodeBlock = false
			isTable = false

			headline = submatch[2]
			position = len(submatch[1])

			node = &Node{Headline: headline, Position: position}

			node.parent = node.findParent(nodes)
			nodes = append(nodes, node)
		} else {
			codeStartReg := regexp.MustCompile(`\A(\#\+BEGIN_SRC)(.*)`)
			codeEndReg := regexp.MustCompile(`\A(\#\+END_SRC)`)

			tableReg := regexp.MustCompile(`\A\s*\|.*`)

			if codeStartReg.MatchString(line) {
				isCodeBlock = true
			} else if codeEndReg.MatchString(line) {
				isCodeBlock = false
			}

			isTable = tableReg.MatchString(line) && !isCodeBlock
			isBlock = isCodeBlock || isTable

			section += line

			if !isBlock {
				if len(nodes) == 0 {
					nodes = []*Node{&Node{Position: 1, Section: []string{section}}}
				} else {
					lastNode := nodes[len(nodes)-1]
					lastNode.Section = append(lastNode.Section, section)
				}
				section = ""
			}
		}
	}

	return nodes
}
