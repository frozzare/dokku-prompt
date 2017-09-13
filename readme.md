# dokku-prompt

Simple dokku prompt (no suggestion exists).

## Installation

```
go get -u github.com/frozzare/dokku-prompt
```

## Usage

```
dokku-prompt [user@]hostname [cmd]
```

For example if you use the dokku user you don't need to set the base command since it's dokku.

```
dokku-prompt dokku@example.com
```

If you use another user, for example root you also need to set a command.

```
dokku-prompt root@example.com dokku
```

# License

MIT © [Fredrik Forsmo](https://github.com/frozzare)
