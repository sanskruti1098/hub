// Code generated by goa v3.15.2, DO NOT EDIT.
//
// resource HTTP client CLI support package
//
// Command:
// $ goa gen github.com/tektoncd/hub/api/v1/design

package client

import (
	"encoding/json"
	"fmt"
	"strconv"

	resource "github.com/tektoncd/hub/api/v1/gen/resource"
	goa "goa.design/goa/v3/pkg"
)

// BuildQueryPayload builds the payload for the resource Query endpoint from
// CLI flags.
func BuildQueryPayload(resourceQueryName string, resourceQueryCatalogs string, resourceQueryCategories string, resourceQueryKinds string, resourceQueryTags string, resourceQueryPlatforms string, resourceQueryLimit string, resourceQueryMatch string) (*resource.QueryPayload, error) {
	var err error
	var name string
	{
		if resourceQueryName != "" {
			name = resourceQueryName
		}
	}
	var catalogs []string
	{
		if resourceQueryCatalogs != "" {
			err = json.Unmarshal([]byte(resourceQueryCatalogs), &catalogs)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for catalogs, \nerror: %s, \nexample of valid JSON:\n%s", err, "'[\n      \"tekton\",\n      \"openshift\"\n   ]'")
			}
		}
	}
	var categories []string
	{
		if resourceQueryCategories != "" {
			err = json.Unmarshal([]byte(resourceQueryCategories), &categories)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for categories, \nerror: %s, \nexample of valid JSON:\n%s", err, "'[\n      \"build\",\n      \"tools\"\n   ]'")
			}
		}
	}
	var kinds []string
	{
		if resourceQueryKinds != "" {
			err = json.Unmarshal([]byte(resourceQueryKinds), &kinds)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for kinds, \nerror: %s, \nexample of valid JSON:\n%s", err, "'[\n      \"task\",\n      \"pipelines\"\n   ]'")
			}
		}
	}
	var tags []string
	{
		if resourceQueryTags != "" {
			err = json.Unmarshal([]byte(resourceQueryTags), &tags)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for tags, \nerror: %s, \nexample of valid JSON:\n%s", err, "'[\n      \"image\",\n      \"build\"\n   ]'")
			}
		}
	}
	var platforms []string
	{
		if resourceQueryPlatforms != "" {
			err = json.Unmarshal([]byte(resourceQueryPlatforms), &platforms)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for platforms, \nerror: %s, \nexample of valid JSON:\n%s", err, "'[\n      \"linux/s390x\",\n      \"linux/amd64\"\n   ]'")
			}
		}
	}
	var limit uint
	{
		if resourceQueryLimit != "" {
			var v uint64
			v, err = strconv.ParseUint(resourceQueryLimit, 10, strconv.IntSize)
			limit = uint(v)
			if err != nil {
				return nil, fmt.Errorf("invalid value for limit, must be UINT")
			}
		}
	}
	var match string
	{
		if resourceQueryMatch != "" {
			match = resourceQueryMatch
			if !(match == "exact" || match == "contains") {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError("match", match, []any{"exact", "contains"}))
			}
			if err != nil {
				return nil, err
			}
		}
	}
	v := &resource.QueryPayload{}
	v.Name = name
	v.Catalogs = catalogs
	v.Categories = categories
	v.Kinds = kinds
	v.Tags = tags
	v.Platforms = platforms
	v.Limit = limit
	v.Match = match

	return v, nil
}

// BuildListPayload builds the payload for the resource List endpoint from CLI
// flags.
func BuildListPayload(resourceListLimit string) (*resource.ListPayload, error) {
	var err error
	var limit uint
	{
		if resourceListLimit != "" {
			var v uint64
			v, err = strconv.ParseUint(resourceListLimit, 10, strconv.IntSize)
			limit = uint(v)
			if err != nil {
				return nil, fmt.Errorf("invalid value for limit, must be UINT")
			}
		}
	}
	v := &resource.ListPayload{}
	v.Limit = limit

	return v, nil
}

// BuildVersionsByIDPayload builds the payload for the resource VersionsByID
// endpoint from CLI flags.
func BuildVersionsByIDPayload(resourceVersionsByIDID string) (*resource.VersionsByIDPayload, error) {
	var err error
	var id uint
	{
		var v uint64
		v, err = strconv.ParseUint(resourceVersionsByIDID, 10, strconv.IntSize)
		id = uint(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be UINT")
		}
	}
	v := &resource.VersionsByIDPayload{}
	v.ID = id

	return v, nil
}

