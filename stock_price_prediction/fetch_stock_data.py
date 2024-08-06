import yfinance as yf
import pandas as pd

# Define the stock symbol and period
symbol = 'AAPL'
period = '5y'

# Fetch the data
data = yf.download(symbol, period=period)

# Save the closing prices to CSV
data = data[['Close']]
data.to_csv('stock_data.csv')
print("Stock data saved to 'stock_data.csv'.")





