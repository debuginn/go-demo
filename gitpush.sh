echo "------ job start ------"
#!/bin/bash
echo "- set local git config -"
git config --local user.name 'debuginn' && git config --local user.email 'debuginn@icloud.com'
echo "- git pull origin main -"
git pull origin main
echo "- git push origin main -"
git add .
git commit -m "study"
git push -u origin master
echo "------ job end ------"