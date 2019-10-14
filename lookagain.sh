#!/bin/bash
#find -name "*.sh" | sed 's/.//' | sed 's/\///'  |sed 's/.sh//g'
find . -name '*.sh' | cut -f2 -d '.' | sed  's/test//g' | cut -f2 -d 'h' |sed 's#/##g' 
