# Bot

This is a simple bot to play captrivia.

## Usage

To install, set up a virtualenv and install requirements in it.  We use `pyenv` and `pyenv-virtualenv` to manage our Python environments:

```sh
pyenv virtualenv bot
pyenv activate bot
pip install -r requirements.txt
```

Make sure you're already running the captrivia backend (instructions in the main README).  Then start the bot, using Ctrl-C to cancel:
```sh
python main.py
```
It will print out its progress every 100 games.
