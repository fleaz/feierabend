# Some ideas for the future

## Split "traverse folders to find repositories" and "Check repo state"
The first command produces a lot I/O and takes a lot of time, the second one is fast. By splitting this you could run
them seperately. The list of found repositories could be saved in a small file somewhere in the users home dir e.g.

## Exit code
By returning a meaningful error code, one could run this tool in a script/cronjob

## Check also for uncommited commits and not only for dirty workdir

## Branches
Get the branches of a repository and check all of them and not only the currently checked out one

## Stash
Maybe one would also be interested in stashed code? Probably a flag would make sense here to (de)activate this

