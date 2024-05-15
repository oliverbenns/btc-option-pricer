# BTC Option Pricer

App that calculates the value of Bybit's BTC options by using the Black-Scholes Model.

You can learn more about the model here: [https://www.investopedia.com/terms/b/blackscholes.asp](https://www.investopedia.com/terms/b/blackscholes.asp).

## Running

To run:

```
cd cmd/app
go run . -riskFreeRate=0.04
```

The `riskFreeRate` parameter is an optional floating point type input. This can be set to a fixed rate investment you have access to, such as your bank's interest rate. When not provided, it defaults to 0.05 the Bank of England rate @ 2nd August 2023.

## Disclaimer

This software is for educational purposes only and should not be used for trading purposes. All financial decisions are solely your responsibility.
