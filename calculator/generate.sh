#!/bin/bash

protoc calcpb/calcpb.proto --go_out=plugins=grpc:.
