## CLI Domain finder
Instead of sitting in front of our favourite domain name provider for hours on end trying different domain name, I created this command-line tools to help me find the right one.
This project contains five small programs that are coupled together by a bash script.
# Sprinkle:
This program will add some web-friendly sprinkle words to increase the chances of finding the available domain names.

# Domainify:
This program will ensure words are acceptable for a domain name by removing unacceptable characters. Once this is done, it will replace spaces with hyphens and add an appropriate top-level domain (i have just added .com, .net and .org) to the end.

# Coolify: 
This program will change a boring old normal word to web 2.0 by fiddling aroun with vowels.

# Synonyms:
This program uses a third party API to find synonyms

# Available:
This program will check to see whether the domain is available or not using an appropriate WHOIS server.

# To SET UP AND RUN:
Fork the repo and clone into your local machine. Cd into the domainfinder, Run the `build.sh` file. if you are on windows, rename it to `build.bat`. Notice if the lib folder has the build file of the sub-programs. If it was successful, Build the domainfinder program by running `go build` and start the program by running `./domainfinder` or `go run .`