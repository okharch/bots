# Telegram Hello Time Bot

The Telegram Hello Time Bot is a Go application designed to interact with users through Telegram. It offers functionalities to report the current time with the `/hello` command, set up periodic time notifications with the `/interval` command, and cancel ongoing notifications with the `/close` command.

## Features

- Respond to the `/hello` command with the current server time.
- Allow users to set up time notifications at specified intervals using the `/interval` command.
- Provide a `/close` command to cancel ongoing interval notifications.

## Prerequisites

- Go 1.15 or later.
- A Telegram bot token. You can obtain one by chatting with [BotFather](https://t.me/botfather) on Telegram.

## Setup

1. Clone this repository to your local machine.
2. Set your Telegram bot token as an environment variable:

    ```bash
    export TELEGRAM_TIME_HELLO_BOTAPI_KEY='your_bot_token_here'
    ```

3. Install the required Go package:

    ```bash
    go get -u github.com/go-telegram-bot-api/telegram-bot-api/v5
    ```

4. Build the Go program:

    ```bash
    go build -o telegram-hello-time
    ```

5. Run the bot:

    ```bash
    ./telegram-hello-time
    ```

## Usage

To interact with the bot, users can navigate to [t.me/TimeHelloBot](https://t.me/TimeHelloBot) and use the following commands:

- `/hello` - The bot replies with the current server time.
- `/interval <seconds>` - The bot starts sending the current time every specified number of seconds. For example, `/interval 60` will cause the bot to send the current time every minute.
- `/close` - Stops any ongoing time notifications set by the `/interval` command.

## Environmental Variables

This application uses the following environment variables:

- `TELEGRAM_TIME_HELLO_BOTAPI_KEY` - The Telegram bot token used for authenticating API requests.

Ensure that this variable is set in your environment before running the bot.

## Contributing

We welcome contributions to the Telegram Hello Time Bot project. Feel free to submit pull requests or open issues to propose new features or report bugs.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
