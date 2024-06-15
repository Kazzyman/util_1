# This is needed for the gls command
eval "$(/opt/homebrew/bin/brew shellenv)"

# This is for the native ls command, it changes the color of the output to a lighter shade of blue for the dirs, I like it.
# ~/.bashrc contains: export LSCOLORS="exfxcxdxbxegedabagacad"
# older lscolors:
#export LSCOLORS=gxfxcxdxbxegedabagacad
# newer and favored colors:
 export LSCOLORS=gxfxbxdxcxegedabagacad
# Ex,Fx,Bx,Dx,Cx,eg,ed,ab,ag,ac,ad
# in the above, there are 11 pairs, e.g., Ex and Fx are the first two pairs (bold blue, and bold magenta, both with default BG)
# x means default color (often black) and as the second member of a pair it refers to the background field
# : a: Black, b: Red, c: Green, d: Brown, e: Blue, f: Magenta, g: Cyan, h: Light gray
# : A: Bold black, usually shows up as dark gray
# ... B: Bold red, C: Bold green, D: Bold brown, usually shows up as yellow
# : E: Bold blue, F: Bold magenta, G: Bold cyan
# H: Bold light gray; looks like bright white when used as foreground color


# This is also for the gls command, it reads your .dircolors configuration file
eval "$(gdircolors -b ~/.dircolors)"

# This is for the gls command, creates an initial .dircolors file, only run this once
# it actually clobbers my precious .dircolors file if it is run post edits to .dircolors
# gdircolors -p > ~/.dircolors

# PS1 formats the prompt, and CLICOLOR turns on colorations 
export PS1="\[\033[36m\]\u\[\033[m\]@\[\033[32m\]\h:\[\033[33;1m\]\w\[\033[m\]\$ "
export CLICOLOR=1

alias l='ls -oTFGtr,acs'
alias g='gls -oaFGtpscr --color=auto --time-style=long-iso --group-directories-first'
alias cdm='cd ~/Music/Music/Media.localized/Music'

# gls works similar to ls but is missing the , option and the T option
#    gls3='gls -oFGr    --color=auto --time-style=long-iso'
# alias g='gls -oFGtpsr --color=auto --time-style=long-iso --group-directories-first'

