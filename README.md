## CLI Domain finder
Instead of sitting in front of our favourite domain name provider for hours on end trying different domain name, I created this command-line tools to help me find the right one.
This project contains five small programs that are coupled together by a bash script.

# The domain name samples for words like "blockbuster", "deep", "cloud"
![sample for blockbuster](https://github.com/bontusss/domainfinder/blob/main/sample.png?raw=true)

![sample for deep and cloud](https://github.com/bontusss/domainfinder/blob/main/sample2.png?raw=true)


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
Fork the repo and clone into your local machine. Cd into the domainfinder, Set execute permission for build.sh script using the chmod command by running `chmod +x build.sh` (PS: Find how to give this permission if you are setting up on windows). Run the `./build.sh` file. if you are on windows, rename it to `build.bat`. Notice if the lib folder has the build file of the sub-programs. If it was successful, Build the domainfinder program by running `go build` and start the program by running `./domainfinder` or `go run .`
N/B: I included the build versions. So you can skip some steps.

![sample for successful build](https://github.com/bontusss/domainfinder/blob/main/build.png?raw=true)
