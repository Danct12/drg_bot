package main

import (
  "context"
  "log"
  "math/rand"
  "os"
  "strings"

  "github.com/diamondburned/arikawa/v3/gateway"
  "github.com/diamondburned/arikawa/v3/session"
)

var rockandstone = [...]string{
  "Come on guys! Rock and Stone!",
  "If you don't Rock and Stone, you ain't comin' home!",
  "We fight for Rock and Stone!",
  "Did I hear a Rock and Stone?",
  "Rock and Stone to the Bone!",
  "That's it lads! Rock and Stone!",
  "Rock and Stone, Brother!",
  "Rock and Stone like there's no tomorrow!",
  "For those about to Rock and Stone, we salute you!",
  "Rock and Stone... Yeeaaahhh!",
}

var bismor = [...]string{
  "Bismor!",
  "We got Bismor here!",
  "Bismor! It feels so good to say it!",
  "I'm so glad to announce that I found some Bismor!",
  "Bismor here!",
  "Bismor, Bismor, Bismor!",
  "BISMOOOOR!",
  "A bit more Bismor! Over here!",
  "I found some Bismor!",
}

func main() {
  var token = os.Getenv("BOT_TOKEN")
  if token == "" {
    log.Fatalln("No BOT_TOKEN specified.")
  os.Exit(1)
  }

  s := session.New("Bot " + token)

  // Add the needed Gateway intents.
  s.AddIntents(gateway.IntentGuildMessages)

  if err := s.Open(context.Background()); err != nil {
    log.Fatalln("Failed to connect:", err)
  }
  defer s.Close()

  u, err := s.Me()
  if err != nil {
    log.Fatalln("Failed to get myself:", err)
  }

  s.AddHandler(func(c *gateway.MessageCreateEvent) {
    if u.ID == c.Author.ID {
      return
    }

    var lower_msg = strings.ToLower(c.Content)

    if strings.Contains(lower_msg, "rock and stone") {
      log.Println(c.Author.Username, "mentioned ROCK AND STONE!")
      s.SendMessage(c.ChannelID,
        ":pick: **" + strings.ToUpper(rockandstone[rand.Intn(len(rockandstone)-1)]) + "** :pick:")
    }

    if strings.Contains(lower_msg, "bismor") {
      log.Println(c.Author.Username, "mentioned BISMOR!")
      s.SendMessage(c.ChannelID,
        "**" + strings.ToUpper(bismor[rand.Intn(len(rockandstone)-1)]) + "**")
    }
  })

  log.Println("Started as", u.Username)

  log.Print("Bot can be invited using the following URL: https://discord.com/api/oauth2/authorize?client_id=", u.ID, "&scope=bot\n")

  // Block forever.
  select {}
}