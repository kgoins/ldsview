# ldsview
Offline search tool for LDAP directory dumps in LDIF format.

## Features
  * Fast and memory efficient parsing of LDIF files
  * Build `ldapsearch` commands to extract an LDIF from a directory
  * Show directory structure
  * UAC and directory time format translation

## Config
Config options can be passed as CLI flags, environment variables, or via a config file courtsey of [viper](https://github.com/spf13/viper). Reference the project's documentation for all of the different ways you can supply configuration. 
  * By default, `ldsview` will look for a file called `.ldsview.{json,toml,yaml}` in the user's home directory
  * Environment variables with a prefix of `LDSVIEW` will be read in by the application

## Usage
Detailed usage information is available via the `--help` flag or the `help` command for `ldsview` and all subcommands.

### Search Syntax
`ldsview`'s search mechanism is based on the [entityfilter](https://github.com/kgoins/entityfilter) project. Detailed information about search filter syntax can be found in that project's [README](https://github.com/kgoins/entityfilter/blob/master/README.md).

### Examples
  * Build `ldapsearch` command to extract LDIF files from a directory: `ldsview cmdbuilder`
    * The command will prompt you for any information needed
    * Have the following ready: 
      * Directory host FQDN or IP
      * Domain DN
      * User to run as
      * User's password
  * Quickly find a specific entity in an LDIF file: `ldsview -f myfile.ldif entity myuser`
  * Parse UAC flag from AD: `ldsview uac 532480`
  * Search LDIF file: `ldsview -f myfile.ldif search "adminCount:=1,sAMAccountName:!=krbtgt"`
    * This command will return all entities with an `adminCount` of 1 that are not `krbtgt`
    * `-i` can be used to limit which attributes are returned from matching entities
    * `--tdc` will translate directory timestamps into a human readable format


### Tools Directory

Additional tools and utilities for managing LDIFs:

***Makefile***: Place the Makefile in the same directory as your exported LDIF and run make.

```sh
>> make -j9 LDIF=./my.domain.ldif
```
This will split and create the following default LDIFs:

* users.ldif
* computers.ldif
* groups.ldif
* domain_admin.ldif
* poss_svc_accnts.ldif
* pass_not_reqd.ldif
* pass_cant_change.ldif
* users_dont_expire.ldif
* trusted_4_delegation.ldif
* preauth_not_reqd.ldif
* password_expired.ldif
* trust2auth4delegation.ldif