// BuildByCatalogKindNameVersionPayload builds the payload for the resource
// ByCatalogKindNameVersion endpoint from CLI flags.
func BuildByCatalogKindNameVersionPayload(resourceByCatalogKindNameVersionCatalog string, resourceByCatalogKindNameVersionKind string, resourceByCatalogKindNameVersionName string, resourceByCatalogKindNameVersionVersion string) (*resource.ByCatalogKindNameVersionPayload, error) {
	var err error
	var catalog string
	{
		catalog = resourceByCatalogKindNameVersionCatalog
	}
	var kind string
	{
		kind = resourceByCatalogKindNameVersionKind
		if !(kind == "task" || kind == "pipeline") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("kind", kind, []any{"task", "pipeline"}))
		}
		if err != nil {
			return nil, err
		}
	}
	var name string
	{
		name = resourceByCatalogKindNameVersionName
	}
	var version string
	{
		version = resourceByCatalogKindNameVersionVersion
	}
	v := &resource.ByCatalogKindNameVersionPayload{}
	v.Catalog = catalog
	v.Kind = kind
	v.Name = name
	v.Version = version

	return v, nil
}

// BuildByCatalogKindNameVersionReadmePayload builds the payload for the
// resource ByCatalogKindNameVersionReadme endpoint from CLI flags.
func BuildByCatalogKindNameVersionReadmePayload(resourceByCatalogKindNameVersionReadmeCatalog string, resourceByCatalogKindNameVersionReadmeKind string, resourceByCatalogKindNameVersionReadmeName string, resourceByCatalogKindNameVersionReadmeVersion string) (*resource.ByCatalogKindNameVersionReadmePayload, error) {
	var err error
	var catalog string
	{
		catalog = resourceByCatalogKindNameVersionReadmeCatalog
	}
	var kind string
	{
		kind = resourceByCatalogKindNameVersionReadmeKind
		if !(kind == "task" || kind == "pipeline") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("kind", kind, []any{"task", "pipeline"}))
		}
		if err != nil {
			return nil, err
		}
	}
	var name string
	{
		name = resourceByCatalogKindNameVersionReadmeName
	}
	var version string
	{
		version = resourceByCatalogKindNameVersionReadmeVersion
	}
	v := &resource.ByCatalogKindNameVersionReadmePayload{}
	v.Catalog = catalog
	v.Kind = kind
	v.Name = name
	v.Version = version

	return v, nil
}

// BuildByCatalogKindNameVersionYamlPayload builds the payload for the resource
// ByCatalogKindNameVersionYaml endpoint from CLI flags.
func BuildByCatalogKindNameVersionYamlPayload(resourceByCatalogKindNameVersionYamlCatalog string, resourceByCatalogKindNameVersionYamlKind string, resourceByCatalogKindNameVersionYamlName string, resourceByCatalogKindNameVersionYamlVersion string) (*resource.ByCatalogKindNameVersionYamlPayload, error) {
	var err error
	var catalog string
	{
		catalog = resourceByCatalogKindNameVersionYamlCatalog
	}
	var kind string
	{
		kind = resourceByCatalogKindNameVersionYamlKind
		if !(kind == "task" || kind == "pipeline") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("kind", kind, []any{"task", "pipeline"}))
		}
		if err != nil {
			return nil, err
		}
	}
	var name string
	{
		name = resourceByCatalogKindNameVersionYamlName
	}
	var version string
	{
		version = resourceByCatalogKindNameVersionYamlVersion
	}
	v := &resource.ByCatalogKindNameVersionYamlPayload{}
	v.Catalog = catalog
	v.Kind = kind
	v.Name = name
	v.Version = version

	return v, nil
}

// BuildByVersionIDPayload builds the payload for the resource ByVersionId
// endpoint from CLI flags.
func BuildByVersionIDPayload(resourceByVersionIDVersionID string) (*resource.ByVersionIDPayload, error) {
	var err error
	var versionID uint
	{
		var v uint64
		v, err = strconv.ParseUint(resourceByVersionIDVersionID, 10, strconv.IntSize)
		versionID = uint(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for versionID, must be UINT")
		}
	}
	v := &resource.ByVersionIDPayload{}
	v.VersionID = versionID

	return v, nil
}

