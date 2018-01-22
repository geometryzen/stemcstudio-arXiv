package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudsearchdomain"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// ProjectRefBundle is the
type ProjectRefBundle struct {
	Found int64
	Start int64
	Refs  []Submission
}

// Submission is the thing we are looking for.
type Submission struct {
	HRef     string
	Owner    string
	GistID   string
	Title    string
	Author   string
	Keywords []string
}

// SearchService is ...
type SearchService interface {
	Search(query string, size int) (*ProjectRefBundle, error)
	Submit(packet *Submission) (interface{}, error)
}

type defaultService struct {
	sess *session.Session
}

// NewSearchService is a factory for creating a search service.
func NewSearchService() SearchService {
	// AWS configuration
	sess, _ := session.NewSession(&aws.Config{Region: aws.String("us-east-1")})
	return &defaultService{sess}
}

func (s *defaultService) Search(query string, size int) (*ProjectRefBundle, error) {
	csd := cloudsearchdomain.New(s.sess, &aws.Config{Endpoint: aws.String("search-doodle-ref-xieragrgc2gcnrcog3r6bme75u.us-east-1.cloudsearch.amazonaws.com")})
	size64 := int64(size)
	data, err := csd.Search(&cloudsearchdomain.SearchInput{Query: aws.String(query), Size: &size64})
	if err != nil {
		fmt.Println("err  : ", err)
		return nil, err
	}
	response := ProjectRefBundle{Found: *data.Hits.Found, Start: *data.Hits.Start, Refs: []Submission{}}
	for _, record := range data.Hits.Hit {
		href := *record.Id
		owner := mapToString(record.Fields, "ownerkey")
		gistID := mapToString(record.Fields, "resourcekey")
		title := mapToString(record.Fields, "title")
		author := mapToString(record.Fields, "author")
		keywords := aws.StringValueSlice(record.Fields["keywords"])
		response.Refs = append(response.Refs, Submission{Author: author, GistID: gistID, HRef: href, Owner: owner, Title: title, Keywords: keywords})
	}
	return &response, nil
}

func (s *defaultService) Submit(packet *Submission) (interface{}, error) {
	db := dynamodb.New(s.sess)
	params := &dynamodb.PutItemInput{TableName: aws.String("DoodleRef"), Item: map[string]*dynamodb.AttributeValue{
		"OwnerKey": {
			S: aws.String(packet.Owner),
		},
		"ResourceKey": {
			S: aws.String(packet.GistID),
		},
		"Type": {
			S: aws.String("Gist"),
		},
		"Title": {
			S: aws.String(packet.Title),
		},
		"Author": {
			S: aws.String(packet.Author),
		},
		"Keywords": {
			SS: aws.StringSlice(packet.Keywords),
		},
	}}
	_, err := db.PutItem(params)
	if err != nil {
		fmt.Println("err  : ", err)
		return nil, nil
	}
	return nil, err
}

func mapToString(fields map[string][]*string, name string) string {
	xs := fields[name]
	if len(xs) > 0 {
		return aws.StringValue(xs[0])
	}
	return ""
}

func mapToStrings(fields map[string][]*string, name string) []string {
	return aws.StringValueSlice(fields[name])
}
