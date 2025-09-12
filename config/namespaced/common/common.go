package common

import "github.com/pkg/errors"

const (
	// ErrFmtNoAttribute is an error string for not-found attributes
	ErrFmtNoAttribute = `"attribute not found: %s`
	// ErrFmtUnexpectedType is an error string for attribute map values of unexpected type
	ErrFmtUnexpectedType = `unexpected type for attribute %s: Expecting a string`
)

// GetNameFromProjectID get a projectID/resourceID
func GetNameFromProjectID(tfstate map[string]interface{}) (string, error) {
	idStr, err := GetAttributeValue(tfstate, "project_id")
	if err != nil {
		return "", err
	}
	pathStr, err := GetAttributeValue(tfstate, "id")
	if err != nil {
		return "", err
	}
	res := idStr + "/" + pathStr
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
