#!/bin/bash

LDIF_FILE="$1"
USAGE="Splits an ldif file into users.ldif, computers.ldif, and groups.ldif\nUsage: split.sh path/to/file.ldif\n"

if [ -z $LDIF_FILE ] || [ "$LDIF_FILE" == "--help" ]; then
    printf "$USAGE"
    exit 1
fi


if [ ! -f $LDIF_FILE ]; then
    echo "$LDIF_FILE not found"
    exit 1
fi

echo "Extracting users"
ldsview -f $LDIF_FILE search "objectClass:=user,objectClass:!=computer" > users.ldif

echo "Extracting computers"
ldsview -f $LDIF_FILE search "objectClass:=computer" > computers.ldif

echo "Extracting groups"
ldsview -f $LDIF_FILE search "objectClass:=group" > groups.ldif
