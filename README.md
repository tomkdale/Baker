#  Baker tool
Quick and easy tool to move files to bak, un-bak, and bak flop files.


 ## Installation
Clone this repo then move the script to your path and just run `baker <filename>`.  

## Basic Use

* Run `baker foo` to move `foo` to baker file `.foo.bak`
* Run `baker .foo.bak` to move `.foo.bak` to `foo`
* Run `baker foo` while a `.foo.bak` already exists and the two files will swap places.

## Flags

* `-h` shows help page.
* `-c file` copies file to file.bak or vise-versa. 
* `-d file` shows diff on file and file.bak

