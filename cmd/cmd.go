package cmd

import (
	"bufio"
	"fmt"
	"image/color"
	"os"
	"strconv"
	"sync"

	"github.com/kballard/go-shellquote"

	"github.com/magicmonkey/go-streamdeck"
	_ "github.com/magicmonkey/go-streamdeck/devices"
	"github.com/spf13/cobra"
)

var deck *streamdeck.Device

func getDeck() *streamdeck.Device {
	if deck == nil {
		var err error
		if deck, err = streamdeck.OpenWithoutReset(); err != nil {
			panic(err)
		}
	}

	return deck
}

var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear the deck",
	Run: func(cmd *cobra.Command, args []string) {
		getDeck().ClearButtons()
	},
}

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Rest the deck",
	Run: func(cmd *cobra.Command, args []string) {
		getDeck().ResetComms()
	},
}

var pollCmd = &cobra.Command{
	Use:   "poll",
	Short: "Poll for button presses until you exit with ctrl+c",
	Run: func(cmd *cobra.Command, args []string) {
		getDeck().ButtonPress(func(btn int, d *streamdeck.Device, err error) {
			if err != nil {
				panic(err)
			}

			fmt.Println(btn)
		})

		var wg sync.WaitGroup
		wg.Add(1)
		wg.Wait()
	},
}

var textInvert = false
var textCmd = &cobra.Command{
	Use:   "text <id> <string>",
	Short: "Put text on a button",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			panic(err)
		}

		fg := color.RGBA{255, 255, 255, 255}
		bg := color.RGBA{0, 0, 0, 255}

		if textInvert {
			fg, bg = bg, fg
		}

		getDeck().WriteTextToButton(id, args[1], fg, bg)
	},
}

var imageCmd = &cobra.Command{
	Use:   "image <id> <filename>",
	Short: "Put an image on a button",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			panic(err)
		}

		if err := getDeck().WriteImageToButton(id, args[1]); err != nil {
			panic(err)
		}
	},
}

var colCmd = &cobra.Command{
	Use:   "col <id> <red> <blue> <green>",
	Short: "Set a button to a colour. The red, green, and blue values must all be 0-255",
	Args:  cobra.ExactArgs(4),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			panic(err)
		}

		red, err := strconv.Atoi(args[1])
		if err != nil {
			panic(err)
		}
		if red < 0 || red > 255 {
			panic("Red value is outside the range 0-255")
		}

		blue, err := strconv.Atoi(args[2])
		if err != nil {
			panic(err)
		}
		if blue < 0 || blue > 255 {
			panic("Blue value is outside the range 0-255")
		}

		green, err := strconv.Atoi(args[3])
		if err != nil {
			panic(err)
		}
		if green < 0 || green > 255 {
			panic("Green value is outside the range 0-255")
		}

		if err := getDeck().WriteColorToButton(id, color.RGBA{uint8(red), uint8(blue), uint8(green), 255}); err != nil {
			panic(err)
		}
	},
}

var parseCmd = &cobra.Command{
	Use:   "parse",
	Short: "Listen on stdin for commands separated by newlines. Each command should be a valid invocation of deck but with the word \"deck\" removed.",
	Run: func(cmd *cobra.Command, args []string) {
		s := bufio.NewScanner(os.Stdin)
		for s.Scan() {
			args, err := shellquote.Split(s.Text())
			if err != nil {
				panic(err)
			}

			switch args[0] {
			case "clear", "reset", "text", "image", "col":
				Root.SetArgs(args)
				Root.Execute()
			default:
				fmt.Fprintf(os.Stderr, "Refusing to parse '%s'\n", args[0])
			}
		}
	},
}

var Root = &cobra.Command{
	Use:  "deck",
	Long: "deck is a CLI for interacting with an Elgato Stream Deck",
}

func init() {
	textCmd.Flags().BoolVarP(&textInvert, "invert", "i", false, "Invert the text (black on white)")

	Root.AddCommand(clearCmd)
	Root.AddCommand(resetCmd)
	Root.AddCommand(pollCmd)
	Root.AddCommand(textCmd)
	Root.AddCommand(imageCmd)
	Root.AddCommand(colCmd)
	Root.AddCommand(parseCmd)
}

func Execute() {
	Root.Execute()
}
