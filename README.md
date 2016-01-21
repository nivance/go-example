#go-example
  
   
##Debug in LiteIDE
* 1 Open your project in LiteIDE
* 2 Select "Build" menu and under it select "Build Configuration..."
* 3 A dialog shows up. Select "Custom" tab. It contains a list of key/value pairs
* 4 Double click the space reserved for the value of "BUILDARGS"
* 5 write -gcflags "-N -l"
* 6 close the dialog by clicking Ok button
* 7 rebuild your project (Ctrl+B)
* 8 press F5 to start debugging
* 9 put some breakpoints wherever you wish
* 10 press f5 again to reach the first breakpoint