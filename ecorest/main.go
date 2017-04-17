package ecorest

import (
	"github.com/OpenTreeMap/otm-ecoservice/ecorest/cache"
	"github.com/OpenTreeMap/otm-ecoservice/ecorest/config"
	"github.com/OpenTreeMap/otm-ecoservice/ecorest/endpoints"
	"net/url"
)

type restManager struct {
	ITreeCodesGET       (func() *endpoints.ITreeCodes)
	EcoGET              (func(url.Values) (*endpoints.BenefitsWrapper, error))
	EcoSummaryPOST      (func(*endpoints.SummaryPostData) (*endpoints.BenefitsWrapper, error))
	EcoFullBenefitsPOST (func(*endpoints.FullBenefitsPostData) (*endpoints.FullBenefitsWrapper, error))
	EcoScenarioPOST     (func(*endpoints.ScenarioPostData) (*endpoints.Scenario, error))
	InvalidateCacheGET  (func())
}

func GetManager(cfg config.Config) *restManager {
	ecoCache, invalidateCache := cache.Init(cfg)
	invalidateCache()

	return &restManager{endpoints.ITreeCodesGET(ecoCache),
		endpoints.EcoGET(ecoCache),
		endpoints.EcoSummaryPOST(ecoCache),
		endpoints.EcoFullBenefitsPOST(ecoCache),
		endpoints.EcoScenarioPOST(ecoCache),
		invalidateCache}
}
