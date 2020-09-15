package internal

import (
	"fmt"
	"strings"
)

type LdapsearchCmdOptions struct {
	host     string
	port     int
	domainDN string
	filter   string
	user     string
	password string
}

func NewLdapsearchCmdOptions(host string, domainDN string, user string, password string) LdapsearchCmdOptions {
	return LdapsearchCmdOptions{
		host:     host,
		port:     389,
		domainDN: domainDN,
		user:     user,
		password: password,
		filter:   "(objectClass=*)",
	}
}

func BuildLdapsearchCmd(options LdapsearchCmdOptions) string {
	var builder strings.Builder

	builder.WriteString("ldapsearch \\ \n")
	builder.WriteString("	-o ldif-wrap=no \\ \n")
	builder.WriteString("	-E pr=1000/noprompt \\ \n")

	builder.WriteString(fmt.Sprintf("	-h %s \\ \n", options.host))
	builder.WriteString(fmt.Sprintf("	-b %s \\ \n", options.domainDN))
	builder.WriteString(fmt.Sprintf("	-D '%s' \\ \n", options.user))
	builder.WriteString(fmt.Sprintf("	-w %s \\ \n", options.password))
	builder.WriteString(fmt.Sprintf("	\"%s\" \n", options.filter))

	return builder.String()
}
