#!/bin/sh
##
## The service for our baseimage
##
cd /markdown-renderer

exec /sbin/setuser markdown-renderer /usr/local/bin/go-markdown-renderer
