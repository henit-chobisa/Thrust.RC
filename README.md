
<p align="center">
  <a href="https://www.gitpod.io">
    <img src="https://user-images.githubusercontent.com/72302948/215353380-af7a74e4-e3cc-446c-b853-d1c12bc275ef.png" alt="thrust Logo" height="250" />
    <br />
    <strong>thrust</strong>
  </a>
  <br />
  <span>Test Companion for Rocket.Chat Apps</span>
</p>

Stop configuring your workspace everytime you wanna test Rocket.Chat Apps and waste the initial 20 mins of yours.
Let thrust handle it, thrust downloads manages and runs container based on Rocket.Chat and Rocket.Chat apps, and sets up your Rocket.Chat App workspace under 1 minute,just place the binary in your directory and execute it.
On launch it sets up everything which you would need and launches an RC Server and installs the app in that for you to test, making your life wayyyy easier that before.

# Prerequisites
- Docker in running state, that's it.
- No need to even do `npm install`, thrust will take care of that

# How to use it ?
- Start Docker
- Download the binary using `wget` from github and provide executable permissions
```bash
 wget https://github.com/henit-chobisa/RC-Test-Environment-Companion/releases/download/v2.0.1/thrust_linux
 chmod +x thrust_linux
```
- Now execute the binary in your shell 
```sh
./thrust_linux start <path to your app directory>

# ./thrust_linux start ./
```
Use mac binary if you're a mac user, That's it sit back & relax!

<img align="right" width="400" alt="image" src="https://user-images.githubusercontent.com/72302948/215354509-722bd660-7a87-4dbc-afee-f243b7f36ee0.png">

## How it looks like ?
The execution will confirm the necessary assets which will be needed to run the apps companion, such as
- Docker running or not
- Are all the initial dependencies full filled
- What images are present in the system and what to pull, it will pull the required images automatically don't worry

<img align="left" width="400" alt="image" src="https://user-images.githubusercontent.com/72302948/215354665-7b54dbde-2140-46ab-a6d2-e5a4d3be9a4f.png">

Next it will start necessary containers such as Rocket.Chat and mongodb using `dockerd` and wait until the container are fully started, it won't start a new container everytime you start, it will look for the running container and just create the ones which are not present.
Along with that, it also creates an admin user for the initialised Rocket.Chat Server, which will later be helpful for installing the apps inside, you will soon have an option to configure everything manually with some flags and yaml files.

<br/>

<img align="right" width="400" alt="image" src="https://user-images.githubusercontent.com/72302948/215354929-fe6266da-d90a-4b89-adaf-37f37922ba81.png">

## What would be the end result ?
In the end you can see the companion container's logs, which will show you two of the essential things, 
- Node Module Installation, which is done inside the container only, so you can expect a frest install everytime.
- Rocket.Chat CLI logs


## Using on EC2 Instance
Please have a look on this [short manual](https://henitchobisa.notion.site/Setting-up-App-s-Companion-in-EC2-fdde72b19afc40ed93c9ded5887a641c) for end to end configuration of your ec2 with Apps' conpanion.

### debugging
- Docker-Compose "executable file not found in $PATH"
```
sudo chmod +x /usr/local/bin/docker
```
<p align="center">
 <img width="400" alt="Icon - Red" src="https://user-images.githubusercontent.com/72302948/215355019-2779af9c-14bb-453c-a56a-b60156390916.png">
 <br />
    <strong>Made with ♥️ for Rocket.Chat Apps by Henit.Chobisa</strong>
</p>

## Note
If you find any bug please open an issue, your contributions would be appreciated.
