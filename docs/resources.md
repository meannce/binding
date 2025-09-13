# Deployment
Defines how an application should be run, keeping its state consistent. It handles rolling updates, ensures a desired number of replicas are running, allows for easy rollbacks alongside other related things.

# Service
As pods can be recreated with new IPs, a service gives them a fixed DNS name/IP and load-balances requests across them for reliability purposes.
