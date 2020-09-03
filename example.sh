#!/bin/bash

# Blank out all the buttons
deck clear

# Put numbers on all the buttons 1 to 15
for i in {0..14}; do
    deck text $i $((i+1))
done

# Reset the deck when we exit
function cleanup() {
    deck reset
}
trap cleanup EXIT

deck poll | while read button; do
    case $button in
    *)
        # Print out which button got pressed
        echo "You pressed button $((button+1))"
        ;;
    esac
done
