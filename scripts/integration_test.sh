#!/bin/bash
set -x
gbookshelf-server &
rm -f mydb.pb
gbsctl add hoge fuga -p 100
gbsctl add foo bar
gbsctl add Designing Data-Intensive Applications --page 624
gbsctl list
ls -lh mydb.pb
xxd mydb.pb
gbsctl remove foo bar
ls -lh mydb.pb
xxd mydb.pb
gbsctl list
pkill gbookshelf-server
