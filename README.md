sudo ./webhook-adapter -platformURL https://platform.clearblade.com:443 -messagingURL platform.clearblade.com:1883 -deviceName slackAdapter -activeKey bH11570h39kk63OyTXf -receiverPort 80 -systemKey bce281e40c0b29f88fa01 -systemSecret BCE281E40BF09BA7EE4D > webhook-adapter.log

# webhook-adapter Adapter

The __webhook-adapter__ adapter provides the ability for the ClearBlade platform to receive incoming HTTP requests via MQTT.

# MQTT Topic Structure
The __webhook-adapter__ adapter utilizes MQTT messaging to communicate with the ClearBlade Platform. The __webhook-adapter__ adapter will publish messages to a MQTT topic in order to provide the ClearBlade Platform/Edge with data received from a HTTP request. The topic structures utilized by the __webhook-adapter__ adapter is as follows:

  * Send HTTP request data to Clearblade: __webhook-adapter/received__ (default topic)


## ClearBlade Platform Dependencies
The __webhook-adapter__ adapter was constructed to provide the ability to communicate with a _System_ defined in a ClearBlade Platform instance. Therefore, the adapter requires a _System_ to have been created within a ClearBlade Platform instance.

Once a System has been created, artifacts must be defined within the ClearBlade Platform system to allow the adapters to function properly. At a minimum: 

  * A device needs to be created in the Auth --> Devices collection. The device will represent the adapter account. The _name_ and _active key_ values specified in the Auth --> Devices collection will be used by the adapter to authenticate to the ClearBlade Platform or ClearBlade Edge. 

## Usage

### Executing the adapter

`./webhook-adapter -systemKey <SYSTEM_KEY> -systemSecret <SYSTEM_SECRET> -platformURL <PLATFORM_URL> -messagingURL <MESSAGING_URL> -deviceName <DEVICE_NAME> -activeKey <DEVICE_ACTIVE_KEY> -receiverPort <LISTENING_PORT> `

   __*Where*__ 

   __systemKey__
  * REQUIRED
  * The system key of the ClearBLade Platform __System__ the adapter will connect to

   __systemSecret__
  * REQUIRED
  * The system secret of the ClearBLade Platform __System__ the adapter will connect to

   __platformURL__
  * REQUIRED
  * The url of the ClearBlade Platform instance the adapter will connect to

   __messagingURL__
  * REQUIRED
  * The MQTT url of the ClearBlade Platform instance the adapter will connect to

   __deviceName__
  * REQUIRED 
  * The device name the adapter will use to authenticate to the ClearBlade Platform
  * Requires the device to have been defined in the _Auth - Devices_ collection within the ClearBlade Platform __System__
   
   __activeKey__
  * REQUIRED
  * The active key the adapter will use to authenticate to the platform
  * Requires the device to have been defined in the _Auth - Devices_ collection within the ClearBlade Platform __System__

   __receiverPort__
  * REQUIRED 
  * The port the adapter will listen on

   __topicName__
  * The MQTT topic the adapter will publish to
  * OPTIONAL
  * Defaults to __webhook-adapter/received__

   __enableTLS__
  * Whether or not the adapter should utilize TLS
  * OPTIONAL
  * Defaults to __false__

   __tlsCertPath__
  * The path to the TLS .crt file
  * REQUIRED if __enableTLS__ is set to __true__

   __tlsKeyPath__
  * The path to the TLS .key file
  * REQUIRED if __enableTLS__ is set to __true__


## Setup
---
The __webhook-adapter__ adapter is dependent upon the ClearBlade Go SDK and its dependent libraries being installed. The __webhook-adapter__ adapter was written in Go and therefore requires Go to be installed (https://golang.org/doc/install).


### Adapter compilation
In order to compile the adapter for execution, the following steps need to be performed:

 1. Retrieve the adapter source code  
    * ```git clone git@github.com:ClearBlade/webhook-adapter.git```
 2. Navigate to the __webhook-adapter__ directory  
    * ```cd webhook-adapter```
 3. ```go get -u github.com/ClearBlade/Go-SDK.git```
    * This command should be executed from within your Go workspace
 4. Compile the adapter
    * ```go build```



