mc-avatar
=========

A martini application for minecraft avatars! Functional but needs improvements :)

Installation
==========
* You first will need Go installed of course. For more on this check out [Go's installation tutorial](http://golang.org/doc/install)

## Dependencies
 This project depends on two libraries:
    * [Martini](http://www.github.com/go-martini/martini)
    * [Imaging](http://www.github.com/disintegration/imaging)
follow the installation guides on the respective repositories.

If all goes well you will have all the dependencies needed to build the project:
  To build the project simply:
   ```
     go build server.go mc.go
   ```
 and this will compile into a executable.

 Then to run it simply:
 ```
   sudo ./server
 ```
 and you should be good to go :)

 Enjoy!
