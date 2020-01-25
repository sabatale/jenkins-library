#!/bin/sh -l
set -x

find /

/piper mavenExecute --defines -Dmaven.test.skip=true --goals package
