#!/bin/bash
# set -x
gbookshelf-server &
rm -f mydb.pb
gbsctl add hoge fuga -p 100
gbsctl add foo bar
gbsctl add Designing Data-Intensive Applications --page 624
gbsctl update Designing Data-Intensive Applications --current 345
gbsctl list
gbsctl remove foo bar
gbsctl list
gbsctl update hoge fuga -c 50
gbsctl list
gbsctl update hoge fuga -c 60 -p 111
gbsctl list
gbsctl update hoge fuga -c 222 -p 222 -s done
gbsctl list
gbsctl update hoge fuga -s 'incomplete'
gbsctl update Designing Data-Intensive Applications --status done
gbsctl list --incomplete_only
gbsctl update hoge fuga -c 555 -p 333
gbsctl list
pkill gbookshelf-server
