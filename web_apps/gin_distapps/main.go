package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	handlers "github.com/jimsyyap/gin_distapps/handlers"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"os"
)

var recipesHandler *handlers.recipesHandler
