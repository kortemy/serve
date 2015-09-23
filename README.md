# serve
Dead simple file server.

Instead of configuring nginx or apache to serve me web content from different places, as a developer I find that a nuisance.

# Install

You must have Go installed.

`$ go get -u github.com/kortemy/serve`

# Usage

Navigate to the folder you want to server content from:

`$ cd ~/projects/secret/notPr0n/web`

Serve:

`$ serve`

`Serving from /home/kortemy/projects/secret/notPr0n/web on 9090`

Open your favorite web browser and go to `localhost:9090`. That is all.

# Configuring

Not to make this only one LOC, you can define a custom port, because... yes.

`$ serve -port 80`

Yeah, screw you, Skype!
