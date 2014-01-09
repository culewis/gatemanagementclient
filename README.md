this simple repo is a proof of concept to see if etcd is mature enough to fit our current needs.
we know what they are and since this is a public repo (for now) we won't mention them :)

in order to run the samples i highly suggest the reader become familiar with the go programming
language upon installing it.  an excellent guide/tutorial to installing and cutting your teeth
in this language can be found at:  http://golang.org/doc/install and 
http://golang.org/doc/code.html

this prototype also relies on a running etcd server.  you can download and install this in literally less than a minute here:  https://github.com/coreos/etcd/releases/.  you'll need mac osx or to run it out of a docker container.  i think these dependencies will hopefully be able to be removed but i have a mac air book and am just prototyping atm so no need to hack at it yet or explore removing that dependency... if you don't have a mac air book from what i've read the docker container is not hard to setup (on linux):  https://www.docker.io/  if you are running windoze then - HA i laugh at you...

after getting an etcd server up it should be trivial to use go's test capabilities to watch the HTTP calls "get 'er done"

note you might need to adjust the limit on open files for an executable running from the command line.  if you see this error from the etcd process (or log):
	http: Accept error: accept tcp[::]:<etcd port>: too many open files; retrying in <number here>ms
this can be mitigated by increasing the number of concurrent "files" open by a process by observing the following article:
	http://superuser.com/questions/433746/is-there-a-fix-for-the-too-many-open-files-in-system-error-on-os-x-10-7-1