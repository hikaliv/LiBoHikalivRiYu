#! /bin/sh

msg='step'
step='1'

if [[ -n $1 ]]; then
  if [[ "$1" =~ ^[1-9]\d*$ ]]; then
    step=$1
    if [[ -n $2 ]]; then
      msg=$2
    fi
  else
    msg=$1
  fi
fi

git reset HEAD~$step
git add .
git ci -m $msg