#!/bin/bash

# Required
# Param 1 = platform URL
# Param 2 = messaging URL
# Param 3 = device name
# Param 4 = device active key
# Param 5 = listen port
# Param 6 = system key
# Param 7 = system secret

# Optional (if using TLS, both 8 and 9 are required)
# Param 8 = TLS Cert path
# Param 9 = TLS Key path

if [ "$#" -eq 7 ]; then
    nohup ./webhook-adapter-amd64 -platformURL "$1" -messagingURL "$2" -deviceName "$3" -activeKey "$4" -receiverPort "$5" -systemKey "$6" -systemSecret "$7" > webhook-adapter.log 2>&1 &
elif [ "$#" -eq 9 ]; then
    nohup ./webhook-adapter-amd64 -platformURL "$1" -messagingURL "$2" -deviceName "$3" -activeKey "$4" -receiverPort "$5" -systemKey "$6" -systemSecret "$7" -enableTLS -tlsCertPath "$8" -tlsKeyPath "$9" > webhook-adapter.log 2>&1 &
else
    echo "Unexpected number of parameters, 7 or 9 are expected"
fi
