package ldsview

import (
	"strings"
	"testing"
)

var EntityStr = `
# MYUSR, ContosoUsers, contoso.com
dn: CN=MYUSR,OU=ContosoUsers,DC=contoso,DC=com
objectClass: top
objectClass: person
objectClass: organizationalPerson
objectClass: user
cn: MYUSR
givenName: MYUSR
distinguishedName: CN=MYUSR,OU=ContosoUsers,DC=contoso,DC=com
instanceType: 4
whenCreated: 20120423175240.0Z
whenChanged: 20190225044802.0Z
displayName: MYUSR
uSNCreated: 793245
memberOf: CN=vault_users,OU=Global,OU=Security,OU=Groups,DC=contoso,DC=com
memberOf: CN=PWD Complexity,OU=Security,OU=Groups,DC=contoso,DC=com
uSNChanged: 1076364863
name: MYUSR
objectGUID:: 7OBfD10nQkSVYY8UHCV2aQ==
userAccountControl: 66048
codePage: 0
countryCode: 0
pwdLastSet: 129857191591306845
primaryGroupID: 805306368
objectSid:: AQUAAAAAAAUVAAAAa9ZiBBbA6jKDPStVYiIMAA==
accountExpires: 9223372036854775807
sAMAccountName: MYUSR
sAMAccountType: 805306368
userPrincipalName: MYUSR@contoso.com
lockoutTime: 0
objectCategory: CN=Person,CN=Schema,CN=Configuration,DC=contoso,DC=com
dSCorePropagationData: 20190827201429.0Z
dSCorePropagationData: 20190529140155.0Z
dSCorePropagationData: 20190407190910.0Z
dSCorePropagationData: 20190311163932.0Z
dSCorePropagationData: 16010714223649.0Z
lastLogonTimestamp: 130674899604502606

`

func TestEntity_BuildEntity(t *testing.T) {
	entityLines := strings.Split(EntityStr, "\n")[1:]

	entity := BuildEntity(entityLines)

	if entity.IsEmpty() {
		t.Fatalf("failed to parse entity")
	}

	dn, _ := entity.GetDN()
	if dn.Value.GetSingleValue() != "CN=MYUSR,OU=ContosoUsers,DC=contoso,DC=com" {
		t.Errorf("Failed to parse entity DN")
	}

	if len(entity.Groups()) != 2 {
		t.Errorf("failed to parse entity groups")
	}
}

func TestEntity_BuildEntityWithAttrFilter(t *testing.T) {
	entityLines := strings.Split(EntityStr, "\n")[1:]

	attrsList := []string{
		"samaccountname",
		"userPrincipalName",
		"objectClass",
	}
	attrFilter := BuildAttributeFilter(attrsList)

	entity := BuildEntityFromAttrList(entityLines, &attrFilter)

	if entity.IsEmpty() {
		t.Fatalf("failed to parse entity")
	}

	dn, _ := entity.GetDN()
	if dn.Value.GetSingleValue() != "CN=MYUSR,OU=ContosoUsers,DC=contoso,DC=com" {
		t.Fatalf("Failed to parse entity DN")
	}

	// 3 user specified, plus the 2 dn attributes
	if entity.Size() != 5 {
		t.Fatalf("Failed to parse all attributes")
	}

	objClassAttr, found := entity.GetAttribute("objectClass")
	if !found || objClassAttr.Value.Size() != 4 {
		t.Fatalf("failed to parse multi-valued attribute")
	}
}
