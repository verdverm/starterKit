#!/bin/bash

./passprompt.sh > tmp.txt

grep "sha" tmp.txt > pass.txt

passwd=`cat pass.txt`
echo "passwd: $passwd"

# fileloc="/ipythonbin/ipython_notebook_config.py"
fileloc="ipython_notebook_config.py"

sed -e "s/IPYTHON_PASSWORD/${passwd}/" $fileloc > tmp.py

cp tmp.py $fileloc
# rm tmp.txt