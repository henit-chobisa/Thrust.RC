# What is it ?
While making Rocket.Chat apps people have to go through a lots of steps of setting up an RC Server, signing up, enabling different modes etc just for have a look on the app.
This App does all that work for you using a config file, which only has general info like `AppDirectory`, `DockerComposePath` ( If you want to give a custom file ), `admin` info etc.
On launch it installs all the things which you would need and launches an RC Server with the admin user that you have provided and installs the app from your app directory to the server making testing of RC Apps way more easier.

# Prerequisites
- Docker ( start it while running the app)
- If you're running rocket.chat apps make sure to install the necessary requirements like npm, node etc. 

# How to use it ?
- Start Docker
- Download Docker Compose file from [Rocket.Chat Developer Docs](https://developer.rocket.chat), in case you want to provide any additional configurations, else the companion downloads the preconfigured docker-compose file by itself, you don't have download anything.
- If you want to override the default configuration, make a `config.json` file with the same directory as of companion, with a minimum configuration of these fields, the below are default fields.
```json
{
    "admin" : {
      "username": "user0",
      "email": "a@b.com",
      "pass": "123456",
      "name": "user"
    },
    "appDir" : "./",
    "composeFilePath" : "./docker-compose.yml"
}
```
- Download the binary using `wget` from github and provide executable permissions
```bash
 wget https://github.com/henit-chobisa/RC-Test-Environment-Companion/releases/download/0.1/RC_AppTestCompanion
 chmod +x RC_AppTestCompanion
```
- Now execute the binary in your shell 
```sh
./RC_AppTestCompanion
```

# How it looks like ?
The execution takes place in `3 Phases` and once the app detects your `app.json` it will show your app name, below Rocket.Chat, you can see the logs on the completion of each step.
<img width="1166" alt="image" src="https://user-images.githubusercontent.com/72302948/211493665-55ccb522-29ea-4e23-9eba-0596a52c6060.png">

# What would be the end result ?
In the end when every thing would be completed, you can open your `http://localhost:3000`, login with the `username` and `password` provided in the `config.json` file and get inside the workspace without have to configure any organization and cloud.
<img width="993" alt="image" src="https://user-images.githubusercontent.com/72302948/211494438-f0dcab91-4ab8-4e07-b615-f7756b465a37.png">
![DemoCompanion](https://user-images.githubusercontent.com/72302948/211494912-abb1a8b4-dee2-4036-adef-3d7f1f7b4b04.gif)

### Made with ♥️ for [Rocket.Chat](https://www.rocket.chat) by [Henit Chobisa](https://twitter.com/henit_chobisa)

## Note
If you find any bug please open an issue, your contributions would be appreciated.
