# Contribution Guide

## Checklist
These are loose checklist for contributors to follow when submitting patches.

1. Small commits

   You can contribute by a lot of ways from improving README, fixing typos, coding style, specific bugs, performance optimizations. However, it is preferred that you break up your commits to single logical change using `git add -p` so it is easier to review the patch. The larger the change, the more necessary it is for the commit to be broken up to tiny little pieces. If your change is large but consistent throughout (e.g. fixing a specific coding style that happens on almost every file), that can be counted as single logical change.

2. File and directory-based commit message

   We're starting to use commit messages that looks like this: `sdl: fixed some typos in render.go` or `examples/render_goroutine: fixed a dereferenced nil pointer` where it starts with folder hierarchy. It's not something strictly required but we would prefer it to be followed.

3. Compatibility with SDL 2.0.0

   The binding should compile with the oldest version SDL2. If there's a function added to the binding but is not supported by the older SDL2 version, a stub function must be provided. See `sdl/filesystem.go` for an example.

## How to contribute
You can start by forking the repository, modify the fork, push the change to your fork, and then send pull requests. [Here](http://blog.campoy.cat/2014/03/github-and-go-forking-pull-requests-and.html) are some instructions on how you can work with your own fork.

