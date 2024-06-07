package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Map to store counts of different file types
var fileTypeCounts = map[string]int{
	"total": 0,
	"mp3":   0,
	"aac":   0,
	"m4a":   0,
	"wav":   0,
	"flac":  0,
	"ogg":   0,
	"aiff":  0,
	"alac":  0,
	"mp4":   0,
	"m4v":   0,
	"mov":   0,
	"avi":   0,
	"m4b":   0,
	"mpg":   0,
	"mpeg":  0,
	"docx":  0,
	"txt":   0,
	"go":    0,
	"mkv":   0,
}

/*
.mpg
.mpeg
These extensions are commonly used for MPEG-1 and MPEG-2 video files.
MPEG-1 is typically used for video CDs, while MPEG-2 is used for DVDs and digital television broadcasting.

.mp4: The most common extension for MPEG-4 video files. It can contain video, audio, subtitles, and images.
.m4v: A video file format developed by Apple, similar to MP4 but often used for video content that may be protected by DRM.
.m4a: Audio files in the MPEG-4 format. This extension is often used for audio-only files encoded in the AAC or ALAC codecs.
.m4b: Similar to M4A, but specifically used for audiobooks and podcasts, supporting bookmarking.
*/

func main() {
	fmt.Println()
	// Variable to store Rick's total count of all interesting file types
	var ricksTotal int

	// Get the current working directory
	rootDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}

	/*
		Next we will be using an anonymous function, also known as a lambda or closure.
		An anonymous function is a function that is defined without a name. In Go, we define anonymous functions
		and use them inline where they're needed, often as arguments to higher-order functions.
	*/
	// Anonymous Function to be called for each file/directory found by Walk; which is, in this case, the higher-order function.
	err = filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		// filepath.Walk: is a standard library function in Go that traverses a directory tree,
		// calling a provided function for each file or directory found. Refer to the end of this file for more on all of this.
		if err != nil {
			return err // exit early in the event of an error
		}

		// Check if it is a regular file (not a directory or special file)
		if info.Mode().IsRegular() {
			fileTypeCounts["total"]++
			// Get the file extension and convert it to lowercase
			ext := strings.ToLower(filepath.Ext(info.Name()))
			// Increment the count for the specific file type if it matches one of the supported types
			switch ext {
			case ".mp3":
				fileTypeCounts["mp3"]++
			case ".aac":
				fileTypeCounts["aac"]++
			case ".m4a":
				fileTypeCounts["m4a"]++
			case ".wav":
				fileTypeCounts["wav"]++
				fmt.Printf("pathName of %swav file%s: %s\n\n", colorYellow, colorReset, path)
			case ".flac":
				fileTypeCounts["flac"]++
			case ".ogg":
				fileTypeCounts["ogg"]++
			case ".aiff":
				fileTypeCounts["aiff"]++
			case ".alac":
				fileTypeCounts["alac"]++
			case ".mp4":
				fileTypeCounts["mp4"]++
				fmt.Printf("pathName of %smp4%s file: %s\n\n", colorCyan, colorReset, path)
			case ".m4v":
				fileTypeCounts["m4v"]++
				fmt.Printf("pathName of %sm4v%s file: %s\n\n", colorCyan, colorReset, path)
			case ".m4b":
				fileTypeCounts["m4b"]++
				fmt.Printf("pathName of %sm4b%s file: %s\n\n", colorCyan, colorReset, path)
			case ".mpg":
				fileTypeCounts["mpg"]++
				fmt.Printf("pathName of %smpg%s file: %s\n\n", colorCyan, colorReset, path)
			case ".mpeg":
				fileTypeCounts["mpeg"]++
				fmt.Printf("pathName of %smpeg%s file: %s\n\n", colorCyan, colorReset, path)
			case ".mov":
				fileTypeCounts["mov"]++
				fmt.Printf("pathName of %smov%s file: %s\n\n", colorCyan, colorReset, path)
			case ".avi":
				fileTypeCounts["avi"]++
				fmt.Printf("pathName of %savi%s file: %s\n\n", colorCyan, colorReset, path)
			case ".docx":
				fileTypeCounts["docx"]++
				fmt.Printf("pathName of %sdocx file%s: %s\n\n", colorYellow, colorReset, path)
			case ".txt":
				fileTypeCounts["txt"]++
				fmt.Printf("pathName of %stxt file%s: %s\n\n", colorYellow, colorReset, path)
			case ".go":
				fileTypeCounts["go"]++
				fmt.Printf("pathName of %sgo file%s: %s\n\n", colorGreen, colorReset, path)
			case ".mkv":
				fileTypeCounts["mkv"]++
			// todo, add other file types, e.g., doc, numbers, pages, jpg, PNG, png, webp, bmp, html, pdf, (Jpg, Mpg, mpeg, etc.)
			default:
				fmt.Printf("pathName of a mystery file: %s%s%s\n\n", colorRed, path, colorReset)
			}
		}

		return nil
	})

	if err != nil {
		fmt.Println("Error walking the path:", err)
		return
	}

	// Print the total number of files found
	fmt.Printf("Total number of files incl all: %d\n", fileTypeCounts["total"])
	// Print the total number of files for each type and calculate Rick's total
	for fileType, countOfFilesOfOneType := range fileTypeCounts {
		if fileType != "total" {
			ricksTotal += countOfFilesOfOneType
			/*
					if countOfFilesOfOneType > 0 && countOfFilesOfOneType < 10 {
						fmt.Printf("pathNames of the docx files: %s\n\n", path)

					}
				Burning Down the House	3:40	The Cardigans & Tom Jones		Alternative	0	1

			*/
			if countOfFilesOfOneType > 0 {
				fmt.Printf("Total number of %s files: %d\n", fileType, countOfFilesOfOneType)
			}
		}
	}

	fmt.Printf("\nThere were %s%d%s total interesting files.\n\n", colorRed, ricksTotal, colorReset)
}

/*
(path string, info os.FileInfo, err error): The parameters of the function.
filepath.Walk will pass the path of the file or directory, information about the file, and any error encountered.

error: The return type of the function. The function must return an error, which filepath.Walk will handle appropriately.

Why Use Anonymous Functions?
	Convenience: Anonymous functions are convenient when you need to pass a simple function as an argument. They allow you to define the function inline,
		keeping related code together.
	Closure: Anonymous functions can capture variables from their surrounding scope. This is useful when you need to maintain state or pass additional context
		without modifying the function signature.
	Readability: When the function logic is short or only relevant in a specific context, an anonymous function can make the code more readable by keeping the
		function definition close to where it is used.
	Flexibility: They allow you to write more flexible and modular code by enabling higher-order functions (functions that take other functions as arguments).


// ===============

ffmpeg -i input.flac -codec:a libmp3lame -qscale:a 2 output.mp3

 Create a new file named convert_flac.sh and add the following

#!/bin/bash
mkdir -p mp3 m4a

for file in *.flac; do
    filename=$(basename "$file" .flac)
    ffmpeg -i "$file" -codec:a libmp3lame -qscale:a 2 "mp3/${filename}.mp3"
    ffmpeg -i "$file" -codec:a aac -b:a 192k "m4a/${filename}.m4a"
done


Make the script executable and run it:
chmod +x convert_flac.sh
./convert_flac.sh


The script will convert all .flac files in the current directory to both MP3 and M4A formats,
placing the output files in mp3 and m4a directories respectively.
*/

// Constants:
const colorReset = "\033[0m"
const colorRed = "\033[31m"
const colorGreen = "\033[32m"
const colorCyan = "\033[36m"
const colorPurple = "\033[35m"
const colorYellow = "\033[33m"
