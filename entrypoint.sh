#!/bin/sh -l
set -x

id


ls -la /bin/piper

find . -name pipeline_config.yml
find . -name pom.xml

/bin/piper mavenExecute --defines -Dmaven.test.skip=true --goals package
