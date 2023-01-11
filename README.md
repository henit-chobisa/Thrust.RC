# Test Companion for Rocket.Chat Apps

Are you a Rocket.Chat App Developer? 

Stop! configuring your workspace everytime you wanna test Rocket.Chat Apps and waste the initial 20 mins of yours.
This App does all that work for you, just place the binary in your directory and execute it.
On launch it sets up everything which you would need and launches an RC Server and installs the app in that for you to test.

# Prerequisites
- Docker in running state, that's it.

# How to use it ?
- Start Docker
- Download the binary using `wget` from github and provide executable permissions
```bash
 wget https://github.com/henit-chobisa/RC-Test-Environment-Companion/releases/download/0.1/RC_AppTestCompanion
 chmod +x RC_AppTestCompanion
```
- Now execute the binary in your shell 
```sh
./RC_AppTestCompanion
```
That's it sit back & relax!

# How it looks like ?
The execution takes place in `3 Phases` and once the app detects your `app.json` it will show your app name, below Rocket.Chat, you can see the logs on the completion of each step.
<img width="1166" alt="image" src="https://user-images.githubusercontent.com/72302948/211493665-55ccb522-29ea-4e23-9eba-0596a52c6060.png">

# What would be the end result ?
In the end when every thing would be completed, you can open your `http://localhost:3000`, login with the `username` and `password` provided in the `config.json` file and get inside the workspace without have to configure any organization and cloud.
<img width="993" alt="image" src="https://user-images.githubusercontent.com/72302948/211494438-f0dcab91-4ab8-4e07-b615-f7756b465a37.png">
![DemoCompanion](https://user-images.githubusercontent.com/72302948/211494912-abb1a8b4-dee2-4036-adef-3d7f1f7b4b04.gif)

## Additional Configuration

```json
{
    "admin" : {             // Admin user info used for starting rocket.chat server
      "username": "user0",
      "email": "a@b.com",
      "pass": "123456",
      "name": "user"
    },
    "appDir" : "./", // path to your app directory
    "composeFilePath" : "./docker-compose.yml", // docker-compose file that you want to use, companion automatically downloads it, if you won't give.
    "installDependencies" : "false", // Installs the dependencies of your app, using npm install
    "watcher" : "true",
    "watcherMode" : "appDir-deep" // this means that the watcher will look at changes for all the files and folders and folder changes, while you can use the "appdir-shallow" option which will only look for for only files in appDir, won't look for subdirectories.
}
```
- If you want to override the configuration of the companion, make a `config.json` file in the same directory as the binary.
- The above is the default configuration used by the companion.
- Hot-Reloading in the companion is dependent upon watcher and watcher mode, watcher looks for the file changes in the directory and performs hot-reloading for your apps.


### Made with ♥️ for [Rocket.Chat](https://www.rocket.chat) by [Henit Chobisa](https://twitter.com/henit_chobisa)

## Note
If you find any bug please open an issue, your contributions would be appreciated.
