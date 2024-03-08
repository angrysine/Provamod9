package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

var count =0

var msgList []string = []string{}
var Time = time.Now().UnixNano() / 1000000
var Qos string
var MessageRate float64
var ExecutionTime string
var Mensagens []string 
var messagePubHandler MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {

	count++
	msgList = append(msgList, string(msg.Payload()[:]), ",")


	if count == 10 {
		
		for _,msg2 := range msgList {
			msgList =append(msgList,fmt.Sprintf("%v",msg.Qos()))
			value := strings.Split(msg2, ",")
			valueint,_  :=strconv.ParseInt(value[2],10,64)
			if (value[1] =="Freezer") && ((valueint< -25))  {
				msgList = append(msgList,string("[ALERTA: Temperatura BAIXA"))
			} else if (value[1] =="Freezer") && ((-15< valueint) )  {
				msgList = append(msgList,string("[ALERTA: Temperatura ALTA"))
			}
			if (value[1] =="Geladeira") && ((valueint< 2))  {
				msgList = append(msgList,string("[ALERTA: Temperatura BAIXA"))
			} else if (value[1] =="Geladeira") && ((10< valueint) )  {
				msgList = append(msgList,string("[ALERTA: Temperatura ALTA"))
			}
		}
		
		
		
		for index,_ := range msgList {
			Writer("./output.txt",fmt.Sprintf("%v",msgList[index]))
		}
		
		defer client.Disconnect(250)
	}
}


func TestSubscriber(t *testing.T) {
	
	opts := MQTT.NewClientOptions().AddBroker("tcp://localhost:3000")
	opts.SetClientID("go_subscriber")
	opts.SetDefaultPublishHandler(messagePubHandler)

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	if token := client.Subscribe("test/topic", 1, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		return
	}
	select {}
	// Bloqueia indefinidamente
}


func Writer(filepath string,text string) {
    // If the file doesn't exist, create it, or append to the file
    f, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal(err)
    }
    if _, err := f.Write([]byte(text +"\n")); err != nil {
        log.Fatal(err)
    }
    if err := f.Close(); err != nil {
        log.Fatal(err)
    }
}