#!/usr/bin/env expect

set timeout 9
set $password [lindex $argv 0]
log_user 0

spawn python -c "from IPython.lib import passwd; print passwd()"
expect "Enter password:"
send "$password\r"
expect "Verify password:"
send "$password\r"

interact
