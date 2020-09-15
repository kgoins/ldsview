#!/bin/bash

RELEASE_DIR="release"
BIN="ldsview"

if [ -d "$RELEASE_DIR" ]; then
    rm -r $RELEASE_DIR
fi

mkdir $RELEASE_DIR

BUILDCMD="govvv build -pkg github.com/kgoins/ldsview/internal"
GOOS=darwin ${BUILDCMD} -o "$RELEASE_DIR/${BIN}_macos"
GOOS=linux ${BUILDCMD} -o "$RELEASE_DIR/${BIN}_linux"
