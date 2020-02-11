package main

import (
	"encoding/json"
	"flag"
	cb "github.com/clearblade/Go-SDK"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

var (
	platURL      string
	messURL      string
	sysKey       string
	sysSec       string
	deviceName   string
	activeKey    string
	listenPort   string
	topicName    string
	enableTLS    bool
	tlsCertPath  string
	tlsKeyPath   string
	deviceClient *cb.DeviceClient
	InboundURL string
)

func init() {
	flag.StringVar(&sysKey, "systemKey", "", "system key (required)")
	flag.StringVar(&sysSec, "systemSecret", "", "system secret (required)")
	flag.StringVar(&platURL, "platformURL", "", "platform url (required)")
	flag.StringVar(&messURL, "messagingURL", "", "messaging URL")
	flag.StringVar(&deviceName, "deviceName", "", "name of device (required)")
	flag.StringVar(&activeKey, "activeKey", "", "active key (password) for device (required)")
	flag.StringVar(&listenPort, "receiverPort", "", "receiver port for adapter (required)")
	flag.StringVar(&topicName, "topicName", "webhook-adapter/received", "topic name to publish received HTTP requests to (defaults to webhook-adapter/received)")
	flag.BoolVar(&enableTLS, "enableTLS", false, "enable TLS on http listener (must provide tlsCertPath and tlsKeyPath params if enabled)")
	flag.StringVar(&tlsCertPath, "tlsCertPath", "", "path to TLS .crt file (required if enableTLS flag is set)")
	flag.StringVar(&tlsKeyPath, "tlsKeyPath", "", "path to TLS .key file (required if enableTLS flag is set)")
	flag.StringVar(&InboundURL, "inboundURL", "/", "URL Path for inbound webhook URL, ex /abcdef/endpoint1")
}

type requestJSONBody struct {
	RequestURL   string      `json:"request_url"`
	URLParams    url.Values  `json:"url_params"`
	Headers      http.Header `json:"headers"`
	Method       string      `json:"method"`
	Body         interface{} `json:"body"`
	TimeReceived string      `json:"time_received"`
}

type requestStringBody struct {
	RequestURL   string      `json:"request_url"`
	URLParams    url.Values  `json:"url_params"`
	Headers      http.Header `json:"headers"`
	Method       string      `json:"method"`
	Body         string      `json:"body"`
	TimeReceived string      `json:"time_received"`
}

func usage() {
	log.Printf("Usage: webhook-adapter [options]\n\n")
	flag.PrintDefaults()
}

func handleRequest(rw http.ResponseWriter, r *http.Request) {
	timeReceived := time.Now().UTC().Format(time.RFC3339)

	qp := r.URL.Query()

	log.Println("Received a http request!")
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("error reading body of request: %s\n", err.Error())
		return
	}

	var bodyJSON interface{}
	var b []byte

	if err := json.Unmarshal(body, &bodyJSON); err != nil {
		log.Println("Unable to parse body into JSON object, sending as raw string instead.")
		msg := &requestStringBody{
			RequestURL:   r.Host + r.RequestURI,
			URLParams:    qp,
			Headers:      r.Header,
			Method:       r.Method,
			Body:         string(body),
			TimeReceived: timeReceived,
		}

		b, err = json.Marshal(msg)
		if err != nil {
			log.Printf("Failed to convert request structure into a string: %s\n", err.Error())
			return
		}
	} else {
		msg := &requestJSONBody{
			RequestURL:   r.Host + r.RequestURI,
			URLParams:    qp,
			Headers:      r.Header,
			Method:       r.Method,
			Body:         bodyJSON,
			TimeReceived: timeReceived,
		}

		b, err = json.Marshal(msg)
		if err != nil {
			log.Printf("Failed to convert request structure into a string: %s\n", err.Error())
			return
		}
	}

	if err := deviceClient.Publish(topicName, b, 2); err != nil {
		log.Printf("Unable to publish request: %s\n", err.Error())
		return
	}

}

func validateFlags() {
	flag.Parse()
	if sysKey == "" || sysSec == "" || platURL == "" || deviceName == "" || activeKey == "" || listenPort == "" {
		log.Printf("Missing required flags\n\n")
		flag.Usage()
		os.Exit(1)
	}

	if _, err := strconv.Atoi(listenPort); err != nil {
		log.Printf("receiverPort must be numeric\n\n")
		flag.Usage()
		os.Exit(1)
	}

	if enableTLS && (tlsCertPath == "" || tlsKeyPath == "") {
		log.Printf("tlsCertPath and tlsKeyPath are required if TLS is enabled\n")
		flag.Usage()
		os.Exit(1)
	}

}

func main() {
	flag.Usage = usage
	validateFlags()

	deviceClient = cb.NewDeviceClient(sysKey, sysSec, deviceName, activeKey)

	if platURL != "" {
		log.Println("Setting custom platform URL to: ", platURL)
		deviceClient.HttpAddr = platURL
	}

	if messURL != "" {
		log.Println("Setting custom messaging URL to: ", messURL)
		deviceClient.MqttAddr = messURL
	}

	log.Println("Authenticating to platform with device: ", deviceName)

	if err := deviceClient.Authenticate(); err != nil {
		log.Fatalf("Error authenticating: %s\n", err.Error())
	}

	if err := deviceClient.InitializeMQTT("webhookadapter_"+deviceName, "", 30, nil, nil); err != nil {
		log.Fatalf("Unable to initialize MQTT: %s\n", err.Error())
	}
	log.Printf("MQTT connected and adapter about to listen on port: %s\n", listenPort)

	http.HandleFunc(InboundURL, handleRequest)

	if enableTLS {
		log.Fatal(http.ListenAndServeTLS(":"+listenPort, tlsCertPath, tlsKeyPath, nil))
	} else {
		log.Fatal(http.ListenAndServe(":"+listenPort, nil))
	}
}
