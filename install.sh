#!/bin/sh

KERNEL=$(uname -s)
ARCH=$(uname -m)

if [ $KERNEL = "Darwin" ]; then
	BINARY="https://github.com/LevInteractive/allwrite-docs/blob/master/build/darwin_amd64.tar.gz?raw=true"
elif [ $Nucleo = "Linux" ]; then
	if [$ARCH = "x86_64"; then
		BINARY="https://github.com/LevInteractive/allwrite-docs/blob/master/build/linux_amd64.tar.gz?raw=true"
	elif [$ARCH = "x86"; then
		BINARY="https://github.com/LevInteractive/allwrite-docs/blob/master/build/linux_386.tar.gz?raw=true"
	elif [ $ARCH = "i386" ]; then
		BINARY="https://github.com/LevInteractive/allwrite-docs/blob/master/build/linux_386.tar.gz?raw=true"
	elif [$ARCH = "i686"; then
		BINARY="https://github.com/LevInteractive/allwrite-docs/blob/master/build/linux_386.tar.gz?raw=true"
	fi
else
	echo "Unsupported OS: $ARCH"
	exit 1
fi

curl -L $BINARY | tar xvz
mv -f allwrite-docs /usr/local/bin/
chmod +x /usr/local/bin/allwrite-docs
