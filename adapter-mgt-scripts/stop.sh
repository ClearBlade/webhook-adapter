#!/bin/bash
ps -ef | grep "webhook-adapter" | grep -v grep | awk '{print $2}' | xargs kill