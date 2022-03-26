Using teamgen
=============

The binaries are named for the operating system and the architecture you
are using. For example, "teamgen_windows.amd64" should run on most 
64 bit Windows systems. If not, please let me know.

The teamgen.zip file has multiple binaries and a data directory. To use, 
open up a terminal window and create a directory where you want to use teamgen.
Go into that directory, copy teamgen.zip in, and then unzip it.

At that point, choose the binary you want to run, and use the help option to
see what can be done:

  ./teamgen_darwin.amd64 -help


Using datafiles
===============

The binaries refer to a "data" directory relative to the binary. So, if you
place the binary in /home/me/bin/<binary>, then you need to put the datafiles
into /home/me/bin/data/. If you unzip teamgen.zip and use the binaries from
there, the "data" directory is already in place.

The datafiles can have blank lines and comments that begin with "#". Any other
line will be consumed and used as data.

There is a feature recommendation for allowing the program to be given a 
specific directory for datafiles. For the moment, the best solution is to 
make a copy of the current datafile, and then edit it to your heart's content.

The name files are just one unicode name per line. The careers and jobs files
have explanations in the file header.


