package azp1_prototype_a_res

import (
	"encoding/json"
	"math/rand"
	"time"
)

type PrototypeGenerator struct {
    // game settings
    Width, Height int
    Theme         string
    // game mechanics
    Mechanics []GameMechanic
    // game layout
    Layout [][]GameElement
}

type GameMechanic struct {
    Name        string
    Description string
    Script      string // script to execute when mechanic is triggered
}

type GameElement struct {
    Type        string // e.g. "platform", "obstacle", "coin"
    X, Y        int
    Width, Height int
    Properties   map[string]interface{}
}

func (pg *PrototypeGenerator) Generate() ([]byte, error) {
    // todo: implement game prototype generation logic
    // for now, just return a JSON representation of the generator state
    return json.Marshal(pg)
}

func NewPrototypeGenerator(width, height int, theme string) *PrototypeGenerator {
    return &PrototypeGenerator{
        Width:  width,
        Height: height,
        Theme:  theme,
        Mechanics: []GameMechanic{},
        Layout:   make([][]GameElement, height),
    }
}

func main() {
    rand.Seed(time.Now().UnixNano())
    pg := NewPrototypeGenerator(800, 600, "sci-fi")
    pg.Mechanics = []GameMechanic{
        {
            Name:        "coincollector",
            Description: "Collect coins to score points",
            Script:      "console.log('coin collected!')",
        },
    }
    pg.Layout = [][]GameElement{
        {
            {
                Type:   "platform",
                X:      0,
                Y:      0,
                Width:  100,
                Height: 20,
            },
        },
    }
    json, err := pg.Generate()
    if err != nil {
        panic(err)
    }
    println(string(json))
}