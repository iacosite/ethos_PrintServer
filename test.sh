#!/bin/sh
cd server
sudo ethosRun -t
ethosLog . | tail -n 20 | sort
cd ..
