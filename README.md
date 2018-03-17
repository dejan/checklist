# checklist

Poor man's checklist app.

<img src="https://raw.githubusercontent.com/dejan/checklist/master/screenshot.png" width="200">

## Install

    go get github.com/dejan/checklist

The app was developed and tested on Mac OS X, but it should also work on Linux and Windows.

## Usage

Program reads stdin and creates checkable menu item for every line:

    $ cat | checklist
    Implement the integration test
    Take out the trash
    Order pizza
    ^D

While this is good for ad-hoc to-do lists, consider storing your commonly used checklist somewhere in the filesystem and just pass their content via stdin when you need to used them:

    checklist < ~/checklists/travel.txt
