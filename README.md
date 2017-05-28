# Cagliostro (Discord)

Granblue Fantasy bot for Discord, based on [risend/vampy][].

[risend/vampy]: <https://github.com/risend/vampy>

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

A `cagliostro.json` file must exist in the working directory.

## Commands

| Command | Arguments | Description |
|---|---|---|
| `events` | | Display current and upcoming events. |
| `emo` | `name` | Display emoji `name`. |

## Available emoji

Check `media/emoji` directory.

If you want to use `whoa.png`, then the command is `!emo whoa`.

## Contact

* [Discord](https://discord.gg/FH5zuJh)

## License

All the files under the `media` are from third-parties and may be subject to
different licensing terms.

Everything else is licensed under **Apache 2.0**.
