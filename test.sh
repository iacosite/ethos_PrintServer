#!/bin/sh
cd server
sudo ethosRun -t
ethosLog .
cd ..
sudo make clean >> /dev/null
