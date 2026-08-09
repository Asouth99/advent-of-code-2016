#!/usr/bin/env bash

dayNumber=$1

if [[ -z $dayNumber ]]; then
    echo "Day Number required" >&2
    exit 1
fi

day=$((10#$dayNumber))

# Make new folder
if [ -d ./day$dayNumber ] ; then
    echo "Error: ./day$dayNumber already exists" >&2
    exit 1
fi
cp -r template/dayXX ./day$dayNumber

# Update file names
mv day${dayNumber}/dayXX_test.go day${dayNumber}/day${dayNumber}_test.go
mv day${dayNumber}/dayXX.go day${dayNumber}/day${dayNumber}.go

# Update references to dayXX within the files
sed -i -e 's/dayXX/day'${dayNumber}'/g' ./day${dayNumber}/day${dayNumber}.go
sed -i -e 's/dayXX/day'${dayNumber}'/g' ./day${dayNumber}/day${dayNumber}_test.go
sed -i -e 's/template\/day'${dayNumber}'/day'${dayNumber}'/g' ./day${dayNumber}/day${dayNumber}_test.go

# Modify main.go file
sed -i "/\/\/ ADD IMPORT HERE/a \ \ \ \ \"aoc2016/day${dayNumber}\"" main.go
sed -i "/\/\/ ADD SOLUTION HERE/a \ \ \ \ ${day}: day${dayNumber}.Solve," main.go