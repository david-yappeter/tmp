## Summary
You need to implement a function that simulates a bank account system. Multiple users can simultaneously access and update their account balance. Your system must ensure that concurrent access does not result in race conditions.
Implement the function that:
- Deposits money into an account.
- Withdraws money from an account (ensuring there’s enough balance).
- Ensures thread-safety while handling concurrent deposits and withdrawals.

----
In this case we make an example of `BankAccount` struct which hold the amount of money, and then we create several `Deposit` and `Withdraw` within a goroutines. This will create a sample case where many deposit and withdraw are happening at the same time for a single `BankAccount` which can lead to `race` condition, it is a condition where a value are updated `twice` from the original value.

For example: when a `deposit` happen at Rp.5000, but also a `withdraw` is happening when the balance is Rp.5000, the result will be wrong if the calculation and data replace happen at the same time, which lead to `race` condition.

To prevent this things, we need to block or queue the `action` while an `action` is happening. We used `sync.Mutex` to prevent the race conditions.


With `sync.Mutex` the Final balance will always be `Rp.3000`.
