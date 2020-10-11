# found on https://gist.github.com/pcgeek86/0206d688e6760fe4504ba405024e887c

cd $HOME
file='go1.15.2.linux-arm64.tar.gz'
wget "https://dl.google.com/go/$file"
sudo tar -C /usr/local -xvf "$file"
cat >> ~/.bashrc << 'EOF'
export GOPATH=$HOME/go
export PATH=/usr/local/go/bin:$PATH:$GOPATH/bin
EOF
source ~/.bashrc
