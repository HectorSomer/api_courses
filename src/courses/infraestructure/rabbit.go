package infraestructure

import (
    "context"
    "encoding/json"
    "event_driven/src/config"
    "event_driven/src/courses/application/models"
    "log"

    amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQPublisher struct {
    conn *amqp.Connection
}

func NewRabbitMQPublisher() *RabbitMQPublisher {
    conn := config.GetRabbitMQConnection()
    return &RabbitMQPublisher{conn}
}

func (p *RabbitMQPublisher) Publish(ctx context.Context, event models.CourseRegistredEvent) (*models.CourseRegistredEvent, error) {
    ch, err := p.conn.Channel()
    if err != nil {
        return nil, err
    }
    defer ch.Close()

    // Declarar el intercambio `direct`
    err = ch.ExchangeDeclare(
        "courses", // nombre del intercambio
        "direct",  // tipo de intercambio
        true,      // duradero
        false,     // autodelete
        false,     // interno
        false,     // noWait
        nil,       // argumentos
    )
    if err != nil {
        return nil, err
    }

    courseEvent := models.CourseRegistredEvent{
        PersonEmit:    event.PersonEmit,
        Message:       event.Message,
        IDUserTeacher: event.IDUserTeacher,
    }

    body, err := json.Marshal(courseEvent)
    if err != nil {
        return nil, err
    }

    // Publicar el mensaje en el intercambio `direct`
    err = ch.PublishWithContext(ctx,
        "courses",    // nombre del intercambio
        "course_key", // clave de enrutamiento
        false,        // mandatory
        false,        // immediate
        amqp.Publishing{
            ContentType: "application/json",
            Body:        body,
        })
    if err != nil {
        return nil, err
    }

    log.Printf("[x] Sent %s", body)
    return &event, nil
}
