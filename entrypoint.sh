#!/bin/sh -l
set -x

/piper mavenExecute --defines -Dmaven.test.skip=true --goals package
