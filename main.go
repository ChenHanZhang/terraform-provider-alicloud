package main

import (
	"encoding/json"
	"fmt"
	"github.com/aliyun/terraform-provider-alicloud/alicloud"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"net/http"
	"strings"
)

func main() {

	provider := alicloud.Provider().(*schema.Provider)

	resourceMaps := provider.ResourcesMap

	ephc := resourceMaps["alicloud_ehpc_cluster"]
	specResourceDefine := ResourceType{
		ResourceTypeCode: "alicloud_ehpc_cluster",
		Properties:       CoreSpecSchema(ephc.Schema),
	}

	b, err := json.Marshal(specResourceDefine)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%s\n", b)
	// postSpecSchema("alicloud_ehpc_cluster", "ehpc", "Cluster", string(b))

	traversalProviderResources()

}

func postSpecSchema(tfResource string, namespace string, resourceCode string, specSchema string) {

	targetUrl := "https://pre-acube.aliyun-inc.com/api/v1/terraform/resource/spec/create"

	postContent := map[string]string{}
	postContent["tfResource"] = tfResource
	postContent["tfVersion"] = ""
	postContent["namespace"] = namespace
	postContent["resourceCode"] = resourceCode
	postContent["meta"] = specSchema

	jsonContent, _ := json.Marshal(postContent)
	payload := strings.NewReader(string(jsonContent))

	req, _ := http.NewRequest("POST", targetUrl, payload)

	req.Header.Add("Content-Type", "application/json")

	_, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println(fmt.Errorf("==== post error: %s, %s, %s", tfResource, namespace, resourceCode))
		fmt.Println(err)
	}
}

func traversalProviderResources() {

	provider := alicloud.Provider().(*schema.Provider)
	resourceMaps := provider.ResourcesMap

	for resourceCode, resourceDefine := range resourceMaps {
		// 资源类型
		fmt.Println(resourceCode)

		specResourceDefine := ResourceType{
			ResourceTypeCode: resourceCode,
			Properties:       CoreSpecSchema(resourceDefine.Schema),
		}

		b, err := json.Marshal(specResourceDefine)
		if err != nil {
			fmt.Println("error:", err)
		}
		codes := strings.Split(resourceCode, "_")
		namespace := codes[1]
		resourceType := camelString(strings.Join(codes[2:], "_"))
		fmt.Println(namespace, resourceType, string(b))
		postSpecSchema(resourceCode, namespace, resourceType, string(b))
	}
}

func camelString(s string) string {
	data := make([]byte, 0, len(s))
	j := false
	k := false
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if k == false && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || k == false) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}
	return string(data[:])
}

func CoreSpecSchema(m map[string]*schema.Schema) Properties {
	if len(m) == 0 {
		// We return an actual (empty) object here, rather than a nil,
		// because a nil result would mean that we don't have a schema at
		// all, rather than that we have an empty one.
		return Properties{}
	}

	ret := map[string]*Property{}

	for name, attribute := range m {
		if attribute.Elem == nil {
			specAttribute := coreSpecAttributeStruct(attribute)
			specAttribute.Name = name
			specAttribute.Optional = attribute.Optional
			specAttribute.Required = attribute.Required
			specAttribute.ForceNew = attribute.ForceNew
			specAttribute.Computed = attribute.Computed
			specAttribute.Default = attribute.Default
			specAttribute.Deprecated = attribute.Deprecated
			specAttribute.Sensitive = attribute.Sensitive
			specAttribute.ConflictsWith = attribute.ConflictsWith
			specAttribute.ExactlyOneOf = attribute.ExactlyOneOf
			specAttribute.AtLeastOneOf = attribute.AtLeastOneOf
			specAttribute.Removed = attribute.Removed
			ret[name] = specAttribute
			continue
		}

		switch attribute.Elem.(type) {
		case *schema.Schema, schema.ValueType, *schema.Resource:
			specAttribute := coreSpecAttributeStruct(attribute)
			specAttribute.Name = name
			specAttribute.Optional = attribute.Optional
			specAttribute.Required = attribute.Required
			specAttribute.ForceNew = attribute.ForceNew
			specAttribute.Computed = attribute.Computed
			specAttribute.Default = attribute.Default
			specAttribute.Deprecated = attribute.Deprecated
			specAttribute.Sensitive = attribute.Sensitive
			specAttribute.ConflictsWith = attribute.ConflictsWith
			specAttribute.ExactlyOneOf = attribute.ExactlyOneOf
			specAttribute.AtLeastOneOf = attribute.AtLeastOneOf
			specAttribute.Removed = attribute.Removed
			ret[name] = specAttribute
		default:
			// Should never happen for a valid schema
			panic(fmt.Errorf("invalid Schema.Elem %#v; need *Schema or *Resource", attribute.Elem))
		}

	}
	return ret
}

