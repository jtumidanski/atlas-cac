package consumers

import (
   "atlas-cac/character"
   "atlas-cac/kafka/handler"
   "github.com/sirupsen/logrus"
)

func CreateEventConsumers(l *logrus.Logger) {
   cec := func(topicToken string, emptyEventCreator handler.EmptyEventCreator, processor handler.EventHandler) {
      createEventConsumer(l, topicToken, emptyEventCreator, processor)
   }
   cec("TOPIC_CHARACTER_ATTACK_COMMAND", character.EmptyAttackCommandCreator(), character.HandleAttackCommand())
}

func createEventConsumer(l *logrus.Logger, topicToken string, emptyEventCreator handler.EmptyEventCreator, processor handler.EventHandler) {
   go NewConsumer(l, topicToken, "Character Attack Coordinator Service", emptyEventCreator, processor)
}
