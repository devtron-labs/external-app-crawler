telemetry-user-analytics this application receive events from devtron (oss app) and process the payload and write it into DB.
check in db for UPID already exists or not if not create an entry in platform table. and update platform_installation_history table for matrix.
if UPID already exists than update platform table for last active use. 


schema for telemetry data below

platform

UPID : devtron-12356
Active Since: 1 April 2021
Last Active: 4 April 2021


platform_installation_history

Install Count : <sum of UPID>
Fail Count: <included those which have no normal event received yet, check interval time> 
Success Count: 1
