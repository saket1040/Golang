package main

import "fmt"

type Song struct {
	Title    string
	Artist   string
	Duration int64
}

type Playlist struct {
	Songs map[string]*Song
}

func (p *Playlist) AddSong(song *Song) {
	if _, exists := p.Songs[song.Title]; exists {
		fmt.Println("Song already exists")
		return
	}

	p.Songs[song.Title] = song
}

func (p *Playlist) RemoveSong(title string) {
	if _, exists := p.Songs[title]; !exists {
		fmt.Println("Song doesn't exist")
		return
	}

	delete(p.Songs, title)
}

func (p *Playlist) TotalDurations() int64 {
	var res int64 = 0
	for _, val := range p.Songs {
		res += val.Duration
	}
	return res
}

func main() {
	playlist := &Playlist{Songs: make(map[string]*Song)}

	song1 := &Song{Title: "Song One", Artist: "Artist A", Duration: 300}
	song2 := &Song{Title: "Song Two", Artist: "Artist B", Duration: 200}

	playlist.AddSong(song1)
	playlist.AddSong(song2)
	playlist.AddSong(song1) // Attempt to add duplicate

	fmt.Println("Total Duration:", playlist.TotalDurations())

	playlist.RemoveSong("Song One")
	fmt.Println("Total Duration after removal:", playlist.TotalDurations())

	playlist.RemoveSong("Nonexistent Song") // Attempt to remove non-existent song
}