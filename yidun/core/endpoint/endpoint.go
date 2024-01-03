package endpoint

import (
	"encoding/json"
	"errors"
	"fmt"
	"sync"
)

type EndpointConfigEntry struct {
	ProductCode string
	RegionCode  string
	Domains     []string
}

// EndpointResolver 用于解析域名
type EndpointResolver interface {
	Resolve(productCode, regionCode string) ([]string, error)
}

// EndpointResolverImpl 实现 EndpointResolver 接口
type DefaultEndpointResolver struct {
	domainMap map[string][]string
}

var (
	endpointResolver *DefaultEndpointResolver
	once             sync.Once
)

func GetEndpointResolver() *DefaultEndpointResolver {
	once.Do(func() {
		endpointResolver = newDefaultEndpointResolver()
	})
	return endpointResolver
}

func newDefaultEndpointResolver() *DefaultEndpointResolver {
	return &DefaultEndpointResolver{
		domainMap: GenerateFlatKeyMap(LoadFromResourceFiles()),
	}
}

type CompoundEndpointResolver struct {
	DefaultEndpointResolver
}

func NewCompoundEndpointResolver(productCode string, regionCode string, domain string) *CompoundEndpointResolver {
	return NewCompoundEndpointResolverWithCollection(productCode, regionCode, []string{domain})
}

func NewCompoundEndpointResolverWithCollection(productCode string, regionCode string, domains []string) *CompoundEndpointResolver {
	configEntry := EndpointConfigEntry{productCode, regionCode, domains}
	return &CompoundEndpointResolver{DefaultEndpointResolver: DefaultEndpointResolver{assembleDomainMap([]EndpointConfigEntry{configEntry})}}
}

func NewCompoundEndpointResolverWithConfigEntry(configEntry EndpointConfigEntry) *CompoundEndpointResolver {
	return &CompoundEndpointResolver{DefaultEndpointResolver: DefaultEndpointResolver{assembleDomainMap([]EndpointConfigEntry{configEntry})}}
}

func NewCompoundEndpointResolverWithConfigEntries(configEntries []EndpointConfigEntry) *CompoundEndpointResolver {
	return &CompoundEndpointResolver{DefaultEndpointResolver: DefaultEndpointResolver{assembleDomainMap(configEntries)}}
}

func assembleDomainMap(configEntries []EndpointConfigEntry) map[string][]string {
	tmpMap := LoadFromResourceFiles()
	mergeConfigWithOption(tmpMap, configEntries, true)
	return GenerateFlatKeyMap(tmpMap)
}

func (aer *DefaultEndpointResolver) Resolve(productCode, regionCode string) ([]string, error) {
	domainKey := genDomainKey(productCode, regionCode)
	domains, ok := aer.domainMap[domainKey]
	if !ok {
		return nil, errors.New(
			fmt.Sprintf("No available domain. productCode=%s, regionCode=%s", productCode, regionCode),
		)
	}
	return domains, nil
}

func genDomainKey(productCode, regionCode string) string {
	return productCode + "_" + regionCode
}

func ParseJsonToMap(jsonStr string) map[string]map[string][]string {
	// 定义目标类型
	m := make(map[string]map[string][]string)

	// 解码JSON字符串
	err := json.Unmarshal([]byte(jsonStr), &m)
	if err != nil {
		fmt.Println("JSON解码失败：", err)
		return nil
	}
	return m
}

// GenerateFlatKeyMap 将域名的索引（product code + region code）组合成单个字符串，方便查找。
// 新索引字符串由 genDomainKey 生成。
// 参数：existingDomainMap - key: product code; value: map - key: region code; value: domain list
// 返回值：key: 新索引字符串; value: domain list
func GenerateFlatKeyMap(existingDomainMap map[string]map[string][]string) map[string][]string {
	// create a map for the flat key
	flatKeyMap := make(map[string][]string)

	// loop through the existingDomainMap and generate the flat key map
	for productCode, regionMap := range existingDomainMap {
		for regionCode, domains := range regionMap {
			flatKey := genDomainKey(productCode, regionCode)
			flatKeyMap[flatKey] = domains
		}
	}
	return flatKeyMap
}
