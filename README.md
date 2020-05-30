# sus-server
Server for the [Simple URL Shortner client](https://github.com/raduschirliu/sus).

## Requirements
- Go 1.13
- A PostgreSQL database

## Running
To run the server, the following environment variables need to be set first:
```bash
DATABASE_URL=postgres://xxxxxxxxxxxxxxxxxxxxxxxx
PORT=xxxx
```
These can also be loaded from a `.env` file placed in the root directory.  

The server can then be built and started by running
```bash
./run.sh
```