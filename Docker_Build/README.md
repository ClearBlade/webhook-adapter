# Docker Image Creation

## Prerequisites

- Building the image requires internet access

### Creating the Docker image for the Webhook Adapter

Clone this repository and execute the following commands to create a docker image for the webhook-adapter adapter:  

- ```GOOS=linux GOARCH=amd64 go build```
- ```cd Docker_Build```
- ```docker build -f Dockerfile -t clearblade_webhook_adapter ..```


# Using the adapter

## Deploying the adapter image

When the docker image has been created, it will need to be saved and imported into the runtime environment. Execute the following steps to save and deploy the adapter image

- On the machine where the ```docker build``` command was executed, execute ```docker save clearblade_webhook_adapter:latest > webhook_adapter.tar``` 

- On the server where docker is running, execute ```docker load -i webhook_adapter.tar```

## Executing the adapter

Once you create the docker image, start the webhook-adapter adapter using the following command:


```docker run -d --name webhook-adapter --network cb-net --restart always clearblade_webhook_adapter -systemKey <YOUR_SYSTEMKEY> -systemSecret <YOUR_SYSTEMSECRET> -platformURL <YOUR_PLATFORMURL> -messagingURL <YOUR_MESSAGINGURL> -deviceName <YOUR_DEVICE_NAME> -activeKey <DEVICE_ACTIVE_KEY> -receiverPort <PORT>```

```
--systemKey The System Key of your System on the ClearBlade Platform
--systemSecret The System Secret of your System on the ClearBlade Platform
--platformURL The address of the ClearBlade Platform (ex. https://platform.clearblade.com)
--messagingURL The MQTT broker address (ex. platform.clearblade.com:1883)
--deviceName The name of a device created on the ClearBlade Platform. Optional, defaults to gcpPubSubAdapter
--activeKey The active key of a device created on the ClearBlade Platform
--receiverPort The port the adapter should listen on
```

Ex.
```docker run --name webhook_adapter clearblade_webhook_adapter -platformURL https://platform.clearblade.com -messagingURL platform.clearblade.com:1883 -deviceName MicroAide_CWR40B -activeKey 01234567890 -receiverPort 80 -systemKey d090d0900ba4afbd90fa95a79858 -systemSecret D090D0900BA4D6D3C9B9EC9DCDD701```

