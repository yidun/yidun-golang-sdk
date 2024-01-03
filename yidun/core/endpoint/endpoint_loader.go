package endpoint

const CONFIG_FILE_NAME = "yidun_regional_endpoints.json"

/*
* 从配置文件中加载域名映射表
@formatter:on
*
*/
func LoadFromResourceFiles() map[string]map[string][]string {
	return ParseJsonToMap(EndpointJSON)
}

/*
*

将新配置合并（去重）到已有映射表中
@param newDomainFirst 是否将新域名添加到原列表头部
*/
func mergeConfigWithOption(existingDomainMap map[string]map[string][]string, newEntries []EndpointConfigEntry, newDomainFirst bool) {
	for _, entry := range newEntries {
		// key: region code; value: domain list
		productCode := entry.ProductCode
		regionCode := entry.RegionCode
		newDomains := entry.Domains

		regionDomainMap, ok := existingDomainMap[productCode]
		if !ok {
			regionDomainMap = make(map[string][]string)
			existingDomainMap[productCode] = regionDomainMap
		}

		existingDomains, ok := regionDomainMap[regionCode]
		if !ok {
			existingDomains = make([]string, 0)
			regionDomainMap[regionCode] = existingDomains
		}

		if newDomainFirst {
			mergedDomains := make([]string, len(newDomains))
			copy(mergedDomains, newDomains)
			for _, domain := range existingDomains {
				if !contains(mergedDomains, domain) {
					mergedDomains = append(mergedDomains, domain)
				}
			}
			regionDomainMap[regionCode] = mergedDomains
		} else {
			for _, domain := range newDomains {
				if !contains(existingDomains, domain) {
					existingDomains = append(existingDomains, domain)
				}
			}
			regionDomainMap[regionCode] = existingDomains
		}
	}
}

func contains(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}
