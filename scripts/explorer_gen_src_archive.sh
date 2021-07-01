#!/usr/bin/bash

PROJECT=explorer
DEST="build/${PROJECT}.src.tar.gz"

# Check that git-archive-all is installed
if ! git-archive-all --help &> /dev/null; then
  echo "# This script uses 'git-archive-all', please install:"
  echo "# pip install git-archive-all"
  exit 1
fi

mkdir build
git-archive-all --prefix="explorer" --force-submodules $DEST

echo "# $0: DONE!"
ls -lh $DEST
