#!/bin/sh

cd /app

printf "\033[0;32m✓ \033[37mInstalling and Auditing Node Modules\n"
npm install

printf "\033[0;32m✓ \033[37mStarting to deploy and watch your Rocket.Chat App\n"
rc-apps "$@"