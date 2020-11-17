# Pterodactyl SFTP Server

This package serves as a SFTP server to run alongside the [Pterodactyl Daemon](https://github.com/pterodactyl/wings).
Previous versions of this software included a standalone mode, however this repository now
serves to provide API level access to the Wings Daemon for SFTP access.

## Configuration

Create a `config.json` in the top directory with the following properties:

```json
{
  "readOnly": false,
  "port": 2022,
  "bind": "",
  "files": "/path/to/folder/for/SFTP",
  "osUser": { "uid": 10, "gid": 10 },
  "users": {
    "<username>": {
      "password": "<SHA-256 hash>",
      "permissions": ["file.read", "file.read-content", "file.create", "file.update", "file.delete"]
    }
  }
}
```

## License

```txt
Copyright (c) 2019 Dane Everitt <dane@daneeveritt.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```