func coreSpecAttributeStruct(v *schema.Schema) *Property {
	switch v.Type {
	case schema.TypeBool:
		return &Property{
			Type: TypeBool,
		}
	case schema.TypeString:
		return &Property{
			Type: TypeString,
		}
	case schema.TypeInt:
		return &Property{
			Type: TypeInteger,
		}
	case schema.TypeFloat:
		return &Property{
			Type: TypeNumber,
		}

	// 复合数据结构，需要进一步解析
	case schema.TypeList, schema.TypeSet, schema.TypeMap:
		var basicType bool
		var property *Property
		var properties = Properties{}

		// 首先根据Elem的类型进行解析
		switch set := v.Elem.(type) {
		case schema.ValueType:
			// 当Elem直接使用基本数据类型
			return &Property{
				Type: convertToSpecTypeKind(set),
			}
		case *schema.Schema:
			// 当Elem使用Schema，即一个基本数据类型节点
			property = coreSpecAttributeStruct(set)
			basicType = true
		case *schema.Resource:
			// 当Elem使用Resource，即一个复合数据类型节点
			properties = CoreSpecSchema(set.Schema)
			basicType = false
		default:
			// 当不存在Elem时，即Map类型（不定向key）
			return &Property{
				Type: TypeMap,
			}
		}

		if basicType {
			if v.MaxItems == 1 {
				return property.Object()
			} else {
				return property.Array()
			}
		} else {
			if v.MaxItems == 1 {
				return properties.Object()
			} else {
				return properties.Array()
			}
		}

	default:
		// 预期之外的类型
		panic(fmt.Errorf("invalid Schema.Type %s", v.Type))
	}
}

type ResourceType struct {
	ResourceTypeCode string                      `json:"resourceTypeCode"` // 资源名
	Properties       `json:"resourceProperties"` // 资源属性定义
}

type Property struct {
	Name                 string                        `json:"name"`                           // 属性code
	Type                 primitiveTypeKind             `json:"type"`                           // 属性类型(int,string..)
	Items                *Property                     `json:"items,omitempty"`                // 属性为array类型时，子项类型定义
	AdditionalProperties *Property                     `json:"additionalProperties,omitempty"` // 属性为 map(object的特例) 类型时，项类型定义
	Required             bool                          `json:"isRequired"`
	Optional             bool                          `json:"isOptional"`
	Computed             bool                          `json:"isComputed"`
	ForceNew             bool                          `json:"isForceNew"`
	DescriptionEn        string                        `json:"descriptionEn"`
	Description          string                        `json:"description"`
	Default              interface{}                   `json:"default"`
	Deprecated           string                        `json:"deprecated"`
	Sensitive            bool                          `json:"sensitive"`
	Removed              string                        `json:"removed"`
	RequiredWith         bool                          `json:"requiredWith"`
	ConflictsWith        []string                      `json:"conflictsWith"`
	AtLeastOneOf         []string                      `json:"atLeastOneOf"`
	ExactlyOneOf         []string                      `json:"exactlyOneOf"`
	*Properties          `json:"properties,omitempty"` // 属性为object类型时，子属性结构定义
}

type Properties map[string]*Property

func (elem *Property) Array() *Property {
	return &Property{
		Type:  TypeArray,
		Items: elem,
	}
}

func (elem *Properties) Array() *Property {
	return &Property{
		Type:  TypeArray,
		Items: elem.Object(),
	}
}

func (elem *Properties) Object() *Property {
	return &Property{
		Type:       TypeObject,
		Properties: elem,
	}
}

type primitiveTypeKind string

const (
	TypeBool    primitiveTypeKind = "boolean"
	TypeNumber  primitiveTypeKind = "number"
	TypeString  primitiveTypeKind = "string"
	TypeInteger primitiveTypeKind = "integer"
	TypeMap     primitiveTypeKind = "map"

	TypeArray  primitiveTypeKind = "array"
	TypeObject primitiveTypeKind = "object"
)

func convertToSpecTypeKind(t schema.ValueType) primitiveTypeKind {
	switch t {
	case schema.TypeBool:
		return TypeBool
	case schema.TypeString:
		return TypeString
	case schema.TypeInt:
		return TypeInteger
	case schema.TypeFloat:
		return TypeNumber
	default:
		// 基本类型默认为string
		return TypeString
	}
}
