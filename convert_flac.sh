#!/bin/bash
mkdir -p mp3 m4a

for file in *.flac; do
    filename=$(basename "$file" .flac)
    ffmpeg -i "$file" -codec:a libmp3lame -qscale:a 2 "mp3/${filename}.mp3"
    ffmpeg -i "$file" -codec:a aac -b:a 192k "m4a/${filename}.m4a"
done