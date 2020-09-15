package cmd

import (
	"errors"
	"fmt"
	"strings"

	ldsview "github.com/kgoins/ldsview/pkg"
)

func buildTitleDomain(domainParts []string) (string, error) {
	domains := []string{}
	for _, domainPart := range domainParts {
		domainSplit := strings.Split(domainPart, "=")
		if len(domainSplit) < 2 {
			return "", errors.New("malformed domain")
		}

		domain := domainSplit[1]
		domains = append(domains, domain)
	}

	return strings.Join(domains, "."), nil
}

func buildTitleObj(objParts []string) (string, error) {
	objComponents := []string{}
	for _, objPart := range objParts {
		objSplit := strings.Split(objPart, "=")
		if len(objSplit) < 2 {
			return "", errors.New("malformed object")
		}

		objComponent := objSplit[1]
		objComponents = append(objComponents, objComponent)
	}

	return strings.Join(objComponents, ", "), nil
}

func BuildTitleLine(entity ldsview.Entity) (string, error) {
	dn, dnFound := entity.GetDN()
	if !dnFound {
		return "", errors.New("Unable to find DN in entity")
	}

	dnParts := strings.Split(dn.Value.GetSingleValue(), ",")

	domainParts := []string{}
	objParts := []string{}

	for _, part := range dnParts {
		if strings.HasPrefix(part, "DC=") {
			domainParts = append(domainParts, part)
		} else {
			objParts = append(objParts, part)
		}
	}

	domain, err := buildTitleDomain(domainParts)
	if err != nil {
		return "", err
	}

	obj, err := buildTitleObj(objParts)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("# %s, %s", obj, domain), nil
}
