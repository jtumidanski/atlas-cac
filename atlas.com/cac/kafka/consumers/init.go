package consumers

import (
	"atlas-cac/character"
	"atlas-cac/kafka/handler"
	"context"
	"github.com/sirupsen/logrus"
	"sync"
)

const (
	CharacterAttackCommand = "character_attack_command"
)

func CreateEventConsumers(l *logrus.Logger, ctx context.Context, wg *sync.WaitGroup) {
	cec := func(topicToken string, name string, emptyEventCreator handler.EmptyEventCreator, processor handler.EventHandler) {
		createEventConsumer(l, ctx, wg, name, topicToken, emptyEventCreator, processor)
	}
	cec("TOPIC_CHARACTER_ATTACK_COMMAND", CharacterAttackCommand, character.EmptyAttackCommandCreator(), character.HandleAttackCommand())
}

func createEventConsumer(l *logrus.Logger, ctx context.Context, wg *sync.WaitGroup, name string, topicToken string, emptyEventCreator handler.EmptyEventCreator, processor handler.EventHandler) {
	wg.Add(1)
	go NewConsumer(l, ctx, wg, name, topicToken, "Character Attack Coordinator Service", emptyEventCreator, processor)
}
