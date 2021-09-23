# webcomics-fetcher2-go #

It's just a POC to test the use of plugins in a Go project...

## related projects (plugins) ##

- [Order Of The Stick plugin](https://github.com/Eldius/webcomics-fetcher2-oots)

## comments ##

### thoughts and decisions ###

- Tryied to use the stdout to pass data from plugin to main app, but it makes harder to show pregress to user (and to show debug data). The file approach looks easier to accomplish this task. An other possible change may be generate the file name and let plugin creat it (instead of create the temp file from main app and trunk file from plugin).
