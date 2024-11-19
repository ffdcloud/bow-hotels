#!/bin/bash

BUILD_FILENAME="build"

# Stop the existing application (if running)
pkill $BUILD_FILENAME || true
