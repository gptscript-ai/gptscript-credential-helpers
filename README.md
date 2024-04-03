This is a fork of https://github.com/docker/docker-credential-helpers for use with GPTScript.

## Building

1 - Download the source.

```shell
$ git clone https://github.com/gptscript-ai/gptscript-credential-helpers.git
$ cd gptscript-credential-helpers
```

2 - Use `make` to build the program you want. That will leave an executable in the `bin` directory inside the repository.

```shell
$ make osxkeychain
```

3 - Put that binary in your `$PATH`, so GPTScript can find it.

```shell
$ cp bin/build/gptscript-credential-osxkeychain /usr/local/bin/
```

### Available programs

1. osxkeychain: Provides a helper to use the macOS keychain as credentials store.
2. secretservice: Provides a helper to use the D-Bus secret service as credentials store.
3. wincred: Provides a helper to use Windows credentials manager as store.
4. pass: Provides a helper to use `pass` as credentials store.

#### Note

`pass` needs to be configured for `gptscript-credential-pass` to work properly.
It must be initialized with a `gpg2` key ID. Make sure your GPG key exists is in `gpg2` keyring as `pass` uses `gpg2` instead of the regular `gpg`.

## License

MIT. See [LICENSE](LICENSE) for more information.
