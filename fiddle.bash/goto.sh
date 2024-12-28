#!/bin/bash -xe
# include this boilerplate
function goto {
  label=$1
  cmd=$(sed -n "/^#label#:/{:a;n;p;ba};" $0 | grep -v ':$')
  eval "$cmd"
  exit
}

start=${1:-"start"}

goto $start

#start#
echo "start"
goto bing

#boom#
echo boom
goto eof

#bang#
echo bang
goto boom

#bing#
echo bing
goto bang

#eof#
echo "the end mother-hugger..."
