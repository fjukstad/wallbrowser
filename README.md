# Wallbrowser
Web Browser for the Tromsø Display Wall

## What
A prototype to figure out if it is possible to get a web browser on the Tromsø
Display Wall. It is. 

## How to use it 
Use the `cf` command in `/share/apps/bin/` to start chromium on the tiles, e.g.: 

```
/share/apps/bin/cf 'export DISPLAY=:0 && /usr/bin/chromium-browser --user-data-dir=/tmp/chromer --window-position=-%px,-%py --window-size=7168,3072 --user-data-dir=/tmp/chromera --app="https://maps.google.com &'
```

The `window-size` flag opens a browser window as big as the display wall, and 
the `window-position` moves the window so that the tile only shows the relevant
part of the window. This strategy is similar to what the VNC viewers are doing. 

It is possible for users to send keyboard events to the display wall by
visiting the /controller url (10.1.1.60:9191/controller in my case). These
events are sent using websockets and triggered on the tiles using  `xdotool`.   

## How to use it. 
- Run the
  [wallbrowser.go](https://github.com/fjukstad/wallbrowser/blob/master/wallbrowser.go)
  script on your local machine, to start the service that orchestrates the user
  inputs. 
- Run the [inputman.go](https://github.com/fjukstad/wallbrowser/blob/master/inputman/inputman.go) script on all the tiles: ```ansible tiles -B 6000 -m shell -a 'export DISPLAY=:0 && go run inputman.go'```
- Run a `cf` command to start the web page that you want to view. 

## Future development
We need to think of a better way of providing user-input and that. 
