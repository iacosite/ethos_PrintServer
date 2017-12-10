#!/bin/bash
echo "Clean.."
sudo make clean >> /dev/null
echo "Make.."
sudo make 
echo "Install.."
sudo -E make install >> /dev/null
