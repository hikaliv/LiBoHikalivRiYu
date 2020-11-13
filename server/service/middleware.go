package service

import (
	dtype "RiYu/data/type"
	"RiYu/server/grammar/tense"
	"context"
	"math/rand"
	"net/url"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func getone() *dtype.Words {
	return getoneWith(words)
}

func getoneWith(words [][]*dtype.Words) *dtype.Words {
	rand.Seed(time.Now().UnixNano())
	itor := rand.Intn(len(words))
	sheet := words[itor]
	itor = rand.Intn(len(sheet))
	return sheet[itor]
}

func getoneTense() (forms *tense.Forms, err error) {
	word := getoneWith([][]*dtype.Words{verbs, adjs})
	for _, label := range word.Label {
		if label == "形" || label == "形动" || label == "五段" || label == "一段" || label == "サ变" || label == "カ变" {
			forms, err = tense.GetTenseForms(word.Ri, label)
			break
		}
	}
	return
}

func search(form url.Values) (words []*dtype.Words, err error) {
	filter := make(bson.M, len(form))
	for key, value := range form {
		filter[key] = bson.M{"$regex": value[0]}
	}
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return
	}
	err = cursor.All(context.Background(), &words)
	return
}
