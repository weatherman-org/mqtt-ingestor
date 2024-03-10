package mqtt

import (
	paho "github.com/eclipse/paho.mqtt.golang"
)

type Session struct {
	Client    paho.Client
	BrokerUrl string
	Username  string
	Password  string
}

func NewSession(brokerUrl, username, password string) (Session, error) {
	opts := paho.NewClientOptions()
	opts.AddBroker(brokerUrl)
	opts.SetUsername(username)
	opts.SetPassword(password)

	client := paho.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return Session{}, token.Error()
	}

	return Session{
		Client:    client,
		BrokerUrl: brokerUrl,
		Username:  username,
		Password:  password,
	}, nil
}

func (s *Session) Subscribe(topic string, callback paho.MessageHandler) error {
	forever := make(chan bool)
	if token := s.Client.Subscribe(topic, 0, callback); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	<-forever
	return nil
}

func (s *Session) Publish(message []byte, topic string) error {
	token := s.Client.Publish(topic, 0, false, message)
	token.Wait()
	if token.Error() != nil {
		return token.Error()
	}
	return nil
}

func (s *Session) Disconnect(millis uint) {
	s.Client.Disconnect(millis)
}
