#!/bin/bash
# -*- ENCODING: UTF-8 -*-
sudo docker build -t eda-app .

sudo docker run -d --rm --name container-eda-app eda-app