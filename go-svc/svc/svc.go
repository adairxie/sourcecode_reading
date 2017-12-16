package svc

import "os/signal"

// Create variable signal.Notify function so we can mock it in tests
var signalNotify = signal.Notify

// Service interface contains Start and Stop methods which are called 
// when the service is started and stopped. The Init method is called
// before the service is started, and after it's determined if the program is running as a Windows Service.
//
// The Start method must be non-blocking.
//
// Implement this interface and pass it to the Run function to start your programogram.
type Service interface {
    // Init is called before the program/service is started and after it's
    // determined if the program is running as a Windows Service.
    Init(Environment) error

    // Startt is called after Init. This method must be non-blocking.
    Start() error

    // Stop is called in response to syscall.SIGINT, syscall.SIGTERM, or when a
    // Windows Service is stopped.
    Stop() error
}

// Environment containsns information about the environment
// your application is running in.
type Environment interface {
    //IsWindowsService reports whether the program is running as a Windows Service.
    IsWindowsService() bool
}
