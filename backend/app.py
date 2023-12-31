from flask import Flask, send_from_directory, request

app = Flask(__name__)
app.config["DEBUG"] = True

# Index route
@app.route("/test")
def test():
    if(request.headers.get('Authorization') is None):
        return '<h1>Unauthorized</h1>'
    else: 
        return '<h1>Authorized</h1>'

#Main
if __name__ == "__main__":
    app.run(host='0.0.0.0', port=3300)
