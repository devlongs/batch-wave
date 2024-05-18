# Batch-wave

Batch-wave is a REST API for creating, managing, and sending batches of transactions on the Ethereum network. It is built usivng the Gin framework and designed to be user-friendly for both technical and non-technical users. This tool is ideal for applications needing to handle multiple transactions efficiently, such as batch payments or airdrops.

## Features
- Create and Manage Batches: Create batches of transactions with specified recipient addresses and amounts.
- Send Transactions: Send all transactions in a batch to the Ethereum network with optimized gas prices.
- Transaction Retries: Automatically retry failed transactions with updated gas prices or other adjustments.
- Reporting and Logging: Generate detailed reports summarizing batch processing results, including success and failure rates, gas usage, and total cost.
