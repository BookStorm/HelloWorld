import wsgiref.simple_server

def application(environment, response):
    response('200 OK', [('Content-Type', 'text/html')])
    return ['Hello World'.encode("utf-8")]

server = wsgiref.simple_server.make_server('', 80, application)
server.serve_forever()