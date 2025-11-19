package common

import (
	"encoding/json"
	"log"

	"app/src/configs"

	"github.com/rabbitmq/amqp091-go"
)

// Publicacion del resultado de la validación al exchange de RabbitMQ
func PublishValidationResult(msg BillingValidationMessage) error {
	rabbitMQConfig := configs.VarConfig().RabbitMQ
	rabbitURL := rabbitMQConfig.URL
	exchange := rabbitMQConfig.Exchange

	conn, err := amqp091.Dial(rabbitURL)
	if err != nil {
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	// Convert JSON
	body, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	// Routing key based on result
	routingKey := ""
	if msg.Resultado == "VALID" {
		routingKey = "invoice.valid"
	} else {
		routingKey = "invoice.invalid"
	}

	// Publish
	err = ch.Publish(
		exchange,
		routingKey,
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		return err
	}

	log.Printf("Mensaje enviado a EXCHANGE='%s' con ROUTING_KEY='%s'", exchange, routingKey)
	return nil
}
