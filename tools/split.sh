#!/bin/bash

LDIF_FILE="$1"

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
