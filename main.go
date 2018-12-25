package main

import (
  "fmt"
  "os"

  "github.com/urfave/cli"
)

// Mapping
var shogunNameMap = map[string]string{
  "1": "徳川家康",
  "2": "徳川秀忠",
  "3": "徳川家光",
  "4": "徳川家綱",
  "5": "徳川綱吉",
  "6": "徳川家宣",
  "7": "徳川家継",
  "8": "徳川吉宗",
  "9": "徳川家重",
  "10": "徳川家治",
  "11": "徳川家斉",
  "12": "徳川家慶",
  "13": "徳川家定",
  "14": "徳川家茂",
  "15": "徳川慶喜",
}

func main() {
  app := cli.NewApp()

  app.Name = "go-edo"
  app.Usage = "Display shogun Name in Edo Bakufu"
  app.Version = "0.0.1"

  app.Action = func (c *cli.Context) error {
    str := c.Args().Get(0)
    fmt.Println(shogunNameMap[str])
    return nil
  }

  app.Run(os.Args)
}
