# 📌 GO TASK 1 ASSIGNMENTS

This repository contains a simple **Go (Golang)** program that demonstrates basic input handling and transaction validation.

---

## 🧾 Project Overview

This is a beginner-level Go project that:

- Prompts the user for transaction details  
- Checks whether the payment is sufficient  
- Calculates and displays change or an error message  

The program is structured into separate packages for better organization:

- `main.go` — Application entry point  
- `input/` — Handles user input  
- `validation/` — Verifies the transaction outcome  

---

## 🛠️ How It Works

1. The program asks for:
   - Item price  
   - Amount paid by the customer  

2. It checks whether the payment is sufficient.

3. If the payment is enough:
   - Displays a success message
   - Shows the calculated change

4. If the payment is insufficient:
   - Displays a rejection message
   - Shows how much money is missing

---

## 📋 Example Output

### ✅ Successful Transaction
```
Harga Barang: 15000
Uang Pembeli: 20000
[SISTEM]: Transaksi Berhasil. Kembalian Anda: 5000
```

### ❌ Failed Transaction
```
Harga Barang: 20000
Uang Pembeli: 15000
[SISTEM]: Transaksi Ditolak! Uang kurang 5000
```

---

## 🗂️ Project Structure
```
gophers_1/
├── go.mod
├── main.go
├── input/
│   └── input.go
└── validation/
    └── check.go
```
