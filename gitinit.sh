#! /bin/sh

echo "======== Git 初始化开始 ========"

rm -rf .git
git init
chmod +x gitconfig.sh
chmod +x gitpush.sh
chmod +x githead.sh
./gitconfig.sh

git remote add github git@github.com:hikaliv/LiBoHikalivRiYu.git

echo "======== Git 远端仓库 ========"
git remote -v

echo "======== Git 初始化结束 ========"

git add .
git st
git ci -m 'LiBo Hikaliv RiYu'