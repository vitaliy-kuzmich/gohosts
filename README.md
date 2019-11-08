light weight alternative of pihole script. It just merge different hosts files. You can take hosts links from here : https://github.com/StevenBlack/hosts. If you have your own hosts file, just mention it at /etc/gohosts.json at hosts. It removes duplicates, and replaces 127.* with 0.*, skips empty lines, etc...  Will create default settings on first run.

hot to use :
add host files at  /etc/gohosts.json, modify output merged result file path, after just run it.
