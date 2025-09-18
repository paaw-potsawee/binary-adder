#  Binary Adder

A small web app that generates two random 8-bit numbers and a random operation (add, subtract, shift left, or XOR).
You enter the correct answer for the given operation to test your skills.

---

##  How to Run

1. **Clone the repository:**

   ```bash
   git clone https://github.com/yourusername/binary-adder.git
   cd binary-adder
   ```

2. **Set up environment variables:**

   * Copy `.env.example` to `.env` and adjust values for your environment, **or**
   * Export them directly as environment variables.

3. **Install dependencies:**

   ```bash
   go mod tidy
   ```

4. **Run the server:**

   ```bash
   go run main.go
   ```

5. **Open your browser:**

   ```
   http://localhost:8080
   ```

   (or whichever port your server is configured to use)

---

## How to Play

* The app will generate:

  * Two random **8-bit numbers**
  * One random operation: **Add, Subtract, Shift Left, or XOR**
* Enter the correct answer for the operation you’re given.
* **Note:**

  * The **Shift** operation means *shift left on the first number*.

---

## Requirements

* [Go installed](https://golang.org/doc/install) (latest stable version recommended)

