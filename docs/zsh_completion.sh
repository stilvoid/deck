#compdef _deck deck


function _deck {
  local -a commands

  _arguments -C \
    '(-h --help)'{-h,--help}'[help for deck]' \
    "1: :->cmnds" \
    "*::arg:->args"

  case $state in
  cmnds)
    commands=(
      "clear:Clear the deck"
      "col:Set a button to a colour. The red, green, and blue values must all be 0-255"
      "help:Help about any command"
      "image:Put an image on a button"
      "loop:Output button presses to stdout and read commands from stdin until you press ctrl+c"
      "reset:Rest the deck"
      "text:Put text on a button"
      "wait:Wait for a single button press"
    )
    _describe "command" commands
    ;;
  esac

  case "$words[1]" in
  clear)
    _deck_clear
    ;;
  col)
    _deck_col
    ;;
  help)
    _deck_help
    ;;
  image)
    _deck_image
    ;;
  loop)
    _deck_loop
    ;;
  reset)
    _deck_reset
    ;;
  text)
    _deck_text
    ;;
  wait)
    _deck_wait
    ;;
  esac
}

function _deck_clear {
  _arguments \
    '(-h --help)'{-h,--help}'[help for clear]'
}

function _deck_col {
  _arguments \
    '(-h --help)'{-h,--help}'[help for col]'
}

function _deck_help {
  _arguments
}

function _deck_image {
  _arguments \
    '(-h --help)'{-h,--help}'[help for image]'
}

function _deck_loop {
  _arguments \
    '(-h --help)'{-h,--help}'[help for loop]'
}

function _deck_reset {
  _arguments \
    '(-h --help)'{-h,--help}'[help for reset]'
}

function _deck_text {
  _arguments \
    '(-h --help)'{-h,--help}'[help for text]' \
    '(-i --invert)'{-i,--invert}'[Invert the text (black on white)]'
}

function _deck_wait {
  _arguments \
    '(-h --help)'{-h,--help}'[help for wait]'
}

