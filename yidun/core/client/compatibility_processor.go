package client

import (
	"reflect"
)

// CompatibilityProcessor 兼容性处理器
// 用于处理 Antispam 结构体中新旧字段的兼容性同步
type CompatibilityProcessor struct{}

// NewCompatibilityProcessor 创建新的兼容性处理器
func NewCompatibilityProcessor() *CompatibilityProcessor {
	return &CompatibilityProcessor{}
}

// ProcessForDeserialization 反序列化后的兼容性处理
// 将新字段值同步到老字段，保持向后兼容
func (cp *CompatibilityProcessor) ProcessForDeserialization(data interface{}) error {
	if data == nil {
		return nil
	}

	v := reflect.ValueOf(data)
	if v.Kind() != reflect.Ptr || v.IsNil() {
		return nil
	}

	cp.processValue(v.Elem())
	return nil
}

// processValue 递归处理值
func (cp *CompatibilityProcessor) processValue(v reflect.Value) {
	if !v.IsValid() {
		return
	}

	switch v.Kind() {
	case reflect.Ptr:
		if !v.IsNil() {
			cp.processValue(v.Elem())
		}
	case reflect.Struct:
		cp.processStruct(v)
	case reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			cp.processValue(v.Index(i))
		}
	}
}

// processStruct 处理结构体
func (cp *CompatibilityProcessor) processStruct(v reflect.Value) {
	t := v.Type()

	// 检查是否是 Antispam 结构体（包含需要同步的字段）
	if cp.isAntispamStruct(t) {
		cp.syncAntispamFields(v)
	}

	// 递归处理导出的字段
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)

		// 只处理导出字段
		if !fieldType.IsExported() {
			continue
		}

		// 只递归处理指针、结构体和切片类型
		switch field.Kind() {
		case reflect.Ptr, reflect.Struct, reflect.Slice:
			cp.processValue(field)
		}
	}
}

// isAntispamStruct 检查是否是包含兼容性字段的 Antispam 结构体
func (cp *CompatibilityProcessor) isAntispamStruct(t reflect.Type) bool {
	_, hasRelateContent := t.FieldByName("RelateContent")
	_, hasRelatedContents := t.FieldByName("RelatedContents")
	_, hasHitSource := t.FieldByName("HitSource")
	_, hasHitSources := t.FieldByName("HitSources")

	return hasRelateContent && hasRelatedContents && hasHitSource && hasHitSources
}

// syncAntispamFields 同步 Antispam 结构体中的新旧字段
func (cp *CompatibilityProcessor) syncAntispamFields(v reflect.Value) {
	// RelateContent -> RelatedContents
	relateContent := v.FieldByName("RelateContent")
	relatedContents := v.FieldByName("RelatedContents")
	if relateContent.IsValid() && relatedContents.IsValid() && relatedContents.CanSet() {
		if !relateContent.IsNil() {
			relatedContents.Set(relateContent)
		} else if !relatedContents.IsNil() && relateContent.CanSet() {
			relateContent.Set(relatedContents)
		}
	}

	// HitSource -> HitSources
	hitSource := v.FieldByName("HitSource")
	hitSources := v.FieldByName("HitSources")
	if hitSource.IsValid() && hitSources.IsValid() && hitSources.CanSet() {
		if !hitSource.IsNil() {
			hitSources.Set(hitSource)
		} else if !hitSources.IsNil() && hitSource.CanSet() {
			hitSource.Set(hitSources)
		}
	}
}

// GetGlobalCompatibilityProcessor 获取全局兼容性处理器实例
func GetGlobalCompatibilityProcessor() *CompatibilityProcessor {
	return globalCompatibilityProcessor
}

// 全局兼容性处理器实例
var globalCompatibilityProcessor = NewCompatibilityProcessor()