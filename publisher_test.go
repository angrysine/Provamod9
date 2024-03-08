package main

import (
	"math/rand/v2"
	"strconv"
	"testing"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)


func TestPublisher(t *testing.T) {
	opts := MQTT.NewClientOptions().AddBroker("tcp://localhost:3000")
	opts.SetClientID("go_publisher")

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	t.Logf("client.IsConnected() returned value: %v\n", client.IsConnected())

	for {
		count := rand.IntN(2)
		loja := rand.IntN(4)
		tipo := ""
		tipo = DefineType(&tipo,count)
		temperatura := DefineTemperature(rand.IntN(35))
		id := "lj" +"0"+strconv.Itoa(loja) +tipo
		text := strconv.Itoa(loja) +","+tipo+","+strconv.Itoa(temperatura)+","+id

		t.Log(text)
		token := client.Publish("temperatura/topic", 1, true, text)
		// fmt.Printf("Publicado: %s\n", text)
		token.Wait()
		time.Sleep(100 * time.Millisecond)	
	}
}


func DefineType(pointer *string, count int) string{
	tipo :=""
	if count ==0{
		tipo = "Geladeira" 
		} else {
			tipo= "Freezer"
		}
	return tipo
}

func DefineTemperature(temperature int) int {
	value := rand.IntN(2)
	if value ==0 {
		return temperature
	}else {
		return -temperature
	}
}