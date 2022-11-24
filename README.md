<div align="center">
  <img src="logo.svg" width="200">
  <p>A command-line tool to jump to a directory</p>
</div>

## Commands

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

## Usage

**Install**

``` sh
# Install
go get github.com/xxx002/gt

# Build
go build
```

**Configure**

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

> This project comes from [yuanchuan/jd](https://github.com/yuanchuan/jd/), for personal use and study only