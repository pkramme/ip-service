# IP Service
This service is used inside my infastructure to make an easy way to find out
ones IP address.

## Usage
`$ ip`  

`$ ip -listen :8000`  
If you want to specify the listening address, use the
`listen` flag. If not launched with this flag, IP
Service will listen on port 80.  

`$ ip -log /var/somewhere/logfile.log`  
If you want to change the log position, use the `log`
flag. Normally, the log is in the same directory and is
called `ipservice.log`.  


## License
Simplified BSD. See `LICENSE` file for more.