// BuildByCatalogKindNamePayload builds the payload for the resource
// ByCatalogKindName endpoint from CLI flags.
func BuildByCatalogKindNamePayload(resourceByCatalogKindNameCatalog string, resourceByCatalogKindNameKind string, resourceByCatalogKindNameName string, resourceByCatalogKindNamePipelinesversion string) (*resource.ByCatalogKindNamePayload, error) {
	var err error
	var catalog string
	{
		catalog = resourceByCatalogKindNameCatalog
	}
	var kind string
	{
		kind = resourceByCatalogKindNameKind
		if !(kind == "task" || kind == "pipeline") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("kind", kind, []any{"task", "pipeline"}))
		}
		if err != nil {
			return nil, err
		}
	}
	var name string
	{
		name = resourceByCatalogKindNameName
	}
	var pipelinesversion *string
	{
		if resourceByCatalogKindNamePipelinesversion != "" {
			pipelinesversion = &resourceByCatalogKindNamePipelinesversion
			err = goa.MergeErrors(err, goa.ValidatePattern("pipelinesversion", *pipelinesversion, "^\\d+(?:\\.\\d+){0,2}$"))
			if err != nil {
				return nil, err
			}
		}
	}
	v := &resource.ByCatalogKindNamePayload{}
	v.Catalog = catalog
	v.Kind = kind
	v.Name = name
	v.Pipelinesversion = pipelinesversion

	return v, nil
}

// BuildByIDPayload builds the payload for the resource ById endpoint from CLI
// flags.
func BuildByIDPayload(resourceByIDID string) (*resource.ByIDPayload, error) {
	var err error
	var id uint
	{
		var v uint64
		v, err = strconv.ParseUint(resourceByIDID, 10, strconv.IntSize)
		id = uint(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be UINT")
		}
	}
	v := &resource.ByIDPayload{}
	v.ID = id

	return v, nil
}

// BuildGetRawYamlByCatalogKindNameVersionPayload builds the payload for the
// resource GetRawYamlByCatalogKindNameVersion endpoint from CLI flags.
func BuildGetRawYamlByCatalogKindNameVersionPayload(resourceGetRawYamlByCatalogKindNameVersionCatalog string, resourceGetRawYamlByCatalogKindNameVersionKind string, resourceGetRawYamlByCatalogKindNameVersionName string, resourceGetRawYamlByCatalogKindNameVersionVersion string) (*resource.GetRawYamlByCatalogKindNameVersionPayload, error) {
	var err error
	var catalog string
	{
		catalog = resourceGetRawYamlByCatalogKindNameVersionCatalog
	}
	var kind string
	{
		kind = resourceGetRawYamlByCatalogKindNameVersionKind
		if !(kind == "task" || kind == "pipeline") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("kind", kind, []any{"task", "pipeline"}))
		}
		if err != nil {
			return nil, err
		}
	}
	var name string
	{
		name = resourceGetRawYamlByCatalogKindNameVersionName
	}
	var version string
	{
		version = resourceGetRawYamlByCatalogKindNameVersionVersion
	}
	v := &resource.GetRawYamlByCatalogKindNameVersionPayload{}
	v.Catalog = catalog
	v.Kind = kind
	v.Name = name
	v.Version = version

	return v, nil
}

// BuildGetLatestRawYamlByCatalogKindNamePayload builds the payload for the
// resource GetLatestRawYamlByCatalogKindName endpoint from CLI flags.
func BuildGetLatestRawYamlByCatalogKindNamePayload(resourceGetLatestRawYamlByCatalogKindNameCatalog string, resourceGetLatestRawYamlByCatalogKindNameKind string, resourceGetLatestRawYamlByCatalogKindNameName string) (*resource.GetLatestRawYamlByCatalogKindNamePayload, error) {
	var err error
	var catalog string
	{
		catalog = resourceGetLatestRawYamlByCatalogKindNameCatalog
	}
	var kind string
	{
		kind = resourceGetLatestRawYamlByCatalogKindNameKind
		if !(kind == "task" || kind == "pipeline") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("kind", kind, []any{"task", "pipeline"}))
		}
		if err != nil {
			return nil, err
		}
	}
	var name string
	{
		name = resourceGetLatestRawYamlByCatalogKindNameName
	}
	v := &resource.GetLatestRawYamlByCatalogKindNamePayload{}
	v.Catalog = catalog
	v.Kind = kind
	v.Name = name

	return v, nil
}
