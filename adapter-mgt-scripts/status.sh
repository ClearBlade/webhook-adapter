#!/bin/bash

# If this script is executed, we know the adapter has been deployed. No need to test for that.
STATUS="Deployed"

if [[ $(ps -ef | grep 'webhook-adapter-amd64 ' | grep -v grep) ]]; then
    STATUS="Running"
else
    STATUS="Stopped"
fi

echo $STATUS