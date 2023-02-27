# DNS-GO-MATIC

Service designed to notify your IP to DNS-O-Matic.

In order to work the program requires 2 env variables:
DNS_O_MATIC_USER (overriden by flag "--user -u")
DNS_O_MATIC_PASSWORD (overriden by flag "--password -p")

An optional env variable to specify a file to write logs to can be configured:
DNS_O_MATIC_LOG_FILE (overriden by flag "--log-file -l")
If no log file is configured, logs will be written on the default output

To install cronjob:
(crontab -l 2>/dev/null; echo "* * * * * path/to/dns-o-matic -u user -p password") | crontab -

To uninstall cronjob:
crontab -l| grep -v "path/to/dns-o-matic" | crontab -

The program can run in "no-cron" mode. If the flag --sleep -s is provided the program will run
in an infinite loop with a sleep time of the --sleep -s minutes

````
Usage:
  DNS-GO-MATIC [flags]

Flags:
  -h, --help              help for DNS-GO-MATIC
  -l, --log-file string   File to write logs into
  -p, --password string   DNS-O-MATIC password
  -s, --sleep uint        Enable no-cron execution. Sleep time in minutes
  -u, --user string       DNS-O-MATIC username
```

## TODO: testing