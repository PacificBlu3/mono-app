from flask import Blueprint

backend_api = Blueprint('backend_api',__name__)
@backend_api.route("/test")
def test():
    return '<h1>Test</h1>'