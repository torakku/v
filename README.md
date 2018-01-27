# Cagliostro (Discord)

Granblue Fantasy bot for Discord.

This bot exists as a free and open source alternative to
the proprietary [risend/vampy][] bot.

[risend/vampy]: <https://risend.github.io/vampy/>

## Features

* [risend/vampy][]-compatible emoji.

## Install

```bash
go get github.com/KuroiKitsu/discord-cagliostro/...
```

## Configuration

Create a `cagliostro.json` file with this content:

```json
{
  "token": "YOUR_TOKEN",
  "prefix": "!",
  "emoji_dir": "./media/emoji"
}
```

## Running the bot

Just run it as `cagliostro`.

There must exist a `cagliostro.json` file in the working directory.

## Commands

| Command | Arguments | Description |
|---|---|---|
| `events` | | Display current and upcoming events. |
| `emo` | `name` | Display emoji `name`. |

## Available emoji

Check `media/emoji` directory.

If you want to use `whoa.png`, then the command is `!emo whoa`.

## Contact

* [Discord](https://discord.gg/R8VkY8t)

## License

All the files under the `media` are from third-parties and may be subject to
different licensing terms.

Everything else is licensed under **Apache 2.0**.
