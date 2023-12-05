from flask import Flask, send_from_directory

app = Flask(__name__)
app.config["DEBUG"] = True

# Index route
@app.route("/test")
def test():
    return '<h1>Test</h1>'

#Main
if __name__ == "__main__":
    app.run(host='0.0.0.0', port=3300)
