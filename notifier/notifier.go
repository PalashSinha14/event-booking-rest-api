// Package notifier runs a small background worker pool that simulates
// sending a confirmation notification (e.g. after registering for an
// event) off the HTTP request path, so a handler can enqueue a job and
// respond to the client immediately instead of waiting on it.
//
// NOTE: this does NOT send a real email. There is no email provider
// wired in - each worker just logs that it "sent" a confirmation. The
// point of this package is to demonstrate a concurrent worker-pool
// pattern (channel + goroutines + non-blocking enqueue), not to actually
// deliver notifications. Wiring in a real provider (SendGrid/SES/etc.)
// would replace the body of processJob with a real API call.
package notifier

import (
	"log"
	"time"
)

// ConfirmationJob is one unit of work for a worker to process.
type ConfirmationJob struct {
	UserEmail string
	EventName string
}

// jobs is buffered so a burst of registrations doesn't block callers of
// Enqueue while workers catch up.
var jobs = make(chan ConfirmationJob, 100)

// StartWorkers launches n background workers that pull jobs off the queue
// and process them concurrently. Call this once at startup.
func StartWorkers(n int) {
	for i := 1; i <= n; i++ {
		go worker(i)
	}
}

func worker(id int) {
	for job := range jobs {
		processJob(id, job)
	}
}

// processJob simulates the latency of a real downstream call (e.g. an
// email provider's API) and logs the result. No email is actually sent.
func processJob(workerID int, job ConfirmationJob) {
	time.Sleep(500 * time.Millisecond)
	log.Printf("[notifier worker %d] (simulated) confirmation sent to %s for event %q\n", workerID, job.UserEmail, job.EventName)
}

// Enqueue queues a confirmation job without blocking the caller. If the
// queue is full, the job is dropped rather than blocking the request.
func Enqueue(job ConfirmationJob) {
	select {
	case jobs <- job:
	default:
		log.Printf("[notifier] job queue full, dropping confirmation for %s\n", job.UserEmail)
	}
}
