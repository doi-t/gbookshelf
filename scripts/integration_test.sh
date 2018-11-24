#!/bin/bash
set -x
gbookshelf-server &
rm -f mydb.pb
gbsctl add hoge -p 100
gbsctl add fuga --page 500
gbsctl add foo
gbsctl list
ls -lh mydb.pb
xxd mydb.pb
gbsctl remove fuga
ls -lh mydb.pb
xxd mydb.pb
gbsctl list
pkill gbookshelf-server
