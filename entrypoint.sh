#!/bin/sh -l
set -x

id


ls -la /bin/piper


/bin/piper mavenExecute --defines -Dmaven.test.skip=true --goals package
