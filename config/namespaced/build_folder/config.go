package build_folder

import (
	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/pkg/errors"
)

const (
	// ErrFmtNoAttribute is an error string for not-found attributes
	ErrFmtNoAttribute = `"attribute not found: %s`
	// ErrFmtUnexpectedType is an error string for attribute map values of unexpected type
	ErrFmtUnexpectedType = `unexpected type for attribute %s: Expecting a string`
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azuredevops_build_folder", func(r *config.Resource) {
		r.ShortGroup = "pipelines"
		r.ExternalName.GetExternalNameFn = GetNameFromProjectID
	})
}

// GetNameFromFullyQualifiedID extracts external-name from Azure ID
// using the "id" attribute.
// Examples of fully qualifiers:
// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.DBforMySQL/servers/server1
// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/networkInterfaces/nic1
func GetNameFromProjectID(tfstate map[string]interface{}) (string, error) {
	idStr, err := GetAttributeValue(tfstate, "project_id")
	if err != nil {
		return "", err
	}
	pathStr, err := GetAttributeValue(tfstate, "path")
	if err != nil {
		return "", err
	}
	res := idStr + pathStr
	return res, nil
}

// GetAttributeValue reads a string attribute from the specified map
func GetAttributeValue(attrMap map[string]interface{}, attr string) (string, error) {
	v, ok := attrMap[attr]
	if !ok {
		return "", errors.Errorf(ErrFmtNoAttribute, attr)
	}
	vStr, ok := v.(string)
	if !ok {
		return "", errors.Errorf(ErrFmtUnexpectedType, attr)
	}
	return vStr, nil
}
