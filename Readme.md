
**Steps to run the project on a MAC and an iOS device**

- If you do not have go already installed on your Mac, please download and install it from here: https://golang.org/doc/install?download=go1.10.1.darwin-amd64.pkg.

- Once the installation is complete, create a directory, $HOME/go. This will be our workspace.

- Next, make the directory src inside the workspace.

- Open terminal and change the directory to the workspace $HOME/go and type the following command to clone the github repository.

git clone https://github.com/abhiarya/GoReactTwitchLive.git

- Put your Twitch app developer client ID in the file $HOME/go/src/GoReactTwitchLive/GoTwitchLive/Twitch/api.go at the line #30 and save the file. 

- In the terminal, change the directory to the workspace $HOME/go/src/GoReactTwitchLive. Run the following command:

go run GoTwitchLive/main.go

You should see following messages:

Initializing Twitch API...
Booting the server...

- Now go to the browser and type the following URL
http://localhost:8080/go-twitch-live

You should see a long JSON response. If you see it then you have successfully run the server. 

- Now change "localhost" to your local IP address at the line #43 in "$HOME/go/src/GoReactTwitchLive/ReactTwitchLive/App.js".

Change the URL to "http://a.b.c.d:8080/go-twitch-live"

where a.b.c.d is the local IP address from the Mac System Preferences Network settings. 

- Now open another terminal (do not close the go terminal) and change the directory to $HOME/go/src/GoReactTwitchLive/ReactTwitchLive.

- Run the following command:

npm start

It should show a QR code and a message like this

Your app is now running at URL:  <The URL>

Press the key 's' on the terminal and enter your email address or phone number to get the link. 

- Install the Expo client app on your iOS device from the App Store. Also make sure to be connected to the same wifi network from the iOS device as the Mac machine. 

- Now, tap on the link received in the message/email on the iOS device. It should open the app in the Expo client. It might takes few minutes to load depending upon the network speed and the device. If everything goes well, you should see the app. 

- On the app, tap on the "Fetch Live Streams" button which should load top 100 currently live streaming users.

- All the results will be loaded in the app. It also shows a user name and the preview image of the live streaming. 

- To watch a live stream, tap on the preview image which should open the live stream in the default browser. 

- You can press the "refresh" button to refetch the currently live streaming users.

- Once done, use Ctrl+C command in both the terminals to close the server and the app.



*Future TO DOs*
The Oauth2 implementation in golang is complete where http://localhost:8080/authorize will redirect to Twitch User authentication. When the user approves the requested scopes, golang will get the user name and return as JSON.

React code implementation is to be done to interface with this golang authorize api and provide login capabilities within the app.

