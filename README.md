# ThrottlePoint 🚦🧠  
A smart reverse proxy with built-in rate limiting — soon to be AI-powered.

## 🛠 What is ThrottlePoint?
ThrottlePoint is a lightweight reverse proxy server written in Go, designed to:
- Intercept and forward HTTP requests
- Apply configurable rate limiting (currently using the Token Bucket algorithm)
- Be easily extendable for real-time AI-based intent analysis and dynamic rate limiting

Think of it as your traffic cop at the API junction — smart, fast, and soon, intuitive.

---

## ✅ Features (Implemented)

### 🌐 Reverse Proxy
- Forwards any incoming request (`/*`) to a backend server.
- Preserves method, headers, and query params.
- Pipes the response back to the client.

### ⛔ Rate Limiting (Token Bucket)
- Rate limit per IP address.
- Configurable token rate and capacity.
- Efficient and thread-safe (uses `sync.Map` and per-bucket mutex).
- Background cleanup of inactive buckets.

### ⚙️ Configurable Backend
- Uses environment variable `BACKEND_URL` to forward requests (default: `http://localhost:3000`).
- Easy to adapt for different backends.

---

## 🧠 Upcoming Features (The Real Deal)

### 🤖 AI-Enhanced Intent Prediction (Core Vision)
- Analyze request metadata (IP, headers, payload).
- Predict request “intent” or risk level using lightweight ML models.
- Smart categorization (e.g., normal traffic vs. scrapers vs. potential abuse).

### 📊 Dynamic Rate Limiting
- Adapt limits in real-time based on intent scores.
- Trusted users = more leniency.
- Suspicious traffic = throttled or blocked.

### 🔄 Feedback Loop
- Log request outcomes.
- Fine-tune AI model/rules continuously for improved predictions.
- Enable reinforcement-style learning.

### 📈 Dashboard & Metrics (Stretch)
- Web dashboard to view request logs, AI decisions, and live traffic stats.
- Manual overrides and real-time debugging.

---

## 🚀 Quickstart

```bash
# clone the repo
git clone https://github.com/asutosh2203/throttle-point.git
cd throttle-point

# create .env file
echo "BACKEND_URL=<YOUR BACKEND URL>" > .env

# run the app
go run main.go
```

---

## 📁 Project Structure
.  
├── main.go&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;# Entry point  
├── handlers/  
│   └── proxy.go&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;# Request forwarding logic  
├── middleware/  
│   └── rateLimiter.go  &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;# Rate limiting middleware   
│   └── token_bucket.go    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;# Token Bucket middleware   
├── go.mod  
├── .env                  &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;# Backend target URL  
├── .gitignore  
└── README.md  

---

## 🤝 Contributing

AI brains welcome. If you’ve got ideas on risk scoring, model integration, or smarter rate control — hop in!

---

## ⚡️ Inspiration

Born from frustration with blunt rate-limiters and a love for AI-enhanced infrastructure.

---

Maintained by [Asutosh](https://asutosh2203.netlify.app)  
Drop a ⭐ if you like where this is going!
