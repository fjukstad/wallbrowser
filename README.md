# Wallbrowser
Web Browser for the Tromsø Display Wall

## What
A prototype to figure out if it is possible to get a web browser on the Tromsø Display Wall. It is. 

## How it works
Each tile opens up a url, specified in [index.html](https://github.com/fjukstad/wallbrowser/blob/master/templates/index.html). Since this web page is the same size as the display wall, each tile must scroll down to a x-y location, so that the tiles together dispaly the entire page. This x-y location is retrieved from the [wallbrowser](https://github.com/fjukstad/wallbrowser/blob/master/wallbrowser.go) running on a different computer (NOTE that the ip-address 10.1.1.60 is hard coded in many places). When each tile has retrieved the x-y location, it scrolls there. It is possible for users to send keyboard events to the display wall by visiting the /controller url (10.1.1.60:9191/controller in my case). These events are sent using websockets and triggered on the tiles using  `xdotool`.   

## How to use it. 
- In [index.html](https://github.com/fjukstad/wallbrowser/blob/master/templates/index.html)(at the bottom) you'll find some iframe tags. Modify one of these to the url that you want to display on the display wall. 
- Run the [wallbrowser.go](https://github.com/fjukstad/wallbrowser/blob/master/wallbrowser.go) script on your local machine, to start the service that orchestrates the whole thing. 
- Run the ansible script ```ansible tiles -B 6000 -m shell -a 'export DISPLAY=:0 && chromium-browser --user-data-dir=/tmp/chromer --window-position=0,0 --incognito --window-size=1024,768 --app="http://10.1.1.60:9191" --temp-profile --disable-overlay-scrollbar'``` on `rocksvv` to open up the browser windows
- Run the [inputman.go](https://github.com/fjukstad/wallbrowser/blob/master/inputman/inputman.go) script on all the tiles: ```ansible tiles -B 6000 -m shell -a 'export DISPLAY=:0 && go run inputman.go'```
- Open the url (10.1.1.60:9191/controller) to send keyboard events to the display wall. 

## Future development
Since the current design loads the user-specified url into an iframe, I think it is impossible to implement the features we want. By using an iframe we have no control what goes on in there (because of cross-origin mess), and cannot calculate things such as $(document).height(). We need these functions to work because of user-input on the web page needs to be coordinates such that the tiles e.g. stop scrolling at correct location. 

I believe that the best way of continuing the development of this project is to look into user-scripts and tampermonkey. Using a user-script it is possible to inject any web page with custom javascript. This will allow us to really figure out what is going on etc etc. I haven't got the time for this right now, but will definately look into it if i get the chance later. 
