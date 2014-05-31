# gohttpbin
This project implements a subset of the functionality of httpbin.org in
[Go](golang.org). It is deployed on
[Google App Engine](https://developers.google.com/appengine/).

The production URL for this service is http://gohttpbin.appspot.com.

# Development Environment
You can run the entire development environment, including the
[GAE development server](https://developers.google.com/appengine/docs/go/gettingstarted/devenvironment)
inside of the [Vagrant](http://vagrantup.com) environment provided with this
project.

To get it working,
[Install Vagrant](http://docs.vagrantup.com/v2/installation/index.html),
`cd` into the top folder of this repository, and `vagrant up` and then
`vagrant ssh`. You will start and ssh into a fully configured VM.

When in the VM, `cd /vagrant` and type `goapp serve` to run the GAE dev server.
Open a new terminal window, and `curl localhost:8080/get` to see it working.

# Deploying to GAE
To deploy, use
`appcfg.py --oauth2 --no_cookies --noauth_local_webserver update .`
