package config

import (
	"errors"
	"log"
	"os"

	"github.com/blackmesadev/black-mesa/mongodb"
	"github.com/blackmesadev/black-mesa/structs"
	"github.com/blackmesadev/discordgo"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/yaml.v3"
)

var db *mongodb.DB

func LoadFlatConfig() structs.FlatConfig {
	mongo := structs.MongoConfig{
		ConnectionString: os.Getenv("MONGOURI"),
		Username:         os.Getenv("MONGOUSER"),
		Password:         os.Getenv("MONGOPASS"),
	}

	redis := structs.RedisConfig{
		Host: os.Getenv("REDIS"),
	}

	api := structs.APIConfig{
		Host:  os.Getenv("APIHOST"),
		Port:  os.Getenv("APIPORT"),
		Token: os.Getenv("APITOKEN"),
	}

	return structs.FlatConfig{
		Token: os.Getenv("TOKEN"),
		Mongo: mongo,
		Redis: redis,
		API:   api,
	}
}

func LoadLavalinkConfig() structs.LavalinkConfig {
	return structs.LavalinkConfig{
		Host:     os.Getenv("LAVALINKURI"),
		Password: os.Getenv("LAVALINKPASS"),
	}
}

func StartDB(cfg structs.MongoConfig) {
	db = mongodb.InitDB()
	db.ConnectDB(cfg)
}

func GetDB() *mongodb.DB {
	return db
}

func AddGuild(g *discordgo.Guild, invokedByUserID string) *structs.Config {
	config := MakeConfig(g, invokedByUserID)

	db.AddConfig(&mongodb.MongoGuild{
		GuildID: g.ID,
		Config:  config,
	})
	return config
}

func GetConfig(guildid string) (*structs.Config, error) {
	config, err := db.GetConfig(guildid)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			log.Println(err)
		}
		return nil, err
	}

	if config == nil {
		err = errors.New("config is nil")
	}

	return config, err
}

func ExportConfigYAML(guildid string) ([]byte, error) {
	config, err := db.GetConfig(guildid)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			log.Println(err)
		}
		return nil, err
	}

	if config == nil {
		err = errors.New("config is nil")
	}

	return yaml.Marshal(config)
}

func ImportConfigYAML(guildid string, in []byte) error {
	config := &structs.Config{}
	err := yaml.Unmarshal(in, config)
	if err != nil {
		return err
	}

	_, err = db.SetConfig(guildid, config)
	return err
}
