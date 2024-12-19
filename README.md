# Currency Rates CLI 💰

[![Go Version](https://img.shields.io/github/go-mod/go-version/yigitdemirko/currency-rates)](https://golang.org/)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

A simple command-line tool to fetch and display current currency rates for Turkish Lira from [doviz.com](https://www.doviz.com/).

## ✨ Features

- 📊 Real-time currency rates
- 💱 Multiple currency support (USD, EUR, GBP, etc.)
- 🔍 Filter specific currencies
- 📈 Price and change percentage display
- 🎯 Clean table format output

## 🚀 Installation

### Prerequisites

- Go 1.21 or higher
- Git (for cloning the repository)

### 📥 Local Installation

```bash
# Clone the repository
git clone https://github.com/yigitdemirko/currency-rates

# Navigate to project directory
cd currency-rates

# Install dependencies
go mod tidy

# Build the tool
go build -o currency-rates
```

### 🌍 Global Installation

1. Add Go's bin directory to your PATH (add to ~/.bashrc, ~/.zshrc, or similar):
```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

2. Install using either method:

```bash
# Method 1: From the project directory
git clone https://github.com/yigitdemirko/currency-rates
cd currency-rates
go install

# Method 2: Directly from GitHub
go install github.com/yigitdemirko/currency-rates@latest
```

After installation, `currency-rates` command will be available globally in your terminal.

## 🎯 Usage

```bash
# Display all rates
currency-rates

# Filter specific currency
currency-rates -filter DOLAR

# Show help message
currency-rates -h
# or
currency-rates --help
```

## 🛠️ Options

```
| Flag         | Description             
|--------------|-------------------------|
| -filter      | Filter by currency name |
| -h, --help   | Show help message       |
```

## 📊 Example Output

```
| Ticker       | Price        | Change    |
|--------------|--------------|-----------|
| GRAM ALTIN   | ₺2.930,85    | %0,36     |
| DOLAR        | ₺35,0932     | %0,00     |
| EURO         | ₺36,3746     | %-0,11    |
| STERLİN      | ₺43,9105     | %-0,17    |
| BIST 100     | 9.765,12     | %-1,52    |
| BITCOIN      | $96.962      | %-3,99    |
| GRAM GÜMÜŞ   | ₺32,89       | %-1,02    |
| BRENT        | $72,29       | %0,06     |
```

## 🤝 Contributing

Contributions, issues, and feature requests are welcome! Feel free to check [issues page](https://github.com/yigitdemirko/currency-rates/issues).

## 📝 License

This project is [MIT](LICENSE) licensed. 