from flask import Flask
app = Flask(__name__)

@app.route('/')
def index():
    return '10 bilhões por cento de certeza!!!\n'

if __name__ == '__main__':
    app.run(host='127.0.0.1', port=8001)
