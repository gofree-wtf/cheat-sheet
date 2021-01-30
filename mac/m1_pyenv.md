M1 맥에 pyenv 설치하기
==================

```shell
# install homebrew
$ /usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
$ echo 'PATH=/opt/homebrew/bin:$PATH' >> ~/.zshrc

# install pyenv
$ brew install pyenv-virtualenv
$ echo 'eval "$(pyenv init -)"' >> ~/.zshrc
$ echo 'eval "$(pyenv virtualenv-init -)"' >> ~/.zshrc
$ echo 'export CPPFLAGS=-I/opt/homebrew/opt/openssl/include' >> ~/.zshrc
$ echo 'export LDFLAGS=-L/opt/homebrew/opt/openssl/lib' >> ~/.zshrc

# create virtualenv
$ pyenv install 3.9.1
$ pyenv virtualenv 3.9.1 gofree
$ pyenv global gofree
```
