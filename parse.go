package gowd

import (
	"io"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"strings"
	"errors"
	"fmt"
)

type ElementsMap map[string]*Element

func NewElementMap() ElementsMap{
	return make(ElementsMap)
}

//ParseElements parse an html fragment and return a list of elements
func ParseElements(r io.Reader, em ElementsMap) ([]*Element, error) {
	nodes, err := html.ParseFragment(r, &html.Node{
		Type:     html.ElementNode,
		Data:     "body",
		DataAtom: atom.Body,
	})
	if err != nil {
		return nil, err
	}
	elems := make([]*Element, 0)
	for _, node := range nodes {
		elems = append(elems, NewElementFromNode(node,em))
	}
	return elems, nil
}

//NewElementFromNode creates an element from existing node
func NewElementFromNode(node*html.Node, em ElementsMap) *Element {
	elem := &Element{
		data:          strings.Trim(node.Data, "\n\r\t "),
		Attributes:    node.Attr,
		nodeType:      node.Type,
		Kids:          make([]*Element, 0),
		eventHandlers: make(map[string]EventHandler),
	}
	if em!=nil && elem.GetID()!=""{
		em[elem.GetID()]=elem
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		elem.AddElement(NewElementFromNode(c,em));
	}
	return elem
}

func ParseElement(innerHtml string, em ElementsMap) (*Element, error) {
	elems, err := ParseElements(strings.NewReader(innerHtml),em)
	if err != nil {
		return nil, err
	}
	if len(elems) != 1 {
		tags := "["
		for _, e := range elems {
			tags += fmt.Sprintf("%v:'%s', ",e.nodeType,strings.Trim(e.data,"\t\n\r "))

		}
		tags += ")"
		return nil, errors.New("The provided html must yield only one html element, I have: " + tags)
	}
	return elems[0], nil
}