## Exportify to M3U
A simple piece of code, wherever you run it it will turn every CSV file in that folder into a .m3u file with the same name

### Use Case
I used this to covert my spotify playlists from exportify into m3u files so I could own my playlists again. I was then able to add them to rhythmbox and use them portably.

### Info
The m3u files will contain "$ARTISTNAME - $TRACKNAME.m4a" which works for me, as my library is in one giant folder.
The code is easily adaptable for other usecases, however if your audio files are in their own directories this is outside the scope of this code.

### To run
No dependencies to install, just put your CSV files in the same folder as the go file, make sure go is installed, and run `go run exportify-to-m3u.go` Then copy your outputted m3u files into the same directory as your audio files, and enjoy having your playlists back.

Exportify is avaliable at https://watsonbox.github.io/exportify/
