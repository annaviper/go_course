package main

import (
	"os"
	"testing"
)

func Test_newDeck(t *testing.T) {
	d := newDeck()

	result_int := len(d)
	if result_int != 16 {
		t.Errorf("Expected deck length to be 20, got %v", result_int)
	}

	firstValue := "Ace of Clubs"
	result := d[0]
	if result != firstValue {
		t.Errorf("Expected first card to be %v, got %v", firstValue, result)
	}

	lastValue := "Four of Spades"
	result = d[len(d)-1]
	if result != lastValue {
		t.Errorf("Expected last card to be %v, got %v", lastValue, result)
	}
}

func Test_saveToFileAndNewDeckFromFile(t *testing.T) {
	filename := "_decktesting"
	os.Remove(filename)

	d := newDeck()
	d.saveToFile(filename)

	loadedDeck := newDeckFromFile(filename)
	expected := 16
	result := len(loadedDeck)
	if result != expected {
		t.Errorf("Expected deck length to be %v, got %v", expected, result)
	}

	os.Remove(filename)
}

// func Test_newDeck(t *testing.T) {
// 	tests := []struct {
// 		name string
// 		want deck
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := newDeck(); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("newDeck() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func Test_deck_saveToFile(t *testing.T) {
// 	type args struct {
// 		filename string
// 	}
// 	tests := []struct {
// 		name    string
// 		d       deck
// 		args    args
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if err := tt.d.saveToFile(tt.args.filename); (err != nil) != tt.wantErr {
// 				t.Errorf("deck.saveToFile() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }
