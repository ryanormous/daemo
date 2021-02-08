
SUMMARY:
  daemo is a naïve daemon.  Although not a true daemon in the idiom of W. Richard Stevens,
  it behaves like one when run using the included systemd configuration.

  The functionality of daemo is limited to management of a single pid.  It doesn't really
  do anything.  Its useful only for testing that involves a pid.

  daemo builds to an executable.


USAGE:
  $ daemo -help
  > -conf string
  >       OPTIONAL. PATH TO CONFIGURATION FILE.
  > -help
  >       NEVER BE DAUNTED.


CONFIGURATION:
  The configuration file is in json format.  Its pretty straightforward.

  Here's an example:
  $ jq . configuration.json 
  > {
  >   "Name": "daemo",
  >   "Root": "",
  >   "Logfile": "",
  >   "Pidfile": "",
  >   "Version": "0.1"
  > }

  FIELDS:
  "Name"
      Optional.  Default is "daemo".
  "Root"
      Optional.  Path to root directory of daemo install path.
  "Logfile"
      Optional.  Default logs to stdout.
  "Pidfile"
      Optional.  Defaults to "/tmp/daemo.pid"
  "Version"
      Optional.  Additional info about daemo.
  "Hostname"
      Name of host running daemo.  Added upon startup.
  "Confpath"
      Path to the config file itself.  Added upon startup.


RECOGNIZED SIGNALS:
  Upon SIGHUP daemo will reload configuration from the file path given at startup.


LOGGING:
  daemo will write to a specified log file.  Probably easiest to let journald handle stdout.


INSTALL:
  $ git clone https://github.com/ryanormous/daemo.git
  $ cd daemo
  $ sudo make


NOTES:
  en.wikipedia.org/wiki/Advanced_Programming_in_the_Unix_Environment
  en.wikipedia.org/wiki/%C3%86

