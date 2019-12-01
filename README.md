# ngpio
Nvidia Jetson Nano GPIO library in Go with limited capabiltiy.

The current implementation is targetting only to toggle the output of the Jetson Nano.

![Nvidia Jetson nano header](https://github.com/jvmvik/ngpio/blob/master/jetson_nano_pinout.png)

## More advanced projects.
 * https://github.com/stianeikeland/go-rpio [Golang / Pi support only]
 * https://github.com/NVIDIA/jetson-gpio [python / Nvidia support]
 
Note: 
 * Port number 5,3,8,10,25,27 are not available as a basic output/input but reserved for special usage.
