/*
Package deque provides a fast ring-buffer deque (double-ended queue)
implementation.

Deque generalizes a queue and a stack, to efficiently add and remove items at
either end with O(1) performance. Queue (FIFO) operations are supported using
PushBack and PopFront. Stack (LIFO) operations are supported using PushBack and
PopBack.

# Ring-buffer Performance

The ring-buffer automatically resizes by powers of two, growing when additional
capacity is needed and shrinking when only a quarter of the capacity is used,
and uses bitwise arithmetic for all calculations.

The ring-buffer implementation significantly improves memory and time
performance with fewer GC pauses, compared to implementations based on slices
and linked lists.

For maximum speed, this deque implementation leaves concurrency safety up to
the application to provide, however the application chooses, if needed at all.

# Reading Empty Deque

Since it is OK for the deque to contain the zero-value of an item, it is
necessary to either panic or return a second boolean value to indicate the
deque is empty, when reading or removing an element. The [Front], [Back],
[PopFront], and [PopBack] methods panic when the deque is empty. Use
[Deque.TryFront], [Deque.TryBack], [Deque.TryPopFront], or [Deque.TryPopBack]
to check for an empty deque with a boolean return value instead.

# Generics

Deque uses generics to create a Deque that contains items of the type
specified. To create a Deque that holds a specific type, provide a type
argument with the Deque variable declaration.
*/
package deque
