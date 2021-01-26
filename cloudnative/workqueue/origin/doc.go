// Package workqueue provides a simple queue that supports the following
// features:
// * Fair: items processed in the order in which they are added.
// * Stingy: a single item will not be processed multiple times concurrently,
//     and if an item is added multiple times before it can be precessed, it
//     will only be processed once.
// * Multiple consumers and producers. In particular, it is allowed for an
//     item to be reenqueued while it is being processed.
// * Shutdown notifications
package origin
