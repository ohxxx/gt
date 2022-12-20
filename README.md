<div align="center">
  <img src="logo.svg" width="200">
  <p>A command-line tool to jump to a directory</p>
</div>

## Usage

```sh
# Add an alias
gt -a xxx 

# Rename the alias
gt -r xxx zzz

# Delete alias
gt -d xxx

# Clear all alias
gt -c

# List of Aliases
gt -l    	
```

## Install

``` sh
go get github.com/ohhooxxx/gt
```

## Configuration

```sh
# Create a function in the configuration (.zshrc or .bashrc)
function gt() {
  if [[ $1 == "" || $1 == "-l" || $1 == "-c" ]]; then
    echo $($HOME/go/bin/gt $@)
  else
    builtin cd "$($HOME/go/bin/gt $@)"
  fi
}
```

## Inspiration

> This project is inspired by [yuanchuan/jd](https://github.com/yuanchuan/jd/), for personal use and learning only

> The font of the logo comes from [smiley-sans](https://github.com/atelier-anchor/smiley-sans)
