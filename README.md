This project covers APIs, CLI/web development, and real-time notifications. You can start small with a CLI-based tracker and expand it into a feature-rich web app!
 
* Key Features:-
  
  • Fetch Live Stock Data: Use APIs to get live stock prices and related data.
  
  • Historical Data: Fetch and display historical trends (optional: graph/chart visualization).
  
  • Real-Time Alerts: Notify users when a stock price crosses a certain threshold.
  
  • CLI or Web Interface: Allow users to interact via a command-line tool or a web-based dashboard.

* Tech Stack:-
  
  • Language: Go
  
  • Data Fetching: Stock Market API (e.g., Alpha Vantage, Yahoo Finance, or Twelve Data)
  
  • Visualization: CLI (text-based) or a web framework like Gin for a browser-based UI
  
  • Storage (Optional): SQLite, MongoDB, or a CSV file for storing data.

* For permanent storage I have used Sqlite3 which need to have install on your system.

SQLite provides a command-line interface (CLI) tool to interact with the database.
* Steps:-
  Install SQLite CLI:

Download and install SQLite from SQLite Downloads if it’s not already installed on your system.

• Open the Database: Run the following command in your terminal or command prompt:
  sqlite3 stocks.db

• Run SQL Commands: Inside the SQLite CLI, you can execute SQL commands to view your data:

• View all tables:
  .tables

• View table schema:
  .schema stocks

• Query data:
  SELECT * FROM stocks;

• Exit:
  .exit
