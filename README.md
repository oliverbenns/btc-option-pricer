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

## Example Output

Below is a sample of the table that is output @ 15th May 2024.

```
+----------------------+-----------+------------+----------+----------+---------------+
|        SYMBOL        |  STRIKE   | UNDERLYING | BEST ASK | BEST BID | BLACK SCHOLES |
+----------------------+-----------+------------+----------+----------+---------------+
| BTC-18MAY24-63000-C  |  63000.00 |   66202.26 |  3600.00 |  3195.00 |       3409.52 |
| BTC-18MAY24-63000-P  |  63000.00 |   66202.26 |   245.00 |   215.00 |        189.42 |
| BTC-18MAY24-63250-C  |  63250.00 |   66202.26 |  3275.00 |  2980.00 |       3196.12 |
| BTC-18MAY24-63250-P  |  63250.00 |   66202.26 |   280.00 |   255.00 |        225.94 |
| BTC-18MAY24-63500-C  |  63500.00 |   66202.26 |  3195.00 |  2770.00 |       2988.03 |
| BTC-18MAY24-63500-P  |  63500.00 |   66202.26 |   325.00 |   285.00 |        267.79 |
| BTC-18MAY24-63750-C  |  63750.00 |   66202.26 |  2940.00 |  2570.00 |       2785.72 |
| BTC-18MAY24-63750-P  |  63750.00 |   66202.26 |   375.00 |   335.00 |        315.40 |
| BTC-18MAY24-64000-C  |  64000.00 |   66202.26 |  2775.00 |  2355.00 |       2589.62 |
| BTC-18MAY24-64000-P  |  64000.00 |   66202.26 |   435.00 |   395.00 |        369.24 |
| BTC-18MAY24-64500-C  |  64500.00 |   66202.26 |  2380.00 |  2040.00 |       2217.71 |
| BTC-18MAY24-64500-P  |  64500.00 |   66202.26 |   565.00 |   530.00 |        497.18 |
+----------------------+-----------+------------+----------+----------+---------------+
```
