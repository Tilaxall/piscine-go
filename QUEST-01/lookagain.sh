#!/bin/bash
#find -name "*.sh" | sed 's/.//' | sed 's/\///'  |sed 's/.sh//g'
find . -name '*.sh' | sed 's/^.*\///' | sed 's/.sh//g'
