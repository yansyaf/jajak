#!/bin/sh
### BEGIN INIT INFO
# Provides:          jajakd
# Required-Start:    $local_fs $network $named $time $syslog
# Required-Stop:     $local_fs $network $named $time $syslog
# Default-Start:     2 3 4 5
# Default-Stop:      0 1 6
# Description:       Jajak API
### END INIT INFO

GO_BIN=$GOPATH/bin
SCRIPT=$GO_BIN/jajak
RUNAS=artikow
NAME=jajakd
SWAGGER_FILE=$GOPATH/src/github.com/toshim45/jajak/swagger.json
SWAGGER=$GO_BIN/swagger
SWAGGER_HOST=127.0.0.1
SWAGGER_PORT=8001

PIDFILE=/var/run/$NAME.pid
SWAGGERPIDFILE=/var/run/$NAME-swg.pid
LOGFILE=/var/log/$NAME.log

start() {
  if [ ! -f $SWAGGER_FILE ]; then
    echo 'Swagger json doesnt exists' >&2
    return 1
  fi
  if [ -f $PIDFILE ] && kill -0 $(cat $PIDFILE); then
    echo 'Service already running' >&2
    return 1
  fi
  echo 'Starting service…' >&2
  local CMD="$SCRIPT &> \"$LOGFILE\" & echo \$!"
  local SWGCMD="$SWAGGER serve $SWAGGER_FILE --host $SWAGGER_HOST --port $SWAGGER_PORT &> \"$LOGFILE\" & echo \$!"
  su -c "$CMD" $RUNAS > "$PIDFILE"
  su -c "$SWGCMD" $RUNAS > "$SWAGGERPIDFILE"
  echo 'Service started' >&2
}

stop() {
  if [ ! -f "$PIDFILE" ] || ! kill -0 $(cat "$PIDFILE"); then
    echo 'Service not running' >&2
    return 1
  fi
  echo 'Stopping service…' >&2
  kill -15 $(cat "$PIDFILE") && rm -f "$PIDFILE"
  kill -15 $(cat "$SWAGGERPIDFILE") && rm -f "$SWAGGERPIDFILE"
  echo 'Service stopped' >&2
}

uninstall() {
  echo -n "Are you really sure you want to uninstall this service? That cannot be undone. [yes|No] "
  local SURE
  read SURE
  if [ "$SURE" = "yes" ]; then
    stop
    rm -f "$PIDFILE"
    rm -f "$SWAGGERPIDFILE"
    echo "Notice: log file was not removed: '$LOGFILE'" >&2
    update-rc.d -f jajakd remove
    rm -fv "$0"
  fi
}

status() {
        printf "%-50s" "Checking $NAME..."
    if [ -f $PIDFILE ]; then
        PID=$(cat $PIDFILE)
            if [ -z "$(ps axf | grep ${PID} | grep -v grep)" ]; then
                printf "%s\n" "The process appears to be dead but pidfile still exists"
            else    
                echo "Running, the PID is $PID"
            fi
    else
        printf "%s\n" "Service not running"
    fi
    if [ -f $SWAGGERPIDFILE ]; then
        SWGPID=$(cat $SWAGGERPIDFILE)
            if [ -z "$(ps axf | grep ${SWGPID} | grep -v grep)" ]; then
                printf "%s\n" "The swagger process appears to be dead but pidfile still exists"
            else    
                echo "Swagger Running, the PID is $SWGPID"
            fi
            else
        printf "%s\n" "Swagger Service not running"
    fi
    
}


case "$1" in
  start)
    start
    ;;
  stop)
    stop
    ;;
  status)
    status
    ;;
  uninstall)
    uninstall
    ;;
  restart)
    stop
    start
    ;;
  *)
    echo "Usage: $0 {start|stop|status|restart|uninstall}"
esac

