#!/bin/bash
#
# Quick script to install or update golang on Ubuntu since im lazy
# Needs curl,wget,bash and sudo
#

VERSION=$(curl -s https://go.dev/VERSION?m=text | head -1)

# Lets get that latest version
wget -N https://go.dev/dl/${VERSION}.linux-amd64.tar.gz -P /tmp/

# Do we have a go version already?
if test -d /usr/local/go; then
  echo "GO exists, make a backup."
  sudo rm -rf /usr/local/go-previous
  sudo mv /usr/local/go /usr/local/go-previous
fi

# Install the new version
echo "Installing golang"
sudo tar -C /usr/local -xzf /tmp/${VERSION}.linux-amd64.tar.gz

# Add golang to PATH
if grep -Fxq "export PATH=\$PATH:/usr/local/go/bin" ~/.bashrc 
then
    echo "PATH for golang already present"
else
    echo "PATH for golang not present adding it"
    echo "export PATH=\$PATH:/usr/local/go/bin" >> ~/.bashrc 
fi


# Cleanup
if test -f /tmp/${VERSION}.linux-amd64.tar.gz; then
  echo "Removing ${VERSION}.linux-amd64.tar.gz"
  rm /tmp/${VERSION}.linux-amd64.tar.gz
fi
