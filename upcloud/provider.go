package upcloud

import (
	"errors"
	"os"
	"time"

	"github.com/UpCloudLtd/upcloud-go-api/upcloud"
	"github.com/UpCloudLtd/upcloud-go-api/upcloud/client"
	"github.com/UpCloudLtd/upcloud-go-api/upcloud/request"
	"github.com/UpCloudLtd/upcloud-go-api/upcloud/service"
	"github.com/nanovms/ops/lepton"
)

// Service is the interface implemented by upcloud service
type Service interface {
	service.Server
	service.Storage
	service.Tag
	service.Account
}

// Provider provides access to the Upcloud API.
type Provider struct {
	upcloud Service
	zone    string
}

// NewProvider returns an instance of Upcloud provider
func NewProvider() *Provider {
	return &Provider{}
}

// NewProviderWithService returns an instance of Upcloud provider and initializes upcloud service
func NewProviderWithService(service Service) *Provider {
	return &Provider{upcloud: service}
}

// Initialize checks conditions to use hyper-v
func (p *Provider) Initialize(config *lepton.ProviderConfig) error {
	user := os.Getenv("UPCLOUD_USER")
	if user == "" {
		return errors.New(`"UPCLOUD_USER" not set`)
	}

	password := os.Getenv("UPCLOUD_PASSWORD")
	if password == "" {
		return errors.New(`"UPCLOUD_PASSWORD" not set`)
	}

	p.zone = os.Getenv("UPCLOUD_ZONE")
	if p.zone == "" {
		return errors.New(`"UPCLOUD_ZONE" not set`)
	}

	if p.upcloud == nil {
		c := client.New(user, password)

		c.SetTimeout(time.Second * 30)

		p.upcloud = service.New(c)
	}

	_, err := p.upcloud.GetAccount()

	if err != nil {
		if serviceError, ok := err.(*upcloud.Error); ok {
			return errors.New(serviceError.ErrorMessage)
		}
		return errors.New("Invalid credentials")
	}

	return nil
}

func (p *Provider) findOrCreateOpsTag() (tags *[]upcloud.Tag, err error) {
	tags = &[]upcloud.Tag{}

	tagsResponse, err := p.upcloud.GetTags()
	if err != nil {
		return
	}

	for _, t := range tagsResponse.Tags {
		if t.Name == "OPS" {
			*tags = append(*tags, t)
			return
		}
	}

	tag := upcloud.Tag{
		Name:        "OPS",
		Description: "Created by ops",
	}

	createTagReq := &request.CreateTagRequest{Tag: tag}

	tagDetails, err := p.upcloud.CreateTag(createTagReq)
	if err != nil {
		return
	}

	*tags = append(*tags, *tagDetails)

	return
}
