package repository

import (
	"github.com/opensourceways/xihe-server/domain"
)

type ResourceSummaryListOption struct {
	Owner domain.Account
	Name  domain.ResourceName
}

type ModelPropertyUpdateInfo struct {
	ResourceToUpdate

	Property domain.ModelModifiableProperty
}

type UserModelsInfo struct {
	Models []domain.ModelSummary
	Total  int
}

type Model interface {
	Save(*domain.Model) (domain.Model, error)
	Delete(*domain.ResourceIndex) error
	Get(domain.Account, string) (domain.Model, error)
	GetByName(domain.Account, domain.ResourceName) (domain.Model, error)
	GetSummaryByName(domain.Account, domain.ResourceName) (domain.ResourceSummary, error)

	FindUserModels([]UserResourceListOption) ([]domain.ModelSummary, error)
	ListSummary([]ResourceSummaryListOption) ([]domain.ResourceSummary, error)

	ListAndSortByUpdateTime(domain.Account, *ResourceListOption) (UserModelsInfo, error)
	ListAndSortByFirstLetter(domain.Account, *ResourceListOption) (UserModelsInfo, error)
	ListAndSortByDownloadCount(domain.Account, *ResourceListOption) (UserModelsInfo, error)

	ListGlobalAndSortByUpdateTime(*GlobalResourceListOption) (UserModelsInfo, error)
	ListGlobalAndSortByFirstLetter(*GlobalResourceListOption) (UserModelsInfo, error)
	ListGlobalAndSortByDownloadCount(*GlobalResourceListOption) (UserModelsInfo, error)

	Search(*ResourceSearchOption) (ResourceSearchResult, error)

	AddLike(*domain.ResourceIndex) error
	RemoveLike(*domain.ResourceIndex) error

	AddRelatedDataset(*RelatedResourceInfo) error
	RemoveRelatedDataset(*RelatedResourceInfo) error

	AddRelatedProject(*domain.ReverselyRelatedResourceInfo) error
	RemoveRelatedProject(*domain.ReverselyRelatedResourceInfo) error

	UpdateProperty(*ModelPropertyUpdateInfo) error

	IncreaseDownload(*domain.ResourceIndex) error
}
